package dbutil

import (
  "database/sql"
  "fmt"
  "strconv"
  "github.com/prometheus/log"
  _ "github.com/mattn/go-sqlite3"
)

type Dbutil struct {
  Filename string
}

func New(filename string) *Dbutil {
  return &Dbutil {
    Filename: filename,
  }
}

func (d *Dbutil) Prepare() {
  db, err := sql.Open("sqlite3", "./" + d.Filename)
  rows, err := db.Query("select id, name from products")
  if err != nil {
    db.Exec("create table products (id int primary key, product_name varchar(20))")
  }
  fmt.Println(rows)
}

func (d *Dbutil) Select() *sql.Rows {
  db, err := sql.Open("sqlite3", "./" + d.Filename)
  rows, err := db.Query("select id, product_name from products")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()
  for rows.Next() {
    var id int
    var name string
    err = rows.Scan(&id, &name)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(id, name)
  }
  err = rows.Err()
  if err != nil {
    log.Fatal(err)
  }
  db.Close();
  return rows
}

func (d *Dbutil) SelectCount() int {
  db, err := sql.Open("sqlite3", "./" + d.Filename)
  if err != nil {
    log.Fatal(err)
  }
  var result string
  var intResult int
  err = db.QueryRow("select count(1) from products").Scan(&result)
  if err != nil {
    log.Fatal(err)
  }
  intResult, err = strconv.Atoi(result)
  db.Close();
  return intResult
}

func (d *Dbutil) Insert(id string, product_name string) {
  db, errOpen := sql.Open("sqlite3", "./" + d.Filename)
  if errOpen != nil {
    log.Fatal(errOpen)
  }
  query_string := "insert into products (id, product_name) values ("+ id + ", '" + product_name + "')"
  _, err := db.Exec(query_string)
  if err != nil {
    log.Fatal(err)
  }
  db.Close()
}

func (d *Dbutil) Update(id string, product_name string) {
  db, errOpen := sql.Open("sqlite3", "./" + d.Filename)
  if errOpen != nil {
    log.Fatal(errOpen)
  }
  query_string := "update products set product_name = '" + product_name + "' where id = " + id
  _, err := db.Exec(query_string)
  if err != nil {
    log.Fatal(err)
  }
  db.Close()
}

func (d *Dbutil) SelectOne(id string) string {
  db, err := sql.Open("sqlite3", "./" + d.Filename)
  var result string
  err = db.QueryRow("select product_name from products where id = " + id).Scan(&result)
  if err != nil {
    log.Fatal(err)
  }
  db.Close();
  return result
}

func (d *Dbutil) Delete(id string) {
  db, errOpen := sql.Open("sqlite3", "./" + d.Filename)
  if errOpen != nil {
    log.Fatal(errOpen)
  }
  query_string := "delete from products where id = " + id 
  _, err := db.Exec(query_string)
  if err != nil {
    log.Fatal(err)
  }
  db.Close()
}

func (d *Dbutil) DeleteAll() {
  db, errOpen := sql.Open("sqlite3", "./" + d.Filename)
  if errOpen != nil {
    log.Fatal(errOpen)
  }
  query_string := "delete from products"
  _, err := db.Exec(query_string)
  if err != nil {
    log.Fatal(err)
  }
  db.Close()
}
