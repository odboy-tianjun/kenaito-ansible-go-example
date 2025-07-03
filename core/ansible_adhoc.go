package core

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apenella/go-ansible/v2/pkg/adhoc"
	"github.com/apenella/go-ansible/v2/pkg/execute"
	results "github.com/apenella/go-ansible/v2/pkg/execute/result/json"
	"io"
	"kenaito-ansible-go-example/model"
	"strings"
	"time"
)

// ExecuteCommand
// mode 模式。local、ssh、smart
// inventoryList 主机列表。
// moduleName 模块名称。
// args 模块参数。
// extraVars 变量。
func ExecuteCommand(mode string, inventoryList []string, moduleName string, args string, extraVars map[string]interface{}) model.CsExecResult {
	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Connection: mode,
		Inventory:  strings.Join(inventoryList, ","),
		ModuleName: moduleName,
		Args:       args,
		ExtraVars:  extraVars,
	}
	/**
	"all" 表示该命令将作用于 Inventory（清单）中的所有主机。
	"webservers"（针对 webservers 组的主机）
	"192.168.1.1"（针对特定 IP 的主机）
	"db*"（匹配所有以 db 开头的主机名）
	*/
	adhocCmd := adhoc.NewAnsibleAdhocCmd(
		adhoc.WithPattern("all"),
		adhoc.WithAdhocOptions(ansibleAdhocOptions),
	)

	result := model.CsExecResult{}
	result.Command = adhocCmd.String()

	buff := new(bytes.Buffer)
	defaultExecute := execute.NewDefaultExecute(
		execute.WithCmd(adhocCmd),
		execute.WithEnvVars(map[string]string{
			"ANSIBLE_LOAD_CALLBACK_PLUGINS": "1",
			"ANSIBLE_STDOUT_CALLBACK":       "json",
		}),
		execute.WithWrite(buff),
	)
	startTime := time.Now()
	timeLayout := "2006-01-02 15:04:05"
	result.StartTime = startTime.Format(timeLayout)
	err := defaultExecute.Execute(context.TODO())
	if err != nil {
		fmt.Printf("ansible命令行参数异常: %v\n", err)
		result.Success = false
		result.ErrMsg = err.Error()
		endTime := time.Now()
		result.EndTime = endTime.Format(timeLayout)
		result.Duration = int(endTime.Sub(startTime).Milliseconds())
		return result
	}

	// 解析的结果有问题
	res, err := results.ParseJSONResultsStream(io.Reader(buff))
	if err != nil {
		fmt.Printf("执行结果解析异常: %v\n", err)
		result.Success = false
		result.ErrMsg = err.Error()
		endTime := time.Now()
		result.EndTime = endTime.Format(timeLayout)
		result.Duration = int(endTime.Sub(startTime).Milliseconds())
		return result
	}
	result.Success = true
	result.Data = res.Stats
	endTime := time.Now()
	result.EndTime = endTime.Format(timeLayout)
	result.Duration = int(endTime.Sub(startTime).Milliseconds())
	return result

	// 自行解析，也有问题，降级吧
	//var data map[string]interface{}
	//jsonStr, err := util.ReaderToString(io.Reader(buff))
	//fmt.Printf("jsonStr => %s\n\n", jsonStr)
	//err = json.Unmarshal([]byte(jsonStr), &data)
	//if err != nil {
	//	fmt.Printf("执行结果解析异常: %v\n", err)
	//	result.Success = false
	//	result.ErrMsg = err.Error()
	//	endTime := time.Now()
	//	result.EndTime = endTime.Format(timeLayout)
	//	result.Duration = int(endTime.Sub(startTime).Milliseconds())
	//	return result
	//}
	//result.Success = true
	//result.Data = data["stats"]
	//endTime := time.Now()
	//result.EndTime = endTime.Format(timeLayout)
	//result.Duration = int(endTime.Sub(startTime).Milliseconds())
	//return result
}
