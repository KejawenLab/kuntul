package adapters

import (
	"github.com/KejawenLab/kuntul"
	storage "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type redis struct {
	Pool  *redsync.Redsync
	Mutex *redsync.Mutex
}

func NewRedisAdapter(pool string) kuntul.Adapter {
	client := storage.NewClient(&storage.Options{
		Addr: pool,
	})

	return &redis{
		Pool: redsync.New(goredis.NewPool(client)),
	}
}

func (l *redis) Lock(task *kuntul.Task) error {
	l.Mutex = l.Pool.NewMutex(task.ID, redsync.WithExpiry(task.Estimation))
	return l.Mutex.Lock()
}

func (l *redis) Unlock() error {
	_, err := l.Mutex.Unlock()

	return err
}
