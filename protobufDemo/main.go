package main

import (
	"protobufDemo/pb"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func main() {

	//序列化
	person := &pb.Person{
		Name:"Jack",
		Age:18,
		Hobby:[]string{"sing","dance","basketball","rap"},
	}

	binaryData, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("proto.Marshal err:",err)
	}

	//为方便测试，序列化和反序列化在同一个程序实现

	//反序列化
	newPerson := &pb.Person{}
	err = proto.Unmarshal(binaryData,newPerson)
	if err != nil {
		fmt.Println("proto.Unmarshal err:",err)
	}

	fmt.Println("序列化前的原始数据:",person)
	fmt.Println("反序列化得到数据:",newPerson)
}
