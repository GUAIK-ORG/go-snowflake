package main

import (
	"snowflake/snowflake"

	"github.com/golang/glog"
)

func main() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			s, err := snowflake.NewSnowflake(int64(i), int64(j))
			if err != nil {
				glog.Error(err)
				return
			}
			// 每个机器id生成2个sid
			for z := 0; z < 2; z++ {
				val := s.NextVal()
				if val == 0 {
					return
				}
				datacenterid, workerid := snowflake.GetDeviceID(val)
				glog.Infof("sid:%d datacenterid:%d workerid:%d create_time:%s", val, datacenterid, workerid, snowflake.GetGenTime(val))
			}
		}
	}
	// 获取时间戳字段已经使用的占比（0.0 - 1.0）
	// 默认开始时间为：2020年01月01日 00:00:00
	glog.Infof("Timestamp status: %f", snowflake.GetTimestampStatus())
}
