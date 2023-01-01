package rdb

import "time"

type CacheRepository interface {
	SetDataWithExpiry(key, value string, expiredPeriod time.Duration) error
	GetData(key string) (string, error)
	FlushData() error
	Ping()
}
