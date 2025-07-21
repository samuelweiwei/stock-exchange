package snowflake

import (
	"fmt"
	"sync"
	"time"
)

const (
	epoch         = 1609459200000              // 自定义纪元时间（例如 2021-01-01 00:00:00 UTC）
	machineIDBits = 10                         // 机器 ID 位数
	sequenceBits  = 12                         // 序列号位数
	maxMachineID  = -1 ^ (-1 << machineIDBits) // 最大机器 ID
	sequenceMask  = -1 ^ (-1 << sequenceBits)  // 最大序列号
)

type Snowflake struct {
	mu            sync.Mutex
	machineID     int64
	sequence      int64
	lastTimestamp int64
}

// NewSnowflake 创建新的 Snowflake 实例
func NewSnowflake(machineID int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, ErrInvalidMachineID
	}
	return &Snowflake{
		machineID:     machineID,
		sequence:      0,
		lastTimestamp: -1,
	}, nil
}

// Generate 生成唯一 ID
func (s *Snowflake) Generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	if timestamp == s.lastTimestamp {
		// 在同一毫秒内，增加序列号
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果序列号已满，等待下一毫秒
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	} else {
		// 不同毫秒，重置序列号
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	// 生成 ID
	id := ((timestamp - epoch) << (machineIDBits + sequenceBits)) | (s.machineID << sequenceBits) | s.sequence
	return id
}

// 错误类型
var ErrInvalidMachineID = fmt.Errorf("invalid machine ID")
