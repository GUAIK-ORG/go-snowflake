package snowflake

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
)

const (
	epoch             = int64(1577808000000)                           // 设置起始时间：2020-01-01 00:00:00，有效期69年
	timestampBits     = uint(41)                                       // 纪元占用位数
	datacenteridBits  = uint(5)                                        // 数据标示id所占位数
	workeridBits      = uint(5)                                        // 机器id所占位数
	sequenceBits      = uint(12)                                       // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))              // 时间戳最大值
	datacenteridMax   = int64(-1 ^ (-1 << datacenteridBits))           // 支持的最大数据标示id数量
	workeridMax       = int64(-1 ^ (-1 << workeridBits))               // 支持的最大机器id数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))               // 支持的最大序列id数量
	workeridShift     = sequenceBits                                   // 机器id左移位数
	datacenteridShift = sequenceBits + workeridBits                    // 数据标示左移位数
	timestampShift    = sequenceBits + workeridBits + datacenteridBits // 时间戳左移位数
)

type Snowflake struct {
	sync.Mutex
	timestamp    int64
	workerid     int64
	datacenterid int64
	sequence     int64
}

func NewSnowflake(datacenterid, workerid int64) (*Snowflake, error) {
	if datacenterid < 0 || datacenterid > datacenteridMax {
		return nil, fmt.Errorf("datacenterid must be between 0 and %d", datacenteridMax-1)
	}
	if workerid < 0 || workerid > workeridMax {
		return nil, fmt.Errorf("workerid must be between 0 and %d", workeridMax-1)
	}
	return &Snowflake{
		timestamp:    0,
		datacenterid: datacenterid,
		workerid:     workerid,
		sequence:     0,
	}, nil
}

func (s *Snowflake) NextVal() int64 {
	now := time.Now().UnixNano() / 1000000
	s.Lock()
	if s.timestamp == now {
		// 当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		s.Unlock()
		glog.Errorf("epoch must be between 0 and %d", timestampMax-1)
		return 0
	}
	s.timestamp = now
	r := int64((t)<<timestampShift | (s.datacenterid << datacenteridShift) | (s.workerid << workeridShift) | (s.sequence))
	s.Unlock()
	return r
}
