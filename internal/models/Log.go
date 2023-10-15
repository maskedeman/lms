package models

import (
	"time"

	"gorm.io/gorm"
)

// Priority:=make(map[string]int)
// type Priority struct {
// 	Low      int
// 	Medium   int
// 	High     int
// 	Critical int
// }
type Priority int

type Log struct {
	gorm.Model

	name      string
	priority  Priority
	info      string
	timestamp time.Time
}
