package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/kbg05/pixstash/internal/api/middleware"
	"github.com/kbg05/pixstash/internal/database"
)

type Application struct{
	q *database.Queries
	ctx context.Context
}

func(a *Application) getAssetHandler(w http.ResponseWriter, r *http.Request){
	id:=r.PathValue("id")
	pid, err:=uuid.Parse(id)
	if err!=nil{
		log.Printf("failed to parse uuid: %v", err)
	}
	
//	data:=database.CreateAssetParams{
//		UserID: pid,
//		OriginalName: "kbgii",
//		FileHash: "aflkakk",
//		FileSize: 1,
//		StoragePath: "kbg/kbgg",
//		Status: "pending",
//	}

	
	//err=queries.CreateAsset(ctx,data)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	value, err:=a.q.GetAsset(a.ctx, pid)
	if err!=nil{
		log.Printf("failed to fetch user: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Println(value)
	data, err:=json.Marshal(value)
	if err!=nil{
		log.Printf("failed to convert to json: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	
}

func main(){
	ctx:=context.Background()

	godotenv.Load("../../.env")
	databaseURI:=os.Getenv("DATABASE_URI")

	pool, err:=database.NewPool(ctx, databaseURI)
	if err!=nil{
		log.Printf("failed to create database pool: %v", err)
	}
	defer pool.Close()

	queries:=database.New(pool)
	app:=Application{
		q: queries,
		ctx: ctx ,
	}
	mux:=http.NewServeMux()
	mux.Handle("GET /user/{id}", (http.HandlerFunc(app.getAssetHandler)))

	log.Println("listening on :8000...")

	err=http.ListenAndServe(":8000", middleware.MiddlewareChain(mux, middleware.LoggerMiddleware))
	if err!=nil{
		log.Fatal(err)
	}
	
}

