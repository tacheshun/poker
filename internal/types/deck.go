package types

import (
	"github.com/google/uuid"
	"math/rand"
)

type Deck struct {
	ID    string `json:"id"`
	Cards []Card `json:"cards"`
}

func NewDeck() *Deck {
	var cards = []Card{}
	for _, suit := range []Suit{Clubs, Diamonds, Hearts, Spades} {
		for _, value := range []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15} {
			cards = append(cards, Card{
				Suit:  suit,
				Value: value,
			})
		}
	}
	return &Deck{
		ID:    uuid.NewString(),
		Cards: cards,
	}
}

func (d *Deck) print() {
	for _, card := range d.Cards {
		card.print()
	}
}

func (d *Deck) draw() Card {
	d.Cards = d.Cards[1:]
	return d.Cards[0]
}

func (d *Deck) shuffle() {
	for i := range d.Cards {
		j := rand.Intn(len(d.Cards))
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
