package core

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/measure"
	results "github.com/apenella/go-ansible/v2/pkg/execute/result/json"
	"github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
	"io"
	"kenaito-ansible-go-example/model"
	"kenaito-ansible-go-example/util"
	"strings"
	"time"
)

// ExecutePlaybook
// mode 模式。local、ssh、smart
// inventoryList 主机列表。
// playbookList playbook脚本列表。
// extraVars 变量。
func ExecutePlaybook(mode string, inventoryList []string, playbookList []string, extraVars map[string]interface{}) model.CsExecResult {
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: mode,
		Inventory:  strings.Join(inventoryList, ","),
		ExtraVars:  extraVars,
	}
	/**
	"all" 表示该命令将作用于 Inventory（清单）中的所有主机。
	"webservers"（针对 webservers 组的主机）
	"192.168.1.1"（针对特定 IP 的主机）
	"db*"（匹配所有以 db 开头的主机名）
	*/
	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(playbookList...),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	result := model.CsExecResult{}
	result.Command = playbookCmd.String()

	buff := new(bytes.Buffer)
	exec := measure.NewExecutorTimeMeasurement(
		stdoutcallback.NewJSONStdoutCallbackExecute(
			execute.NewDefaultExecute(
				execute.WithCmd(playbookCmd),
				execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
				execute.WithWrite(io.Writer(buff)),
			),
		),
	)

	startTime := time.Now()
	result.StartTime = util.FormatDateTime(startTime)
	err := exec.Execute(context.TODO())
	if err != nil {
		fmt.Printf("ansible-playbook命令行参数异常: %v\n", err)
		result.Success = false
		result.ErrMsg = err.Error()
		endTime := time.Now()
		result.EndTime = util.FormatDateTime(endTime)
		result.Duration = int(endTime.Sub(startTime).Milliseconds())
		return result
	}

	res, err := results.ParseJSONResultsStream(io.Reader(buff))
	if err != nil {
		fmt.Printf("执行结果解析异常: %v\n", err)
		result.Success = false
		result.ErrMsg = err.Error()
		endTime := time.Now()
		result.EndTime = util.FormatDateTime(endTime)
		result.Duration = int(endTime.Sub(startTime).Milliseconds())
		return result
	}

	result.Success = true
	result.Data = res.Stats
	endTime := time.Now()
	result.EndTime = util.FormatDateTime(endTime)
	result.Duration = int(endTime.Sub(startTime).Milliseconds())
	return result
}
