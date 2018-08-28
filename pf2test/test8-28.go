package main

import(
	"ProtobufTest/pf2test/github.com/golang/protobuf/proto"
	"ProtobufTest/pf2test/protobuf"
	"log"
	"fmt"
)

func main() {
	//初始化test，将test8-28协议中的字段赋值
	test := &protobuf.Test{
		Label : proto.String("hello"),
		Type:proto.Int32(0),
		Reps:[]int64{1,2,3},
	}

	//对test进行编码
	data,err := proto.Marshal(test)
	if err != nil{
		log.Fatal("marshaling error: ",err)
	}
	fmt.Println(*test.Label)
	fmt.Println(*test.Type)
	fmt.Println(test.Reps)
	fmt.Println(data)

	//对newTest进行解码
	newTest := &protobuf.Test{}
	err = proto.Unmarshal(data,newTest)
	if err != nil{
		log.Fatal("unmarshaling error: ",err)
	}
	fmt.Println(*test.Label)
	fmt.Println(*test.Type)
	fmt.Println(test.Reps)
	fmt.Println(newTest)

	if test.GetLabel() != newTest.GetLabel(){
		log.Fatal("data mismatch %q != %q",test.GetLabel(),newTest.GetLabel())
	}
}
