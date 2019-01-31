package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "os"
// import "fmt"

var db *sql.DB

func main() {
  mysql_server_host := os.Args[1]
  mysql_server_port := os.Args[2]
  mysql_server_user := ""
  if len(os.Args) > 3 {
    mysql_server_user = os.Args[3]
  }
  mysql_server_pass := ""
  if len(os.Args) > 4 {
    mysql_server_pass = os.Args[4]
  }

  var err error

  mysql_url := ""
  if mysql_server_user != "" {
    mysql_url += mysql_server_user
  }
  if mysql_server_pass != "" {
    mysql_url += ":" + mysql_server_pass
  }
  if mysql_server_user != "" {
    mysql_url += "@"
  }
  mysql_url += "tcp(" + mysql_server_host + ":" + mysql_server_port + ")/"

  db, err = sql.Open("mysql", mysql_url)
  if err != nil {
    panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
  }

  err = db.Ping()
  if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
  }

  var wsrep_local_state_label string
  var wsrep_local_state int64
  row := db.QueryRow("show status like 'wsrep_local_state'")
  err = row.Scan(&wsrep_local_state_label, &wsrep_local_state)

  if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
  }

  if wsrep_local_state == 4 {
    os.Exit(mainReturnWithCode(0))
  }

  os.Exit(mainReturnWithCode(1))
}

func mainReturnWithCode(exitcode int) int {
  // do stuff, defer functions, etc.
  defer db.Close()

  return exitcode // a suitable exit code
}
