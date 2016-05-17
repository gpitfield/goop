package goop

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func getTestNode(t *testing.T) (node *html.Node) {
	data, err := os.Open("test.html")
	defer data.Close()
	checkErr(err, t)
	node, err = html.Parse(data)
	checkErr(err, t)
	return
}

func checkErr(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

// Attr returns the node's value for the named attr, otherwise empty.
func TestAttr(t *testing.T) {
	main := getTestNode(t)
	emptyDiv := FindElementType("div", main)
	noAttr := Attr("class", emptyDiv)
	if noAttr != "" {
		t.Fatalf("TestAttr: attr found where should be none.")
	}
}

// HasAttr returns true if the named attr of the given node contains the supplied value, otherwise false.
func TestHasAttr(t *testing.T) {
	main := getTestNode(t)
	noClassDiv := FindElementType("div", main)
	noAttr := HasAttr("class", "none", noClassDiv)
	if noAttr != false {
		t.Fatalf("TestHasAttr: attr found where there should be none.")
	}
	classNone := FindElementType("div", noClassDiv.NextSibling.NextSibling)
	noneAttr := HasAttr("class", "none", classNone)
	if noneAttr != true {
		t.Fatalf("TestHasAttr: class attr `none` not found.")
	}
}

// ParseNodeAttr traverses the dom tree, and executes the parse function on each node with the matching
// value for the given attr.
func TestParseNodeAttr(t *testing.T) {
	main := getTestNode(t)
	parse := func(node *html.Node) {
		if c := Attr("class", node); c != "some" {
			t.Fatalf("TestParseNodeAttr has wrong class", c)
		} else {
			fmt.Println("TestParseNodeAttr: correct class found.")
		}
	}
	ParseNodeAttr("class", "some", main, parse)
}

// FindNodeAttr finds and returns the first node in the input node matching the supplied attribute and value.
// attr can be any html attr.
func TestFindNodeAttr(t *testing.T) {
	main := getTestNode(t)
	span := FindNodeAttr("data-test", "test", main)
	if id := Attr("id", span); id != "span3" {
		t.Fatalf("TestFindNodeAttr failed", id)
	}
	style := FindNodeAttr("style", "display:none", main)
	if id := Attr("id", style); id != "styleDiv" {
		t.Fatalf("TestFindNodeAttr failed", id)
	}
}

// FindElementType finds and returns the first element in the input node with the supplied type.
func TestFindElementType(t *testing.T) {
	main := getTestNode(t)
	span := FindElementType("span", main)
	if id := Attr("id", span); id != "span1" {
		t.Fatalf("TestFindElementType failed.")
	}
}

// HasClass returns true if the supplied node has the given html class.
func TestHasClass(t *testing.T) {
	main := getTestNode(t)
	div := FindNodeAttr("class", "another", main)
	if one := HasClass("one", div); one != true {
		t.Fatalf("TestHasClass failed; missing class `one`")
	}
}

// FindNodeClass finds and returns the first node in the input node with the given html class.
func TestFindNodeClass(t *testing.T) {
	main := getTestNode(t)
	div := FindNodeClass("some", main)
	if c := HasClass("some", div); c != true {
		t.Fatalf("TestFindNodeClass failed - wrong class")
	}
}

// HasId returns true if the supplied node has the given html id.
func TestHasId(t *testing.T) {
	main := getTestNode(t)
	span := FindNodeId("span4", main)
	if id := HasId("span4", span); id != true {
		t.Fatalf("TestHasId failed.")
	}
}

func TestFindNodeId(t *testing.T) {
	main := getTestNode(t)
	span := FindNodeId("span4", main)
	if id := HasId("span4", span); id != true {
		t.Fatalf("TestHasId failed.")
	}
}
