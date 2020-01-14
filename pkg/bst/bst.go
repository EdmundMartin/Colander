package bst

import (
	"bytes"
	"encoding/gob"
	"os"
)

type Node struct {
	Key   int
	Value interface{}
	Left  *Node
	Right *Node
}

func compareStrings(first, second string) int {
	if first > second {
		return -1
	} else if first == second {
		return 0
	}
	return 1
}

func (n *Node) Compare(other interface{}) int {
	switch v := n.Value.(type) {
	case string:
		oth := other.(string)
		return compareStrings(v, oth)
	}
	return 0
}

func NewNode(key int, val interface{}) *Node {
	return &Node{Key: key, Value: val}
}

type BST struct {
	Root  *Node
	Items int
}

type Result struct {
	PK int
	Value interface{}
}

func NewBST() *BST {
	return &BST{}
}

func LoadFromFile(filename string) (*BST, error) {
	bt := &BST{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(bt)
	if err != nil {
		return nil, err
	}
	return bt, nil
}

func (bt *BST) Insert(pk int, val interface{}) {
	// Insert root node
	if bt.Root == nil {
		bt.Root = NewNode(bt.Items, val)
		return
	}
	node := bt.Root
	for {
		if node.Compare(val) == -1 {
			if node.Left == nil {
				node.Left = NewNode(pk, val)
				return
			} else {
				node = node.Left
			}
		} else if node.Compare(val) == 0 {
			if node.Right == nil {
				node.Right = NewNode(pk, val)
				return
			} else {
				node = node.Right
			}
		} else {
			if node.Right == nil {
				node.Right = NewNode(pk, val)
				return
			} else {
				node = node.Right
			}
		}
	}
}

func (bt *BST) All() []Result {
	results := []Result{}
	if bt.Root == nil {
		return results
	}
	stack := []*Node{bt.Root}
	var node *Node
	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		results = append(results, Result{node.Key, node.Value})
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return results
}

func (bt *BST) Save(filename string) error {
	buf := &bytes.Buffer{}
	fo, _ := os.Create(filename)
	defer fo.Close()
	err := gob.NewEncoder(buf).Encode(bt)
	if err != nil {
		return err
	}
	fo.Write(buf.Bytes())
	return nil
}
