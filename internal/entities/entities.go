package entities

import "github.com/jmoiron/sqlx"

type Entities struct {
	Games GameEntity
}

func NewEntities(db *sqlx.DB) Entities {
	return Entities{
		Games: GameEntity{DB: db},
	}
}
