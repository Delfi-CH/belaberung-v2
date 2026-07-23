package main

import (
	"fmt"
	"net/http"
	"os"

	"delfi.dev/belaberung-v2/db"
	"delfi.dev/belaberung-v2/route"
	"delfi.dev/belaberung-v2/ws"
	"github.com/gin-contrib/cors"
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

	hub := ws.NewHub()
	go hub.Run()

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

	store.Options(sessions.Options{
		HttpOnly: true,
		MaxAge:   86400,
		Path:     "/",
	})

	r.Use(sessions.Sessions("redis", store))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	userRouter := r.Group("/users")
	route.InitUserRouter(userRouter, db)

	roomRouter := r.Group("/rooms")
	route.InitRoomRouter(roomRouter, db)

	authRouter := r.Group("/auth")
	route.InitAuthRouter(authRouter, db)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, world")
	})

	r.GET("/ws", ws.Handler(hub, db))

	r.Run()
}

/* todo
func CreateMessage(
	hub *ws.Hub,
	message models.Message,
) error {


	err := db.NewInsert().
		Model(&message).
		Scan(context.Background())


	if err != nil {
		return err
	}


	hub.Broadcast(
		ws.Message{
			ID: message.ID,
			Content: message.Content,
			UserID: message.UserID,
			RoomID: message.RoomID,
			Timestamp: message.Timestamp,
		},
	)


	return nil
}


*/
