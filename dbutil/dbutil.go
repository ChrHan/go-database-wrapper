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
  return intResult
}
/*func main() {
  // -> config from env
  cfg := &config{}
  if err := envconfig.InitWithPrefix(&cfg, "APP"); err != nil {
    log.Fatalf("init config: err=%s\n", err)
    syscall.Exit(1)
  }

  // get os args
  fmt.Println(len(os.Args))
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
  db := connect_db(db_filename)
  prep_sql(db)
  switch command {
    case "select":
      rows := select_sql(db)
      fmt.Println(rows)
    case "insert":
      insert_sql(db, id, product_name)
    case "delete_all":
      delete_all_sql(db)
    case "delete":
      delete_sql(db, id)
    case "update":
      update_sql(db, id, product_name)
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

func insert_sql(db *sql.DB, id string, product_name string) {
  query_string := "insert into products (id, product_name) values ("+ id + ", '" + product_name + "')"
  fmt.Println(query_string)
	_, err := db.Exec(query_string)
	if err != nil {
		log.Fatal(err)
	}
}

func delete_all_sql(db *sql.DB) {
	_, err := db.Exec("delete from products")
	if err != nil {
		log.Fatal(err)
	}
}

func delete_sql(db *sql.DB, id string) {
	_, err := db.Exec(fmt.Sprintf("delete from products where id = %s", id))
	if err != nil {
		log.Fatal(err)
	}
}

func update_sql(db *sql.DB, id string, product_name string) {
	_, err := db.Exec("update products set product_name = '" + product_name + "' where id = " + id )
	if err != nil {
		log.Fatal(err)
	}
} */
