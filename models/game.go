package models

import (
	"database/sql"
	"errors"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type gameState int

const (
	Started gameState = iota
	WaitingForPlayers
	InProgress
	End
)

const tableName = "games"

type Game struct {
	ID         int       `json:"id" db:"id"`
	State      gameState `json:"state" db:"state"`
	Phase      int       `json:"phase" db:"phase"`
	PhaseCount int       `json:"phase_count" db:"phase_count"`
	ChannelID  string    `json:"channel_id" db:"channel_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// Finds an active game on the channel (unique)
func ActiveGameOnChannel(channelID string) (*Game, error) {
	game := Game{
		ChannelID: channelID,
	}

	err := db.Get(&game, "SELECT * FROM games WHERE state != $1 AND channel_id = $2", End, game.ChannelID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &game, nil
}

func NewGame(channelID string) error {
	game, err := ActiveGameOnChannel(channelID)
	if game != nil {
		return errors.New("Game is currently in progress")
	} else if err != nil && err != sql.ErrNoRows {
		return err
	}

	// Skip Starting state and move immediately to WaitingForPlayers state
	stmt, err := db.Prepare("INSERT INTO	games(channel_id, state) VALUES($1, $2)")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(channelID, WaitingForPlayers)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (game *Game) Create() error {
	stmt, err := db.Prepare("INSERT INTO	games(state, phase, phase_count, channel_id) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(game.State, game.Phase, game.PhaseCount, game.ChannelID)

	return err
}

// NextEvent is a sad excuse for a State Machine
func (game *Game) NextEvent() error {
	game.State++

	_, err := db.NamedExec("UPDATE games SET state=:state", game.State)

	return err
}

func Find(builder sq.SelectBuilder) (Game, error) {
	var game Game
	err := builder.RunWith(db).QueryRow().Scan(&game.ID, &game.State, &game.ChannelID)

	if err != nil {
		return game, err
	}

	return game, nil
}

// Join creates a player on an active game that is on WaitingForPlayer state
func Join(channelID string, player Player) (Game, Player, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// Search for game that is in WaitingForPlayers state
	// condition := sq.Eq{"channel_id", channelID, "state", WaitingForPlayers}
	sqBuilder := psql.Select("id, state, channel_id").From("games").Where("channel_id = ? AND state = ?", channelID, WaitingForPlayers)
	game, err := Find(sqBuilder)

	if err != nil {
		return game, Player{}, err
	}

	// Create a player and associate it on the game
	player.GameID = game.ID
	err = player.Create()
	if err != nil {
		return game, player, err
	}

	return game, player, nil
	// If there's an error, return an error

	// validations
	// If no game is found - Error
	// If player create fails - Error
	// If player is joining a game but he is not in that channel - error
}
