package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

const (

	//发生攻击的类型，仅在攻击日志中出现。

	State_default         = "default"         //：默认
	State_sqli            = "sqli"            //：SQL注入攻击
	State_xss             = "xss"             //：跨站脚本攻击
	State_webshell        = "webshell"        //：WebShell攻击
	State_robot           = "robot"           //：恶意爬虫
	State_cmdi            = "cmdi"            //：命令注入攻击
	State_rfi             = "rfi"             //：远程文件包含
	State_lfi             = "lfi"             //： 本地文件包含
	State_illegal         = "illegal"         //：非法请求
	State_vuln            = "vuln"            //：漏洞攻击
	State_cc              = "cc"              //：命中CC防护规则
	State_custom_custom   = "custom_custom"   //：命中精准防护规则
	State_custom_whiteip  = "custom_whiteip"  //：命中IP黑白名单规则
	State_custom_geoip    = "custom_geoip"    //：命中地理位置控制规则
	State_antitamper      = "antitamper"      //： 命中网页防篡改规则
	State_anticrawler     = "anticrawler"     //：命中JS挑战反爬虫规则
	State_leakage         = "leakage"         //：命中敏感信息泄露规则
	State_followed_action = "followed_action" //：攻击惩罚
)

type WafAccessLogDto struct {
	dto.PageDto
	RequestId      string    `json:"requestId" sql:"-" `
	WafIp          string    `json:"wafIp" sql:"-" `
	HostId         string    `json:"hostId"  sql:"-"`
	XRealIp        string    `json:"xRealIp"  sql:"-"`
	Scheme         string    `json:"scheme" `
	ResponseCode   string    `json:"responseCode"  sql:"-"`
	Method         string    `json:"method" `
	HttpHost       string    `json:"httpHost"  sql:"-"`
	UpstreamAddr   string    `json:"upstreamAddr"  sql:"-"`
	Url            string    `json:"url" `
	RequestLength  string    `json:"requestLength"  sql:"-"`
	ResponseLength string    `json:"responseLength"  sql:"-"`
	State          string    `json:"state" `
	Message        string    `json:"message" `
	CreateTime     time.Time `json:"createTime" sql:"-"`
	StartTime      string    `json:"startTime" `

	StatusCd  string `json:"statusCd" sql:"-"`
	IPAddress string `json:"ipAddress"`
	Country   string `json:"country"`
}

type WafAccessLogMapDto struct {
	dto.PageDto
	WafIp      string `json:"wafIp" sql:"-" `
	WafCountry string `json:"wafCountry"`
	XRealIp    string `json:"xRealIp"  sql:"-"`
	IPAddress  string `json:"ipAddress"`
	Country    string `json:"country"`
}
