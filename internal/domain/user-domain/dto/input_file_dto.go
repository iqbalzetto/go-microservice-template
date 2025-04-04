package dto

import "io"

type InputFileDTO struct {
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	Type      string    `json:"type"`
	Extension string    `json:"extension"`
	Reader    io.Reader `json:"reader"`
}
