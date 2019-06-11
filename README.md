# 说明

[Demo源码](https://github.com/meetbetter/protocol-buffer-demo)

protocol buffer是谷歌推出的高效率序列化反序列化工具，可以自定义数据结构，然后使用对应语言的代码生成器生成的代码读写这个数据结构。虽然在和前端打交道时还是要配合使用JSON，但是在后端内部可以尝试使用protocol buffer改进性能。

下面总结下在GoLang中使用protocol buffer的方法。

# Go环境配置

## 下载protobuf

```shell
git clone https://github.com/protocolbuffers/protobuf.git
```

## 安装(Linux Ubuntu)

### (1)安装依赖工具

```shell
sudo apt-get install autoconf automake libtool curl make g++ unzip libffi-dev -y
```

### (2)进入protobuf文件

```shell
cd protobuf/
```

### (3)进行安装检测 并生成自动安装脚本

```shell
./autogen.sh
./configure
```

### (4)进行编译C代码和安装

```shell
make
sudo make install
```

### (5)刷新linux共享库关系

```shell
sudo ldconfig
```

### (6)测试protobuf编译工具

```shell
protoc -h
```

如果正常输出 相关指令 没有报任何error，为安装成功。



## 获取 GoLang的protobuf包

由于protocol buffer原生不支持go语言，需要下载golang的protobuf插件。

###  (1)下载

```shell
go get -v -u github.com/golang/protobuf/proto
```

### (2)进入到文件夹内进行编译

```shell
cd $GOPATH/src/github.com/golang/protobuf/protoc-gen-go/
go build
```

### (3)拷贝可执行文件

将生成的 protoc-gen-go可执行文件，放在/bin目录下。

```shell
sudo cp protoc-gen-go /bin/
```

尝试补齐protoc-gen-go 如果可以补齐代表成功，如果执行不报错 代表工具成功。

# Go使用protobuf

## 新建.proto文件

基本格式如下：

```protobuf
syntax = "proto3"; //必须指定protobuf协议版本号
package pb; //包名

//定义一个protobuf协议
message Person {
    string name = 1; //数字表示序号，并不是变量值.
    int32 age = 2;
    repeated string hobby = 3; //对应go中[]string

}
```

## 生成Go数据结构

在.proto所在目录执行如下命令，

```shell
protoc --go_out=.  *.proto
```

在当前目录下会生成对应的.go文件，可以在其中找到go的数据结构，

```go
type Person struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Hobby                []string `protobuf:"bytes,3,rep,name=hobby,proto3" json:"hobby,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```

**注意**，该文件只能生成不能手动修改。

## 使用protobuf

在main函数中新建Person对象并进行序列化和反序列化，

```go
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

	//反序列化
	newPerson := &pb.Person{}
	err = proto.Unmarshal(binaryData,newPerson)
	if err != nil {
		fmt.Println("proto.Unmarshal err:",err)
	}

	fmt.Println("序列化前的原始数据:",person)
	fmt.Println("反序列化得到数据:",newPerson)
}

```



执行后可以看到person和newPerson都喜欢**唱、跳、篮球和rap**。

```shell
序列化前的原始数据: name:"Jack" age:18 hobby:"sing" hobby:"dance" hobby:"basketball" hobby:"rap" 
反序列化得到数据: name:"Jack" age:18 hobby:"sing" hobby:"dance" hobby:"basketball" hobby:"rap"
```

