// main project main.go
package main

import (
	"fmt"
	"ProtobufTest/pf2test/protobuf"
	"ProtobufTest/pf2test/github.com/golang/protobuf/proto"
)

func main() {
	bodyData1 := "chengdu/tianfu/feiyu/company"
	bodyData2 := "chengdu/tianfu/feiyu/company"

	p := &protobuf.StringMessage{
		Body: proto.String(bodyData1),
		Header: &protobuf.Header{
			MessageId: proto.String("123456  "),
			Topic:     proto.String("Golang"),
		},
	}

	p3 := &protobuf.StringMessage{
		Body: proto.String(bodyData2),
		Header: &protobuf.Header{
			MessageId: proto.String("456789  "),
			Topic:     proto.String("Clang"),
		},
	}

	pData1, err := proto.Marshal(p)
	pData2, err := proto.Marshal(p3)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(pData1))

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(pData2))

	p2 := &protobuf.StringMessage{}
	proto.Unmarshal(pData1, p2)

	p4 := &protobuf.StringMessage{}
	proto.Unmarshal(pData2, p4)

	fmt.Println(p2.GetHeader().GetTopic())
	fmt.Println(p4.GetHeader().GetTopic())
}
