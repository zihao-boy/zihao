package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/menu"
	"github.com/zihao-boy/zihao/zihao-service/entity/vo"
)

const(
	query_menu string = `
		SELECT mm.m_id ,mm.name menu_name,mm.g_id,mm.url,mm.seq menu_seq,
			mm.description menu_description, mmg.name menu_group_name,mmg.icon,mmg.label,mmg.seq menu_group_seq,
			mmg.description menu_group_description,mm.is_show,mm.description
		FROM menu mm,menu_group mmg
		WHERE mm.g_id = mmg.g_id
		  and mmg.group_type = 'P_WEB'
		  AND mm.m_id IN (
			SELECT pp.m_id FROM privilege_user ppu,privilege pp
			WHERE ppu.p_id = pp.p_id
				AND ppu.privilege_flag = '0'
				AND ppu.user_id = ?
				AND ppu.status_cd = '0'
				AND pp.status_cd = '0'
				UNION
				SELECT pp.m_id FROM privilege_user ppu,privilege_group ppg,privilege pp,privilege_rel ppr
				WHERE ppu.p_id = ppr.pg_id
				AND ppr.pg_id = ppg.pg_id
				AND ppr.p_id = pp.p_id
				AND ppu.privilege_flag = '1'
				AND ppu.user_id = ?
				AND ppu.status_cd = '0'
				AND pp.status_cd = '0'
				AND ppg.status_cd = '0'
				AND ppr.status_cd = '0'
			)
			AND mm.status_cd = '0'
			AND mmg.status_cd = '0'
	`
)

type MenuDao struct {

}

/**
查询用户
*/
func (*MenuDao) GetMenu(userVo vo.LoginUserVo) ([]*menu.MenusDto,error){
	var menusDto []*menu.MenusDto
	db := mysql.G_DB.Raw(query_menu,userVo.UserId,userVo.UserId)
	if err:=db.Scan(&menusDto).Error; err !=nil{
		return nil,err
	}

	return menusDto,nil
}
