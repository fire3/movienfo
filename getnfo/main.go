package main

import (
    "fmt"
    "encoding/json"
    "net/url"
    "strings"
    "log"
    //"strconv"
    "database/sql"
    _ "github.com/lib/pq"
    "encoding/xml"
    //"time"
    "os"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {

    if len(os.Args) != 2 {
        fmt.Printf("Usage: %s Name\n",os.Args[0])
        os.Exit(0)
    }
    db, err := sql.Open("postgres", "user=fire3 password=theone dbname=douban sslmode=disable")
    checkErr(err)

    var id int
    var title string
    var rawinfo []byte
    err = db.QueryRow("SELECT * FROM movie WHERE title=$1",os.Args[1]).Scan(&id,&title,&rawinfo)

    switch {
    case err == sql.ErrNoRows:
        log.Fatal("No movie with title :" + os.Args[1])
    case err != nil:
        log.Fatal(err)
    default:
        var m MovieInfo
        json.Unmarshal(rawinfo, &m)

        var x MovieXml

        x.Title = title
        x.Originaltitle = m.OriginalTitle
        x.Rating = fmt.Sprintf("%.1f", m.Rating.Value)
        x.Year = m.Year
        x.Votes = fmt.Sprintf("%d",m.Rating.Count)
        x.Outline = m.Cover.Description
        x.Plot = m.Intro
        x.Genres = m.Genres
        x.Playcount = 0
        u, _ := url.Parse(m.Cover.Image.Large.URL)
        x.Thumb = u.Scheme + "://" + u.Host + u.Path

        for _, d := range m.Directors {
            x.Directors =  append(x.Directors, d.Name)
        }

        for _, a := range m.Actors {
            roles := strings.Join(a.Roles,", ")
            var r Actor 
            r.Name = a.Name
            r.Role = roles
            x.Actors = append(x.Actors, r)
        }

        output,_ := xml.MarshalIndent(x,"  ", "    ")
        os.Stdout.Write(output)
        fmt.Println("")


    }

    db.Close()
}
