package models

import (
	"errors"
	"log"
	"time"

	"github.com/carlqt/internal/entities"
)

type GameModel struct {
	GameEntity entities.GameEntity
}

// NewGame creates a new game row if there is no active game for that channel
func (g GameModel) NewGame(channelID string) (*entities.Game, error) {
	game, err := g.GameEntity.FindActiveGameByChannelID(channelID)
	if err != nil {
		return game, err
	} else if game != nil {
		return game, errors.New("game is in progress")
	}

	game = &entities.Game{
		State:     entities.Started,
		ChannelID: channelID,
		CreatedAt: time.Now(),
	}

	err = g.GameEntity.Create(game)

	if err != nil {
		log.Println(err)
		return game, err
	}

	return game, nil
}
