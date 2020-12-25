package Boo

import "strings"

type node struct {
	path     string
	children []*node
	isWild   bool
	handler  HandlerFunc
}

func splitPath(path string) []string {
	vs := strings.Split(path, "/")
	var parts = make([]string, 0)
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func isWild(path string) bool {
	return path[0] == ':' || path[0] == '*'
}

func (n *node) matchChild(path string) *node {
	for _, child := range n.children {
		if child.path == path || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) insert(fullPath string, handler HandlerFunc) {
	parts := splitPath(fullPath)

	for i, part := range parts {
		child := n.matchChild(part)
		if child == nil {
			child = &node{path: part, isWild: isWild(part), handler: handler}
			n.children = append(n.children, child)
		} else if child.isWild && child.path != parts[i] {
			prefix := fullPath[:strings.Index(fullPath, parts[i])] + child.path
			panic("新路径" + fullPath + "中的" + parts[i] + "与现有前缀" + prefix + "中的通配符" + child.path + "冲突")
		}

		n = child
	}
}

func (n *node) search(fullPath string) (*node, []Param) {
	parts := splitPath(fullPath)
	params := make([]Param, 0)

	for i, part := range parts {
		child := n.matchChild(part)
		if child == nil {
			return nil, nil
		}
		n = child
		param := Param{key: child.path[1:], value: parts[i]}
		params = append(params, param)
	}

	return n, params
}
