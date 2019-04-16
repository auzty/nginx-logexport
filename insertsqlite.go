package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "regexp"
    "database/sql"
    "strings"
    "time"
    _ "github.com/mattn/go-sqlite3"
//    "encoding/json"
)

func insertdata(logpath string,dbpath string) {

    //regex for split www.example.com 202.67.33.36 - - [03/Feb/2019:10:56:27 +0700] "GET /stylesheets/full-modal.css HTTP/1.1" 200 542 "http://www.example.com/" "Mozilla/5.0 (Linux; Android 5.1; A1601) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.99 Mobile Safari/537.36" "-"
    //to become (( www.example.com 2016-02-15 17:29:24 ))
    resubtime := regexp.MustCompile("(^.*?) ([0-9].* )\\[(.*?) +.*?\\](.*$)")

    //default logpath are /var/log/nginx/access.log.1 (yesterday full log)
    if(logpath == ""){
        logpath = "/var/log/nginx/access/log"
    }
    file, err := os.Open(logpath)

    db, err := sql.Open("sqlite3", dbpath)
    check(err)

    stmt,err := db.Prepare("CREATE TABLE IF NOT EXISTS  nginxlog (id INTEGER PRIMARY KEY, vhost TEXT, logdate TEXT)")
    check(err)

    stmt.Exec()
    // insert
    stmt, err = db.Prepare("INSERT INTO nginxlog(vhost, logdate) values(?,?)")
    check(err)

    datelayout := "02/Jan/2006:15:04:05" //dd/Mmm/yyyy:hh:mm:ss
    dateformat := "15"  // only show hours instead full datetime
    if err != nil {
    fmt.Println(err)
    return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    db.Exec("BEGIN EXCLUSIVE TRANSACTION;") //start transaction
    for scanner.Scan() {
      newstring := resubtime.ReplaceAllString(scanner.Text(),"$1 $3")
      logstring := strings.Split(newstring," ")
      vhost := logstring[0]
      datelog := logstring[1]
      
      tmpdatetime,_ := time.Parse(datelayout,datelog)
      tmpdatetimeformat := tmpdatetime.Format(dateformat)
        stmt.Exec(vhost, tmpdatetimeformat)

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }


    db.Exec("COMMIT;") // commit the transaction
        db.Close()

}

// pakai sqlite saja, nanti diquery -1hari lebih gampang (baca log dari hari sebelumnya, terserah di recreate atau enggak sqlite db nya)