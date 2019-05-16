package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func (s *Service) makeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, gin.H{
		"error": 0,
		"msg":   "success",
		"data":  data,
	}
}

func (s *Service) makeErrJSON2(code int, msg interface{}) (int, interface{}) {
	return code / 100, gin.H{
		"error": code,
		"msg":   fmt.Sprint(msg),
	}
}

func (s *Service) makeErrJSON3(httpCode int, code int, msg string) (int, interface{}) {
	return httpCode, gin.H{
		"error": code,
		"msg":   fmt.Sprint(msg),
	}
}

func (s *Service) getDayBeginAndEnd() (*time.Time, *time.Time) {
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02", timeStr)
	timeDayBegin := time.Unix(t.Unix(), 0)
	timeNextDayBegin := time.Unix(t.Unix()+86400-1, 0)
	return &timeDayBegin, &timeNextDayBegin
}
