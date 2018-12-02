package service

import(
	"net/http"
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
	router.HandleFunc("/singerapi/api/season/{id}", handleSeasonId(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/season/{id}/singers", handleSingers(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/season/{id}/songs", handleSongs(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/singers/?name={singer}", handleSingerName(format)).Methods("GET")
	router.HandleFunc("/singerapi/api/songs/?name={song}", handleSongName(format)).Methods("GET")

	// 添加中间件
	server.UseHandler(router)

	return server
}

// 请求处理函数，闭包函数
func handler(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 获得请求
		params := mux.Vars(request)
		// 解析请求内容
		name := params["name"]
		// 渲染模板
		format.JSON(writer, http.StatusOK, map[string]string{"name" : name})
	}
}
