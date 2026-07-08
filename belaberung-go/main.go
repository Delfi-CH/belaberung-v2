package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"delfi.dev/belaberung-v2/db"
	"delfi.dev/belaberung-v2/model"
	"delfi.dev/belaberung-v2/route"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("No .env found...")
	}

	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := gin.Default()

	secret, exists := os.LookupEnv("BELABERUNG_SESSION_SECRET")

	if !exists {
		fmt.Println("No session secret provided!")
		os.Exit(1)
	}

	redisHost, exists := os.LookupEnv("BELABERUNG_REDIS_HOST")

	if !exists {
		fmt.Println("No session secret provided!")
		os.Exit(1)
	}

	redisUser, exists := os.LookupEnv("BELABERUNG_REDIS_USERNAME")

	if !exists {
		fmt.Println("No session secret provided!")
		os.Exit(1)
	}

	redisPass, exists := os.LookupEnv("BELABERUNG_REDIS_PASSWORD")

	if !exists {
		fmt.Println("No session secret provided!")
		os.Exit(1)
	}

	store, err := redis.NewStore(10, "tcp", redisHost, redisUser, redisPass, []byte(secret))

	if err != nil {
		panic(err)
	}

	r.Use(sessions.Sessions("redis", store))

	userRouter := r.Group("/users")
	route.InitUserRouter(userRouter, db)

	authRouter := r.Group("/auth")
	route.InitAuthRouter(authRouter, db)

	r.GET("/", func(c *gin.Context) {
		model.CreateUser(context.Background(), db, "demo", "example.com", "1234")
		c.String(http.StatusOK, "hello, world")
	})

	r.Run()
}
