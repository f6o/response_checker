package util

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type DBUtil struct {
	Tx *sql.Tx
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS FOO (id integer not null primary key, name text); DELETE FROM FOO;")
	if err != nil {
		return err
	}
	return nil
}

func (util *DBUtil) AddRequest(req http.Request) {
	log.Println(fmt.Sprintf("%v", req))
}

func (util *DBUtil) AddResponse(res http.Response) {
	log.Println(fmt.Sprintf("%v", res))
}

func (util *DBUtil) InsertToFoo(id int, text string) {
	stmt, err := util.Tx.Prepare("INSERT INTO FOO (ID, NAME) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(id, text)
	if err2 != nil {
		log.Fatal(id)
	}
}
