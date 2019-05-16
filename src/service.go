package main

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Service struct {
	DB     *gorm.DB
	Router *gin.Engine
	Conf   *conf
}

type conf struct {
	Port string
	DB   struct {
		Dsn string
	}
	Cors bool
}

func (s *Service) init() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	s.initConf()
	s.initDB()
	s.initTable()
	s.initRouter()
}

func (s *Service) initDB() {
	db, err := gorm.Open("mysql", s.Conf.DB.Dsn)
	if err != nil {
		panic(err)
	}
	s.DB = db
}

func (s *Service) initTable() {
	//insert DB struct here
	iss := []interface{}{&Lucky{}, &LuckyItem{}, &LuckyRecord{}}
	for _, v := range iss {
		if !s.DB.HasTable(v) {
			s.DB.CreateTable(v)
		}
	}
}

func (s *Service) initConf() {
	c := new(conf)
	_, err := toml.DecodeFile("config_lucky/conf.toml", c)
	if err != nil {
		panic(err)
	}
	s.Conf = c
}
