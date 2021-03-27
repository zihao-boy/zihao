(function(vc){
    vc.extends({
        propTypes: {
           emitChooseMonitorHost:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseMonitorHostInfo:{
                monitorHosts:[],
                _currentMonitorHostName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseMonitorHost','openChooseMonitorHostModel',function(_param){
                $('#chooseMonitorHostModel').modal('show');
                vc.component._refreshChooseMonitorHostInfo();
                vc.component._loadAllMonitorHostInfo(1,10,'');
            });
        },
        methods:{
            _loadAllMonitorHostInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('monitorHost.listMonitorHosts',
                             param,
                             function(json){
                                var _monitorHostInfo = JSON.parse(json);
                                vc.component.chooseMonitorHostInfo.monitorHosts = _monitorHostInfo.monitorHosts;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseMonitorHost:function(_monitorHost){
                if(_monitorHost.hasOwnProperty('name')){
                     _monitorHost.monitorHostName = _monitorHost.name;
                }
                vc.emit($props.emitChooseMonitorHost,'chooseMonitorHost',_monitorHost);
                vc.emit($props.emitLoadData,'listMonitorHostData',{
                    monitorHostId:_monitorHost.monitorHostId
                });
                $('#chooseMonitorHostModel').modal('hide');
            },
            queryMonitorHosts:function(){
                vc.component._loadAllMonitorHostInfo(1,10,vc.component.chooseMonitorHostInfo._currentMonitorHostName);
            },
            _refreshChooseMonitorHostInfo:function(){
                vc.component.chooseMonitorHostInfo._currentMonitorHostName = "";
            }
        }

    });
})(window.vc);
