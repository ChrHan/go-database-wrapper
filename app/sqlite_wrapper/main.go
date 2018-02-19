package main

import (
	"os"
  "fmt"
  "syscall"
  "github.com/prometheus/log"
  "github.com/vrischmann/envconfig"
  _ "github.com/mattn/go-sqlite3"
  dbutil "github.com/ChrHan/go-sqlite-utility/dbutil"
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
  if len(os.Args) < 2 {
    fmt.Println("Please input program argument (select/insert/delete/update).")
    fmt.Println("No argument found. exiting program.")
    log.Fatalf("Please input program argument (select/insert/delete/update).")
    log.Fatalf("No argument found. exiting program.")
    syscall.Exit(1)
  }
  command := os.Args[1]
  var id string
  var product_name string
  if len(os.Args) >= 3 {
    id = os.Args[2]
    fmt.Println("id = " + id)
  }
  if len(os.Args) == 4 {
    product_name = os.Args[3]
  }
  db_filename := cfg.Filename

  // to access filename from config : cfg.Filename
  db := dbutil.New(db_filename)
  db.Prepare()
  switch command {
    case "select":
      rows := db.Select()
      fmt.Println(rows)
    case "insert":
      db.Insert(id, product_name)
    case "delete_all":
      db.DeleteAll()
    case "delete":
      db.Delete(id)
    case "update":
      db.Update(id, product_name)
  }
}
