package Boo

// 前缀树节点
type node struct {
	path      string
	wildChild bool
	children  []*node
	handler   HandlerFunc
}

func (n *node) matchChild(path string) *node {
	for _, child := range n.children {
		if child.path == path {
			return child
		}
	}
	return nil
}

// 插入子节点
func (n *node) insert(pattern string, parts []string, handler HandlerFunc) {
	if len(parts) == 0 && pattern == "/" {
		n.path = pattern
		n.handler = handler
		return
	}

	for i := 0; i < len(parts); i++ {
		child := n.matchChild(parts[i])
		if child == nil {
			child = &node{path: parts[i], wildChild: parts[i][0] == ':' || parts[i][0] == '*'}
			n.children = append(n.children, child)
		}
		n = child
	}
	n.handler = handler
}

// 查找节点
func (n *node) getNode() {

}
