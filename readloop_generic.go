// +build !linux

package KICMP

// 从udp 读取信息
func (s *ICMPSession) readLoop() {
	s.defaultReadLoop()
}

func (l *Listener) monitor() {
	l.defaultMonitor()
}
