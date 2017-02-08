/*
Copyright 2017 yasukun

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package jstree

import (
	"fmt"
	"os"
	"path/filepath"
)

const jspre = "jst"

type (
	Node struct {
		ID     string `json:"id"`
		Parent string `json:"parent"`
		Text   string `json:"text"`
	}
	Nodes struct {
		Data []Node `json:"data"`
	}
	Root struct {
		Core Nodes `json:"core"`
	}
)

// DirWalk ...
func DirWalk(root string) ([]Node, error) {
	parents := make(map[string]int)
	jstree := []Node{}
	idx := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		rel, _ := filepath.Rel(root, path)
		if _, ok := parents[rel]; !ok {
			parents[rel] = idx

		}
		if val, ok := parents[filepath.Dir(rel)]; ok {
			node_id := fmt.Sprintf("%s_%d", jspre, idx)
			parent_id := fmt.Sprintf("%s_%d", jspre, val)
			if node_id == parent_id {
				parent_id = "#"
			}
			// log.Printf("id: %s, parent: %s, text: %s\n", node_id, parent_id, info.Name())
			jstree = append(jstree, Node{ID: node_id, Parent: parent_id, Text: info.Name()})
		}
		idx += 1
		return nil
	})

	if err != nil {
		return jstree, err
	}

	return jstree, nil
}

// Jstree ...
func Jstree(root string) (Root, error) {
	jstree, err := DirWalk(root)
	if err != nil {
		return Root{Core: Nodes{Data: []Node{}}}, err
	}
	return Root{Core: Nodes{Data: jstree}}, nil
}
