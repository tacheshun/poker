package types

type Player struct {
	Name string `json:"name"`
	Hand *Hand  `json:"hand"`
}
