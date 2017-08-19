package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "strconv"
    "database/sql"
    _ "github.com/lib/pq"
    "time"
    "os"
)

var apikey string = "0dad551ec0f84ed02907ff5c42e8ec70"

func makesearchurl(start int) string {
    path := "https://frodo.douban.com"
    path += "/api/v2/movie/tag?"
    path +="count=1&"
    path +="start="
    path += strconv.Itoa(start)
    path += "&sort=T&score_range=0,10&os_rom=android&apikey="
    path += apikey
    //path += "&q="
    //path += url.QueryEscape("电影")
    //fmt.Printf("Encoded URL is %q\n", path)
    return path
}

func makeinfourl(id string, typ string) string {
    Url, err := url.Parse("https://frodo.douban.com")
    if err != nil {
        panic("boom")
    }

    Url.Path += "/api/v2/"
    Url.Path += typ
    Url.Path += "/"
    Url.Path += id
    parameters := url.Values{}
    parameters.Add("os_rom", "android")
    parameters.Add("apikey", "0dad551ec0f84ed02907ff5c42e8ec70")
    Url.RawQuery = parameters.Encode()
    //fmt.Printf("Encoded URL is %q\n", Url.String())
    return Url.String()
}

func getidtype(num int) (string , string) {

    ua := "api-client/1 com.douban.frodo/5.2.0(103) Android/23 product/MT7-CL00 vendor/HUAWEI model/HUAWEI MT7-CL00 rom/android network/wifi"
    tr := &http.Transport{
        DisableKeepAlives: true,
    }
    client := &http.Client {
        Transport: tr,
        Timeout: 5*time.Second,
    }
    req, err := http.NewRequest("GET", makesearchurl(num), nil)
    if err != nil {
        return "",""
    }

    req.Header.Set("User-Agent",ua)

    resp, err := client.Do(req)
    if err != nil {
        return "",""
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "",""
    }
    var s SearchResult
    json.Unmarshal(body, &s)

    if len(s.Data) == 0 {
        return "",""
    }

    //fmt.Println(s.Data[0])
    return s.Data[0].ID, s.Data[0].Type

}

func gettvinfo(id string) (TvInfo, error) {

    if len(id) == 0 {
        return TvInfo{},fmt.Errorf("No valid Tv")
    }
    ua := "api-client/1 com.douban.frodo/5.2.0(103) Android/23 product/MT7-CL00 vendor/HUAWEI model/HUAWEI MT7-CL00 rom/android network/wifi"
    tr := &http.Transport{
        DisableKeepAlives: true,
    }
    client := &http.Client {
        Transport: tr,
        Timeout: 5*time.Second,
    }
    req, err := http.NewRequest("GET", makeinfourl(id,"tv"), nil)
    if err != nil {
        return TvInfo{}, err
    }

    req.Header.Set("User-Agent",ua)

    resp, err := client.Do(req)
    if err != nil {
        return TvInfo{}, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return TvInfo{}, err
    }
    var tv TvInfo
	json.Unmarshal(body, &tv)

    return tv,nil
}

func getmovieinfo(id string) (MovieInfo, error) {

    if len(id) == 0 {
        return MovieInfo{},fmt.Errorf("No valid Movie")
    }
    ua := "api-client/1 com.douban.frodo/5.2.0(103) Android/23 product/MT7-CL00 vendor/HUAWEI model/HUAWEI MT7-CL00 rom/android network/wifi"
    tr := &http.Transport{
        DisableKeepAlives: true,
    }
    client := &http.Client {
        Transport: tr,
        Timeout: 5*time.Second,
    }
    req, err := http.NewRequest("GET", makeinfourl(id,"movie"), nil)
    if err != nil {
        return MovieInfo{}, err
    }

    req.Header.Set("User-Agent",ua)

    resp, err := client.Do(req)
    if err != nil {
        return MovieInfo{}, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return MovieInfo{}, err
    }
    var movie MovieInfo
	json.Unmarshal(body, &movie)

    //fmt.Println(movie)

    return movie,nil

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    var movie MovieInfo
    var err error

    if len(os.Args) != 3 {
        fmt.Printf("Usage: %s movie|tv ID\n",os.Args[0])
        os.Exit(0)
    }
    db, err := sql.Open("postgres", "user=fire3 password=theone dbname=douban sslmode=disable")
    checkErr(err)

    //插入数据
    stmt_movie, err := db.Prepare("INSERT INTO movie(id,title,info) VALUES($1,$2,$3) ON CONFLICT (id) DO NOTHING")
    checkErr(err)

    stmt_tv, err := db.Prepare("INSERT INTO tv(id,title,info) VALUES($1,$2,$3) ON CONFLICT (id) DO NOTHING")
    checkErr(err)

    id := os.Args[2]
    typ := os.Args[1]

		if typ == "movie" {
			movie, err = getmovieinfo(id)
			if err == nil {
				if len(movie.ID) == 0 {
                    fmt.Println(id + " has no info.")
                }
				if len(movie.ID) > 0 {
					fmt.Println(movie.ID + " " + movie.Title)
					movieid ,_ := strconv.Atoi(movie.ID)
					moviejson, _ := json.Marshal(movie)
					_, err := stmt_movie.Exec(movieid,movie.Title , moviejson)
					checkErr(err)
				}
			} else {
				log.Println(err)
				time.Sleep(10*time.Second)
			}

		}

		if typ == "tv" {
            tv, err := gettvinfo(id)
			if err == nil {
				if len(tv.ID) == 0 {
                    fmt.Println(id + " has no info.")
                }
				if len(tv.ID) > 0 {
                    fmt.Println(tv.ID + " " + tv.Title)
                    tvid ,_ := strconv.Atoi(tv.ID)
                    tvjson, _ := json.Marshal(tv)
                    _, err := stmt_tv.Exec(tvid,tv.Title , tvjson)
                    checkErr(err)
                }
			} else {
				log.Println(err)
				time.Sleep(5*time.Second)
			}

		}
        time.Sleep(1*time.Second)
        if id == "" && typ == "" {
            fmt.Println("No id/type returned.")
        }

    db.Close()
}
