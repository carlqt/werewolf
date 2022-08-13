package entities

import (
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type PlayerEntity struct {
	DB *sqlx.DB
}

type Player struct {
	ID         int       `json:"id" db:"id"`
	GameID     int       `json:"game_id"`
	Name       string    `json:"name"`
	ExternalID string    `json:"external_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

func (p PlayerEntity) Create(player *Player) error {
	stmt, err := p.DB.Prepare("INSERT INTO players(game_id, name, external_id) VALUES($1, $2, $3)")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(player.GameID, player.Name, player.ExternalID)

	return err
}

func (p PlayerEntity) ExistsInGame(gameID int, externalID string) (bool, error) {
	var id int

	err := p.DB.QueryRowx("SELECT game_id FROM players WHERE external_id = $1 AND game_id = $2", externalID, gameID).Scan(&id)

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
