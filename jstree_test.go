package jstree

import (
	"log"
	"testing"
	"encoding/json"
)

// TestDirWalk ...
func TestDirWalk(t *testing.T) {
	jstree, err := DirWalk("testdata/results")
	if err != nil {
		t.Error(err)
	}
  log.Println(jstree)
}

// TestJstree ...
func TestJstree(t *testing.T)  {
  root, err:= Jstree("testdata/results")
  if err != nil {
    t.Error(err)
  }
  bytes, err := json.MarshalIndent(root, "", "\t")
  if err != nil {
    t.Error(err)
  }
  log.Println(string(bytes))
}
