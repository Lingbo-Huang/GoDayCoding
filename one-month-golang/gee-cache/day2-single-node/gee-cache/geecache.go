package gee_cache

import "sync"

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

type Getter interface {
	Get(key string) ([]byte, error)
}

/*
如果缓存不存在，应从数据源（文件，数据库等）获取数据并添加到缓存中。
为了扩展性，如何从源头获取数据，应该是用户决定的事情，我们就把这件事交给用户好了。
将一个带有相同签名的函数直接包装成接口的实现。
不需要为每种具体实现定义一个结构体。只要一个函数符合接口方法的签名，便可以创建该接口的实现。
因此，我们设计了一个回调函数(callback)，在缓存不存在时，调用这个函数，得到源数据。
*/

/*
结构体方式：适用于需要管理内部状态、拥有多个复杂成员函数的情况。经典的例子包括数据库连接管理、缓存管理等。

函数方式（适配器模式）：适合于简单的功能实现，以及需要将不同的行为以参数传递的场景。例如，处理HTTP请求的中间件函数，经常使用这种模式，因为它们通常是无状态的，并且每个函数都只处理单一任务。
*/

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}
