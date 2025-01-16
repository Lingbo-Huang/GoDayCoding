package gee

import (
	"fmt"
	"strings"
)

/*
内存优化：可以采用一些压缩技术，如路径压缩（将只有一个子节点的节点合并）
		或者字符编码优化（减少字符存储的空间）来减少内存占用。
更新操作优化：设计更合理的更新和删除算法，例如，在删除节点时，采用标记删除而不是立即删除的方法，
			等到合适的时机（如系统空闲或者下一次路由重建）再进行真正的删除操作，以减少对路由匹配的即时影响。
			同时，提供更方便的更新接口，允许批量更新或者动态加载新的路由规则。
*/

type node struct {
	pattern  string  // 到该节点的带匹配路由
	part     string  // 该节点处的值
	children []*node // 子节点
	isWild   bool    // 是否模糊匹配
}

// String 节点信息
func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

// pattern 是完整的路由模式，parts 是将路由拆分成的字符串切片，height 表示当前正在处理的路由部分的索引。
func (n *node) insert(pattern string, parts []string, height int) {
	// 如果已经处理完所有的路由部分，将 pattern 存储在当前节点中并返回。
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 获取当前高度对应的路由部分
	part := parts[height]
	child := n.matchChild(part)
	if child == nil { // 没匹配上就插入
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 匹配上之后以该子节点为node继续向下遍历，直到height抵达路由长度
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	// 如果递归到底部或者当前节点的路径部分是通配符
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" { // 如果该节点不是有效的匹配终点
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	// 对所有匹配的子节点挨个递归遍历，一旦路径匹配到终点就直接返回
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

// 第一个匹配成功的子节点，用于插入(开发服务时，注册路由规则，映射handler)
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找(访问时，匹配路由规则，查找到对应的handler)
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
