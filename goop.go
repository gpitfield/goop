// Package goop adds a more semantic set of parsing functions to golang.org/x/net/html. Used in conjunction with
// the traversal functions in golang.org/x/net/html, it enables the selection of elements based on html attributes,
// with specific helper functions for class and id.
//
// All of the FindNodeX functions consider the supplied node as well as all of its subnodes.
package goop

import (
	"strings"

	"golang.org/x/net/html"
)

// Attr returns the node's value for the named attr, otherwise empty.
func Attr(attr string, node *html.Node) string {
	for _, at := range node.Attr {
		if at.Key == attr {
			return at.Val
		}
	}
	return ""
}

// HasAttr returns true if the named attr of the given node contains the supplied value, otherwise false.
func HasAttr(attr string, value string, node *html.Node) bool {
	if node.Attr == nil || node == nil {
		return false
	}
	for _, at := range node.Attr {
		if at.Key == attr {
			vals := []string{at.Val}
			if attr == "style" {
				vals = strings.Split(at.Val, ";")
			} else {
				vals = strings.Split(at.Val, " ")
			}
			for _, val := range vals {
				if val == value {
					return true
				}
			}
		}
	}
	return false
}

// ParseNodeAttr traverses the dom tree, and executes the parse function on each node with the matching
// value for the given attr.
func ParseNodeAttr(attr string, value string, in *html.Node, parse func(*html.Node)) {
	if HasAttr(attr, value, in) {
		parse(in)
	}
	for c := in.FirstChild; c != nil; c = c.NextSibling {
		ParseNodeAttr(attr, value, c, parse)
	}
	return
}

// FindNodeAttr finds and returns the first node in the input node matching the supplied attribute and value.
// attr can be any html attr.
func FindNodeAttr(attr string, value string, in *html.Node) (out *html.Node) {
	if HasAttr(attr, value, in) {
		return in
	}
	for c := in.FirstChild; c != nil; c = c.NextSibling {
		if match := FindNodeAttr(attr, value, c); match != nil {
			return match
		}
	}
	return
}

// FindElementType finds and returns the first element in the input node with the supplied type.
func FindElementType(nodeType string, in *html.Node) (out *html.Node) {
	if in.Type == html.ElementNode && in.Data == nodeType {
		return in
	}
	for c := in.FirstChild; c != nil; c = c.NextSibling {
		if match := FindElementType(nodeType, c); match != nil {
			return match
		}
	}
	return
}

// HasClass returns true if the supplied node has the given html class.
func HasClass(c string, node *html.Node) bool {
	return HasAttr("class", c, node)
}

// FindNodeClass finds and returns the first node in the input node with the given html class.
func FindNodeClass(class string, in *html.Node) (out *html.Node) {
	return FindNodeAttr("class", class, in)
}

// HasId returns true if the supplied node has the given html id.
func HasId(id string, node *html.Node) bool {
	return HasAttr("id", id, node)
}

// FindNodeId finds and returns the first node in the input node with the given html id.
func FindNodeId(id string, in *html.Node) (out *html.Node) {
	return FindNodeAttr("id", id, in)
}
