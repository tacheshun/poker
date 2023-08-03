package usecases

import "github.com/tacheshun/poker/internal/types"

type GameRound struct {
	Players []*types.Player `json:"players"`
	Deck    *types.Deck     `json:"deck"`
}

func NewGameRound() *GameRound {
	return &GameRound{
		Players: []*types.Player{},
		Deck:    types.NewDeck(),
	}
}

func (g *GameRound) play() {
	g.Deck.Shuffle()
	for i := 0; i < 5; i++ {
		for _, player := range g.Players {
			player.Hand.AddCard(g.Deck.Draw())
		}
	}
}
