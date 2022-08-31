package zentao

// UserType 用户类型
type UserType string

var (
	InSideUser  UserType = "inside"  // 内部用户
	OutSideUser UserType = "outside" // 外部用户
)

type UserGender string

var (
	ManGender   UserGender = "m" // 男
	WomanGender UserGender = "f" // 女
)

// CustomResp 通用Resp
type CustomResp struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
