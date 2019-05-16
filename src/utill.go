package main

import (
	"fmt"
	"math/rand"
	"time"
)

func (s *Service) int64ToTime(i int64) *time.Time {
	vt := time.Unix(i, 0)
	return &vt
}

func (s *Service) isEmpty(input ...string) error {
	for _, v := range input {
		if v == "" {
			return fmt.Errorf("missing param")
		}
	}
	return nil
}

func (s *Service) recordListGen(d []*LuckyItem, baseInfo *LuckyCreateBody) ([]*LuckyRecord, error) {
	listN := make([]*LuckyRecord, 0, 1000)
	//init array
	count := 0
	for _, v := range d {
		for i := 0; i < v.Count; i++ {
			listN = append(listN, &LuckyRecord{LuckyID: v.LuckyID, ItemID: v.ID})
		}
		count += v.Count
	}
	timeScrap := (baseInfo.StartTime - baseInfo.EndTime) / int64(count)
	fmt.Println(timeScrap)
	//resort
	for i := 0; i < count; i++ {
		j := rand.Int() % count
		listN[i], listN[j] = listN[j], listN[i]
	}
	for i := 0; i < count; i++ {
		j := rand.Int() % count
		listN[i], listN[j] = listN[j], listN[i]
	}

	//give time
	for i := 0; i < count; i++ {
		listN[i].ShouldTime = s.int64ToTime(int64(baseInfo.StartTime + int64(i)*timeScrap))
	}
	return listN, nil
}
