package service

import (
	"net/http"

	"github.com/singerapi/singer-server/database"
	"github.com/singerapi/singer-server/entity"
	"github.com/singerapi/singer-server/render"
)

// handleSeasonsID 处理 /singerapi/api/seasons/{id} 请求
func handleSeasonsID(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)

	sea := entity.Season{}
	if err := database.Query(id, &sea); err != nil {
		notFindErrorHandler(w, r)
		return
	}
	render.ResourceItemJSONRender(w, http.StatusOK, &sea, pageOpt.Prefix)
}

// 处理 /singerapi/api/seasons/ 请求
func handleSeasonsAll(w http.ResponseWriter, r *http.Request) {
	_, pageOpt := parseRequest(r)
	handleSeasons(w, r, func(s *entity.Season) bool {
		return true
	}, &pageOpt)
}

// 处理 /singerapi/api/seasons/{id}/songs/ 请求
func handleSeasonsSongs(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleSongs(w, r, func(s *entity.Song) bool {
		return s.HasSeason(id)
	}, &pageOpt)
}

// 处理 /singerapi/api/seasons/{id}/singers/ 请求
func handleSeasonsSingers(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleSingers(w, r, func(s *entity.Singer) bool {
		return s.HasSeason(id)
	}, &pageOpt)
}

// 处理 /singerapi/api/seasons/{id}/albums/ 请求
func handleSeasonsAlbums(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleAlbums(w, r, func(s *entity.Album) bool {
		return s.HasSeason(id)
	}, &pageOpt)
}
