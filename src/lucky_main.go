package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

func (s Service) getLuckyHandler(c *gin.Context) (int, interface{}) {
	err := s.Auth(c)
	if err != nil {
		return s.makeErrJSON2(40100, "unauthorized")
	}

	user := c.Query("user")
	ids := c.Query("id")
	id, err := strconv.Atoi(ids)
	if err != nil || id <= 0 {
		return s.makeErrJSON2(40330, "id invalid")
	}
	uid := uint(id)
	name, errCode, err := s.getLucky(user, uid)
	if err != nil {
		return s.makeErrJSON2(errCode, err)
	}
	return s.makeSuccessJSON(gin.H{
		"rewards": name,
	})
}

func (s *Service) getLucky(user string, id uint) (name string, errCode int, err error) {
	l := new(Lucky)
	s.DB.Where(&Lucky{Model: gorm.Model{ID: id}}).Find(l)
	// not found
	if l.Name == "" {
		return "", 40400, fmt.Errorf("lucky not found")
	}
	// not in time
	if l.StartTime.Unix() > time.Now().Unix() {
		return "", 40300, fmt.Errorf("lucky not in range")
	}
	if l.EndTime.Unix() < time.Now().Unix() {
		return "", 40301, fmt.Errorf("lucky not in range")
	}

	//out of time
	if l.TimesPerDay > 0 {
		records := make([]LuckyRecordAll, 0, 100)
		tb, te := s.getDayBeginAndEnd()
		s.DB.Model(&LuckyRecordAll{}).Where(&LuckyRecordAll{Owner: user}).Where("created_at > ? AND created_at < ?", tb, te).Find(&records)
		if len(records) >= l.TimesPerDay {
			return "", 40320, fmt.Errorf("lucky out of time")
		}
	}

	//out of times
	if !l.AllowReLuckyWhenGet {
		r := new(LuckyRecord)
		s.DB.Model(&LuckyRecord{}).Where(&LuckyRecord{Owner: user, LuckyID: id}).Find(r)
		if r.ShouldTime != nil {
			return "", 40321, fmt.Errorf("lucky out of times")
		}
	}

	//check if lucky
	item := new(LuckyRecord)
	s.DB.Where(&LuckyRecord{Model: gorm.Model{ID: id}}).Where("should_time < ? AND owner = ''", time.Now()).Find(item)
	//log
	err = s.luckyRecord(user)
	if err != nil {
		return "", 50011, fmt.Errorf("DB error")
	}

	itemName := ""
	if item.ItemID <= 0 {
		//unlucky
	} else {
		//lucky
		itemID := item.ItemID
		lt := new(LuckyItem)
		s.DB.Where(&LuckyItem{Model: gorm.Model{ID: itemID}}).Find(lt)
		if lt.Name == "" {
			return "", 50012, fmt.Errorf("DB error")
		}
		//done and insert data
		err = s.luckyRecordWithItem(item.ID, user)
		if err != nil {
			return "", 50013, fmt.Errorf("DB error")
		}
		itemName = lt.Name
	}
	return itemName, 0, nil
}
