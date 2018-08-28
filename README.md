## 学习Protobuf做的小实验，有proto3和proto2两个版本  
软件环境：Windows7 + liteide/Goland.注意：   
- Goland需要加src包下的项目源文件夹名，而liteide不需要。  
#### 学习日志  
- 2018.8.28中午：完成protobuf编码解码的逻辑练习；

#### 注意点Attention  
- 每次将proto协议文件生成.pb.go文件时记得把.pb.go文件里的fileDescriptor依次编号；  

#### 收获  
- 编码(Marshal)后的数据如下[10 5 104 101 108 108 111 16 0 24 1 24 2 24 3]，如果再加入数据报头，协议号，然后通过Write函数传递给接收方，接收方用Read函数接收并读取数据。

