package models

import (
	"database/sql"
	"fmt"
	"go_line_bot/entity"
	"os"
)

var (
	name     = os.Getenv("DATABASE_USERNAME")
	password = os.Getenv("DATABASE_PASSWORD")
	database = os.Getenv("DATABASE_NAME")
)

func Find(identifier int) (admin entity.Admin) {
	// ToDo: デザインパターン等でDB処理をディレクトリ分類させたい
	DbConnection, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", name, password, database))
	if err != nil {
		panic(err.Error())
	}
	defer DbConnection.Close()
	row, err := DbConnection.Query("SELECT * FROM admins WHERE id = ?", identifier)
	if err != nil {
		return
	}
	defer row.Close()
	var id int
	var line_messaging_id int
	var line_messaging_secret string
	var line_messaging_token string
	var line_login_id int
	var line_login_secret string
	row.Next() //行での繰り返し処理を実行
	// DBで取得した値をカラム毎に変数にセット
	if err = row.Scan(&id, &line_messaging_id, &line_messaging_secret, &line_messaging_token, &line_login_id, &line_login_secret); err != nil {
		return
	}
	// 上の変数を通してエンティティに値をセット
	admin.Id = id
	admin.LineMessagingId = line_messaging_id
	admin.LineMessagingSecret = line_messaging_secret
	admin.LineMessagingToken = line_messaging_token
	admin.LineLoginId = line_login_id
	admin.LineLoginSecret = line_login_secret
	return
}

// func (sqlHandler *SqlHandler) Find(id int) {

// }
