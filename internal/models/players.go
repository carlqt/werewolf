package models

import (
	"errors"
	"log"

	"github.com/carlqt/internal/entities"
)

type IPlayer interface {
	Create(*entities.Player) error
	ExistsInGame(int, string) (bool, error)
}

type PlayerModel struct {
	PlayerEntity IPlayer
	GameEntity   GameInterface
}

func (p PlayerModel) JoinGame(gameChannelID string, player *entities.Player) error {
	game, err := p.GameEntity.FindJoinableGame(gameChannelID)

	if err != nil || game == nil {
		return errors.New("no games found")
	}

	player.GameID = game.ID

	exists, err := p.PlayerEntity.ExistsInGame(game.ID, player.ExternalID)
	if err != nil {
		log.Println(err)
		return errors.New("could not connect to DB")
	} else if exists {
		return errors.New("player has already joined")
	}

	err = p.PlayerEntity.Create(player)

	if err != nil {
		return errors.New("unable to join")
	}

	return nil
}
