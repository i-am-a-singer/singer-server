package service

import (
	"net/http"

	"github.com/singerapi/singer-server/database"
	"github.com/singerapi/singer-server/entity"
	"github.com/singerapi/singer-server/render"
)

// handleSongsID 处理 /singerapi/api/songs/{id} 请求
func handleSongsID(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)

	sea := entity.Song{}
	if err := database.Query(id, &sea); err != nil {
		notFindErrorHandler(w, r)
		return
	}
	render.ResourceItemJSONRender(w, http.StatusOK, &sea, pageOpt.Prefix)
}

// 处理 /singerapi/api/songs/ 请求
func handleSongsAll(w http.ResponseWriter, r *http.Request) {
	_, pageOpt := parseRequest(r)
	handleSongs(w, r, func(s *entity.Song) bool {
		return true
	}, &pageOpt)
}

// 处理 /singerapi/api/songs/{id}/seasons/ 请求
func handleSongsSeasons(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleSeasons(w, r, func(s *entity.Season) bool {
		return s.HasSong(id)
	}, &pageOpt)
}

// 处理 /singerapi/api/songs/{id}/singers/ 请求
func handleSongsSingers(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleSingers(w, r, func(s *entity.Singer) bool {
		return s.HasSong(id)
	}, &pageOpt)
}

// 处理 /singerapi/api/songs/{id}/albums/ 请求
func handleSongsAlbums(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleAlbums(w, r, func(s *entity.Album) bool {
		return s.HasSong(id)
	}, &pageOpt)
}
