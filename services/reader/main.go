package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/rs/cors"
	"golang.org/x/net/context"  
	"log"
)

func main() {
	redis_host := "localhost" 
	redis_port := "6379"
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "OPTIONS" {
			writer.WriteHeader(http.StatusOK)
			return
		}
		fmt.Fprintf(writer, "up")
	})

	mux.HandleFunc("/data", func(writer http.ResponseWriter, request *http.Request) {
		rdb := redis.NewClient(&redis.Options{
			Addr: redis_host + ":" + redis_port,
			DB:   0,
		})
		ctx := context.Background()
		key := rdb.Get(ctx, "SHAREDKEY")
		if err := key.Err(); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(writer, key.Val())
	})

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	}).Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
