package types

import "fmt"

type Card struct {
	Suit  Suit `json:"suit"`
	Value int  `json:"value"`
}

func (c *Card) Print() {
	fmt.Println(c.Value, "of", c.Suit)
}
