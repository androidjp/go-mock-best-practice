package adapter

import (
	"encoding/json"
)

// 使用 adapter包来存放为了测试打桩而声明的变量
// 指向第三方库函数的对象（为了打桩而使用）
var Marshal = json.Marshal
var UnMarshal = json.Unmarshal
