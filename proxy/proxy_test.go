package proxy

import "testing"

func TestProxy(t *testing.T) {
	proxy := &Proxy{}
	proxy.Request() // 通过代理类访问目标类
}
