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
	Status      uint
}

func CreateNewRequestTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS REQUEST(
		id number,
		method text,
		url text,
		contenttype text,
		body text,
		reqtime text
	); DELETE FROM REQUEST;`)
	if err != nil {
		return err
	}
	return nil
}

func (r *Request) Insert(tx *sql.Tx) error {
	stmt, err := tx.Prepare("INSERT INTO REQUEST (method,body,contenttype,url) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(r.Method, r.Body, r.ContentType, r.URL.EscapedPath())
	if err2 != nil {
		return err2
	}
	return nil
}
