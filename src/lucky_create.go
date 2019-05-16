package main

import (
	"fmt"
	"math/rand"
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
	//build struct Lucky
	dLucky := &Lucky{Name: data.Name, StartTime: s.int64ToTime(data.StartTime), EndTime: s.int64ToTime(data.EndTime)}

	//begin tx
	tx := s.DB.Begin()
	//create Object Lucky
	if tx.Create(dLucky).RowsAffected != 1 {
		tx.Rollback()
		return 0, 50000, fmt.Errorf("error when insert data")
	}
	tempData1 := new(Lucky)
	//get newest ID
	tx.Model(&Lucky{}).Order("create_at DESC").Find(tempData1)
	if tempData1.Name == "" {
		tx.Rollback()
		return 0, 50001, fmt.Errorf("error when insert data")
	}
	LuckyID := tempData1.ID

	//build Objects LuckyItem
	items := make([]*LuckyItem, 0, 100)
	for _, v := range data.Item {
		items = append(items, &LuckyItem{Name: v.Name, Count: v.Count, LuckyID: LuckyID})
	}
	//Create Objects LuckyItem
	if tx.Create(items).RowsAffected != int64(len(items)) {
		tx.Rollback()
		return 0, 50002, fmt.Errorf("error when insert data")
	}

	//get items data
	items2 := make([]*LuckyItem, 0, len(items))
	tx.Model(&LuckyItem{}).Where(&LuckyItem{LuckyID: LuckyID}).Find(&items2)
	if len(items2) != len(items) {
		tx.Rollback()
		return 0, 50003, fmt.Errorf("error when insert data")
	}

	records, err := s.recordListGen(items2, data)

	if err != nil {
		tx.Rollback()
		return 0, 50004, fmt.Errorf("error when insert data")
	}

	if tx.Create(records).RowsAffected != int64(len(records)) {
		tx.Rollback()
		return 0, 50005, fmt.Errorf("error when insert data")
	}

	return LuckyID, 0, nil

}
