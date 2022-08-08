package entities

import (
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type GameState int

const (
	Started GameState = iota
	WaitingForPlayers
	InProgress
	End
)

type GameEntity struct {
	DB *sqlx.DB
}

type Game struct {
	ID         int       `json:"id" db:"id"`
	State      GameState `json:"state" db:"state"`
	Phase      int       `json:"phase" db:"phase"`
	PhaseCount int       `json:"phase_count" db:"phase_count"`
	ChannelID  string    `json:"channel_id" db:"channel_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

func (g GameEntity) Create(game *Game) error {
	stmt, err := g.DB.Prepare("INSERT INTO	games(state, phase, phase_count, channel_id) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(game.State, game.Phase, game.PhaseCount, game.ChannelID)

	return err
}

func (g GameEntity) FindActiveGameByChannelID(channelID string) (*Game, error) {
	game := Game{
		ChannelID: channelID,
	}

	err := g.DB.Get(&game, "SELECT * FROM games WHERE state != $1 AND channel_id = $2", End, game.ChannelID)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}

	return &game, nil
}
