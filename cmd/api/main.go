package main

import (
	"context"
	"fmt"
	"os"
	"github.com/kbg05/pixstash/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

func main(){
	ctx:=context.Background()

	godotenv.Load("../../.env")
	databaseURI:=os.Getenv("DATABASE_URI")

	pool, err:=database.NewPool(ctx, databaseURI)
	if err!=nil{
		fmt.Printf("failed to create database pool: %v", err)
	}
	defer pool.Close()

	queries:=database.New(pool)


	
	var pid pgtype.UUID
	pid.Scan("285a8149-56fa-441f-9e95-fae970558a1e")

	data:=database.CreateAssetParams{
		UserID: pid,
		OriginalName: "kbgii",
		FileHash: "aflkakk",
		FileSize: 1,
		StoragePath: "kbg/kbgg",
		Status: "pending",
	}
	
	err=queries.CreateAsset(ctx,data)
	if err!=nil{
		fmt.Println(err)
	}
}

