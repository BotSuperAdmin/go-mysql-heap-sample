package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"
)

// ...

type s struct {
	Passwd string
}

func main() {

	s1 := s{
		Passwd: "password_string_in_stack",
	}
	log.Println(s1.Passwd)

	db, err := sql.Open("mysql", "root:password_string_in_heap@tcp(127.0.0.1:3306)/information_schema")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// go func() {
	// 	time.Sleep(time.Minute * 3)
	// 	f, err := os.Create("heapdump")
	// 	if err != nil {
	// 		log.Fatalf("Could not open file for writing: %v\n", err)
	// 	}

	// 	debug.WriteHeapDump(f.Fd())
	// 	f.Close()
	// }()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		err := db.Ping()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong"))
		}
	})

	http.ListenAndServe(":8080", nil)
}
