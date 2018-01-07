package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "encoding/xml"
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

func do_tv_nfo (c *gin.Context) {
    id := c.Query("id") // shortcut for c.Request.URL.Query().Get("lastname")
    x := querytv_by_id(id)

    output,_ := xml.MarshalIndent(x,"  ", "    ")
    //c.Data(http.StatusOK,"application/xml; charset=utf-8",output)
    c.Header("Content-Disposition","attachment;filename=tvshow.nfo")
    c.Data(http.StatusOK,"application/octet-stream; charset=utf-8",output)
}


func do_movie_nfo (c *gin.Context) {
    id := c.Query("id") // shortcut for c.Request.URL.Query().Get("lastname")
    x := querymovie_by_id(id)

    output,_ := xml.MarshalIndent(x,"  ", "    ")
    //c.Data(http.StatusOK,"application/xml; charset=utf-8",output)
    c.Header("Content-Disposition","attachment;filename=movie.nfo")
    c.Data(http.StatusOK,"application/octet-stream; charset=utf-8",output)
}

func do_tv_search (c *gin.Context) {
    var titles string
    search := c.Query("q") // shortcut for c.Request.URL.Query().Get("lastname")
    infos := querytv_by_name(search)

    titles = `
    <style>
    .tablemovie table {
        width:100%;
        margin:15px 0;
        border:0;
    }
    .tablemovie th {
        background-color:#1E90FF;
        color:#FFFFFF
    }
    .tablemovie,.tablemovie th,.tablemovie td {
        font-size:0.95em;
        text-align:center;
        padding:4px;
        border-collapse:collapse;
    }
    .tablemovie th,.tablemovie td {
        border: 1px solid #206bfe;
        border-width:1px 0 1px 0;
        border:2px outset #ffffff;
    }
    .tablemovie tr {
        border: 1px solid #ffffff;
    }
    .tablemovie tr:nth-child(odd){
        background-color:#b4dafe;
    }
    .tablemovie tr:nth-child(even){
        background-color:#ffffff;
    }
    </style>
    `

    titles += "<table class=tablemovie>"
    titles += "<tr><th>名称</th><th>NFO下载</th></tr>"
    for _,v := range infos {
        titles += "<tr>"
        titles += fmt.Sprintf("<td>%s</td>",v.Title)
        titles += fmt.Sprintf("<td><a href=\"/tvnfo?id=%s\">NFO下载</a></td>",v.ID)
        titles += "</tr>"
    }
    titles += "</table>"

    c.Data(http.StatusOK, "text/html", []byte(titles))

}

func do_movie_search (c *gin.Context) {
    //firstname := c.DefaultQuery("firstname", "Guest")
    var titles string
    search := c.Query("q") // shortcut for c.Request.URL.Query().Get("lastname")
    infos := querymovie_by_name(search)

    titles = `
    <style>
    .tablemovie table {
        width:100%;
        margin:15px 0;
        border:0;
    }
    .tablemovie th {
        background-color:#1E90FF;
        color:#FFFFFF
    }
    .tablemovie,.tablemovie th,.tablemovie td {
        font-size:0.95em;
        text-align:center;
        padding:4px;
        border-collapse:collapse;
    }
    .tablemovie th,.tablemovie td {
        border: 1px solid #206bfe;
        border-width:1px 0 1px 0;
        border:2px outset #ffffff;
    }
    .tablemovie tr {
        border: 1px solid #ffffff;
    }
    .tablemovie tr:nth-child(odd){
        background-color:#b4dafe;
    }
    .tablemovie tr:nth-child(even){
        background-color:#ffffff;
    }
    </style>
    `

    titles += "<table class=tablemovie>"
    titles += "<tr><th>名称</th><th>原名</th><th>NFO下载</th></tr>"
    for _,v := range infos {
        titles += "<tr>"
        titles += fmt.Sprintf("<td>%s</td>",v.Title)
        titles += fmt.Sprintf("<td>%s</td>",v.Originaltitle)
        titles += fmt.Sprintf("<td><a href=\"/movienfo?id=%s\">NFO下载</a></td>",v.ID)
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
    router.GET("/movie",do_movie_search)
    router.GET("/tv",do_tv_search)
    router.GET("/movienfo",do_movie_nfo)
    router.GET("/tvnfo",do_tv_nfo)
    //    c.Data(http.StatusOK,"application/xml; charset=utf-8",xml)
    router.Run(":8081")
}
