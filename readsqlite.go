package main

import (
//    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
//    "reflect"
//    "strconv"
    "strings"
)

func readingdata(vhost string,hour string) int {

//    rows, _ := db.Query("select * from nginxlog where vhost like ? and logdate = ?",vhost,hour)

    var totalwww,totalnonwww int
   
    if(strings.HasPrefix(vhost,"www")){
       totalwww = queryvhostcount(vhost,hour)
       totalnonwww = queryvhostcount(strings.TrimPrefix(vhost,"www."),hour)
    }else {
       totalnonwww = queryvhostcount(vhost,hour)
       totalwww = queryvhostcount("www."+vhost,hour)
    }

    totalcount := totalnonwww + totalwww

    return totalcount;
//    fmt.Println("Total Keseluruhan : ",totalcount)
}


func queryvhostcount(vhost string,hour string) int{
    var count int
    db, err := sql.Open("sqlite3", "./data.db")
    check(err)

    // insert
    rows, _ := db.Query("select count(*) from nginxlog where vhost like ? and logdate = ?",vhost,hour)

    for rows.Next() {
        rows.Scan(&count)
//        fmt.Println("Jumlah row : ",strconv.Itoa(count))
    }

    db.Close()
    
    return count;
}