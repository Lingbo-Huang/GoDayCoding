package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node       // 存储每种请求方式的Trie 树根节点
	handlers map[string]HandlerFunc // 存储每种请求方式的 HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 把路径解析为有效的路由部分切片
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' { // 遇到*开头的通配符就不需要继续解析了，直接返回带有这个路由部分的parts
				break
			}
		}
	}
	return parts
}

// 注册路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	// 检查这个方法的前缀树是不是空的
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	// 往前缀树里插入
	r.roots[method].insert(pattern, parts, 0)
	// 注册HandlerFunc到map里
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	// path: /users/123 -> searchParts: ["users", "123"]
	// path: /files/documents/report.pdf -> searchParts: ["files", "documents", "report.pdf"]
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	// n.pattern: /users/:id
	// n.pattern: /files/*filepath
	// 找到匹配的路由的终点节点
	n := root.search(searchParts, 0)

	if n != nil {
		// parts: ["users", ":id"]
		// parts: ["files", "*filepath"]
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				// params["id"] = "123"
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				// params["filepath"] = "documents/report.pdf"
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 page not found: %s\n", c.Path)
	}
}
