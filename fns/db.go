package fns

import (
  //"database/sql"
  "log"
  "os"
  "fmt"
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

func GetTodosForUser(db *sqlx.DB, username string) []Todo {
  todos := []Todo {}

  err := db.Select(&todos, "SELECT * FROM TODOS WHERE USERNAME = $1;", username)
  if err != nil {
    log.Fatal(err)
  }

  return todos
}

func AddTodo(db *sqlx.DB, username, text string) string/*sql.Result*/ {
  res, err := db.NamedExec(`INSERT INTO TODOS (USERNAME, TEXT) VALUES (:username,:text)`, 
    map[string]interface{} {
      "username": username,
      "text": text,
    })
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(res)

  return "Successfully added todo."
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
