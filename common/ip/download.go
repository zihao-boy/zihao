package ip


import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"github.com/zihao-boy/zihao/config"
	"io/ioutil"
	"net/http"
)

// @ref https://zhangzifan.com/update-qqwry-dat.html

// getKey from cz88.net
func getKey() (uint32, error) {
	resp, err := http.Get(config.G_AppConfig.QqwryUrl)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return 0, err
	} else {
		// @see https://stackoverflow.com/questions/34078427/how-to-read-packed-binary-data-in-go
		return binary.LittleEndian.Uint32(body[5*4:]), nil
	}
}

// GetOnline get data from online server
func GetOnline() ([]byte, error) {
	resp, err := http.Get(config.G_AppConfig.QqwryUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		if key, err := getKey(); err != nil {
			return nil, err
		} else {
			for i := 0; i < 0x200; i++ {
				key = key * 0x805
				key++
				key = key & 0xff

				body[i] = byte(uint32(body[i]) ^ key)
			}

			reader, err := zlib.NewReader(bytes.NewReader(body))
			if err != nil {
				return nil, err
			}

			return ioutil.ReadAll(reader)
		}
	}
}