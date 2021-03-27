(function(vc){
    vc.extends({
        propTypes: {
           emitChooseAppVar:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseAppVarInfo:{
                appVars:[],
                _currentAppVarName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseAppVar','openChooseAppVarModel',function(_param){
                $('#chooseAppVarModel').modal('show');
                vc.component._refreshChooseAppVarInfo();
                vc.component._loadAllAppVarInfo(1,10,'');
            });
        },
        methods:{
            _loadAllAppVarInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('appVar.listAppVars',
                             param,
                             function(json){
                                var _appVarInfo = JSON.parse(json);
                                vc.component.chooseAppVarInfo.appVars = _appVarInfo.appVars;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseAppVar:function(_appVar){
                if(_appVar.hasOwnProperty('name')){
                     _appVar.appVarName = _appVar.name;
                }
                vc.emit($props.emitChooseAppVar,'chooseAppVar',_appVar);
                vc.emit($props.emitLoadData,'listAppVarData',{
                    appVarId:_appVar.appVarId
                });
                $('#chooseAppVarModel').modal('hide');
            },
            queryAppVars:function(){
                vc.component._loadAllAppVarInfo(1,10,vc.component.chooseAppVarInfo._currentAppVarName);
            },
            _refreshChooseAppVarInfo:function(){
                vc.component.chooseAppVarInfo._currentAppVarName = "";
            }
        }

    });
})(window.vc);
