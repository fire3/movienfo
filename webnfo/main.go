package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "encoding/json"
    "net/url"
    "strings"
    "database/sql"
    _ "github.com/lib/pq"
    "encoding/xml"
    //"log"
    //"strconv"
    //"time"
    "os"
    "io"
    //"time"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func querydb_by_id(id string) MovieXml {
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

func querydb_by_name(name string) []MovieXml {

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

func downloadFile(filename string, url string) (err error) {

  // Create the file
  out, err := os.Create(cover_dir + filename)
  if err != nil  {
    return err
  }
  defer out.Close()

  // Get the data
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // Writer the body to file
  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    return err
  }

  return nil
}

func do_nfo (c *gin.Context) {
    id := c.Query("id") // shortcut for c.Request.URL.Query().Get("lastname")
    x := querydb_by_id(id)

    output,_ := xml.MarshalIndent(x,"  ", "    ")
    //c.Data(http.StatusOK,"application/xml; charset=utf-8",output)
    c.Header("Content-Disposition","attachment;filename=movie.nfo")
    c.Data(http.StatusOK,"application/octet-stream; charset=utf-8",output)
}

func do_search (c *gin.Context) {
    //firstname := c.DefaultQuery("firstname", "Guest")
    var titles string
    search := c.Query("q") // shortcut for c.Request.URL.Query().Get("lastname")
    infos := querydb_by_name(search)

    titles = `
    <style>
    .table12_11 table {
        width:100%;
        margin:15px 0;
        border:0;
    }
    .table12_11 th {
        background-color:#1E90FF;
        color:#FFFFFF
    }
    .table12_11,.table12_11 th,.table12_11 td {
        font-size:0.95em;
        text-align:center;
        padding:4px;
        border-collapse:collapse;
    }
    .table12_11 th,.table12_11 td {
        border: 1px solid #206bfe;
        border-width:1px 0 1px 0;
        border:2px outset #ffffff;
    }
    .table12_11 tr {
        border: 1px solid #ffffff;
    }
    .table12_11 tr:nth-child(odd){
        background-color:#b4dafe;
    }
    .table12_11 tr:nth-child(even){
        background-color:#ffffff;
    }
    </style>
    `

    titles += "<table class=table12_11>"
    titles += "<tr><th>名称</th><th>原名</th><th>NFO下载</th></tr>"
    for _,v := range infos {
        titles += "<tr>"
        titles += fmt.Sprintf("<td>%s</td>",v.Title)
        titles += fmt.Sprintf("<td>%s</td>",v.Originaltitle)
        titles += fmt.Sprintf("<td><a href=\"/nfo?id=%s\">NFO下载</a></td>",v.ID)
        //u, _ := url.Parse(v.Cover.Image.Normal.URL)
        //thumb  := u.Scheme + "://" + u.Host + u.Path
        //titles += fmt.Sprintf("<td> <img src=%s> </td>",thumb)
        //titles += fmt.Sprintf("<td> <img src=images/%s.jpg width=200px> </td>",v.ID)
        titles += "</tr>"
    }
    titles += "</table>"

    c.Data(http.StatusOK, "text/html", []byte(titles))
}

var cover_dir string = "/home/fire3/douban_cover/"

func main() {

    router := gin.Default()
    router.StaticFile("/","./index")
    router.Static("/images", cover_dir)
    router.GET("/search",do_search)
    router.GET("/nfo",do_nfo)
    //    c.Data(http.StatusOK,"application/xml; charset=utf-8",xml)
    router.Run(":8081")
}
