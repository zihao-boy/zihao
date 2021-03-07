/**
    入驻小区
**/
(function(vc){
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data:{
            tenantSettingManageInfo:{
                tenantSettings:[],
                total:0,
                records:1,
                moreCondition:false,
                settingId:'',
                conditions:{
                    specCd:'',

                }
            }
        },
        _initMethod:function(){
            vc.component._listTenantSettings(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent:function(){
            
            vc.on('tenantSettingManage','listTenantSetting',function(_param){
                  vc.component._listTenantSettings(DEFAULT_PAGE, DEFAULT_ROWS);
            });
             vc.on('pagination','page_event',function(_currentPage){
                vc.component._listTenantSettings(_currentPage,DEFAULT_ROWS);
            });
        },
        methods:{
            _listTenantSettings:function(_page, _rows){

                vc.component.tenantSettingManageInfo.conditions.page = _page;
                vc.component.tenantSettingManageInfo.conditions.row = _rows;
                var param = {
                    params:vc.component.tenantSettingManageInfo.conditions
               };

               //发送get请求
               vc.http.apiGet('/tenant/getTenantSettings',
                             param,
                             function(json,res){
                                var _tenantSettingManageInfo=JSON.parse(json);
                                vc.component.tenantSettingManageInfo.total = _tenantSettingManageInfo.total;
                                vc.component.tenantSettingManageInfo.records = _tenantSettingManageInfo.records;
                                vc.component.tenantSettingManageInfo.tenantSettings = _tenantSettingManageInfo.data;
                                vc.emit('pagination','init',{
                                     total:vc.component.tenantSettingManageInfo.records,
                                     currentPage:_page
                                 });
                             },function(errInfo,error){
                                console.log('请求失败处理');
                             }
                           );
            },
            _openAddTenantSettingModal:function(){
                vc.emit('addTenantSetting','openAddTenantSettingModal',{});
            },
            _openEditTenantSettingModel:function(_tenantSetting){
                vc.emit('editTenantSetting','openEditTenantSettingModal',_tenantSetting);
            },
            _openDeleteTenantSettingModel:function(_tenantSetting){
                vc.emit('deleteTenantSetting','openDeleteTenantSettingModal',_tenantSetting);
            },
            _queryTenantSettingMethod:function(){
                vc.component._listTenantSettings(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition:function(){
                if(vc.component.tenantSettingManageInfo.moreCondition){
                    vc.component.tenantSettingManageInfo.moreCondition = false;
                }else{
                    vc.component.tenantSettingManageInfo.moreCondition = true;
                }
            }

             
        }
    });
})(window.vc);
