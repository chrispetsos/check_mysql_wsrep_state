package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "os"

var db *sql.DB

func main() {
  mysql_server_host := os.Args[1]
  mysql_server_port := os.Args[2]

  var err error

  db, err = sql.Open("mysql", "haproxy_check@tcp(" + mysql_server_host + ":" + mysql_server_port + ")/")
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
