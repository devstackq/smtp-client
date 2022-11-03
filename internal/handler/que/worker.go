package que

import (
	"time"

	"github.com/devstackq/smtp-mailer/internal/service"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

func CelerySendMail(mailer service.Mailer) {
	// create redis connection pool
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://")
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	// initialize celery client
	cli, _ := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		5, // number of workers
	)

	// register task
	cli.Register("worker.Send", mailer.Send)

	// start workers (non-blocking call)
	cli.StartWorker()

	// wait for client request
	time.Sleep(10 * time.Second) //get from config

	// stop workers gracefully (blocking call)
	cli.StopWorker()
}
