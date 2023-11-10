package controllers_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/urakawa-jinsei/myapi/controllers"
	"github.com/urakawa-jinsei/myapi/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
