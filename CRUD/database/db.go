// note: i pull a docker image of redis latest version to start my db

package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ExampleDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0, // Use default db
	})

	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ping)

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}


	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key ", val)

	val2, err := rdb.Get(ctx, "key").Result()
	if err == redis.Nil {
		fmt.Println("key not found")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2 ", val2)
	}
}