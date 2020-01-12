package main

import (
	"snowflake/snowflake"

	"github.com/golang/glog"
)

func main() {
	s, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		glog.Error(err)
		return
	}
	for i := 0; i < 1024; i++ {
		val := s.NextVal()
		if val == 0 {
			return
		}
		glog.Info(val)
	}
}
