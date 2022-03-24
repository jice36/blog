package models

import "time"

type Note struct {
	Header     string    `json:"header"`
	Text       string    `json:"text"`
	Tags       []string  `json:"tags"`
	TimeCreate time.Time `json:"time_create"`
}

type NotesResp struct {
	Notes []*Note `json:"notes"`
}

type SendNoteReq struct {
	UserID string   `json:"user_id"`
	Header string   `json:"header"`
	Text   string   `json:"text"`
	Tags   []string `json:"tags"`
}
