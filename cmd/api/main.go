package main

import (
	"context"
	"fmt"
	"os"
	"github.com/kbg05/pixstash/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

func main(){
	ctx:=context.Background()
	godotenv.Load("../../.env")
	databaseUri:=os.Getenv("DATABASE_URI")
	fmt.Println(databaseUri)
	conn, err:=pgx.Connect(ctx, databaseUri)
	if err!=nil{
		fmt.Print(err)
	}
	defer conn.Close(ctx)
	queries:=database.New(conn)
	var pid pgtype.UUID
	pid.Scan("285a8149-56fa-441f-9e95-fae970558a1e")

	data:=database.CreateAssetParams{
		UserID: pid,
		OriginalName: "kbg",
		FileHash: "aflka",
		FileSize: 1,
		StoragePath: "kbg/kbgg",
		Status: "pending",
	}
	err=queries.CreateAsset(ctx,data)
	if err!=nil{
		fmt.Println(err)
	}


}

