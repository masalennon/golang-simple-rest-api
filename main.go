package main

import (
	"io/ioutil"
	"time"
	"encoding/json"
	"fmt"
	"net/http"
)

type article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allArticles []article

var articles = allArticles{
	{
		ID:          "12",
		Title:       "Introduction to golang",
		Description: "Go言語の基本",
	},
	{
		ID:          "13",
		Title:       "Introduction to Algorithm",
		Description: "アルゴリズムの基本",
	},	
	{
		ID:          "14",
		Title:       "Introduction to Programming",
		Description: "プログラミングの基本",
	},
}

func main() {
	http.HandleFunc("/articles", handleAllArticlesRequest)
	http.Handle("/articles/", http.StripPrefix("/articles/", http.HandlerFunc(handleSingleArticleRequest)))
	http.ListenAndServe(":8081", nil)
}

func handleAllArticlesRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getArticles(w, r)
	case http.MethodPost:
		addArticle(w, r)
	default:
		http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
		t := time.NewTicker(100)
		for range t.C {
			t.Stop()
		}
	}
}

func handleSingleArticleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getArticle(w, r)
	case "DELETE":
		deleteArticle(w, r)
	case "PATCH":
		updateArticle(w, r)
	default:
		http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
	}
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles) //JSONにエンコードしたtrustsをレスポンスボディ（w）に書き込む
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Path //URIを返す。articles/test/13 にアクセスしたら articles/test/13を返す。
	fmt.Printf("r : %+v\n", r)
	fmt.Println("eventID: ", eventID)
	for _, article := range articles {
		if article.ID == eventID {
			json.NewEncoder(w).Encode(article)
			break
		}
	}
}

func addArticle(w http.ResponseWriter, r *http.Request) {
	var article article
	json.NewDecoder(r.Body).Decode(&article)
	articles = append(articles, article)

	json.NewEncoder(w).Encode(articles)
}

func handlePost(w http.ResponseWriter, r *http.Request) {

}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Path
	var updatedArticle article

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedArticle)

	for i, singleArticle := range articles {
		if singleArticle.ID == eventID {
			singleArticle.Title = updatedArticle.Title
			singleArticle.Description = updatedArticle.Description
			tmpArticle := articles[i+1:]
			articles = append(articles[:i], singleArticle)
			articles = append(articles, tmpArticle...)

			json.NewEncoder(w).Encode(singleArticle)
		}
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Path
	for i, article := range articles {
		fmt.Println("articles: ", articles)
		if eventID == article.ID {
			articles = append(articles[:i], articles[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully\n", eventID)
			fmt.Println("articles after delete", articles)
		}
	}
}