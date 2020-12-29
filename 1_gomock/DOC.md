# 简单的mock--gomock
gomock + mockgen

# 安装
```shell script
go get github.com/golang/mock/gomock
```
运行完后你会发现，在$GOPATH/src目录下有了github.com/golang/mock子目录，且在该子目录下有GoMock包和mockgen工具。
```shell script
cd $GOPATH/src/github.com/golang/mock/mockgen
go build
```
则在当前目录下生成了一个可执行程序mockgen。

将mockgen程序移动到$GOPATH/bin目录下：
```
mv mockgen $GOPATH/bin
```
最后，尝试敲一下命令行：
```shell script
mockgen help
```
如果出现`-bash: mockgen: command not found`，表示你的环境变量PATH中没有配置`$GOPATH/bin`。


# 文档
```
go doc github.com/golang/mock/gomock
```
[在线参考文档](https://link.jianshu.com/?t=http://godoc.org/github.com/golang/mock/gomock)

# 使用
1. 在项目根目录打开命令行
2. 找到对应目录下的某个将你要mock的接口所在的.go文件，生成对应的mock文件
    ```
    mockgen -source=1_gomock/db/repository.go  > test/1_gomock/db/mock_repository.go
    ```
    当然，前提是你这个 mock_spider目录已经存在。

3. 然后，使用这个mock文件中的 `MockXxx(t)` 方法

# 参考文章
https://www.jianshu.com/p/f4e773a1b11f



# GoStub
https://www.jianshu.com/p/70a93a9ed186
