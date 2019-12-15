package util

import (
	"net/http"
	"log"
	"fmt"
	"database/sql"
)

type DBUtil struct {
	Tx *sql.Tx
}

func (util * DBUtil) AddRequest(req http.Request) {
	log.Println(fmt.Sprintf("%v", req))
}

func (util * DBUtil) AddResponse(res http.Response) {
	log.Println(fmt.Sprintf("%v", res))
}

func (util * DBUtil) InsertToFoo(id int, text string) {
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