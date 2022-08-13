package models

import (
	"testing"

	"github.com/carlqt/internal/entities"
)

type mockEntity struct{}

// Game exists
func (m mockEntity) FindActiveGameByChannelID(channelID string) (*entities.Game, error) {
	return new(entities.Game), nil
}

func (m mockEntity) FindJoinableGame(channelID string) (*entities.Game, error) {
	return new(entities.Game), nil
}

func (m mockEntity) Create(*entities.Game) error {
	return nil
}

// scenario: Successful, when active game exists, when error
func TestNewGame(t *testing.T) {
	entity := mockEntity{}
	game := GameModel{GameEntity: entity}

	_, err := game.NewGame("abc")

	if err == nil {
		t.Errorf("Shoud return `Game is in progress` error")
	}
}
