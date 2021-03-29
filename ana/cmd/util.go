package cmd

import (
	"fmt"

	"github.com/go-redis/redis"
)

const (
	commandHistoryKey = "ana#command#history"
	batch             = 1000
)

func ScanKeys(rds *redis.Client, pattern string, needSize *int) ([]string, error) {
	var keys, vals []string
	var cursor uint64
	var err error
	fmt.Printf("Scanning %s...\n", pattern)
	for {
		vals, cursor, err = rds.Scan(cursor, pattern, batch).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, vals...)
		if needSize != nil && len(keys) >= *needSize {
			break
		}
		if cursor == 0 {
			break
		}
	}

	return keys, nil
}
