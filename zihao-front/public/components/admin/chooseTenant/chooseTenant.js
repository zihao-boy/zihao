(function(vc){
    vc.extends({
        propTypes: {
           emitChooseTenant:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseTenantInfo:{
                tenants:[],
                _currentTenantName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseTenant','openChooseTenantModel',function(_param){
                $('#chooseTenantModel').modal('show');
                vc.component._refreshChooseTenantInfo();
                vc.component._loadAllTenantInfo(1,10,'');
            });
        },
        methods:{
            _loadAllTenantInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('tenant.listTenants',
                             param,
                             function(json){
                                var _tenantInfo = JSON.parse(json);
                                vc.component.chooseTenantInfo.tenants = _tenantInfo.tenants;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseTenant:function(_tenant){
                if(_tenant.hasOwnProperty('name')){
                     _tenant.tenantName = _tenant.name;
                }
                vc.emit($props.emitChooseTenant,'chooseTenant',_tenant);
                vc.emit($props.emitLoadData,'listTenantData',{
                    tenantId:_tenant.tenantId
                });
                $('#chooseTenantModel').modal('hide');
            },
            queryTenants:function(){
                vc.component._loadAllTenantInfo(1,10,vc.component.chooseTenantInfo._currentTenantName);
            },
            _refreshChooseTenantInfo:function(){
                vc.component.chooseTenantInfo._currentTenantName = "";
            }
        }

    });
})(window.vc);
