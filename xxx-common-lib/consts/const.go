package consts

const AuthTokenHeaderKey = "Authorization"

type UserCtxKey struct {
}

// 是否删除：1-否，2-是
const (
	// DeleteNot 1-未删除
	DeleteNot = 1
	// DeleteYes 2-已删除
	DeleteYes = 2
)
