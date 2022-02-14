package config

import (
	"database/sql"
	"strconv"

	_ "github.com/godror/godror"
	_ "github.com/mattn/go-sqlite3"
)

/*
	*sql.Db tipinde bir connection döndürür.
	Bağlantıtyı açar
	Oracle ve Sqlite3 için bağlantıların açıldığı metodları içerir
*/

func OracleConn(conn OracleConfig) (*sql.DB, error) {

	db, err := sql.Open("godror", conn.Username+"/"+conn.Password+conn.Host+":"+strconv.Itoa(conn.Port)+"/"+conn.ServiceName)

	db.Ping()

	if err != nil {
		panic(err)
	}
	return db, err
}

// func Sqlite3Conn(dbPath string) (*sql.DB, error) {
// 	db, err := sql.Open("sqlite3", dbPath)
// 	db.Ping()

// 	if err != nil {
// 		panic(err)
// 	}

// 	return db, err
// }
