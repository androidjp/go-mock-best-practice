使用场景：
* 基本场景：为一个函数打桩
* 基本场景：为一个过程打桩
* 基本场景：为一个方法打桩
* 复合场景：由任意相同或不同的基本场景组合而成
* 特殊场景：桩中桩的一个案例

局限：
* Monkey不是线程安全的，不要将Monkey用于并发的测试

安装：
```sh
go get -v bou.ke/monkey
```

运行单元测试：
1. 命令行运行：
```sh
go test -gcflags=-l -v
```
2. IDE 运行：在 Go tool Arguments 栏加上这一个选项：`-gcflags=-l`




mockgen -source=repository/repo.go  > mock/mock_repository/mock_repo.go