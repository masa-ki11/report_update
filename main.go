package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	server := "" //サーバー
	port := 1433
	user := ""     //ユーザー名
	password := "" //パスワード
	database := "" //DB名

	// SQL Server接続文字列
	connString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s",
		server, port, user, password, database)

	// SQL Serverに接続
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 接続テスト
	err = db.PingContext(context.Background())
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}

	// SQLコマンド作成
	query := "UPDATE tbl_DairyReport SET Syunin = m.kyouryoku_id FROM tbl_DairyReport AS d JOIN master_accounts AS m ON d.Syunin = m.account_id WHERE LEFT(d.Syunin, 1) = 'K' "

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query: ", err.Error())
	}
	fmt.Println("Success!")
	defer rows.Close()

}
