package u_crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {

	var (
		md = md5.New()
	)

	md.Write([]byte(str))
	cipherStr := md.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
