package models

import (
	"github.com/jinzhu/gorm"
)

// Stuff Type
type Stuff struct {
	gorm.Model
	Name  string `json:"name"`
	Image string `json:"image"`
}

// TableName stuff
func (b *Stuff) TableName() string {
	return "stuff"
}
