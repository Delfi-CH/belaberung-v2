package main

import (
	"delfi.dev/belaberung-v2/db"
	"delfi.dev/belaberung-v2/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"github.com/gin-contrib/sessions"
  	"github.com/gin-contrib/sessions/cookie"
	"os"
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

	secret, exists := os.LookupEnv("BELABERUNG_COOKIE_SECRET")

	if !exists {
		fmt.Println("No cookie secret provided!")
		os.Exit(1)
	}

	store := cookie.NewStore([]byte(secret))

	r.Use(sessions.Sessions("cookie", store))

	userRouter := r.Group("/users")
	route.InitUserRouter(userRouter, db)

	authRouter := r.Group("/auth")
	route.InitAuthRouter(authRouter, db)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, world")
	})

	r.Run()
}
