package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (s *Service) Auth(c *gin.Context) error {
	// set up your own auth model here

	auth := c.GetHeader("Authorization")
	if "sign "+s.Conf.Key != auth {
		return fmt.Errorf("unauthorized")
	}

	return nil
}
