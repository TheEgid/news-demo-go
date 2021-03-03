package views

import (
	"encoding/json"
	"fmt"
	"github.com/TheEgid/news-demo-go/models"
	"github.com/TheEgid/news-demo-go/utils"
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

var templ = template.Must(template.ParseFiles("index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_ = templ.Execute(w, nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := utils.GoDotEnvVariable("APIKEY")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal server error"))
		return
	}

	params := u.Query()
	searchKey := params.Get("q")
	page := params.Get("page")
	if page == "" {
		page = "1"
	}

	search := &models.Search{}
	search.SearchKey = searchKey

	next, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Unexpected server error", http.StatusInternalServerError)
		return
	}

	search.NextPage = next
	pageSize := 20

	endpoint := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&pageSize=%d&page=%d&apiKey=%s&sortBy=publishedAt&language=ru", url.QueryEscape(search.SearchKey), pageSize, search.NextPage, apiKey)
	resp, err := http.Get(endpoint)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&search.Results)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	search.TotalPages = int(math.Ceil(float64(search.Results.TotalResults / pageSize)))
	if ok := !search.IsLastPage(); ok {
		search.NextPage++
	}

	err = templ.Execute(w, search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
