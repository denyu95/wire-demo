## 操作
1. 进入 wire_set.go 文件，InitApp方法中按照注释切换MySQL/Redis源
2. 执行完毕在项目根目录执行wire命令，控制台输出`wire: wire-demo: wrote /Users/denyu/wire-demo/wire_gen.go`即成功
3. 执行 `go run wire_gen.go`

## 单元测试
在单元测试中可以自定义实现DataSource接口，来mock数据库取数据的行为。具体参照 wire_gen_test.go 文件