(function(vc){
    vc.extends({
        propTypes: {
           emitChooseWafIpBlackWhite:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseWafIpBlackWhiteInfo:{
                wafIpBlackWhites:[],
                _currentWafIpBlackWhiteName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseWafIpBlackWhite','openChooseWafIpBlackWhiteModel',function(_param){
                $('#chooseWafIpBlackWhiteModel').modal('show');
                vc.component._refreshChooseWafIpBlackWhiteInfo();
                vc.component._loadAllWafIpBlackWhiteInfo(1,10,'');
            });
        },
        methods:{
            _loadAllWafIpBlackWhiteInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('wafIpBlackWhite.listWafIpBlackWhites',
                             param,
                             function(json){
                                var _wafIpBlackWhiteInfo = JSON.parse(json);
                                vc.component.chooseWafIpBlackWhiteInfo.wafIpBlackWhites = _wafIpBlackWhiteInfo.wafIpBlackWhites;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseWafIpBlackWhite:function(_wafIpBlackWhite){
                if(_wafIpBlackWhite.hasOwnProperty('name')){
                     _wafIpBlackWhite.wafIpBlackWhiteName = _wafIpBlackWhite.name;
                }
                vc.emit($props.emitChooseWafIpBlackWhite,'chooseWafIpBlackWhite',_wafIpBlackWhite);
                vc.emit($props.emitLoadData,'listWafIpBlackWhiteData',{
                    wafIpBlackWhiteId:_wafIpBlackWhite.wafIpBlackWhiteId
                });
                $('#chooseWafIpBlackWhiteModel').modal('hide');
            },
            queryWafIpBlackWhites:function(){
                vc.component._loadAllWafIpBlackWhiteInfo(1,10,vc.component.chooseWafIpBlackWhiteInfo._currentWafIpBlackWhiteName);
            },
            _refreshChooseWafIpBlackWhiteInfo:function(){
                vc.component.chooseWafIpBlackWhiteInfo._currentWafIpBlackWhiteName = "";
            }
        }

    });
})(window.vc);
