package main

import (
    "encoding/xml"
)

type MovieXml struct {
    XMLName xml.Name `xml:"movie"`
    Title           string `xml:"title"`
    Originaltitle   string `xml:"originaltitle"`
    Sorttitle       string `xml:"sorttitle"`
    Set             string `xml:"set"`
    Rating          string `xml:"rating"`
    Year            string `xml:"year"`
    Top250          string `xml:"top250"`
    Votes           string `xml:"votes"`
    Outline         string `xml:"outline"`
    Plot            string `xml:"plot"`
    Tagline         string `xml:"tagline"`
    Runtime         string `xml:"runtime"`
    Thumb           string `xml:"thumb"`
    Mpaa            string `xml:"mpaa"`
    Playcount       int `xml:"playcount"`
    ID              string `xml:"id"`
    Filenameandpath string `xml:"filenameandpath"`
    Trailer         string `xml:"trailer"`
    Genres           []string    `xml:"genre"`
    Credits         string `xml:"credits"`
    Fileinfo        struct {
        Streamdetails struct {
        } `xml:"streamdetails"`
    } `xml:"fileinfo"`
    Studio   string `xml:"studio"`
    Directors []string `xml:"director"`
    Actors    []struct {
        Name string `xml:"name"`
        Role string `xml:"role"`
    } `xml:"actor"`
}

type Actor   struct {
			Name string `xml:"name"`
			Role string `xml:"role"`
		} 

type MovieInfo struct {
	Rating struct {
		Count     int     `json:"count"`
		Max       int     `json:"max"`
		StarCount int     `json:"star_count"`
		Value     float64 `json:"value"`
	} `json:"rating"`
	VendorCount int      `json:"vendor_count"`
	Pubdate     []string `json:"pubdate"`
	Cover       struct {
		Liked       bool   `json:"liked"`
		Description string `json:"description"`
		Author      struct {
			Loc struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				UID  string `json:"uid"`
			} `json:"loc"`
			Kind   string `json:"kind"`
			Name   string `json:"name"`
			URL    string `json:"url"`
			URI    string `json:"uri"`
			Avatar string `json:"avatar"`
			Type   string `json:"type"`
			ID     string `json:"id"`
			UID    string `json:"uid"`
		} `json:"author"`
		LikersCount int `json:"likers_count"`
		Image       struct {
			Large struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"large"`
			Small struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"small"`
			Normal struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"normal"`
		} `json:"image"`
		URI           string `json:"uri"`
		URL           string `json:"url"`
		CreateTime    string `json:"create_time"`
		CommentsCount int    `json:"comments_count"`
		AllowComment  bool   `json:"allow_comment"`
		Position      int    `json:"position"`
		OwnerURI      string `json:"owner_uri"`
		Type          string `json:"type"`
		ID            string `json:"id"`
		SharingURL    string `json:"sharing_url"`
	} `json:"cover"`
	Pic struct {
		Large  string `json:"large"`
		Normal string `json:"normal"`
	} `json:"pic"`
	LineticketURL    string      `json:"lineticket_url"`
	BodyBgColor      string      `json:"body_bg_color"`
	IsTv             bool        `json:"is_tv"`
	Intro            string      `json:"intro"`
	TicketPriceInfo  string      `json:"ticket_price_info"`
	NullRatingReason string      `json:"null_rating_reason"`
	Year             string      `json:"year"`
	Webisode         interface{} `json:"webisode"`
	ID               string      `json:"id"`
	Genres           []string    `json:"genres"`
	CanInteract      bool        `json:"can_interact"`
	ReviewCount      int         `json:"review_count"`
	Title            string      `json:"title"`
	Languages        []string    `json:"languages"`
	IsReleased       bool        `json:"is_released"`
	CommentCount     int         `json:"comment_count"`
	Actors           []struct {
		Name     string      `json:"name"`
		Roles    []string    `json:"roles"`
		Title    string      `json:"title"`
		URL      string      `json:"url"`
		Abstract string      `json:"abstract"`
		Author   interface{} `json:"author"`
		URI      string      `json:"uri"`
		CoverURL string      `json:"cover_url"`
		Avatar   struct {
			Large  string `json:"large"`
			Normal string `json:"normal"`
		} `json:"avatar"`
		Type       string `json:"type"`
		ID         string `json:"id"`
		SharingURL string `json:"sharing_url"`
	} `json:"actors"`
	Interest      interface{} `json:"interest"`
	HeaderBgColor string      `json:"header_bg_color"`
	Type          string      `json:"type"`
	Linewatches   []struct {
		URL    string `json:"url"`
		Source struct {
			Literal string `json:"literal"`
			Pic     string `json:"pic"`
			Name    string `json:"name"`
		} `json:"source"`
		SourceURI string `json:"source_uri"`
		Free      bool   `json:"free"`
	} `json:"linewatches"`
	InfoURL       string   `json:"info_url"`
	HasLinewatch  bool     `json:"has_linewatch"`
	Durations     []string `json:"durations"`
	EpisodesCount int      `json:"episodes_count"`
	IsDoubanIntro bool     `json:"is_douban_intro"`
	SharingURL    string   `json:"sharing_url"`
	Countries     []string `json:"countries"`
	URL           string   `json:"url"`
	ReleaseDate   string   `json:"release_date"`
	OriginalTitle string   `json:"original_title"`
	URI           string   `json:"uri"`
	WebisodeCount int      `json:"webisode_count"`
	Directors     []struct {
		Name     string      `json:"name"`
		Roles    []string    `json:"roles"`
		Title    string      `json:"title"`
		URL      string      `json:"url"`
		Abstract string      `json:"abstract"`
		Author   interface{} `json:"author"`
		URI      string      `json:"uri"`
		CoverURL string      `json:"cover_url"`
		Avatar   struct {
			Large  string `json:"large"`
			Normal string `json:"normal"`
		} `json:"avatar"`
		Type       string `json:"type"`
		ID         string `json:"id"`
		SharingURL string `json:"sharing_url"`
	} `json:"directors"`
	InBlacklist bool     `json:"in_blacklist"`
	Aka         []string `json:"aka"`
	Trailer     struct {
		SharingURL string `json:"sharing_url"`
		VideoURL   string `json:"video_url"`
		Title      string `json:"title"`
		URI        string `json:"uri"`
		CoverURL   string `json:"cover_url"`
		TermNum    int    `json:"term_num"`
		NComments  int    `json:"n_comments"`
		CreateTime string `json:"create_time"`
		Runtime    string `json:"runtime"`
		Type       string `json:"type"`
		ID         string `json:"id"`
		Desc       string `json:"desc"`
	} `json:"trailer"`
}


type TvInfo struct {
	Rating struct {
		Count     int     `json:"count"`
		Max       int     `json:"max"`
		StarCount float64 `json:"star_count"`
		Value     float64 `json:"value"`
	} `json:"rating"`
	VendorCount int      `json:"vendor_count"`
	Pubdate     []string `json:"pubdate"`
	Cover       struct {
		Liked       bool   `json:"liked"`
		Description string `json:"description"`
		Author      struct {
			Loc struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				UID  string `json:"uid"`
			} `json:"loc"`
			Kind   string `json:"kind"`
			Name   string `json:"name"`
			URL    string `json:"url"`
			URI    string `json:"uri"`
			Avatar string `json:"avatar"`
			Type   string `json:"type"`
			ID     string `json:"id"`
			UID    string `json:"uid"`
		} `json:"author"`
		LikersCount int `json:"likers_count"`
		Image       struct {
			Large struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"large"`
			Small struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"small"`
			Normal struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"normal"`
		} `json:"image"`
		URI           string `json:"uri"`
		URL           string `json:"url"`
		CreateTime    string `json:"create_time"`
		CommentsCount int    `json:"comments_count"`
		AllowComment  bool   `json:"allow_comment"`
		Position      int    `json:"position"`
		OwnerURI      string `json:"owner_uri"`
		Type          string `json:"type"`
		ID            string `json:"id"`
		SharingURL    string `json:"sharing_url"`
	} `json:"cover"`
	Pic struct {
		Large  string `json:"large"`
		Normal string `json:"normal"`
	} `json:"pic"`
	LineticketURL    string      `json:"lineticket_url"`
	BodyBgColor      string      `json:"body_bg_color"`
	IsTv             bool        `json:"is_tv"`
	Intro            string      `json:"intro"`
	TicketPriceInfo  string      `json:"ticket_price_info"`
	NullRatingReason string      `json:"null_rating_reason"`
	Year             string      `json:"year"`
	Webisode         interface{} `json:"webisode"`
	ID               string      `json:"id"`
	Genres           []string    `json:"genres"`
	CanInteract      bool        `json:"can_interact"`
	ReviewCount      int         `json:"review_count"`
	Title            string      `json:"title"`
	Languages        []string    `json:"languages"`
	IsReleased       bool        `json:"is_released"`
	CommentCount     int         `json:"comment_count"`
	Actors           []struct {
		Name     string      `json:"name"`
		Roles    []string    `json:"roles"`
		Title    string      `json:"title"`
		URL      string      `json:"url"`
		Abstract string      `json:"abstract"`
		Author   interface{} `json:"author"`
		URI      string      `json:"uri"`
		CoverURL string      `json:"cover_url"`
		Avatar   struct {
			Large  string `json:"large"`
			Normal string `json:"normal"`
		} `json:"avatar"`
		Type       string `json:"type"`
		ID         string `json:"id"`
		SharingURL string `json:"sharing_url"`
	} `json:"actors"`
	Interest      interface{} `json:"interest"`
	HeaderBgColor string      `json:"header_bg_color"`
	Type          string      `json:"type"`
	Linewatches   []struct {
		URL    string `json:"url"`
		Source struct {
			Literal string `json:"literal"`
			Pic     string `json:"pic"`
			Name    string `json:"name"`
		} `json:"source"`
		SourceURI string `json:"source_uri"`
		Free      bool   `json:"free"`
	} `json:"linewatches"`
	InfoURL       string   `json:"info_url"`
	HasLinewatch  bool     `json:"has_linewatch"`
	Durations     []string `json:"durations"`
	EpisodesCount int      `json:"episodes_count"`
	IsDoubanIntro bool     `json:"is_douban_intro"`
	SharingURL    string   `json:"sharing_url"`
	Countries     []string `json:"countries"`
	URL           string   `json:"url"`
	ReleaseDate   string   `json:"release_date"`
	OriginalTitle string   `json:"original_title"`
	URI           string   `json:"uri"`
	WebisodeCount int      `json:"webisode_count"`
	Directors     []struct {
		Name     string      `json:"name"`
		Roles    []string    `json:"roles"`
		Title    string      `json:"title"`
		URL      string      `json:"url"`
		Abstract string      `json:"abstract"`
		Author   interface{} `json:"author"`
		URI      string      `json:"uri"`
		CoverURL string      `json:"cover_url"`
		Avatar   struct {
			Large  string `json:"large"`
			Normal string `json:"normal"`
		} `json:"avatar"`
		Type       string `json:"type"`
		ID         string `json:"id"`
		SharingURL string `json:"sharing_url"`
	} `json:"directors"`
	InBlacklist bool        `json:"in_blacklist"`
	Aka         []string    `json:"aka"`
	Trailer     interface{} `json:"trailer"`
}
