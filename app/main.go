package main

import (
	"os"
	"strconv"
	"github.com/redis/go-redis/v9"
    "github.com/gin-gonic/gin"
)

func main() {
	client := redis.NewClient(&redis.Options{
        Addr:	  os.Getenv("ADDR") ,
        Password: os.Getenv("PASSWORD"),
        DB:		  0,
        Protocol: 2,
    })

    router := gin.Default()
    router.GET("/", func(c *gin.Context){
		nb_requests := 0
		val, err := client.Get(c, "nb_requests").Result()
		if err != nil  {
			nb_requests = 1
		}else{
			nb_requests, err = strconv.Atoi(val)
			nb_requests += 1
		}
		_ = client.Set(c, "nb_requests", nb_requests, 0).Err()

		c.JSON(200, gin.H{
			"version": os.Getenv("VERSION"),
			"nb_requests": nb_requests,
		})
	})

    router.Run()
}