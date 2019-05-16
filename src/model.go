package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Lucky struct {
	gorm.Model
	Name      string
	Items     []LuckyItem
	StartTime *time.Time
	EndTime   *time.Time
}

type LuckyItem struct {
	gorm.Model
	LuckyID uint
	Name    string
	Count   int
}
