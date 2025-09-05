package models

import "time"

type Document struct {
	ID           string    `json:"id"`
	Filename     string    `json:"filename"`
	UploadedAt   time.Time `json:"uploaded_at"`
	OriginalText string    `json:"original_text"`
	Summary      string    `json:"summary"`
}
