
## 模拟性能优化点
- block 模拟阻塞
- cpu 模拟cpu密集型
- goroutine 模拟goroutine泄露
- mem 模拟堆内存泄露
- mutex 模拟锁竞争
## 分析
- 引入 _ "net/http/pprof"包
- 启动http服务
- 访问http://localhost:6060/debug/pprof可以查看pprof的信息
- go tool pprof -h 查看帮助
- go tool pprof http://localhost:6060/debug/pprof/allocs 查看allocs的信息，进入交互模式
- top: 查看cpu使用率前10的函数
  * flat: 当前函数直接分配的内存
  * flat%: 直接分配内存占总内存的百分比
  * sum%: 累积百分比
  * cum: 当前函数及其调用的所有函数分配的内存
  * cum%: 累积内存占总内存的百分比
- list Run: 查看函数Run的源码以及对应的内存分配情况
- web: 生成svg文件，查看函数调用关系(brew install graphviz)
- go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile 可以通过web访问pprof的信息
- 用单元测试的方式测试性能  
```shell 
go test -bench . ./data_test/ -blockprofile block.out -cpuprofile cpu.out -memprofile mem.out -outputdir ./data_test/testout
go tool pprof -http=:8080 /Users/huanglingbo2/GolandProjects/GoDayCoding/pprof/data_test/testout/cpu.out
```
## 注意点
**不是所有的block都是需要优化的**
