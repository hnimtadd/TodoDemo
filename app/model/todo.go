package model

import "time"

type Todo struct {
	Id          string    `json:"id"`
	Content     string    `json:"content"`
	Last_update time.Time `json:"last_update"`
	Is_deleted  bool      `json:"is_deleted"`
	Metadata    Metadata  `json:"metadata"`
}
type Metadata struct {
	Size   int    `json:"size"`
	Format string `json:"format"`
	Source string `json:"source"`
}
