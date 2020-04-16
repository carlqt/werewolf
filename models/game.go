package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Game struct {
	ID         string `json:"id" db:"id"`
	State      int    `json:"state" db:"state"`
	Phase      int    `json:"phase" db:"phase"`
	PhaseCount int    `json:"phase_count" db:"phase_count"`
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
}
