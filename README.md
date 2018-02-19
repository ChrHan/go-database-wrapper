# SQLite Go Wrapper

A simple `Go` program to execute `CRUD` to `sqlite3` database.

Currently programmed only for this table specification:

    CREATE TABLE PRODUCTS (
      id  INT PRIMARY KEY,
      product_name  VARCHAR(20)
    )

# Prerequisite

1. `Go` version 1.94 upwards
1. `dep` tool installed

# Usage

1. Clone this repo
1. Run `dep ensure`
1. cd to `app/sqlite_wrapper`
1. Run `go install`
1. Run program using `sqlite-wrapper {select|insert id product_name|update id product_name|delete id|delete_all}`

