package service

import (
	"errors"
	"fmt"
	"os/exec"
)

/**
Exec函数的实现中调用了库函数exec.LoopPath和exec.Command，因此Exec函数的返回值和运行时的底层环境密切相关。在UT中，如果被测函数调用了Exec函数，则应根据用例的场景对Exec函数打桩。
*/
func Exec(cmd string, args ...string) (string, error) {
	cmdpath, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Errorf("exec.LookPath err: %v, cmd: %s", err, cmd)
		return "", errors.New("ErrExecLookPathFailed")
	}

	var output []byte
	output, err = exec.Command(cmdpath, args...).CombinedOutput()
	if err != nil {
		fmt.Errorf("exec.Command.CombinedOutput err: %v, cmd: %s", err, cmd)
		return "", errors.New("ErrExecCombinedOutputFailed")
	}
	fmt.Println("CMD[", cmdpath, "]ARGS[", args, "]OUT[", string(output), "]")
	return string(output), nil
}

// 过程
func InternalDoSth(mData map[string]interface{}) {
	mData["keyA"] = "valA"
}

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
