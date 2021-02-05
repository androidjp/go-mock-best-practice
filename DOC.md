## 使用场景
* 基本场景：为一个函数打桩
* 基本场景：为一个过程打桩
* 基本场景：为一个方法打桩
* 特殊场景：桩中桩的一个案例

## 局限
* Monkey不是线程安全的，不要将Monkey用于并发的测试
* 对inline函数打桩无效（一般需要：通过命令行参数`-gcflags=-l`禁止inline）
    ```go
    // 像这种函数，很简单短小，在源码层面来看时有函数结构的，但是编译后却不具备函数的性质。
    func IsEqual(a, b string) bool {
        return a == b
    }
    ```
* Monkey只能为首字母大写的方法/函数打桩（当然，这样其实更符合编码规范）。
* API不够简洁优雅，同时不支持多次调用桩函数（方法）而呈现不同行为的复杂情况。

## 安装
```sh
go get -v bou.ke/monkey
```

## 运行单元测试
方式一：命令行运行
```sh
go test -gcflags=-l -v
```

方式二：IDE 运行：在 Go tool Arguments 栏加上这一个选项：`-gcflags=-l`

## 关键用法
### 1. 函数打桩
1. 假设目前有这样一个函数 `Exec`：
    ```go
    package service
    
    func Exec(cmd string, args ...string) (string, error) {
        //...........
    }
    ```
2. 我们直接可以使用`monkey.Patch`将其打桩，不需要声明什么变量：
    ```go
    // 打桩
    guard := Patch(service.Exec, func(cmd string, args ...string) (string, error) {
        return "sss", nil
    })
    defer guard.Unpatch()
    // 调用
    output, err := service.Exec("cmd01", "--conf", "conf/app_test.yaml")
    ```

### 2. 过程打桩
和函数一样，相较于`gostub`好的一点，就是不需要声明变量去指向这个函数，从而减少业务代码的修改。
1. 假设有这样一个过程：
    ```go
    func InternalDoSth(mData map[string]interface{}) {
        mData["keyA"] = "valA"
    }
    ```
2. 一样方式进行打桩
    ```go
    patchGuard := Patch(service.InternalDoSth, func(mData map[string]interface{}) {
        mData["keyA"] = "valB"
    })
    defer patchGuard.Unpatch()
    
    ..............
    ```

### 3. 方法打桩
注意：只能打Public方法的桩，也就是首字母大写的方法
1. 假设有这样一个类以及它的方法定义：
    ```go
    type Etcd struct {
    }
    
    // 成员方法
    func (e *Etcd) Get(id int) []string {
        names := make([]string, 0)
        switch id {
        case 0:
            names = append(names, "A")
        case 1:
            names = append(names, "B")
        }
        return names
    }
    
    func (e *Etcd) Save(vals []string) (string, error) {
        return "存储DB成功", nil
    }
    
    func (e *Etcd) GetAndSave(id int) (string, error) {
        vals := e.Get(id)
        if vals[0] == "A" {
            vals[0] = "C"
        }
        return e.Save(vals)
    }
    ```
2. 通过`PatchInstanceMethod`即可打桩，然后直接调用:
    ```go
    // 打桩
    var e = &service.Etcd{}
    guard := PatchInstanceMethod(reflect.TypeOf(e), "Get", func(e *service.Etcd, id int) []string {
        return []string{"Jasper"}
    })
    defer guard.Unpatch()
    
    // 调用
    res := e.Get(1)
    ```
3. 当我想要一个测试用例里打多个成员方法的桩，这样即可：
    ```go
    var e = &service.Etcd{}
    // stub Get
    theVals := make([]string, 0)
    theVals = append(theVals, "A")
    PatchInstanceMethod(reflect.TypeOf(e), "Get", func(e *service.Etcd, id int) []string {
        return theVals
    })
    // stub Save
    PatchInstanceMethod(reflect.TypeOf(e), "Save", func(e *service.Etcd, vals []string) (string, error) {
        return "", errors.New("occurs error")
    })
    
    // 一键删除所有补丁
    defer UnpatchAll()
    
    .............
    ```

### 4. 配合gomock（桩中桩）
当我需要mock一个接口，并且，重新定义这个mock对象的某个方法时，使用。

详情看demo例子`repo_test.go`

