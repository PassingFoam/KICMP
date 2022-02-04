// +build !linux

package KICMP

import (
	"golang.org/x/net/ipv4"
)

func (s *ICMPSession) tx(txqueue []ipv4.Message) {
	s.defaultTx(txqueue)
}
