package ip

import (
	"errors"
	"fmt"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"io/ioutil"
	"net"
	"net/http"
)

var Service_IP string

func InitIP() {
	IPData.FilePath = config.G_AppConfig.IpData
	IPData.InitIPData()
	Service_IP = GetServerIp(config.G_AppConfig.ServerIpUrl)
}

func GetOnlyIpAddr(ip string) ResultQQwry {
	qqWry := NewQQwry()

	return qqWry.Find(ip)

}

func GetIpAddr(ips []string) map[string]ResultQQwry {
	qqWry := NewQQwry()
	rs := map[string]ResultQQwry{}
	if len(ips) > 0 {
		for _, v := range ips {
			rs[v] = qqWry.Find(v)
		}
	}
	return rs
}

func GetAccessLogByIp(accessLogs []*waf.WafAccessLogDto) {
	qqWry := NewQQwry()
	if len(accessLogs) > 0 {
		for _, v := range accessLogs {
			rs := qqWry.Find(v.XRealIp)
			v.IPAddress = rs.Area
			v.Country = rs.Country
		}
	}
}

func GetAccessLogMapByIp(accessLogs []*waf.WafAccessLogMapDto) {
	qqWry := NewQQwry()
	if len(accessLogs) > 0 {
		for _, v := range accessLogs {
			rs := qqWry.Find(v.XRealIp)
			v.IPAddress = rs.Area
			v.Country = rs.Country

			rs = qqWry.Find(v.WafIp)
			v.WafCountry = rs.Country
		}
	}
}

func GetServerIp(serverIpUrl string) string {
	responseClient, errClient := http.Get(serverIpUrl) // 获取外网 IP
	if errClient != nil {
		fmt.Printf("获取外网 IP 失败，请检查网络\n")
		ip, _ := GetClientIp()
		return ip
	}
	// 程序在使用完 response 后必须关闭 response 的主体。
	defer responseClient.Body.Close()

	body, _ := ioutil.ReadAll(responseClient.Body)
	clientIP := fmt.Sprintf("%s", string(body))
	return clientIP
}

func GetClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("Can not find the client ip address!")
}
