```
通常用于协程之间的通信

通常我们线上的业务场景可能是这样的
FetchMeta------> doSth-------->

FetchMeta 是goroutineA
doSth     是goroutineB 


我们的最终诉求是：
 goroutineA-finished 
 goroutineB 开始执行
 
goroutineA <------ goroutineB



```