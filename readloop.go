package KICMP

import (
	"sync/atomic"

	"github.com/pkg/errors"
)

const ProtocolICMP = 1

func (s *ICMPSession) defaultReadLoop() {
	buf := make([]byte, mtuLimit + 8)
	var src string
	for {
		if n, addr, err := s.conn.ReadFrom(buf); err == nil {
			// make sure the packet is from the same source
			if src == "" { // set source address
				src = addr.String()
			} else if addr.String() != src {
				atomic.AddUint64(&DefaultSnmp.InErrs, 1)
				continue
			}
			//fmt.Println(string(buf[:n]))
			s.packetInput(buf[:n])
		} else {
			s.notifyReadError(errors.WithStack(err))
			return
		}
	}
}

func (l *Listener) defaultMonitor() {
	buf := make([]byte, mtuLimit + 8)
	for {
		if n, from, err := l.conn.ReadFrom(buf); err == nil {

			l.packetInput(buf[:n], from)
		} else {
			l.notifyReadError(errors.WithStack(err))
			return
		}
	}
}
