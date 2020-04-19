package models

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type gameState int

const (
	Started gameState = iota
	WaitingForPlayers
	InProgress
	End
)

type Game struct {
	ID         int       `json:"id" db:"id"`
	State      gameState `json:"state" db:"state"`
	Phase      int       `json:"phase" db:"phase"`
	PhaseCount int       `json:"phase_count" db:"phase_count"`
	ChannelID  string    `json:"channel_id" db:"channel_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// Finds an active game on the channel (unique)
func ActiveGameOnChannel(db *sqlx.DB, channelID string) (*Game, error) {
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

func NewGame(db *sqlx.DB, channelID string) error {
	game, err := ActiveGameOnChannel(db, channelID)
	if game != nil {
		return errors.New("Game is currently in progress")
	} else if err != nil && err != sql.ErrNoRows {
		return err

	stmt, err := db.Prepare("INSERT INTO	games(channel_id) VALUES($1)")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(channelID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (game *Game) Create(db *sqlx.DB) error {
	stmt, err := db.Prepare("INSERT INTO	games(state, phase, phase_count, channel_id) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(game.State, game.Phase, game.PhaseCount, game.ChannelID)

	return err
}
