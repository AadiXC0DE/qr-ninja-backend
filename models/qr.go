package models

import "github.com/jinzhu/gorm"

type QRCode struct {
	gorm.Model
	URL   string `json:"url"`
	Date  string `json:"date"`
	Image string `json:"image"`
}
