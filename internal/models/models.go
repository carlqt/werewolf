package models

import (
	"github.com/carlqt/internal/entities"
)

type Models struct {
	Games GameModel
}

func NewModels(entities entities.Entities) Models {
	return Models{
		Games: GameModel{GameEntity: entities.Games},
	}
}
