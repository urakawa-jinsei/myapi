package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urakawa-jinsei/myapi/controllers/services"
	"github.com/urakawa-jinsei/myapi/models"
)

// 1. コントローラ構造体を定義
type ArticleController struct {
	// 2. フィールドに MyAppService 構造体を含める
	service services.ArticleServicer
}

// コンストラクタの定義
func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

// GET /hello のハンドラ
func (c *ArticleController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// POST /article のハンドラ
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// GET /article/list のハンドラ
func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	// クエリパラメータpageを取得
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	// 暫定でこれを追加することで
	// 「変数pageが使われていない」というコンパイルエラーを回避
	log.Println(page)

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articleList)
}

// GET /article/{id} のハンドラ
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// POST /article/nice のハンドラ
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}