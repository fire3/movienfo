package main

import (
    "fmt"
    "encoding/json"
    "net/url"
    "strings"
    "database/sql"
    _ "github.com/lib/pq"
    //"log"
    //"strconv"
    //"time"
)


func querytv_by_name(name string) []TvXml {

    var infos []TvXml
    var rawinfo []byte

    name,_ = url.QueryUnescape(name)
    name = strings.Replace(name,".","%",-1)
    name = strings.Replace(name," ","%",-1)
    name = strings.Replace(name,"+","%",-1)


    db, err := sql.Open("postgres", "user=fire3 password=theone dbname=douban sslmode=disable")
    checkErr(err)
    rows, err := db.Query("SELECT info FROM tv WHERE (title LIKE $1) OR (info->>'aka' LIKE $1) OR (UPPER(info->>'original_title') LIKE UPPER($1));" ,"%"+name+"%")
    checkErr(err)
    defer rows.Close()

    for rows.Next() {
        rows.Scan(&rawinfo)
        var m TvInfo
        json.Unmarshal(rawinfo, &m)
        var x TvXml
        x = gettvxml(m)
        if len(m.ID) > 0 {
            infos = append(infos,x)
            //u, _ := url.Parse(m.Cover.Image.Large.URL)
            //url := u.Scheme + "://" + u.Host + u.Path
            //downloadFile(m.ID + ".jpg",url)
        }
    }

    db.Close()

    return infos
}

func gettvxml(m TvInfo) (TvXml) {

    var x TvXml

    x.Title = m.Title
    x.Rating = fmt.Sprintf("%.1f", m.Rating.Value)
    x.Year = m.Year
    x.Votes = fmt.Sprintf("%d",m.Rating.Count)
    x.Plot = m.Intro
    x.ID = m.ID

    for _, a := range m.Actors {
        roles := strings.Join(a.Roles,", ")
        var r TvActor
        r.Name = a.Name
        r.Role = roles
        x.Actors = append(x.Actors, r)
    }

    return x

    //output,e := xml.MarshalIndent(x,"  ", "    ")
    //return output, e
}


func querytv_by_id(id string) TvXml {
    var rawinfo []byte
    db, err := sql.Open("postgres", "user=fire3 password=theone dbname=douban sslmode=disable")
    checkErr(err)
    err = db.QueryRow("SELECT info FROM tv WHERE id = $1;" ,id).Scan(&rawinfo)
    checkErr(err)

    var m TvInfo
    json.Unmarshal(rawinfo, &m)
    var x TvXml
    x = gettvxml(m)
    db.Close()

    return x
}
