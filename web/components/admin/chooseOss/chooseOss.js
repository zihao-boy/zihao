(function(vc){
    vc.extends({
        propTypes: {
           emitChooseOss:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseOssInfo:{
                osss:[],
                _currentOssName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseOss','openChooseOssModel',function(_param){
                $('#chooseOssModel').modal('show');
                vc.component._refreshChooseOssInfo();
                vc.component._loadAllOssInfo(1,10,'');
            });
        },
        methods:{
            _loadAllOssInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('oss.listOsss',
                             param,
                             function(json){
                                var _ossInfo = JSON.parse(json);
                                vc.component.chooseOssInfo.osss = _ossInfo.osss;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseOss:function(_oss){
                if(_oss.hasOwnProperty('name')){
                     _oss.ossName = _oss.name;
                }
                vc.emit($props.emitChooseOss,'chooseOss',_oss);
                vc.emit($props.emitLoadData,'listOssData',{
                    ossId:_oss.ossId
                });
                $('#chooseOssModel').modal('hide');
            },
            queryOsss:function(){
                vc.component._loadAllOssInfo(1,10,vc.component.chooseOssInfo._currentOssName);
            },
            _refreshChooseOssInfo:function(){
                vc.component.chooseOssInfo._currentOssName = "";
            }
        }

    });
})(window.vc);
