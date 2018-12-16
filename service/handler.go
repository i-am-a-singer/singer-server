package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/singerapi/singer-server/database"
	"github.com/singerapi/singer-server/entity"
	"github.com/singerapi/singer-server/render"
)

// PageOptions for render Page
type PageOptions struct {
	Prefix  string
	URLPath string
	Page    int
	Limit   int
}

func parseRequest(r *http.Request) (ID int, Opt PageOptions) {
	Opt.Prefix = "http://" + r.Host + APIPrefix
	Opt.URLPath = "http://" + r.Host + r.URL.Path
	params := mux.Vars(r)
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	ID, _ = strconv.Atoi(params["id"])
	Opt.Page, Opt.Limit = 1, 0
	if pageParam != "" {
		Opt.Page, _ = strconv.Atoi(pageParam)
	}
	if limitParam != "" {
		Opt.Limit, _ = strconv.Atoi(limitParam)
	}
	return
}

// handleSeasonsList
func handleSeasons(w http.ResponseWriter, r *http.Request, filter func(s *entity.Season) bool, opt *PageOptions) {

	seas := []entity.Season{}
	database.QuerySeasonList(&seas, filter)

	entityList := []render.Resource{}
	for i := range seas {
		entityList = append(entityList, render.Resource(&seas[i]))
	}
	render.ResourceListJSONRender(w, http.StatusOK, entityList, opt.Prefix, opt.URLPath, opt.Page, opt.Limit)
}

// handleSongsList
func handleSongs(w http.ResponseWriter, r *http.Request, filter func(s *entity.Song) bool, opt *PageOptions) {

	seas := []entity.Song{}
	database.QuerySongList(&seas, filter)

	entityList := []render.Resource{}
	for i := range seas {
		entityList = append(entityList, render.Resource(&seas[i]))
	}
	render.ResourceListJSONRender(w, http.StatusOK, entityList, opt.Prefix, opt.URLPath, opt.Page, opt.Limit)
}

// handleSingersList
func handleSingers(w http.ResponseWriter, r *http.Request, filter func(s *entity.Singer) bool, opt *PageOptions) {

	seas := []entity.Singer{}
	database.QuerySingerList(&seas, filter)

	entityList := []render.Resource{}
	for i := range seas {
		entityList = append(entityList, render.Resource(&seas[i]))
	}
	render.ResourceListJSONRender(w, http.StatusOK, entityList, opt.Prefix, opt.URLPath, opt.Page, opt.Limit)
}

// handleAlbumsList
func handleAlbums(w http.ResponseWriter, r *http.Request, filter func(s *entity.Album) bool, opt *PageOptions) {

	seas := []entity.Album{}
	database.QueryAlbumList(&seas, filter)

	entityList := []render.Resource{}
	for i := range seas {
		entityList = append(entityList, render.Resource(&seas[i]))
	}
	render.ResourceListJSONRender(w, http.StatusOK, entityList, opt.Prefix, opt.URLPath, opt.Page, opt.Limit)
}

// errorHandler 404 msg
func notFindErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "custom 404")
}

// errorHandler 500 msg
func internalErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "custom 500")
}
