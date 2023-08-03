package types

import "fmt"

type Hand struct {
	Cards []Card `json:"cards"`
	Score int    `json:"score"`
}

func NewHand() *Hand {
	return &Hand{
		Cards: []Card{},
		Score: 0,
	}
}

func (h *Hand) AddCard(c Card) {
	h.Cards = append(h.Cards, c)
	h.Score += c.Value
	fmt.Println("Score:", h.Score)
}
