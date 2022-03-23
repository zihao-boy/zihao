package waf

import "time"

type WafAccessLogDto struct {
	RequestId     string    `json:"requestId" sql:"-" `
	WafId     string    `json:"wafId" sql:"-" `
	HostId     string    `json:"hostId"  sql:"-"`
	XrealIp     string    `json:"xRealIp"  sql:"-"`
	Scheme     string    `json:"scheme" `
	ResponseCode     string    `json:"responseCode"  sql:"-"`
	Method     string    `json:"method" `
	HttpHost     string    `json:"httpHost"  sql:"-"`
	UpstreamAddr     string    `json:"upstreamAddr"  sql:"-"`
	Url     string    `json:"url" `
	RequestLength     string    `json:"requestLength"  sql:"-"`
	ResponseLength     string    `json:"responseLength"  sql:"-"`
	State     string    `json:"state" `
	Message     string    `json:"message" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}

