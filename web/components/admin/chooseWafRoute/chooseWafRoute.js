(function(vc){
    vc.extends({
        propTypes: {
           emitChooseWafRoute:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseWafRouteInfo:{
                wafRoutes:[],
                _currentWafRouteName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseWafRoute','openChooseWafRouteModel',function(_param){
                $('#chooseWafRouteModel').modal('show');
                vc.component._refreshChooseWafRouteInfo();
                vc.component._loadAllWafRouteInfo(1,10,'');
            });
        },
        methods:{
            _loadAllWafRouteInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('wafRoute.listWafRoutes',
                             param,
                             function(json){
                                var _wafRouteInfo = JSON.parse(json);
                                vc.component.chooseWafRouteInfo.wafRoutes = _wafRouteInfo.wafRoutes;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseWafRoute:function(_wafRoute){
                if(_wafRoute.hasOwnProperty('name')){
                     _wafRoute.wafRouteName = _wafRoute.name;
                }
                vc.emit($props.emitChooseWafRoute,'chooseWafRoute',_wafRoute);
                vc.emit($props.emitLoadData,'listWafRouteData',{
                    wafRouteId:_wafRoute.wafRouteId
                });
                $('#chooseWafRouteModel').modal('hide');
            },
            queryWafRoutes:function(){
                vc.component._loadAllWafRouteInfo(1,10,vc.component.chooseWafRouteInfo._currentWafRouteName);
            },
            _refreshChooseWafRouteInfo:function(){
                vc.component.chooseWafRouteInfo._currentWafRouteName = "";
            }
        }

    });
})(window.vc);
