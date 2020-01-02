package util

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	Status      int
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

func CreateNewResponseTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS RESPONSE(
		reqid number,
		status number,
		contenttype text,
		body text,
		restime text
	); DELETE FROM RESPONSE;`)
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

func (r *Response) Insert(tx *sql.Tx) error {
	stmt, err := tx.Prepare("INSERT INTO RESPONSE (status,body,contenttype) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(r.Status, r.Body, r.ContentType)
	if err2 != nil {
		return err2
	}
	return nil
}

var EmptyResponse = Response{}

func (r *Request) DoRequest() (Response, error) {
	if r == nil {
		return EmptyResponse, fmt.Errorf("request is nil")
	}

	var resp *http.Response
	var err error

	switch r.Method {
	case "GET":
		resp, err = http.Get(r.URL.String())
	case "POST":
		reader := strings.NewReader(r.Body)
		resp, err = http.Post(r.URL.String(), r.ContentType, reader)
	default:
		err = fmt.Errorf("unexpected request method: %s", r.Method)
		return EmptyResponse, err
	}
	if err != nil {
		return EmptyResponse, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return EmptyResponse, err
	}

	return Response{
		Status:      resp.StatusCode,
		Body:        string(b),
		ContentType: resp.Header.Get("content-type"),
		Header:      resp.Header,
	}, nil
}
