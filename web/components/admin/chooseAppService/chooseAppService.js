(function(vc){
    vc.extends({
        propTypes: {
           emitChooseAppService:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseAppServiceInfo:{
                appServices:[],
                _currentAppServiceName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseAppService','openChooseAppServiceModel',function(_param){
                $('#chooseAppServiceModel').modal('show');
                vc.component._refreshChooseAppServiceInfo();
                vc.component._loadAllAppServiceInfo(1,10,'');
            });
        },
        methods:{
            _loadAllAppServiceInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('appService.listAppServices',
                             param,
                             function(json){
                                var _appServiceInfo = JSON.parse(json);
                                vc.component.chooseAppServiceInfo.appServices = _appServiceInfo.appServices;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseAppService:function(_appService){
                if(_appService.hasOwnProperty('name')){
                     _appService.appServiceName = _appService.name;
                }
                vc.emit($props.emitChooseAppService,'chooseAppService',_appService);
                vc.emit($props.emitLoadData,'listAppServiceData',{
                    appServiceId:_appService.appServiceId
                });
                $('#chooseAppServiceModel').modal('hide');
            },
            queryAppServices:function(){
                vc.component._loadAllAppServiceInfo(1,10,vc.component.chooseAppServiceInfo._currentAppServiceName);
            },
            _refreshChooseAppServiceInfo:function(){
                vc.component.chooseAppServiceInfo._currentAppServiceName = "";
            }
        }

    });
})(window.vc);
