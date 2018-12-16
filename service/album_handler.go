package service

import (
	"net/http"

	"github.com/singerapi/singer-server/database"
	"github.com/singerapi/singer-server/entity"
	"github.com/singerapi/singer-server/render"
)

// handleAlbumsID 处理 /singerapi/api/albums/{id} 请求
func handleAlbumsID(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)

	sea := entity.Album{}
	if err := database.Query(id, &sea); err != nil {
		notFindErrorHandler(w, r)
		return
	}
	render.ResourceItemJSONRender(w, http.StatusOK, &sea, pageOpt.Prefix)
}

// 处理 /singerapi/api/albums/ 请求
func handleAlbumsAll(w http.ResponseWriter, r *http.Request) {
	_, pageOpt := parseRequest(r)
	handleAlbums(w, r, func(s *entity.Album) bool {
		return true
	}, &pageOpt)
}

// 处理 /singerapi/api/albums/{id}/seasons/ 请求
func handleAlbumsSeasons(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleSeasons(w, r, func(s *entity.Season) bool {
		return s.HasAlbum(id)
	}, &pageOpt)
}

// 处理 /singerapi/api/albums/{id}/songs/ 请求
func handleAlbumsSongs(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleSongs(w, r, func(s *entity.Song) bool {
		return s.HasAlbum(id)
	}, &pageOpt)
}

// 处理 /singerapi/api/albums/{id}/singers/ 请求
func handleAlbumsSingers(w http.ResponseWriter, r *http.Request) {
	id, pageOpt := parseRequest(r)
	handleSingers(w, r, func(s *entity.Singer) bool {
		return s.HasAlbum(id)
	}, &pageOpt)
}
