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
    //"os"
    //"time"
)


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func getmoviexml(name string) ([]byte, error) {

    db, err := sql.Open("postgres", "user=fire3 password=theone dbname=douban sslmode=disable")
    checkErr(err)

    var id int
    var title string
    var rawinfo []byte
    err = db.QueryRow("SELECT * FROM movie WHERE title=$1",name).Scan(&id,&title,&rawinfo)

    switch {
    case err == sql.ErrNoRows:
        db.Close()
        return []byte{},err
    case err != nil:
        db.Close()
        return []byte{},err
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

        output,e := xml.MarshalIndent(x,"  ", "    ")
        db.Close()
        return output, e
    }
}

func main() {

    router := gin.Default()
    router.StaticFile("/","./index")
    router.POST("/movie", func(c *gin.Context) {
        movie := c.PostForm("movie")
        xml,_ := getmoviexml(movie)
        c.Data(http.StatusOK,"application/xml; charset=utf-8",xml)
    })
    router.Run(":8081")
}
