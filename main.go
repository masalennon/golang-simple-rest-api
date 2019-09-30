package main

import (
	"io/ioutil"
	"time"
	"encoding/json"
	"fmt"
	"net/http"
)

type article struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
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
	// mux := http.NewServeMux() //デフォルトのマルチプレクサを生成するコード、リクエストをハンドラにリダイレクトする。
	// mux.HandleFunc("/", homeLink)
	// mux.HandleFunc("/events", getAllEvents)
	// mux.Handle("/events/", http.StripPrefix("/events/", http.HandlerFunc(getOneEvent)))

	http.HandleFunc("/articles", handleAllArticlesRequest)//HandldFuncもHandleも役割は同じ。両方URLにハンドラ関数を登録する。
	http.Handle("/articles/", http.StripPrefix("/articles/", http.HandlerFunc(handleSingleArticleRequest))) //アクセスされたURLからarticles部分を削除してハンドリングする。つまり、articles/13の場合、13とhandleSingleArticleRequestメソッドを紐付ける
	//第一引数のURLへのリクエストが到着すると、第二引数に指定されているハンドラ関数にリダイレクトされる。全てのハンドラ関数が第一引数にResponseWriterをとり、第二引数にRequestへのポインタを取るのでハンドラ関数に改めて引数を渡す必要はない
	//ハンドラ関数とは第一引数にResponseWriterをとり、第二引数にRequestへのポインタを取るGoの関数に過ぎない
	//今まで気づかなかったが、HandlerFuncはよく見たら型なので、↑でキャストしているのだな。
	//「func (w http.ResponseWriter, r *http.Request)」関数型はHandlerFunc型にキャストすればHandlerとコンパチになる。
	//http.Handle("/articles/test/", http.StripPrefix("/articles/", http.HandlerFunc(getArticle))) このように書くと、/test/xxx というURLとgetArticleが結び付けられ、getArticle内のr.URL.PathのURLはtest/xxxになる
	//HTTPリクエスト発生時に呼び出される具体的な処理が書かれています。
	http.ListenAndServe(":8081", nil) //なぜnilでいい？
}

func handleAllArticlesRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getArticles(w, r)
	case http.MethodPost:
		addArticle(w, r)
	default:
		http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
		// w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// w.Header().Set("X-Content-Type-Options", "nosniff")
		// w.WriteHeader(http.StatusMethodNotAllowed)
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
	id := r.URL.Query().Get("id")


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