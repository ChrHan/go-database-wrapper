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

  // to access filename from config : cfg.Filename
  os.Remove("./" + cfg.Filename)

  db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  sqlStmt := `
  create table foo (id integer not null primary key, name text);
  delete from foo;
  `
  _, err = db.Exec(sqlStmt)
  if err != nil {
    log.Printf("%q: %s\n", err, sqlStmt)
    return
  }

  tx, err := db.Begin()
  if err != nil {
    log.Fatal(err)
  }
  stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  for i := 0; i < 100; i++ {
    _, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
    if err != nil {
      log.Fatal(err)
    }
  }
  tx.Commit()

  rows, err := db.Query("select id, name from foo")
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

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}
	rows, err = db.Query("select id, name from foo")
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
}
