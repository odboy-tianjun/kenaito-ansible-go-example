package model

type CsExecResult struct {
	// 执行的ansible命令
	Command string
	// 执行成功与否
	Success bool
	// 执行结果
	//ok 表示任务执行成功，没有做任何修改；
	//changed 表示任务执行成功，并且做了一些修改；
	//failed 表示任务执行失败；
	//unreachable 表示任务无法连接到目标主机。
	Data interface{}
	// 异常信息
	ErrMsg string
	// 开始时间
	//StartTime    time.Time
	StartTime string
	// 结束时间
	//EndTime    time.Time
	EndTime string
	// 执行耗时
	Duration int
}
