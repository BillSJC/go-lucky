package main

import (
	"fmt"
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
