package mapping

import "github.com/zihao-boy/zihao/entity/dto"

/**
编码映射实体类
*/
type MappingDto struct {
	dto.PageDto
	Id         string `json:"id"`
	Domain     string `json:"domain"`
	Name       string `json:"name"`
	Zkeys      string `json:"zkeys"`
	Value      string `json:"value"`
	Remark     string `json:"remark"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
}
