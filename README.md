A RESTful HTTP client and server.


Installation
============

Install by running:

	goinstall github.com/nathankerr/rest.go


Usage
=====

See the *examples* folder for coded examples.

rest.go uses the standard http package to provide resources through specified resource routes. Add a new route and corresponding resource using the Resource method:

	rest.Resource("resourcepath", resourcevariable)

The resource URI-path is then */resourcepath/*.

A *resource* is an object that may have any of the following methods which respond to the specified HTTP requests:

	GET /resource/ => Index(http.ResponseWriter)
	GET /resource/id => Find(http.ResponseWriter, id string)
	POST /resource/ => Create(http.ResponseWriter, *http.Request)
	PUT /resource/id => Update(http.ResponseWriter, id string, *http.Request)
	DELETE /resource/id => Delete(http.ResponseWriter, id string)
	OPTIONS /resource/ => Options(http.ResponseWriter, id string)
	OPTIONS /resource/id => Options(http.ResponseWriter, id string)

If you want to add a permission/accessibility-check to the resource, implement

	HasAccess(*http.Request) (bool, os.Error)


The server will then route HTTP requests to the appropriate method call.

The snips example provides a full example of both a client and server.
The weekdays example project implements a simple REST server.

