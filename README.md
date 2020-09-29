# go-learn
golang study


## int 和 int32 的区别


go 官方文档说明 https://golang.org/pkg/builtin/#int32

> int is a signed integer type that is at least 32 bits in size.It is a distinct type,not an alias for int32


SDS(Simple Dynamic String):
- buf 实际数据 
- len 已使用长度 4 字节
- alloc 分配的长度 alloc >= len 4字节

redis string 内存布局(SDS+RedisObject):
- Long 类型整数,RedisObject 指针直接赋值为整数数据
- 保存字符串数据,字符数据小于 44 字节,RedisObject 跟 SDS 连续分布
- 大于 44 字节,
- 元数据 8 字节
- 指针 指向 dicEntry 8 字节
- dicEntry 
  - 
  - 
  - 
