package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"gorm.io/gorm"
)

const (
	query_monitorTaskAttr string = `
				select t.*,tts.spec_name from task_attr t
left join task_template_spec tts on t.spec_cd = tts.spec_cd and tts.status_cd = '0'
				where 
				t.status_cd = '0'

			$if AttrId != '' then
			and t.attr_id = #AttrId#
			$endif
			$if TaskId != '' then
				and t.task_id = #TaskId#
			$endif
				
	`

	insert_monitorTaskAttr string = `	
insert into task_attr(task_id, attr_id, spec_cd, value) 
VALUES (#TaskId#, #AttrId#, #SpecCd#, #Value#)
	`

	update_monitorTaskAttr string = `
			update task_attr t set t.value = #Value#
where t.attr_id = #AttrId#
and t.status_cd = '0'
	`
	delete_monitorTaskAttr string = `
	update task_attr t set
			t.status_cd = '1'
			where
			 1=1 
			 t.attr_id = #AttrId#
			and t.status_cd = '0'
	`
)

type MonitorTaskAttrDao struct {
}

/**
查询用户
*/
func (*MonitorTaskAttrDao) GetMonitorTaskAttrs(monitorTaskAttrDto monitor.MonitorTaskAttrDto) ([]*monitor.MonitorTaskAttrDto, error) {
	var monitorTaskAttrDtos []*monitor.MonitorTaskAttrDto
	sqlTemplate.SelectList(query_monitorTaskAttr, objectConvert.Struct2Map(monitorTaskAttrDto), func(db *gorm.DB) {
		db.Scan(&monitorTaskAttrDtos)
	}, false)

	return monitorTaskAttrDtos, nil
}

/**
保存服务sql
*/
func (*MonitorTaskAttrDao) SaveMonitorTaskAttr(monitorTaskAttrDto monitor.MonitorTaskAttrDto) error {
	return sqlTemplate.Insert(insert_monitorTaskAttr, objectConvert.Struct2Map(monitorTaskAttrDto), false)
}

/**
修改服务sql
*/
func (*MonitorTaskAttrDao) UpdateMonitorTaskAttr(monitorTaskAttrDto monitor.MonitorTaskAttrDto) error {
	return sqlTemplate.Update(update_monitorTaskAttr, objectConvert.Struct2Map(monitorTaskAttrDto), false)
}

/**
删除服务sql
*/
func (*MonitorTaskAttrDao) DeleteMonitorTaskAttr(monitorTaskAttrDto monitor.MonitorTaskAttrDto) error {
	return sqlTemplate.Delete(delete_monitorTaskAttr, objectConvert.Struct2Map(monitorTaskAttrDto), false)
}
