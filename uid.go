// This package provides unique id in distribute system
// the algorithm is inspired by Twitter's famous snowflake
// its link is: https://github.com/twitter/snowflake/releases/tag/snowflake-2010
//

// 0               41	           51			 64
// +---------------+----------------+------------+
// |timestamp(ms)  | worker node id | sequence	 |
// +---------------+----------------+------------+

// Copyright (C) 2016 by zheng-ji.info

// copy from https://github.com/zheng-ji/goSnowFlake/blob/master/uid.go
// by frank

package common

import (
	"errors"
	"github.com/dahaipublic/common/base58"
	"strconv"
	"sync"
	"time"
)

const (
	// 1735660800000 = 2025-01-01
	CEpoch = 1735660800000 // 原代码是 1474802888000，2016-09-25 19:28:08

	CWorkerIdBits  = 10 // Num of WorkerId Bits
	CSenquenceBits = 12 // Num of Sequence Bits

	CWorkerIdShift  = 12
	CTimeStampShift = 22

	CSequenceMask = 0xfff // equal as getSequenceMask()
	CMaxWorker    = 0x3ff // equal as getMaxWorkerId()
)

// IdWorker Struct
type IdWorker struct {
	workerId      int64
	lastTimeStamp int64
	sequence      int64
	maxWorkerId   int64
	lock          *sync.Mutex
}

// NewIdWorker Func: Generate NewIdWorker with Given workerid
func NewIdWorker(workerid int64) (iw *IdWorker, err error) {
	iw = new(IdWorker)

	iw.maxWorkerId = getMaxWorkerId()

	if workerid > iw.maxWorkerId || workerid < 0 {
		return nil, errors.New("worker not fit")
	}
	iw.workerId = workerid
	iw.lastTimeStamp = -1
	iw.sequence = 0
	iw.lock = new(sync.Mutex)
	return iw, nil
}

func getMaxWorkerId() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func getSequenceMask() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (iw *IdWorker) timeGen() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (iw *IdWorker) timeReGen(last int64) int64 {
	ts := time.Now().UnixNano() / 1000 / 1000
	for {
		if ts <= last {
			ts = iw.timeGen()
		} else {
			break
		}
	}
	return ts
}

// NewId Func: Generate next id
func (iw *IdWorker) NextId() (ts int64, err error) {
	iw.lock.Lock()
	defer iw.lock.Unlock()
	ts = iw.timeGen()
	if ts == iw.lastTimeStamp {
		iw.sequence = (iw.sequence + 1) & CSequenceMask
		if iw.sequence == 0 {
			ts = iw.timeReGen(ts)
		}
	} else {
		iw.sequence = 0
	}

	if ts < iw.lastTimeStamp {
		err = errors.New("Clock moved backwards, Refuse gen id")
		panic(err)
		return 0, err
	}
	iw.lastTimeStamp = ts
	ts = (ts-CEpoch)<<CTimeStampShift | iw.workerId<<CWorkerIdShift | iw.sequence
	return ts, nil
}

func (iw *IdWorker) NextBase10Id() string {
	id, _ := iw.NextId()
	return strconv.FormatInt(id, 10)
}

func (iw *IdWorker) NextBase16Id() string {
	id, _ := iw.NextId()
	return strconv.FormatInt(id, 16)
}

func (iw *IdWorker) NextBase58Id() string {
	id, _ := iw.NextId()
	return base58.EncodeXX(id)
}

var IDWorker *IdWorker = nil

/*
 *@note 生成UUID生成器对象
 */
func NewIDWorker(id uint64) {
	worker, err := NewIdWorker(int64(id))
	if err != nil {
		panic("NewIdWorker fail!")
	}

	IDWorker = worker
}

func NewBase58UUID() string {
	return IDWorker.NextBase58Id()
}

func NewBase16UUID() string {
	return IDWorker.NextBase16Id()
}

func NewBase10UUID() string {
	return IDWorker.NextBase10Id()
}
