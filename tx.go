package KICMP

import (
	"golang.org/x/net/icmp"
	"sync/atomic"

	"github.com/pkg/errors"
	"golang.org/x/net/ipv4"
)

func (s *ICMPSession) defaultTx(txqueue []ipv4.Message) {
	nbytes := 0
	npkts := 0
	for k := range txqueue {
		//TODO 把测试id修改
		body := &icmp.Echo{
			ID: s.id,
			//ID:   s.id,
			Seq:  s.seq,
			Data: txqueue[k].Buffers[0],
		}


		msg := &icmp.Message{
			Type: s.icmpProto,
			Code: 0,
			Body: body,
		}
		sbytes, _ := msg.Marshal(nil)
		if n, err := s.conn.WriteTo(sbytes, txqueue[k].Addr); err == nil {
			nbytes += n
			npkts++
			s.seq++
		} else {
			s.notifyWriteError(errors.WithStack(err))
			break
		}
	}
	atomic.AddUint64(&DefaultSnmp.OutPkts, uint64(npkts))
	atomic.AddUint64(&DefaultSnmp.OutBytes, uint64(nbytes))
}
