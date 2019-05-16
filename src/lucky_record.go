package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func (s *Service) luckyRecord(user string) error {
	tx := s.DB.Begin()
	if tx.Create(&LuckyRecordAll{Owner: user}).RowsAffected != 1 {
		return fmt.Errorf("insert error")
	}
	tx.Commit()
	return nil
}

func (s *Service) luckyRecordWithItem(id uint, owner string) error {
	tx := s.DB.Begin()
	if tx.Model(&LuckyRecord{}).Where(&LuckyRecord{Model: gorm.Model{ID: id}}).Updates(map[string]string{"owner": owner}).RowsAffected != 1 {
		tx.Rollback()
		return fmt.Errorf("update faild")
	}
	tx.Commit()
	return nil
}
