(function(vc){
    vc.extends({
        propTypes: {
           emitChooseWaf:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseWafInfo:{
                wafs:[],
                _currentWafName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseWaf','openChooseWafModel',function(_param){
                $('#chooseWafModel').modal('show');
                vc.component._refreshChooseWafInfo();
                vc.component._loadAllWafInfo(1,10,'');
            });
        },
        methods:{
            _loadAllWafInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('waf.listWafs',
                             param,
                             function(json){
                                var _wafInfo = JSON.parse(json);
                                vc.component.chooseWafInfo.wafs = _wafInfo.wafs;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseWaf:function(_waf){
                if(_waf.hasOwnProperty('name')){
                     _waf.wafName = _waf.name;
                }
                vc.emit($props.emitChooseWaf,'chooseWaf',_waf);
                vc.emit($props.emitLoadData,'listWafData',{
                    wafId:_waf.wafId
                });
                $('#chooseWafModel').modal('hide');
            },
            queryWafs:function(){
                vc.component._loadAllWafInfo(1,10,vc.component.chooseWafInfo._currentWafName);
            },
            _refreshChooseWafInfo:function(){
                vc.component.chooseWafInfo._currentWafName = "";
            }
        }

    });
})(window.vc);
