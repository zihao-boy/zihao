package ip

import (
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/waf"
)

func InitIP() {
	IPData.FilePath = config.G_AppConfig.IpData
	IPData.InitIPData()
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

func GetAccessLogByIp(accessLogs []*waf.WafAccessLogDto)  {
	qqWry := NewQQwry()
	if len(accessLogs) > 0 {
		for _, v := range accessLogs {
			rs := qqWry.Find(v.XRealIp)
			v.IPAddress = rs.Area
			v.Country = rs.Country
		}
	}
}
