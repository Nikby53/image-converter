package models

import (
	"time"
)

type Request struct {
	Filename      string    `json:"filename"`
	Status        string    `json:"status"`
	SourceFormat  string    `json:"source_format"`
	TargetFormat  string    `json:"target_format"`
	Ratio         int       `json:"ratio"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
	OriginalImgID string    `json:"original_img_id"`
	TargetImgID   string    `json:"target_img_id"`
}
