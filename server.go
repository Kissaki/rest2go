package rest

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var resources = make(map[string]interface{})

// Generic resource handler
func resourceHandler(c http.ResponseWriter, req *http.Request) {
	uriPath := req.URL.Path

	// try to get resource with full uri path
	resource, ok := resources[uriPath]
	var id string
	if !ok {
		// no resource found, thus check if the path is a resource + ID
		i := strings.LastIndex(uriPath, "/")
		if i == -1 {
			log.Println("No slash found in URIPath ", uriPath)
			NotFound(c)
			return
		}
		// Move index to after slash as that’s where we want to split
		i++
		id = uriPath[i:]
		uriPathParent := uriPath[:i]
		resource, ok = resources[uriPathParent]
		if !ok {
			log.Println("Invalid URIPath-Parent ", uriPathParent)
			NotFound(c)
			return
		}
	}

	var hasAccess bool = false
	if accesschecker, ok := resource.(accessChecker); ok {
		hasAccess, _ = accesschecker.HasAccess(req)
	} else {
		// no checker for resource, so always give access
		log.Println("Resource ", uriPath, " has no accessChecker. Giving access …")
		hasAccess = true
	}
	if hasAccess {
		// call method on resource corresponding to the HTTP method
		switch req.Method {
		case "GET":
			if len(id) == 0 {
				// no ID -> Index
				if resIndex, ok := resource.(indexer); ok {
					resIndex.Index(c)
				} else {
					NotImplemented(c)
				}
			} else {
				// Find by ID
				if resFind, ok := resource.(finder); ok {
					resFind.Find(c, id)
				} else {
					NotImplemented(c)
				}
			}
		case "POST":
			// Create
			if resCreate, ok := resource.(creater); ok {
				resCreate.Create(c, req)
			} else {
				NotImplemented(c)
			}
		case "PUT":
			// Update
			if resUpdate, ok := resource.(updater); ok {
				resUpdate.Update(c, id, req)
			} else {
				NotImplemented(c)
			}
		case "DELETE":
			// Delete
			if resDelete, ok := resource.(deleter); ok {
				resDelete.Delete(c, id)
			} else {
				NotImplemented(c)
			}
		case "OPTIONS":
			// List usable HTTP methods
			if resOptions, ok := resource.(optioner); ok {
				resOptions.Options(c, id)
			} else {
				NotImplemented(c)
			}
		default:
			NotImplemented(c)
		}
	}
	return
}

// Add a resource route
func Resource(path string, res interface{}) {
	// check and warn for missing leading slash
	if fmt.Sprint(path[0:1]) != "/" {
		log.Println("Resource was added with a path with no leading slash. Did you mean to add /", path, " ?")
	}
	// add potentially missing trailing slash (resource always ends with slash)
	pathLen := len(path)
	if pathLen > 1 && path[pathLen-1:pathLen] != "/" {
		log.Println("adding trailing slash to ", path)
		path = fmt.Sprint(path, "/")
	}
	log.Println("Adding resource ", res, " at ", path)
	resources[path] = res
	http.Handle(path, http.HandlerFunc(resourceHandler))
}

// Emits a 404 Not Found
func NotFound(c http.ResponseWriter) {
	http.Error(c, "404 Not Found", http.StatusNotFound)
}

// Emits a 501 Not Implemented
func NotImplemented(c http.ResponseWriter) {
	http.Error(c, "501 Not Implemented", http.StatusNotImplemented)
}

// Emits a 201 Created with the URI for the new location
func Created(c http.ResponseWriter, location string) {
	c.Header().Set("Location", location)
	http.Error(c, "201 Created", http.StatusCreated)
}

// Emits a 200 OK with a location. Used when after a PUT
func Updated(c http.ResponseWriter, location string) {
	c.Header().Set("Location", location)
	http.Error(c, "200 OK", http.StatusOK)
}

// Emits a bad request with the specified instructions
func BadRequest(c http.ResponseWriter, instructions string) {
	c.WriteHeader(http.StatusBadRequest)
	c.Write([]byte(instructions))
}

// Emits a 204 No Content
func NoContent(c http.ResponseWriter) {
	http.Error(c, "204 No Content", http.StatusNoContent)
}
