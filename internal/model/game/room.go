package game

import "time"

type Room struct {
	ID        string    `json:"id"`
	Players   []string  `json:"players"`
	CreatedAt time.Time `json:"createdAt"`
}
