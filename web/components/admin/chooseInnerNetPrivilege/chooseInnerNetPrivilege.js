(function(vc){
    vc.extends({
        propTypes: {
           emitChooseInnerNetPrivilege:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseInnerNetPrivilegeInfo:{
                innerNetPrivileges:[],
                _currentInnerNetPrivilegeName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseInnerNetPrivilege','openChooseInnerNetPrivilegeModel',function(_param){
                $('#chooseInnerNetPrivilegeModel').modal('show');
                vc.component._refreshChooseInnerNetPrivilegeInfo();
                vc.component._loadAllInnerNetPrivilegeInfo(1,10,'');
            });
        },
        methods:{
            _loadAllInnerNetPrivilegeInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('innerNetPrivilege.listInnerNetPrivileges',
                             param,
                             function(json){
                                var _innerNetPrivilegeInfo = JSON.parse(json);
                                vc.component.chooseInnerNetPrivilegeInfo.innerNetPrivileges = _innerNetPrivilegeInfo.innerNetPrivileges;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseInnerNetPrivilege:function(_innerNetPrivilege){
                if(_innerNetPrivilege.hasOwnProperty('name')){
                     _innerNetPrivilege.innerNetPrivilegeName = _innerNetPrivilege.name;
                }
                vc.emit($props.emitChooseInnerNetPrivilege,'chooseInnerNetPrivilege',_innerNetPrivilege);
                vc.emit($props.emitLoadData,'listInnerNetPrivilegeData',{
                    innerNetPrivilegeId:_innerNetPrivilege.innerNetPrivilegeId
                });
                $('#chooseInnerNetPrivilegeModel').modal('hide');
            },
            queryInnerNetPrivileges:function(){
                vc.component._loadAllInnerNetPrivilegeInfo(1,10,vc.component.chooseInnerNetPrivilegeInfo._currentInnerNetPrivilegeName);
            },
            _refreshChooseInnerNetPrivilegeInfo:function(){
                vc.component.chooseInnerNetPrivilegeInfo._currentInnerNetPrivilegeName = "";
            }
        }

    });
})(window.vc);
