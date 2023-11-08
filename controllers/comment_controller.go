package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/urakawa-jinsei/myapi/controllers/services"
	"github.com/urakawa-jinsei/myapi/models"
)

// 1. コントローラ構造体を定義
type CommentController struct {
	// 2. フィールドに CommentService 構造体を含める
	service services.CommentServicer
}

// コンストラクタの定義
func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment のハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
