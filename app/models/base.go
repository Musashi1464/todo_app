package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo_app_heroku/config"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// テーブルを作成するコードを記述

var Db *sql.DB

var err error

/*
// テーブルの名前
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)
*/

func init() {
	url := os.Getenv("DATABASE_URL") // herokuの環境変数の値を取り出せる. herokuのpostgresのurlが入る
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}
	/*
		Db, err = sql.Open(config.Config.SQLDriver, config.Config.Dbname)
		if err != nil {
			log.Fatalln(err)
		}

		cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			name STRING,
			email STRING,
			password STRING,
			created_at DATETIME)`, tableNameUser)

		Db.Exec(cmdU)

		cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT,
			user_id INTEGER,
			created_at DATETIME)`, tableNameTodo)

		Db.Exec(cmdT)

		cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			email STRING,
			user_id INTEGER,
			created_at DATETIME)`, tableNameSession)

		Db.Exec(cmdS)
	*/
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
