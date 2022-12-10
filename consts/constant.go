package consts

const (
	TraceKey = "traceId"
)

const (
	Production  = "Production"
	PreRelease  = "Pre-release"
	Test        = "Test"
	Development = "Development"
)

const (
	Critical  = 1 // 异常
	Passing   = 2 // 健康
	Normal    = 1 // 正常（非隔离）
	Isolation = 2 // 隔离
)

const (
	SyncNSInterval = 5 // 健康更新周期，单位：秒
)
