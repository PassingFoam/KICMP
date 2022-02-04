// +build linux

package KICMP

import (
	"net"
	"os"
	"sync/atomic"

	"github.com/pkg/errors"
	"golang.org/x/net/ipv4"
)

// the read loop for a client session
func (s *ICMPSession) readLoop() {
	// default version
	if s.xconn == nil {
		s.defaultReadLoop()
		return
	}

	// x/net version
	var src string
	msgs := make([]ipv4.Message, batchSize)
	for k := range msgs {
		msgs[k].Buffers = [][]byte{make([]byte, mtuLimit)}
	}

	for {
		if count, err := s.xconn.ReadBatch(msgs, 0); err == nil {
			for i := 0; i < count; i++ {
				msg := &msgs[i]
				// make sure the packet is from the same source
				if src == "" { // set source address if nil
					src = msg.Addr.String()
				} else if msg.Addr.String() != src {
					atomic.AddUint64(&DefaultSnmp.InErrs, 1)
					continue
				}

				// source and size has validated
				s.packetInput(msg.Buffers[0][:msg.N])
			}
		} else {
			// compatibility issue:
			// for linux kernel<=2.6.32, support for sendmmsg is not available
			// an error of type os.SyscallError will be returned
			if operr, ok := err.(*net.OpError); ok {
				if se, ok := operr.Err.(*os.SyscallError); ok {
					if se.Syscall == "recvmmsg" {
						s.defaultReadLoop()
						return
					}
				}
			}
			s.notifyReadError(errors.WithStack(err))
			return
		}
	}
}

// monitor incoming data for all connections of server
func (l *Listener) monitor() {
	var xconn batchConn
	//if _, ok := l.conn.(*icmp.PacketConn); ok {
	//	addr, err := net.ResolveIPAddr("ip", l.conn.LocalAddr().String())
	//	if err == nil {
	//		if addr.IP.To4() != nil {
	//			xconn = l.conn.(*icmp.PacketConn).IPv4PacketConn()
	//		} else {
	//			xconn = l.conn.(*icmp.PacketConn).IPv6PacketConn()
	//		}
	//	}
	//}

	// default version
	if xconn == nil {
		l.defaultMonitor()
		return
	}

	// x/net version
	msgs := make([]ipv4.Message, batchSize)
	for k := range msgs {
		msgs[k].Buffers = [][]byte{make([]byte, mtuLimit)}
	}

	for {
		if count, err := xconn.ReadBatch(msgs, 0); err == nil {
			for i := 0; i < count; i++ {
				msg := &msgs[i]
				l.packetInput(msg.Buffers[0][20:msg.N], msg.Addr)
			}
		} else {
			// compatibility issue:
			// for linux kernel<=2.6.32, support for sendmmsg is not available
			// an error of type os.SyscallError will be returned
			if operr, ok := err.(*net.OpError); ok {
				if se, ok := operr.Err.(*os.SyscallError); ok {
					if se.Syscall == "recvmmsg" {
						l.
							defaultMonitor()
						return
					}
				}
			}
			l.notifyReadError(errors.WithStack(err))
			return
		}
	}
}
