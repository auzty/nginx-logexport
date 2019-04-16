package main

import (
	"flag"
	"time"
	"os"
	"log"
	"fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
	var mode string // generated or getdata
	var hour string // check the hour count
	var vhost string // check specific virtual host
	var loglocation string // the file log location
	var dbfile string // sqlite location

	var totalcount int

	flag.StringVar(&mode,"mode","getdata","write : Will write the sqlite3 db\ngetdata : get the data from data")	
	flag.StringVar(&vhost,"vhost","www.example.com","Will find www.example.com from sqlite")	
	flag.StringVar(&dbfile,"db","/opt/icingavhost/data.db","sqlite database exact path")	
	flag.StringVar(&loglocation,"logfile","/var/log/nginx/access.log.1","Will read access.log.1 instead live streaming the log")	

	flag.Parse()

	timeformat := time.Now()
	hour = timeformat.Format("15") // only get the hours from time.Now()

	if(mode=="write"){
		 _, err := os.Stat(loglocation)
		if err != nil {
			if  os.IsNotExist(err){
				//file not fount
				log.Fatal(err)
			} else {
				//other error
				panic(err)
			}
		}else{
			insertdata(loglocation,dbfile)
		}

	}else {
		totalcount = readingdata(vhost,hour)

	// icinga2 plugins deps

	// #Icinga return codes
	// UNKNOWN = -1
	// OK = 0
	// WARNING = 1
	// CRITICAL = 2
	// example : os.Exit(2) -> critical

		fmt.Println("OK - ",vhost," get ",totalcount, " request(s) | value=",totalcount)
		os.Exit(0)
	}



}

