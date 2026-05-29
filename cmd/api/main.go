package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/kbg05/pixstash/internal/api/middleware"
	"github.com/kbg05/pixstash/internal/database"
)


func main(){
	ctx:=context.Background()

	godotenv.Load("../../.env") //dev
	databaseURI:=os.Getenv("DATABASE_URI")

	pool, err:=database.NewPool(ctx, databaseURI)
	if err!=nil{
		log.Printf("failed to create database pool: %v", err)
	}
	defer pool.Close()

	queries:=database.New(pool)


	mux:=http.NewServeMux()

	log.Println("listening on :8000...")

	err=http.ListenAndServe(":8000", middleware.MiddlewareChain(mux, middleware.LoggerMiddleware))
	if err!=nil{
		log.Fatal(err)
	}
	
}

