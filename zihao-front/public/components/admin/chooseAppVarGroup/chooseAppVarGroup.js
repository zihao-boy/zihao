(function(vc){
    vc.extends({
        propTypes: {
           emitChooseAppVarGroup:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseAppVarGroupInfo:{
                appVarGroups:[],
                _currentAppVarGroupName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseAppVarGroup','openChooseAppVarGroupModel',function(_param){
                $('#chooseAppVarGroupModel').modal('show');
                vc.component._refreshChooseAppVarGroupInfo();
                vc.component._loadAllAppVarGroupInfo(1,10,'');
            });
        },
        methods:{
            _loadAllAppVarGroupInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('appVarGroup.listAppVarGroups',
                             param,
                             function(json){
                                var _appVarGroupInfo = JSON.parse(json);
                                vc.component.chooseAppVarGroupInfo.appVarGroups = _appVarGroupInfo.appVarGroups;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseAppVarGroup:function(_appVarGroup){
                if(_appVarGroup.hasOwnProperty('name')){
                     _appVarGroup.appVarGroupName = _appVarGroup.name;
                }
                vc.emit($props.emitChooseAppVarGroup,'chooseAppVarGroup',_appVarGroup);
                vc.emit($props.emitLoadData,'listAppVarGroupData',{
                    appVarGroupId:_appVarGroup.appVarGroupId
                });
                $('#chooseAppVarGroupModel').modal('hide');
            },
            queryAppVarGroups:function(){
                vc.component._loadAllAppVarGroupInfo(1,10,vc.component.chooseAppVarGroupInfo._currentAppVarGroupName);
            },
            _refreshChooseAppVarGroupInfo:function(){
                vc.component.chooseAppVarGroupInfo._currentAppVarGroupName = "";
            }
        }

    });
})(window.vc);
