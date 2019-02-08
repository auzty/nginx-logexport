// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "regexp"
//    "encoding/json"
)

type Logformat struct{
    Jam string
    Detil []DetailLog 
}

type DetailLog struct{
    DomaiName string
    TotalHits int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    resubtime := regexp.MustCompile("(^.*?) ([0-9].* )\\[(.*?)\\](.*$)")
    file, err := os.Open("samplelog.log")

    var tmpoutput []string

    if err != nil {
    fmt.Println(err)
    return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
//        fmt.Println(scanner.Text()) //print text
      newstring := resubtime.ReplaceAllString(scanner.Text(),"$1 $3")
    fmt.Println(newstring)
        tmpoutput= append(tmpoutput,newstring)

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("hasil akhir : ")
    fmt.Println(tmpoutput)
}

// pakai sqlite saja, nanti diquery -1hari lebih gampang (baca log dari hari sebelumnya, terserah di recreate atau enggak sqlite db nya)