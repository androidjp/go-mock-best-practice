# 使用场景
- 全局变量打桩
- 函数打桩
- 过程打桩
- 第三方库打桩

# 安装
```
go get -v -u github.com/prashantv/gostub
```

# 关键用法

### 1. 全局变量打桩

对于全局变量：

```Go
var (
  GlobalCount int
  Host        string
)
```

可以这样打桩：

```Go
// 全局变量 GlobalCount int 打桩
// 全局变量 Host string 打桩
stubs := gostub.Stub(&GlobalCount, 10).
  Stub(&Host, "https://www.bing.cn")
defer stubs.Reset() 
```

### 2. 函数打桩

假设有个函数：

```Go
func Exec(cmd string, args ...string) (string, error) {
  return "", nil
}
```

那么，首先我要先变成这样的写法：

```Go
var Exec = func(cmd string, args ...string) (string, error) {
  return "", nil
}
```

以上写法不影响业务逻辑使用。

然后再进行打桩：

方式一：`StubFunc` 直接设置返回结果

```Go
stubs := gostub.StubFunc(&gomock_service.Exec, "xxx-vethName100-yyy", nil)
defer stubs.Reset()
```

方式二：`Stub` 还能设置具体逻辑

```Go
stubs := gostub.Stub(&Exec, func(cmd string, args ...string) (string, error) {
      return "xxx-vethName100-yyy", nil
    })
defer stubs.Reset()
```

### 3. 过程打桩

对于一些没有返回值的函数，我们称为“过程”：

```Go
var DestroyResource = func() {
  fmt.Println("清理资源等工作")
}

```

打桩开始：

方式一：`StubFunc` 直接设置返回结果（当你想这个过程啥都不做时，可以这样）

```Go
stubs := gostub.StubFunc(&gomock_service.DestroyResource)
defer stubs.Reset()
```

方式二：`Stub` 还能设置具体逻辑

```Go
stubs := gostub.Stub(&gomock_service.DestroyResource, func() {
      // do sth
    })
defer stubs.Reset()
```

### 4. 第三方库打桩

很多第三方库的函数（注意，是函数，不是某个对象的某个成员方法），我们会经常使用，而在单元测试时不是我们的关注点，或者想他报错等，就可以选择打桩。

1. 假如，我想打桩json的序列化和反序列化函数，那么，先在`adapter`包下定义`json.go`文件，然后声明对象：
    ```Go
    package adapter
    
    import (
      "encoding/json"
    )
    
    var Marshal = json.Marshal
    var UnMarshal = json.Unmarshal 
    ```
2. 单元测试中，就可以直接使用`gostub.StubFunc`来进行打桩了：
    ```Go
    // given
    var mikeStr = `{"name":"Jasper", "age":18}`
    stubs := gostub.StubFunc(&adapter.Marshal, []byte(mikeStr), nil)
    defer stubs.Reset()
    
    stu := &entity.Student{Name: "Mike", Age: 18}
    
    // when
    res, err := stu.Print()
    
    // then
    assert.Equal(t, "{\"name\":\"Jasper\", \"age\":18}", res)
    assert.Nil(t, err)
    ```