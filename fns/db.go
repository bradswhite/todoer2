package fns

import (
  //"database/sql"
  "log"
  "os"
  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
)

func GetAllTodos(db *sqlx.DB) []Todo {
  todos := []Todo {}

  err := db.Select(&todos, "SELECT * FROM TODOS;")
  if err != nil {
    log.Fatal(err)
  }

  return todos
}

func Conn() *sqlx.DB {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading .env file")
  }

  /*var (
    host = os.Getenv("HOST")
    port = os.Getenv("PORT")
    user = os.Getenv("USER")
    password = os.Getenv("PASSWORD")
    dbname = os.Getenv("DBNAME")
  )
  var conninfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)*/
  var conninfo string = os.Getenv("CONNINFO")

  db, err := sqlx.Connect("postgres", conninfo)
  if err != nil {
    panic(err)
  }

  return db
}
