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


func querymovie_by_name(name string) []MovieXml {

    var infos []MovieXml
    var rawinfo []byte

    name,_ = url.QueryUnescape(name)
    name = strings.Replace(name,".","%",-1)
    name = strings.Replace(name," ","%",-1)
    name = strings.Replace(name,"+","%",-1)


    db, err := sql.Open("postgres", "user=fire3 password=theone dbname=douban sslmode=disable")
    checkErr(err)
    rows, err := db.Query("SELECT info FROM movie WHERE (title LIKE $1) OR (info->>'aka' LIKE $1) OR (UPPER(info->>'original_title') LIKE UPPER($1));" ,"%"+name+"%")
    checkErr(err)
    defer rows.Close()

    for rows.Next() {
        rows.Scan(&rawinfo)
        var m MovieInfo
        json.Unmarshal(rawinfo, &m)
        var x MovieXml
        x = getmoviexml(m)
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

func getmoviexml(m MovieInfo) (MovieXml) {

    var x MovieXml

    x.Title = m.Title
    x.Originaltitle = m.OriginalTitle
    x.Rating = fmt.Sprintf("%.1f", m.Rating.Value)
    x.Year = m.Year
    x.Votes = fmt.Sprintf("%d",m.Rating.Count)
    x.Plot = m.Intro
    x.ID = m.ID
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

    return x

    //output,e := xml.MarshalIndent(x,"  ", "    ")
    //return output, e
}


func querymovie_by_id(id string) MovieXml {
    var rawinfo []byte
    db, err := sql.Open("postgres", "user=fire3 password=theone dbname=douban sslmode=disable")
    checkErr(err)
    err = db.QueryRow("SELECT info FROM movie WHERE id = $1;" ,id).Scan(&rawinfo)
    checkErr(err)

    var m MovieInfo
    json.Unmarshal(rawinfo, &m)
    var x MovieXml
    x = getmoviexml(m)
    db.Close()

    return x
}
