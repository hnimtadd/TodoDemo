package model

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"time"
)

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

func (this *Todo) UnmarshallJSON(b []byte) error {
	todo := Todo{}
	if err := json.Unmarshal(b, &todo); err != nil {
		return err
	}
	fmt.Println(todo)
	*this = todo
	return nil
}
