package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Player struct {
	ID     string `json:"id" db:"id"`
	GameID int    `json:"game_id" db:"game_id"`
	RoleID int    `json:"role_id" db:"role_id"`
	Name   int    `json:"name" db:"name"`
	State  int    `json:"state" db:"state"`
}

func (player *Player) Create(db *sqlx.DB) error {
	stmt, err := db.Prepare("INSERT INTO players(game_id, role_id, name, state) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		log.Println(err)
		return err
	}

	return _, stmt.Exec(player.GameID, player.RoleID, player.Name, player.State)
}
