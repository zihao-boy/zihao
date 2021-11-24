package local

import (
	"encoding/json"
	"fmt"

	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/entity/dto/mapping"
	"github.com/zihao-boy/zihao/system/mapper"
)

func (r *Local) SaveMapping(mappingDto mapping.MappingDto) (err error) {
	data, err := json.Marshal(mappingDto)

	r.client.Set(mappingDto.Domain+mappingDto.Zkeys, data,
		-1)
	return err
}

func (r *Local) GetMapping(zKey string) (mapping mapping.MappingDto, err error) {
	data, _ := r.client.Get(constants.DOMAIN_COMMON + zKey)
	json.Unmarshal(data.([]byte), &mapping)
	return mapping, err
}

// Init
func InitMapping() {
	var (
		mappingAllMapper mapper.MappingAllMapper
		mappingDto       mapping.MappingDto = mapping.MappingDto{}
		mappingDtos      []*mapping.MappingDto
		err              error
	)

	mappingDtos, err = mappingAllMapper.GetMappings(mappingDto)

	if err != nil {
		fmt.Print("加载sql 失败", err)
		return
	}

	for _, item := range mappingDtos {
		G_Local.SaveMapping(*item)
	}
}
