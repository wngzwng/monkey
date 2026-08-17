```go
// 跑测试
go test 模块地址
//例如
go test ./parser

// 单独跑某一个测试
go test -v -run TestOperatorPrecedenceParsing ./parser
```