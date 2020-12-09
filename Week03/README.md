## 学习笔记

1. 开启gorouting的时候，func本身不要做隐式的`go func`， 而是在调用方进行 go 调用
2. errgroup是一个比较好的用于多个goroutine 返回错误收集的包，有多种使用模式，需要自己练习掌握
3. channel的buff不会增加性能，只能在一定时间内降低延迟，性能需要多个worker来增加
4. context尽量少用KV， 如果需要可以使用多版本，上游ctx看不到下游的ctx，保持隔离，其他的context类型按需使用
