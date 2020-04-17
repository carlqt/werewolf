package models

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type gameState int

const (
	Started gameState = iota
	WaitingForPlayers
	InProgress
)

type Game struct {
	ID         string    `json:"id" db:"id"`
	State      gameState `json:"state" db:"state"`
	Phase      int       `json:"phase" db:"phase"`
	PhaseCount int       `json:"phase_count" db:"phase_count"`
	ChannelID  string    `json:"channel_id" db:"channel_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// Finds an active game on the channel (unique)
func ActiveGameOnChannel(db *sqlx.DB, channelID string) (*Game, error) {
	game := Game{
		State:     WaitingForPlayers,
		ChannelID: channelID,
	}

	err := db.Get(&game, "SELECT * FROM games WHERE state = $1 AND channel_id = $2", game.State, game.ChannelID)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func NewGame(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO games DEFAULT VALUES")
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (game *Game) Create(db *sqlx.DB) error {
	stmt, err := db.Prepare("INSERT INTO	games(state, phase, phase_count) VALUES($1, $2, $3)")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(game.State, game.Phase, game.PhaseCount)

	return err
}
