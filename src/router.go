package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Service) initRouter() {

	s.Router = gin.Default()

	if s.Conf.Cors {
		cf := cors.DefaultConfig()
		cf.AllowAllOrigins = true
		cf.AddAllowMethods("POST", "GET", "DELETE", "PUT", "OPTION", "PATCH")
		cf.AddAllowHeaders("Authorization,Content-type,User-Agent")
		s.Router.Use(cors.New(cf))
	}

	s.Router.GET("/lucky", func(c *gin.Context) {

	})

	s.Router.POST("/lucky", func(c *gin.Context) {
		c.JSON(s.LuckyCreateHandler(c))
	})

	err := s.Router.Run(s.Conf.Port)
	panic(err)

}
