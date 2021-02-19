package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/menu"
	"github.com/zihao-boy/zihao/zihao-service/entity/vo"
	"github.com/zihao-boy/zihao/zihao-service/system/dao"
)



type MenuService struct {
	menuDao dao.MenuDao
}

/**
查询 系统信息
 */
func (menuService *MenuService) GetMenus(ctx iris.Context)  []map[string]interface{} {
	var (
		err       error
		userVo = vo.LoginUserVo{UserId: ctx.Values().GetString(constants.UID)}
		menusDto []*menu.MenusDto
	)

	menusDto,err = menuService.menuDao.GetMenu(userVo)
	if(err != nil){
		return nil
	}

	tempMenus := refreshMenusInfo(menusDto)
	return tempMenus

}
/**
 * 刷新菜单信息
 * 将 数据 [{
 * "gId": "800201904001",
 * "menuDescription": "添加员工",
 * "menuGroupSeq": 1,
 * "menuSeq": 1,
 * "icon": "fa-desktop",
 * "mId": "700201904001",
 * "menuName": "添加员工",
 * "pId": "500201904001",
 * "menuGroupName": "员工管理",
 * "label": "",
 * "menuGroupDescription": "员工管理",
 * "url": "/"
 * }],
 * 转为：
 * "[{'id':1,'icon':'fa-desktop','name':'我的菜单','label':'HOT','childs':[" +
 * "{'name':'子菜单','href':'http://www.baidu.com'}]}," +
 * "{'id':2,'icon':'fa-flask','name':'我的菜单','childs':[],'href':'/doc'}," +
 * "{'id':3,'icon':'fa-globe','name':'我的菜单','childs':[{'name':'子菜单','href':'http://www.baidu.com'}]}" +
 * "]";
 *
 */
func refreshMenusInfo(menusDto []*menu.MenusDto) []map[string]interface{}{
	tempMenus := make([]map[string]interface{},0)
	//tempMenus.PushBack()
	var tempMenu map[string]interface{}
	for  menuIndex := 0; menuIndex < len(menusDto); menuIndex++ {
		var tMenu = menusDto[menuIndex]
		tempMenu = getMenuFromMenus(tempMenus, tMenu.GId)
		if tempMenu == nil {
			childs := make([]map[string]interface{},0)
			tempMenu = make(map[string]interface{})
			tempMenu["id"] = tMenu.GId
			tempMenu["icon"] = tMenu.Icon
			tempMenu["name"] = tMenu.MenuGroupName
			tempMenu["label"] = tMenu.Label
			tempMenu["seq"] = tMenu.MenuGroupSeq
			tempMenu["childs"] = childs
			tempMenus = append(tempMenus,tempMenu)
		}
		//获取孩子菜单
		childs := tempMenu["childs"].([]map[string]interface{})
		var childMenu = make(map[string]interface{})
		childMenu["name"] =  tMenu.MenuName
		childMenu["href" ] =  tMenu.Url
		childMenu["seq"] =  tMenu.MenuSeq
		childMenu["isShow"] =  tMenu.IsShow
		childs = append(childs,childMenu)

		tempMenu["childs"] = childs

	}



	return tempMenus
}

/**
 * 在菜单列表查询菜单
 *
 * @param gId
 * @return
 */
func getMenuFromMenus( tempMenus []map[string]interface{}, gId string) map[string]interface{} {
	for index := 0;index< len(tempMenus);index++ {
		tempMenu := tempMenus[index]
		if tempMenu["id"] == gId{
			return tempMenu
		}
	}
	return nil;
}
