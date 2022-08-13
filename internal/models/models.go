package models

import (
	"github.com/carlqt/internal/entities"
)

type Models struct {
	Games   GameModel
	Players PlayerModel
}

func NewModels(entities entities.Entities) Models {
	return Models{
		Games:   GameModel{GameEntity: entities.Games},
		Players: PlayerModel{GameEntity: entities.Games, PlayerEntity: entities.Players},
	}
}
