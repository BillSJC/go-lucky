package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Lucky struct {
	gorm.Model
	Name                string
	StartTime           *time.Time
	EndTime             *time.Time
	TimesPerDay         int
	AllowReLuckyWhenGet bool
}

type LuckyItem struct {
	gorm.Model
	LuckyID uint
	Name    string
	Count   int
}

type LuckyRecord struct {
	gorm.Model
	ShouldTime *time.Time
	RealTime   *time.Time
	ItemID     uint
	LuckyID    uint
	Owner      string
}

type LuckyRecordAll struct {
	gorm.Model
	Owner string
}
