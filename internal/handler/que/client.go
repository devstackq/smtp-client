package que

import (
	"fmt"
	"time"

	"github.com/devstackq/smtp-mailer/internal/models"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

func cleint() {

	// create redis connection pool
	redisPool := &redis.Pool{
		MaxIdle:     3,                 // maximum number of idle connections in the pool
		MaxActive:   0,                 // maximum number of connections allocated by the pool at a given time
		IdleTimeout: 240 * time.Second, // close connections after remaining idle for this duration
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://")
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	// initialize celery client
	cli, _ := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		1,
	)

	// prepare arguments
	taskName := "worker.Send"

	mail := models.Mail{
		From:       "8akebaev@gmail.com",
		Subject:    "Thread x",
		Message:    "Hola amiGO!",
		TmplTypeID: "63616c596c7b4bc739130641",
	}

	// run task
	asyncResult, err := cli.Delay(taskName, mail)
	if err != nil {
		panic(err)
	}

	// get results from backend with timeout
	res, err := asyncResult.Get(10 * time.Second)
	if err != nil {
		panic(err)
	}

	fmt.Println(res, "result")
}
