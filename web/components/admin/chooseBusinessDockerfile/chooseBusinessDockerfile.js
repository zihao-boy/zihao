(function(vc){
    vc.extends({
        propTypes: {
           emitChooseBusinessDockerfile:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseBusinessDockerfileInfo:{
                businessDockerfiles:[],
                _currentBusinessDockerfileName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseBusinessDockerfile','openChooseBusinessDockerfileModel',function(_param){
                $('#chooseBusinessDockerfileModel').modal('show');
                vc.component._refreshChooseBusinessDockerfileInfo();
                vc.component._loadAllBusinessDockerfileInfo(1,10,'');
            });
        },
        methods:{
            _loadAllBusinessDockerfileInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('businessDockerfile.listBusinessDockerfiles',
                             param,
                             function(json){
                                var _businessDockerfileInfo = JSON.parse(json);
                                vc.component.chooseBusinessDockerfileInfo.businessDockerfiles = _businessDockerfileInfo.businessDockerfiles;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseBusinessDockerfile:function(_businessDockerfile){
                if(_businessDockerfile.hasOwnProperty('name')){
                     _businessDockerfile.businessDockerfileName = _businessDockerfile.name;
                }
                vc.emit($props.emitChooseBusinessDockerfile,'chooseBusinessDockerfile',_businessDockerfile);
                vc.emit($props.emitLoadData,'listBusinessDockerfileData',{
                    businessDockerfileId:_businessDockerfile.businessDockerfileId
                });
                $('#chooseBusinessDockerfileModel').modal('hide');
            },
            queryBusinessDockerfiles:function(){
                vc.component._loadAllBusinessDockerfileInfo(1,10,vc.component.chooseBusinessDockerfileInfo._currentBusinessDockerfileName);
            },
            _refreshChooseBusinessDockerfileInfo:function(){
                vc.component.chooseBusinessDockerfileInfo._currentBusinessDockerfileName = "";
            }
        }

    });
})(window.vc);
