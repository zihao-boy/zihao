(function(vc){
    vc.extends({
        propTypes: {
           emitChooseMonitorHostGroup:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseMonitorHostGroupInfo:{
                monitorHostGroups:[],
                _currentMonitorHostGroupName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseMonitorHostGroup','openChooseMonitorHostGroupModel',function(_param){
                $('#chooseMonitorHostGroupModel').modal('show');
                vc.component._refreshChooseMonitorHostGroupInfo();
                vc.component._loadAllMonitorHostGroupInfo(1,10,'');
            });
        },
        methods:{
            _loadAllMonitorHostGroupInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('monitorHostGroup.listMonitorHostGroups',
                             param,
                             function(json){
                                var _monitorHostGroupInfo = JSON.parse(json);
                                vc.component.chooseMonitorHostGroupInfo.monitorHostGroups = _monitorHostGroupInfo.monitorHostGroups;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseMonitorHostGroup:function(_monitorHostGroup){
                if(_monitorHostGroup.hasOwnProperty('name')){
                     _monitorHostGroup.monitorHostGroupName = _monitorHostGroup.name;
                }
                vc.emit($props.emitChooseMonitorHostGroup,'chooseMonitorHostGroup',_monitorHostGroup);
                vc.emit($props.emitLoadData,'listMonitorHostGroupData',{
                    monitorHostGroupId:_monitorHostGroup.monitorHostGroupId
                });
                $('#chooseMonitorHostGroupModel').modal('hide');
            },
            queryMonitorHostGroups:function(){
                vc.component._loadAllMonitorHostGroupInfo(1,10,vc.component.chooseMonitorHostGroupInfo._currentMonitorHostGroupName);
            },
            _refreshChooseMonitorHostGroupInfo:function(){
                vc.component.chooseMonitorHostGroupInfo._currentMonitorHostGroupName = "";
            }
        }

    });
})(window.vc);
