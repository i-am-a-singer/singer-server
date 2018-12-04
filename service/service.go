package service

import(
	"net/http"
	"encoding/json"
	"github.com/urfave/negroni"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"

	"github.com/singerapi/singer-server/database"
)

func NewServer() *negroni.Negroni {
	// Classic() 返回一个 Negroni 实例，会添加异常回复，日志和静态文件服务器三个中间件
	server := negroni.Classic()

	// 创建一个路由器
	router := mux.NewRouter()
 
	format := render.New(render.Options {
		IndentJSON : true,
	})

	// 初始化路由器
	router.HandleFunc("/singerapi/api/", handleAPI(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/seasons/", handleSeasons(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/seasons/{id}", handleSeasonsId(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/seasons/{id}/singers/", handleSingers(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/seasons/{id}/songs/", handleSongs(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/singers/?singer={singer}", handleSingerName(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/songs/?song={song}", handleSongName(format)).Methods("GET")

	// 添加中间件
	server.UseHandler(router)

	return server
}


// 处理 /singerapi/api/ 请求
func handleAPI(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		respBody map[string][string] {
			"all seasons":"http://127.0.0.1/singerapi/api/seasons/",
    		"specific season":"http://127.0.0.1/singerapi/api/seasons/{id}/",
    		"singers in season":"http://127.0.0.1/singerapi/api/seasons/{id}/singers",
    		"songs in season":"http://127.0.0.1/singerapi/api/seasons/{id}/songs",
    		"specific singer":"http://127.0.0.1/singerapi/api/singers/?name={name}",
    		"specific song":"http://127.0.0.1/singerapi/api/songs/?name={name}"
		}
		format.JSON(writer, http.StatusOK, respBody)
	}
}

// 处理 /singerapi/api/seasons/ 请求
func handleSeasons(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		jsonStr :=  '{
			"songs":{
				"我是歌手第一季":"http://127.0.0.1:8080/singerapi/api/seasons/1/",
				"我是歌手第一季":"http://127.0.0.1:8080/singerapi/api/seasons/2/",
				"我是歌手第一季":"http://127.0.0.1:8080/singerapi/api/seasons/3/",
				"我是歌手第一季":"http://127.0.0.1:8080/singerapi/api/seasons/4/",
				"我是歌手第一季(歌手2017)":"http://127.0.0.1:8080/singerapi/api/seasons/5/",
				"我是歌手第一季(歌手2018)":"http://127.0.0.1:8080/singerapi/api/seasons/6/"},
			"url":"http://127.0.0.1/singerapi/api/seasons/"
		}'
		respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		}
	}
}

// 处理 /singerapi/api/seasons/{id} 请求
func handleSeasonsId(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		jsonStr := db.query(1, id)
		respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		}
	}
}

// 处理 /singerapi/api/seasons/{id}/singers/ 请求
func handleSingers(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		jsonStr := db.query(2, id)
		respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		}
	}
}

// 处理 /singerapi/api/seasons/{id}/songs/ 请求
func handleSongs(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		jsonStr := db.query(3, id)
		respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		}
	}
}

// 处理 /singerapi/api/singers/?singer={singer} 请求
func handleSingerName(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		singer := params["singer"]

		jsonStr := db.query(4, singer)
		respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		}
	}
}

// 处理 /singerapi/api/songs/?song={song} 请求
func handleSongName(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		song := params["song"]

		jsonStr := db.query(5, song)
		respBody map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &respBody); err == nil {
			format.JSON(writer, http.StatusOK, respBody)
		}
	}
}

