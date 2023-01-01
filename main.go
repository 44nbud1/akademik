package main

import (
	"fmt"
	"github.com/44nbud1/akademik/akademik"
	config "github.com/44nbud1/akademik/pkg/pkgsql"
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	router := mux.NewRouter()

	dbCon := config.NewSqlDB(config.DatabaseModel{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	})

	db, err := dbCon.GetDatabasePostgres()
	if err != nil {
		log.Fatalln("Failed to connect to db")
	}

	flake, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatalln("Failed to connect snowflake")
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln("Failed to init zap log")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})

	if !strings.Contains(client.Ping().String(), "PONG") {
		log.Fatalln("Failed to start Redis")
	}

	akademik.NewAcademicModule(db, flake, router, logger, client)
	serveHTTP(fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")), router)
	fmt.Println("Service Stop !!!")
}

func serveHTTP(addrHTTP string, router *mux.Router) {

	log.Println("http server started. Listening on port: ", addrHTTP)
	server := &http.Server{
		Handler:           router,
		Addr:              addrHTTP,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
		return
	}

}
