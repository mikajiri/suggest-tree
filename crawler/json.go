package crawler

type TreeNode struct {
	Text  string     `json:"text,omitempty"`
	Nodes []TreeNode `json:"nodes,omitempty"`
}

func (t *TreeNode) Add(node TreeNode, indexes []int) {
	if len(indexes) > 0 {
		i := indexes[0]
		for len(t.Nodes) <= i {
			t.Add(TreeNode{}, []int{})
		}
		t.Nodes[i].Add(node, indexes[1:])
	} else {
		t.Nodes = append(t.Nodes, node)
	}
}
