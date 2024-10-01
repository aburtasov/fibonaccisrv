package storage

import (
	"context"
	"math/big"
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
	Get(x, y int) ([]string, error)
}

func NewRedisStorage(address string) *RedisStorage {
	return &RedisStorage{
		rdb: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func (r *RedisStorage) Insert(length int) error {

	var ctx = context.Background()

	fib := make([]*big.Int, length+1)
	fib[1] = big.NewInt(0)
	fib[2] = big.NewInt(1)

	for i := 3; i <= length; i++ {
		fib[i] = new(big.Int).Add(fib[i-1], fib[i-2])
	}

	for i := 1; i <= length; i++ {
		r.mutex.Lock()
		err := r.rdb.Set(ctx, strconv.Itoa(i), fib[i].String(), 0).Err()
		r.mutex.Unlock()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RedisStorage) Get(x, y int) ([]string, error) {

	var ctx = context.Background()
	var fibSlice []string

	for i := x; i <= y; i++ {

		key := strconv.Itoa(i)
		val, err := r.rdb.Get(ctx, key).Result()

		if err != nil {
			return nil, err
		}

		fibSlice = append(fibSlice, val)
	}

	return fibSlice, nil
}

// Close закрывает соединение с Redis
func (r *RedisStorage) Close() error {
	return r.rdb.Close()
}
