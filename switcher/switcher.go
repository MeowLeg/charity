package switcher

import (
	// "database/sql"
	// _ "github.com/mattn/go-sqlite3"
	// "log"
	"net/http"
	"strconv"
	"encoding/json"
	// "strings"
	"time"
)

type Xl map[string]func(*http.Request) (string, interface{})

var apiKey string = "sk_test_KyTS4OTur5mPO08e5SmDCqD0"

func Dispatch(db *sql.DB) Xl {
	return Xl{
		"auth": func(r *http.Request) (string, interface{}) {
			return "验证成功", nil
		},
	}
}

func GetParameter(r *http.Request, key string) string {
	s := r.URL.Query().Get(key)
	if s == "" {
		panic("没有参数" + key)
	}
	return s
}

func today() string {
	return time.Now().Format("2006-01-02")
}

func tommorrow() string {
	return time.Unix(time.Now().Unix()+86400, 0).Format("2006-01-02")
}
