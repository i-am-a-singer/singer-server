package service

import (
	"net/http"

	"github.com/singerapi/singer-server/render"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Root is a Server Root,
// APIPrefix is api resource url prefix
var (
	APIPrefix = "/singerapi/api"
)

// NewServer return Negroni instance
func NewServer() *negroni.Negroni {
	// Classic() 返回一个 Negroni 实例，会添加异常回复，日志和静态文件服务器三个中间件
	server := negroni.Classic()

	// 创建一个路由器
	router := mux.NewRouter()

	// 初始化路由器
	router.HandleFunc("/singerapi/api/", handleAPI).Methods("GET")

	router.HandleFunc("/singerapi/api/seasons/{id}/", handleSeasonsID).Methods("GET")
	router.HandleFunc("/singerapi/api/seasons/", handleSeasonsAll).Methods("GET")
	router.HandleFunc("/singerapi/api/seasons/{id}/songs/", handleSeasonsSongs).Methods("GET")
	router.HandleFunc("/singerapi/api/seasons/{id}/singers/", handleSeasonsSingers).Methods("GET")
	router.HandleFunc("/singerapi/api/seasons/{id}/albums/", handleSeasonsAlbums).Methods("GET")

	router.HandleFunc("/singerapi/api/songs/{id}/", handleSongsID).Methods("GET")
	router.HandleFunc("/singerapi/api/songs/", handleSongsAll).Methods("GET")
	router.HandleFunc("/singerapi/api/songs/{id}/seasons/", handleSongsSeasons).Methods("GET")
	router.HandleFunc("/singerapi/api/songs/{id}/singers/", handleSongsSingers).Methods("GET")
	router.HandleFunc("/singerapi/api/songs/{id}/albums/", handleSongsAlbums).Methods("GET")

	router.HandleFunc("/singerapi/api/singers/{id}/", handleSingersID).Methods("GET")
	router.HandleFunc("/singerapi/api/singers/", handleSingersAll).Methods("GET")
	router.HandleFunc("/singerapi/api/singers/{id}/songs/", handleSingersSongs).Methods("GET")
	router.HandleFunc("/singerapi/api/singers/{id}/seasons/", handleSingersSeasons).Methods("GET")
	router.HandleFunc("/singerapi/api/singers/{id}/albums/", handleSingersAlbums).Methods("GET")

	router.HandleFunc("/singerapi/api/albums/{id}/", handleAlbumsID).Methods("GET")
	router.HandleFunc("/singerapi/api/albums/", handleAlbumsAll).Methods("GET")
	router.HandleFunc("/singerapi/api/albums/{id}/songs/", handleAlbumsSongs).Methods("GET")
	router.HandleFunc("/singerapi/api/albums/{id}/singers/", handleAlbumsSingers).Methods("GET")
	router.HandleFunc("/singerapi/api/albums/{id}/albums/", handleAlbumsSeasons).Methods("GET")

	// 添加静态文件中间件, Prefix必须是"/xxx"，而不是"/xxx/"
	static := negroni.NewStatic(http.Dir("./dist"))
	static.Prefix = "/singerapi"
	server.Use(static)

	// 添加路由
	server.UseHandler(router)
	return server
}

// 处理 /singerapi/api/ 请求
func handleAPI(w http.ResponseWriter, r *http.Request) {
	prefix := "http://" + r.Host + APIPrefix

	var respBody = map[string]string{
		"api":     prefix + "/",
		"seasons": prefix + "/seasons/",
		"songs":   prefix + "/songs/",
		"singers": prefix + "/singers/",
		"albums":  prefix + "/albums/",
	}
	if err := render.MapJSONRender(w, http.StatusOK, respBody); err != nil {
		internalErrorHandler(w, r)
	}
}
