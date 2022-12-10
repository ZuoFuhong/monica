package entity

type ServiceVO struct {
	Id       int    `json:"id"`       // 服务ID
	Name     string `json:"name"`     // 服务名称
	Business string `json:"business"` // 业务名称
	Owners   string `json:"owners"`   // 负责人
	Remark   string `json:"remark"`   // 备注
	Ctime    int64  `json:"ctime"`    // 创建时间
	Mtime    int64  `json:"mtime"`    // 更新时间
}

type ServiceNSVO struct {
	Business  string `json:"business"` // 业务名称
	Name      string `json:"name"`     // 服务名称
	Namespace string `json:"ns"`       // 命名空间
	Token     string `json:"token"`    // 鉴权凭证
	Ctime     int64  `json:"ctime"`    // 创建时间
}

type InstanceVO struct {
	ID          int    `json:"id"`
	Namespace   string `json:"namespace"`    // 命名空间
	ServiceName string `json:"service_name"` // 服务名称
	Healthy     int    `json:"healthy"`      // 健康状态：1-异常 2-健康
	Isolate     int    `json:"isolate"`      // 隔离状态：1-不隔离 2-隔离
	IP          string `json:"ip"`           // 实例IP
	Port        int    `json:"port"`         // 端口
	Weight      int    `json:"weight"`       // 权重
	Metadata    string `json:"metadata"`     // 元数据
	Rtime       int64  `json:"rtime"`        // 健康更新时间
	Ctime       int64  `json:"ctime"`        // 创建时间
	Mtime       int64  `json:"mtime"`        // 更新时间
}

type InstanceNode struct {
	IP       string `json:"ip"`       // 实例IP
	Port     int    `json:"port"`     // 端口
	Weight   int    `json:"weight"`   // 权重
	Metadata string `json:"metadata"` // 元数据
}
