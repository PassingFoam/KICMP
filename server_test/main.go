package main

import (
	"KICMP"
	"fmt"
	"io"
	"net"
	"strconv"
)

func main(){
	fmt.Println("server test")
	fmt.Println("kcp listens")
	lis, err := KICMP.ListenWithOptions("0.0.0.0", nil, 10, 3)
	if err!=nil {
		panic(err)
	}
	for {
		conn, e :=lis.AcceptKCP()
		if e!=nil {
			panic(e)
		}
		fmt.Println("some from remote")
		go func(conn net.Conn){
			var buffer = make([]byte,1024,1024)
			for {
				n,e :=conn.Read(buffer)
				if e!=nil {
					if e == io.EOF {
						break
					}
					fmt.Println(e)
					break
				}

				fmt.Println("receive from client:", string(buffer[:n]))
				for i:=0;i<=10;i++{
					conn.Write([]byte("aaaafkyou" + strconv.Itoa(i)))

				}



			}
		}(conn)
	}
	//fmt.Println("single server test")
	//
	//conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	//if err != nil {
	//	return
	//}
	//
	//for{
	//	var buffer = make([]byte, 1500+8)
	//
	//	n, srcaddr, err := conn.ReadFrom(buffer)
	//
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println(srcaddr.String())
	//	fmt.Println(string(buffer[:n]))
	//}

}



