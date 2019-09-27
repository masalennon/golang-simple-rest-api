package main

import (
	"fmt"
	"net/http"
	"encoding/json"

)

func main() {
	// mux := http.NewServeMux() //デフォルトのマルチプレクサを生成するコード、リクエストをハンドラにリダイレクトする。
    // mux.HandleFunc("/", homeLink)
    // mux.HandleFunc("/events", getAllEvents)
    // mux.Handle("/events/", http.StripPrefix("/events/", http.HandlerFunc(getOneEvent)))

	http.HandleFunc("/articles", handleArticleRequest)
	http.Handle("/articles/", http.StripPrefix("/articles/", http.HandlerFunc(getArticle)))//アクセスされたURLからarticles部分を削除してハンドリングする。つまり、articles/13の場合、13とgetArticleメソッドを紐付ける
	//第一引数のURLへのリクエストが到着すると、第二引数に指定されているハンドラ関数にリダイレクトされる。全てのハンドラ関数が第一引数にResponseWriterをとり、第二引数にRequestへのポインタを取るのでハンドラ関数に改めて引数を渡す必要はない
	//ハンドラ関数とは第一引数にResponseWriterをとり、第二引数にRequestへのポインタを取るGoの関数に過ぎない
	//HTTPリクエスト発生時に呼び出される具体的な処理が書かれています。
    http.ListenAndServe(":8081", nil) //なぜnilでいい？
}

func handleArticleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleArticlesGet(w, r)
	default:
		http.Error(w, r.Method + " method not allowed", http.StatusMethodNotAllowed)
		// w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// w.Header().Set("X-Content-Type-Options", "nosniff")
		// w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

type article struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
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
}

func handleArticlesGet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)//JSONにエンコードしたtrustsをレスポンスボディ（w）に書き込む
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Path //URIの末尾を返す。articles/の後に文字列を追加して確認済み。その時も13だった。
	fmt.Printf("%+v\n", r)
	fmt.Println(eventID)
	for _, article := range articles {
		if article.ID == eventID {
			json.NewEncoder(w).Encode(article)
			break
		}
	}
	

}

func handlePost(w http.ResponseWriter, r *http.Request) {

}