package encrypt

import (
	"crypto/md5"
	"fmt"
	"io"
)

const (
	SLAT_CODE string = "zihaoboyisverygood"
)

/**
md5加密
add by 吴学文
time 2021-01-18
 */
func Md5(srcStr string) string{
	w := md5.New()
	io.WriteString(w, srcStr+SLAT_CODE)
	//将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	return md5str2
}
