package render

import (
	"fmt"
	"io"

	"github.com/unrolled/render"
)

// Resource is an interface can generate a struct that all ID are converted to URL resource
type Resource interface {
	APIFormatConstruct(HostPrefix string) interface{}
}

// ResourceList type format resource list
type ResourceList struct {
	Count       int           `json:"count"`
	NextURL     string        `json:"next"`
	PreviousURL string        `json:"previous"`
	RList       []interface{} `json:"results"`
}

// DefaultLimit is page limit for a resource list
var (
	format       = render.New(render.Options{IndentJSON: true})
	DefaultLimit = 10
)

// ResourceItemJSONRender render resource item to JSON Format
func ResourceItemJSONRender(w io.Writer, status int, r Resource, HostPrefix string) error {
	apiFormatedItem := r.APIFormatConstruct(HostPrefix)
	return format.JSON(w, status, apiFormatedItem)
}

// ResourceListJSONRender render resource item to JSON Format
func ResourceListJSONRender(w io.Writer, status int, resourceList []Resource, HostPrefix string, APIURL string, page, limit int) error {
	if limit <= 0 {
		limit = DefaultLimit
	}
	respBody := ResourceList{
		len(resourceList),
		"",
		"",
		[]interface{}{},
	}
	if page > 1 {
		respBody.PreviousURL = fmt.Sprintf("%s?page=%d", APIURL, page-1)
	}
	if page < respBody.Count/limit+1 {
		respBody.NextURL = fmt.Sprintf("%s?page=%d", APIURL, page+1)
	}
	for i, re := range resourceList {
		if i/limit+1 == page {
			respBody.RList = append(respBody.RList, re.APIFormatConstruct(HostPrefix))
		}
	}
	return format.JSON(w, status, respBody)
}

// MapJSONRender render map to JSON Format
func MapJSONRender(w io.Writer, status int, respBody map[string]string) error {
	return format.JSON(w, status, respBody)
}
