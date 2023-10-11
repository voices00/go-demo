package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil{
		fmt.Println("client dial err=",err)
		return
	}
	defer conn.Close()
	for{
		fmt.Println("请输入信息，回车结束输入")
		reader := bufio.NewReader(os.Stdin)
		//终端读取用户回车，并准备发送给服务器
		line,err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("readString err=",err)
		}
		line = strings.Trim(line,"\r\n")
		if line == "exit"{
			fmt.Println("客户端退出...")
			break
		}
		line = strings.TrimSpace(line)
		//将line发送给服务器
		n,err := conn.Write([]byte(line))
		if err != nil{
			fmt.Println("conn.Write err=",err)
		}
		fmt.Printf("发送的内容了%d文字\n",n)
	}
}
