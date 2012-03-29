/*
	Defines SnipsCollection struct-type and NewSnipsCollection()
	It is used by the example REST server for storing and managing snips
	It has the methods Add, WithId, All and Remove.
*/
package main

import (
	"log"
)

// Snip definintion
type Snip struct {
	Id   int
	Body string
}

// Create a new snippet
func newSnip(id int, body string) *Snip {
	log.Println("Creating new Snip:", id, body)
	return &Snip{id, body}
}

// SnipsCollection definition
type SnipsCollection struct {
	v      []Snip
	nextId int
}

// SnipsCollection creation function
func NewSnipsCollection() *SnipsCollection {
	log.Println("Creating new SnipsCollection")
	return &SnipsCollection{v: make([]Snip, 0), nextId: 0}
}

// Add a new snippet (snipped with passed text as body)
func (snips *SnipsCollection) Add(body string) int {
	log.Println("Adding Snip:", body)
	id := snips.nextId
	snips.nextId++

	snip := newSnip(id, body)
	snips.v = append(snips.v, *snip)

	return id
}

// Get a snippet by ID
func (snips *SnipsCollection) WithId(id int) (*Snip, bool) {
	log.Println("Searching for Snip with id: ", id)
	for _, snip := range snips.v {
		if snip.Id == id {
			return &snip, true
		}
	}
	return nil, false
}

// Get all snippets
func (snips *SnipsCollection) All() []Snip {
	log.Println("Finding all Snips")
	data := snips.v
	all := make([]Snip, len(snips.v))

	for k, snip := range data {
		all[k] = snip
	}

	return all
}

// Remove a snippet by ID
func (snips *SnipsCollection) Remove(id int) {

	newSnips := make([]Snip, len(snips.v))
	for _, snip := range snips.v {
		if snip.Id != id {
			newSnips = append(newSnips, snip)
		}
	}
	snips.v = newSnips
}
