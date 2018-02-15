package main

import (
	"os"
	"database/sql"
  "fmt"
  "syscall"
  "github.com/prometheus/log"
  "github.com/vrischmann/envconfig"
  _ "github.com/mattn/go-sqlite3"
)

type config struct {
  // Filename of SQLite3 database
  Filename string `envconfig:"default=database.db"`

  // LogLevel is a minimal log severity required for the message to be logged.
  // Valid levels: [debug, info, warn, error, fatal, panic].
  LogLevel string `envconfig:"default=info"`
}

func main() {
  // -> config from env
  cfg := &config{}
  if err := envconfig.InitWithPrefix(&cfg, "APP"); err != nil {
    log.Fatalf("init config: err=%s\n", err)
    syscall.Exit(1)
  }

  // get os args
  fmt.Println(len(os.Args))
  if len(os.Args) != 2 {
    fmt.Println("Please input program argument (select/insert/delete/update).")
    fmt.Println("No argument found. exiting program.")
    log.Fatalf("Please input program argument (select/insert/delete/update).")
    log.Fatalf("No argument found. exiting program.")
    syscall.Exit(1)
  }
  command := os.Args[1]
  fmt.Println(command)
  db_filename := cfg.Filename

  // to access filename from config : cfg.Filename
  // os.Remove("./" + db_filename)
  // executeSql(db_filename)
  db := connect_db(db_filename)
  prep_sql(db)
  switch command {
    case "select":
      rows := select_sql(db)
      fmt.Println(rows)
    case "insert":
      insert_sql(db)
    case "delete":
      delete_sql(db)
    case "update":
      update_sql(db)
  }
  close_db(db)
}

func connect_db(db_filename string) *sql.DB {
  db, err := sql.Open("sqlite3", "./" + db_filename)
	if err != nil {
    log.Fatal(err)
  }
  return db
}

func close_db(db *sql.DB) {
  db.Close()
  fmt.Println("db is now closed")
}

func prep_sql(db *sql.DB) {
  rows, err := db.Query("select id, name from products")
  if err != nil {
    db.Exec("create table products (id int primary key, product_name varchar(20))")
  }
  fmt.Println(rows)
}

func select_sql(db *sql.DB) *sql.Rows {
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
  return rows
}

func insert_sql(db *sql.DB) {
  tx, err := db.Begin()
  if err != nil {
    log.Fatal(err)
  }
  stmt, err := tx.Prepare("insert into products(id, product_name) values(?, ?)")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  for i := 0; i < 100; i++ {
    _, err := stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
    if err != nil {
      log.Fatal(err)
    }
  }
  tx.Commit()
}

func delete_sql(db *sql.DB) {
	_, err := db.Exec("delete from products")
	if err != nil {
		log.Fatal(err)
	}
}

func update_sql(db *sql.DB) {
	_, err := db.Exec("update products set product_name = 'a'")
	if err != nil {
		log.Fatal(err)
	}
}
