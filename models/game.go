package models

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type GameState int

const (
	Started GameState = iota
	WaitingForPlayers
	InProgress
	End
)

const tableName = "games"

type Game struct {
	ID         int       `json:"id" db:"id"`
	State      GameState `json:"state" db:"state"`
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

// Join creates a player on an active game that is on WaitingForPlayer state
func Join(channelID string, player Player) error {
	var game Game
	err := db.Get(&game, "SELECT * FROM games WHERE channel_id = $1 AND state = $2", channelID, WaitingForPlayers)

	if err != nil {
		return err
	}

	// Create a player and associate it on the game
	player.GameID = game.ID

	playerExists, _ := PlayerExists(game.ID, player.Name)
	if playerExists {
		return errors.New("Player has already joined this game")
	}

	err = CreatePlayer(&player)
	if err != nil {
		return err
	}

	return nil
}
