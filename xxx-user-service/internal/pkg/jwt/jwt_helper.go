package xjwt

import (
    "encoding/json"
    "example.com/xxx/user-service/internal/conf"
    "github.com/cristalhq/jwt/v5"
)

type Key string

type JwtHelper struct {
    key []byte
}

func NewJwtHelper(config *conf.Bootstrap) *JwtHelper {
    return &JwtHelper{
        key: []byte(config.Jwt.Key),
    }
}

type UserClaims struct {
    jwt.RegisteredClaims
    Role string // 角色
}

// CreateToken 生成 token
func (a *JwtHelper) CreateToken(claims UserClaims) (*jwt.Token, error) {
    signer, _ := jwt.NewSignerHS(jwt.HS256, a.key)
    // create a Builder
    builder := jwt.NewBuilder(signer)
    // and build a Tokenjm
    return builder.Build(claims)
}

// VerifyToken 解析并校验 token
func (a *JwtHelper) VerifyToken(token string) (*jwt.Token, error) {
    verifier, _ := jwt.NewVerifierHS(jwt.HS256, a.key)
    return jwt.Parse([]byte(token), verifier)
}

func (a *JwtHelper) ParseToken(token *jwt.Token) (UserClaims, error) {
    var newClaims UserClaims
    err := json.Unmarshal(token.Claims(), &newClaims)
    return newClaims, err
}
