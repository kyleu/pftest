// Content managed by Project Forge, see [projectforge.md] for details.
package filesystem

import (
	"path"

	"github.com/kyleu/pftest/app/util"
)

type Node struct {
	Name     string `json:"name"`
	Dir      bool   `json:"dir,omitempty"`
	Children Nodes  `json:"children,omitempty"`
}

func (n *Node) Get(path ...string) *Node {
	if len(path) == 0 {
		return n
	}
	for _, node := range n.Children {
		if node.Name == path[0] {
			return node.Get(path[1:]...)
		}
	}
	return nil
}

func (n *Node) Flatten(curr string) []string {
	x := path.Join(curr, n.Name)
	ret := n.Children.Flatten(x)
	if !n.Dir {
		ret = append(ret, x)
	}
	return ret
}

type Nodes []*Node

func (n Nodes) Flatten(curr string) []string {
	ret := make([]string, 0, len(n))
	for _, node := range n {
		ret = append(ret, node.Flatten(curr)...)
	}
	return ret
}

type Tree struct {
	Nodes  Nodes         `json:"nodes,omitempty"`
	Config util.ValueMap `json:"config,omitempty"`
	keys   []string
}

func (t Tree) Flatten() []string {
	return t.Nodes.Flatten("")
}
