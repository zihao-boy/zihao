(function(vc){
    vc.extends({
        propTypes: {
           emitChooseHostGroup:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseHostGroupInfo:{
                hostGroups:[],
                _currentHostGroupName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseHostGroup','openChooseHostGroupModel',function(_param){
                $('#chooseHostGroupModel').modal('show');
                vc.component._refreshChooseHostGroupInfo();
                vc.component._loadAllHostGroupInfo(1,10,'');
            });
        },
        methods:{
            _loadAllHostGroupInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('hostGroup.listHostGroups',
                             param,
                             function(json){
                                var _hostGroupInfo = JSON.parse(json);
                                vc.component.chooseHostGroupInfo.hostGroups = _hostGroupInfo.hostGroups;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseHostGroup:function(_hostGroup){
                if(_hostGroup.hasOwnProperty('name')){
                     _hostGroup.hostGroupName = _hostGroup.name;
                }
                vc.emit($props.emitChooseHostGroup,'chooseHostGroup',_hostGroup);
                vc.emit($props.emitLoadData,'listHostGroupData',{
                    hostGroupId:_hostGroup.hostGroupId
                });
                $('#chooseHostGroupModel').modal('hide');
            },
            queryHostGroups:function(){
                vc.component._loadAllHostGroupInfo(1,10,vc.component.chooseHostGroupInfo._currentHostGroupName);
            },
            _refreshChooseHostGroupInfo:function(){
                vc.component.chooseHostGroupInfo._currentHostGroupName = "";
            }
        }

    });
})(window.vc);
