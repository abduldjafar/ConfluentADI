package model

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// structur of Userhara table
type WordCount struct {
	Word   string `gorm:"primary_key"`
	Jumlah int
}
