package rest

import (
	"http"
)

// Checks if the request is actually permitted to access the resource in the requested way
type accessChecker interface {
	HasAccess(*http.Request) bool
}

// Lists all the items in the resource
// GET /resource/
type indexer interface {
	Index(http.ResponseWriter)
}

// Creates a new resource item
// POST /resource/
type creater interface {
	Create(http.ResponseWriter, *http.Request)
}

// Views a resource item
// GET /resource/id
type finder interface {
	Find(http.ResponseWriter, string)
}

// PUT /resource/id
type updater interface {
	Update(http.ResponseWriter, string, *http.Request)
}

// DELETE /resource/id
type deleter interface {
	Delete(http.ResponseWriter, string)
}

// Return options to use the service. If string is nil, then it is the base URL
// OPTIONS /resource/id
// OPTIONS /resource/
type optioner interface {
	Options(http.ResponseWriter, string)
}
