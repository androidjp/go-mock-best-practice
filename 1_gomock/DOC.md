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

# mockgen使用
1. 在项目根目录打开命令行
2. 找到对应目录下的某个将你要mock的接口所在的.go文件，生成对应的mock文件
    ```
    mockgen -source=1_gomock/db/repository.go  > test/1_gomock/db/mock_repository.go
    ```
    当然，前提是你这个 `test/1_gomock/db/`目录已经存在。

3. 然后，使用这个mock文件中的 `MockXxx(t)` 方法


# 关键用法
## 接口打桩步骤
1. 首先，使用mockgen工具，将对应的接口生成mock文件
2. 然后，开始打桩
    ```go
    // 1. 初始化 mock控制器
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    // 2. 初始化mock对象，并注入控制器
    mockRepo := mock_gomock_db.NewMockRepository(ctrl)
    
    // 3. 设定mock对象的返回值
    mockRepo.EXPECT().Create("name", []byte("jasper")).Return(nil)
    ```
3. 然后，测你要测的逻辑
    ```go
    // when
    demoSvr := &gomock_service.DemoService{Repo: mockRepo}
    data, err := demoSvr.InsertData("name", "jasper")
    // then
    assert.Equal(t, "success", data)
    assert.Nil(t, err)
    ```

## 接口打桩定义前N次返回值
```go
// 前两次返回错误
mockRepo.EXPECT().Create("name", []byte("jasper")).Return(errors.New("db connection error")).Times(2)
// 第三次正常
mockRepo.EXPECT().Create("name", []byte("jasper")).Return(nil)
```


## 断言接口调用顺序
方式一：`After`
```go
// retrieve 先执行
retrieveName := mockRepo.EXPECT().Retrieve("name").Return([]byte("jasper"), nil)
// update 在 retrieve 之后
mockRepo.EXPECT().Update("name", []byte("mike")).Return(nil).After(retrieveName) 
```
方式二：
```go
gomock.InOrder(
    // retrieve 先执行
    mockRepo.EXPECT().Retrieve("name").Return([]byte("jasper"), nil),
    // update 在 retrieve 之后
    mockRepo.EXPECT().Update("name", []byte("mike")).Return(nil),
)
```