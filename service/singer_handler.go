package service

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/singerapi/singer-server/database"
	"github.com/singerapi/singer-server/entity"
	"github.com/singerapi/singer-server/render"
)

// handleSingersID 处理 /singerapi/api/singers/{id} 请求
func handleSingersID(w http.ResponseWriter, r *http.Request) {
	prefix := r.Proto + "://" + r.Host + APIPrefix
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	sea := entity.Singer{}
	if err := database.Query(id, &sea); err != nil {
		notFindErrorHandler(w, r)
		return
	}
	render.ResourceItemJSONRender(w, http.StatusOK, &sea, prefix)
}

// 处理 /singerapi/api/singers/ 请求
func handleSingersAll(w http.ResponseWriter, r *http.Request) {
	_, pageOpt := parseRequest(r)
	handleSingers(w, r, func(s *entity.Singer) bool {
		return true
	}, &pageOpt)
}

// 处理 /singerapi/api/singers/{id}/seasons/ 请求
func handleSingersSeasons(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleSeasons(w, r, func(s *entity.Season) bool {
		return s.HasSinger(id)
	}, &pageOpt)
}

// 处理 /singerapi/api/singers/{id}/songs/ 请求
func handleSingersSongs(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleSongs(w, r, func(s *entity.Song) bool {
		return s.HasSinger(id)
	}, &pageOpt)
}

// 处理 /singerapi/api/singers/{id}/albums/ 请求
func handleSingersAlbums(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleAlbums(w, r, func(s *entity.Album) bool {
		return s.HasSinger(id)
	}, &pageOpt)
}
