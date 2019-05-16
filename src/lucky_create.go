package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

/*
func (s *Service)getLucky(openid string)(result string,status int,err error){

}
*/
type LuckyCreateBody struct {
	Name      string
	StartTime int64
	EndTime   int64
	Item      []struct {
		Name  string
		Count int
	}
}

func (s *Service) createLucky(data *LuckyCreateBody) (id uint, status int, err error) {
	//pre check data
	if s.isEmpty(data.Name) != nil {
		return 0, 40000, fmt.Errorf("missing data")
	}
	if data.StartTime > data.EndTime {
		return 0, 40001, fmt.Errorf("start time should earler then end time")
	}
	if len(data.Item) == 0 {
		return 0, 40002, fmt.Errorf("no items")
	}
	for _, v := range data.Item {
		if s.isEmpty(v.Name) != nil || v.Count < 0 {
			return 0, 40003, fmt.Errorf("item data error")
		}
	}
	dLucky := &Lucky{Name: data.Name, StartTime: s.int64ToTime(data.StartTime), EndTime: s.int64ToTime(data.EndTime)}
}
