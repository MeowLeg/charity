package main

import (
	sw "charity/switcher"
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type Ret struct {
	Success bool        `json:"success"`
	ErrMsg  string      `json:"errMsg"`
	Data    interface{} `json:"data"`
}

func main() {
	rt := httprouter.New()
	rt.GET("/ssp", DlmHandler)

	n := negroni.Classic()
	n.UseHandler(rt)
	n.Run(":7063")
}

func DlmHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer func() {
		err := recover()
		if err != nil {
			rw.Write(GenJsonpResult(r, &Ret{false, err.(string), nil}))
			log.Println(err)
		}
	}()

	db := ConnectDB("./middle.db")
	LogClient(r.RemoteAddr, db)

	switcher := sw.Dispatch(db)
	var ret []byte
	if Authorize(r) {
		msg, data := switcher[sw.GetParameter(r, "cmd")](r)
		ret = GenJsonpResult(r, &Ret{true, msg, data})
	} else {
		panic("Not authorized!")
	}
	rw.Write(ret)
}

func Authorize(r *http.Request) bool {
	token := sw.GetParameter(r, "token")
	// log.Println(token)
	return token == "Jh2044695"
}

func GenJsonpResult(r *http.Request, rt *Ret) []byte {
	bs, err := json.Marshal(rt)
	if err != nil {
		panic(err)
	}
	return []byte(sw.GetParameter(r, "callback") + "(" + string(bs) + ")")
}

func ConnectDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	return db
}

func LogClient(ip string, db *sql.DB) {
	// 如果没有click表，会出现pointer为nil的问题
	stmt, _ := db.Prepare("insert into clicks(ip) values(?)")
	stmt.Exec(ip)
}
