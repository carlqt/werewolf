package models

import (
	"database/sql"
)

type Player struct {
	ID     int           `json:"id" db:"id"`
	GameID int           `json:"game_id" db:"game_id"`
	RoleID sql.NullInt64 `json:"role_id" db:"role_id"`
	Name   string        `json:"name" db:"name"`
	State  int           `json:"state" db:"state"`
}

// CreatePlayer creates a player
func CreatePlayer(player *Player) error {
	_, err := db.NamedExec("INSERT INTO players(game_id, role_id, name, state) VALUES(:game_id, :role_id, :name, :state)", player)

	if err != nil {
		return err
	}

	return nil
}
