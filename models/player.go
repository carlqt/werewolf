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

func PlayerExists(gameID int, name string) (bool, error) {
	var id int
	err := db.QueryRowx("SELECT id FROM players WHERE name = $1 AND game_id = $2", name, gameID).Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		return true, err
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}
