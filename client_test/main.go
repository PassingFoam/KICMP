package main

import (
	"KICMP"
	"fmt"
	"io"
)

func main(){
	kcpconn, err := KICMP.DialWithOptions("10.211.55.3", nil, 10, 3)
	if err!=nil {
		panic(err)
	}

	_, err2 := kcpconn.Write([]byte("hello kcp.emmmmmmmmmmmmmmm"))
	if err2 != nil {
		fmt.Println(err2)
	}
	_, err2 = kcpconn.Write([]byte("h123123123"))


	for i:=0;i<=10;i++{
		var buffer = make([]byte,1024,1024)
		n,e :=kcpconn.Read(buffer)
		if e!=nil {
			if e == io.EOF {
				break
			}
			fmt.Println(e)
			break
		}

		fmt.Println("receive from server:", string(buffer[:n]))
		//kcpconn.Write([]byte(strconv.Itoa(i)))
		//kcpconn.Write([]byte("client test"))
	}




	//addr, err := net.ResolveIPAddr("ip", "1.117.196.158")
	//if err != nil {
	//	return
	//}
	//
	//conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	//if err != nil {
	//	return
	//}
	//
	//
	//body := &icmp.Echo{
	//	ID:   96201,
	//	Seq:  22,
	//	Data: []byte("icmp test2"),
	//}
	//
	//
	//msg := &icmp.Message{
	//	Type: ipv4.ICMPTypeEcho,
	//	Code: 0,
	//	Body: body,
	//}
	//sbytes, err := msg.Marshal(nil)
	//_, err = conn.WriteTo(sbytes,addr)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	select {

	}
}
