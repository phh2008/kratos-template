package crypto

import "errors"

var (
    ErrKeyLength  = errors.New("secret key length is invalid")
    ErrCipherText = errors.New("cipherText is too short")
)
