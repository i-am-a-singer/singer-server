package service

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"

	db "github.com/singerapi/singer-server/database"
)

// NewServer return Negroni instance
func NewServer() *negroni.Negroni {
	// Classic() 返回一个 Negroni 实例，会添加异常回复，日志和静态文件服务器三个中间件
	server := negroni.Classic()

	// 创建一个路由器
	router := mux.NewRouter()

	format := render.New(render.Options{
		IndentJSON: true,
	})

	// 初始化路由器
	router.HandleFunc("/singerapi/api/", handleAPI(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/seasons/", handleSeasons(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/season/{id}/", handleSeasonsID(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/season/{id}/singers/", handleSingers(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/season/{id}/songs/", handleSongs(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/singers/?singer={singer}", handleSingerName(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/songs/?song={song}", handleSongName(format)).Methods("GET")

	// 添加静态文件中间件
	static := negroni.NewStatic(http.Dir("./dist"))
	static.Prefix = "/singerapi"
	server.Use(static)

	// 添加路由
	server.UseHandler(router)

	return server
}

// 处理 /singerapi/api/ 请求
func handleAPI(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var respBody = map[string]string{
			"all seasons":       "http://127.0.0.1/singerapi/api/seasons/",
			"specific season":   "http://127.0.0.1/singerapi/api/season/{id}/",
			"singers in season": "http://127.0.0.1/singerapi/api/season/{id}/singers/",
			"songs in season":   "http://127.0.0.1/singerapi/api/season/{id}/songs/",
			"specific singer":   "http://127.0.0.1/singerapi/api/singers/?singer={singer}",
			"specific song":     "http://127.0.0.1/singerapi/api/songs/?song={song}",
		}
		format.JSON(writer, http.StatusOK, respBody)
	}
}

// 处理 /singerapi/api/seasons/ 请求
func handleSeasons(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		jsonStr := `{
			"songs":{
				"我是歌手第一季":"http://127.0.0.1:8080/singerapi/api/season/1/",
				"我是歌手第二季":"http://127.0.0.1:8080/singerapi/api/season/2/",
				"我是歌手第三季":"http://127.0.0.1:8080/singerapi/api/season/3/",
				"我是歌手第四季":"http://127.0.0.1:8080/singerapi/api/season/4/",
				"我是歌手第五季(歌手2017)":"http://127.0.0.1:8080/singerapi/api/season/5/",
				"我是歌手第六季(歌手2018)":"http://127.0.0.1:8080/singerapi/api/season/6/"},
			"url":"http://127.0.0.1/singerapi/api/seasons/"
		}`
		var respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		} else if err != nil {
			fmt.Printf("Unmarshl json fail with err : %s", err)
		}
	}
}

// handleSeasonsID 处理 /singerapi/api/seasons/{id}/ 请求
func handleSeasonsID(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		jsonStr := db.Query(1, id)
		var respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		} else if err != nil {
			fmt.Printf("Unmarshl json fail with err : %s", err)
		}
	}
}

// 处理 /singerapi/api/seasons/{id}/singers/ 请求
func handleSingers(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		jsonStr := db.Query(2, id)
		var respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		} else if err != nil {
			fmt.Printf("Unmarshl json fail with err : %s", err)
		}
	}
}

// 处理 /singerapi/api/seasons/{id}/songs/ 请求
func handleSongs(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		jsonStr := db.Query(3, id)
		var respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		} else if err != nil {
			fmt.Printf("Unmarshl json fail with err : %s", err)
		}
	}
}

// 处理 /singerapi/api/singers/?singer={singer} 请求
func handleSingerName(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		singer := params["singer"]

		jsonStr := db.Query(4, singer)
		var respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		} else if err != nil {
			fmt.Printf("Unmarshl json fail with err : %s", err)
		}
	}
}

// 处理 /singerapi/api/songs/?song={song} 请求
func handleSongName(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		song := params["song"]

		jsonStr := db.Query(5, song)
		var respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		} else if err != nil {
			fmt.Printf("Unmarshl json fail with err : %s", err)
		}
	}
}
