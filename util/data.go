package util

import (
	"database/sql"
	"net/http"
	"net/url"
)

type Request struct {
	Method      string
	Body        string
	ContentType string
	URL         url.URL
	Header      http.Header
}

type Response struct {
	Body        string
	ContentType string
	Header      http.Header
}

func CreateRequestTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS REQUEST(
		id number,
		method text,
		url text,
		body text,
		reqtime text
	); DELETE FROM REQUEST;`)
	if err != nil {
		return err
	}
	return nil
}

func (*Request) Insert(tx *sql.Tx) error {
	stmt, err := tx.Prepare("INSERT INTO REQUEST (method,url,body) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err2 := stmt.Exec()
	if err2 != nil {
		return err2
	}
	return nil
}
