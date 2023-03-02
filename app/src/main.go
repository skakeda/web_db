package main

import (
  "net/http"
  "html/template"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  _ "fmt"
//   "log"
)


func main() {
    http.HandleFunc("/skakeda", skakeda)
    http.HandleFunc("/myamagata", myamagata)
    http.HandleFunc("/syatani", syatani)
    http.HandleFunc("/tozono", tozono)
    http.HandleFunc("/yichikawa", yichikawa)
    http.ListenAndServe(":8000", nil)
}

func skakeda(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:8jjStzw7@(mysql-container)/test")
                if err != nil {
                        panic(err.Error())
                }
    defer db.Close()

     rows, err := db.Query("SELECT txt FROM data WHERE id=1")
                if err != nil {
                        panic(err.Error())
                }

    defer rows.Close()

    var body string
    rows.Next()
    rows.Scan(&body)
// log.Print(body)
    t , err := template.ParseFiles("main.tpl")
                if err != nil {
                        panic(err.Error())
                }
    t.Execute(w, body)
}

func syatani(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:8jjStzw7@(mysql-container)/syatani")

                if err != nil {
                        panic(err.Error())
                }
    defer db.Close()

     rows, err := db.Query("SELECT txt FROM data WHERE id=1")
                if err != nil {
                        panic(err.Error())
                }

    defer rows.Close()

    var body string
    rows.Next()
    rows.Scan(&body)
// log.Print(body)
    t , err := template.ParseFiles("main.tpl")
                if err != nil {
                        panic(err.Error())
                }
    t.Execute(w, body)
}

func myamagata(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:8jjStzw7@(mysql-container)/myamagata_db")
    
                if err != nil {
                        panic(err.Error())
                }
    defer db.Close()

     rows, err := db.Query("SELECT txt FROM data WHERE id=1")
                if err != nil {
                        panic(err.Error())
                }

    defer rows.Close()

    var body string
    rows.Next()
    rows.Scan(&body)
// log.Print(body)
    t , err := template.ParseFiles("main.tpl")
                if err != nil {
                        panic(err.Error())
                }
    t.Execute(w, body)
}

func tozono(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:8jjStzw7@(mysql-container)/tozono_db")

                if err != nil {
                        panic(err.Error())
                }
    defer db.Close()

     rows, err := db.Query("SELECT txt FROM data WHERE id=1")
                if err != nil {
                        panic(err.Error())
                }

    defer rows.Close()

    var body string
    rows.Next()
    rows.Scan(&body)
// log.Print(body)
    t , err := template.ParseFiles("main.tpl")
                if err != nil {
                        panic(err.Error())
                }
    t.Execute(w, body)
}

func yichikawa(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:8jjStzw7@(mysql-container)/yichikawa")

                if err != nil {
                        panic(err.Error())
                }
    defer db.Close()

     rows, err := db.Query("SELECT txt FROM data WHERE id=1")
                if err != nil {
                        panic(err.Error())
                }

    defer rows.Close()

    var body string
    rows.Next()
    rows.Scan(&body)
// log.Print(body)
    t , err := template.ParseFiles("main.tpl")
                if err != nil {
                        panic(err.Error())
                }
    t.Execute(w, body)
}