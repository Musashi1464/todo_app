package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// テーブルを作成するコードを記述

var Db *sql.DB

var err error

// テーブルの名前
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, "testuser:Musashi/0830634@/test_db?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INT PRIMARY KEY AUTO_INCREMENT,
			uuid TEXT NOT NULL UNIQUE,
			name VARCHAR(100),
			email VARCHAR(100),
			password VARCHAR(100),
			created_at DATETIME)`, tableNameUser)

	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INT PRIMARY KEY AUTO_INCREMENT,
			content TEXT,
			user_id INT,
			created_at DATETIME)`, tableNameTodo)

	_, err = Db.Exec(cmdT)
	if err != nil {
		log.Fatalln(err)
	}

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INT PRIMARY KEY AUTO_INCREMENT,
			uuid TEXT NOT NULL UNIQUE,
			email VARCHAR(100),
			user_id INT,
			created_at DATETIME)`, tableNameSession)

	_, err = Db.Exec(cmdS)
	if err != nil {
		log.Fatalln(err)
	}
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
