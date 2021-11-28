(function(vc){
    vc.extends({
        propTypes: {
           emitChooseTenantSetting:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseTenantSettingInfo:{
                tenantSettings:[],
                _currentTenantSettingName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseTenantSetting','openChooseTenantSettingModel',function(_param){
                $('#chooseTenantSettingModel').modal('show');
                vc.component._refreshChooseTenantSettingInfo();
                vc.component._loadAllTenantSettingInfo(1,10,'');
            });
        },
        methods:{
            _loadAllTenantSettingInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('tenantSetting.listTenantSettings',
                             param,
                             function(json){
                                var _tenantSettingInfo = JSON.parse(json);
                                vc.component.chooseTenantSettingInfo.tenantSettings = _tenantSettingInfo.tenantSettings;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseTenantSetting:function(_tenantSetting){
                if(_tenantSetting.hasOwnProperty('name')){
                     _tenantSetting.tenantSettingName = _tenantSetting.name;
                }
                vc.emit($props.emitChooseTenantSetting,'chooseTenantSetting',_tenantSetting);
                vc.emit($props.emitLoadData,'listTenantSettingData',{
                    tenantSettingId:_tenantSetting.tenantSettingId
                });
                $('#chooseTenantSettingModel').modal('hide');
            },
            queryTenantSettings:function(){
                vc.component._loadAllTenantSettingInfo(1,10,vc.component.chooseTenantSettingInfo._currentTenantSettingName);
            },
            _refreshChooseTenantSettingInfo:function(){
                vc.component.chooseTenantSettingInfo._currentTenantSettingName = "";
            }
        }

    });
})(window.vc);
