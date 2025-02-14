package gee_cache

type ByteView struct {
	b []byte // 存储缓存值
}

func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice 返回字节切片的copy
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
