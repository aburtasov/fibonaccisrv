package storage

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"sync"

	redis "github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	rdb   *redis.Client
	mutex sync.Mutex
}

type Storage interface {
	Insert(length int) error
	Get(x, y int) []int
}

func NewRedisStorage(address string) *RedisStorage {
	return &RedisStorage{
		rdb: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "", // no password set
			DB:       0,  // use default DV
		}),
	}
}

func (r *RedisStorage) Insert(length int) error {

	var ctx = context.Background()

	fib := make(map[int]int)
	fib[1] = 0
	fib[2] = 1

	for i := 3; i <= length; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	for i := range fib {
		r.mutex.Lock()
		err := r.rdb.Set(ctx, strconv.Itoa(i), fib[i], 0).Err()
		if err != nil {
			return err
		}
		r.mutex.Unlock()
	}
	return nil
}

func (r *RedisStorage) Get(x, y int) []int {

	var ctx = context.Background()
	var fibSlice []int

	for i := 1; i <= y; i++ {

		if i >= x && i <= y {
			key := strconv.Itoa(i)
			val, err := r.rdb.Get(ctx, key).Result()

			if err != nil {
				fmt.Printf("failed to get value from Redis:%s\n", err.Error())
			}

			num, _ := strconv.Atoi(val)

			fibSlice = append(fibSlice, num)
		} else {
			continue
		}
	}

	sort.Ints(fibSlice)

	return fibSlice

}
