package encrypt

import (
	"encoding/base64"
	"fmt"
)

func Encode(data string) string {
	// Base64 Standard Encoding
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	//fmt.Println(sEnc) // aGVsbG8gd29ybGQxMjM0NSE/JComKCknLUB+
	return sEnc
}

func Decode(sEnc string) string {
	// Base64 Standard Decoding
	sDec, err := base64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return ""
	}
	return string(sDec) //hello world12345!?$*&()'-@~
}
