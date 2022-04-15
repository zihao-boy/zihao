(function(vc){
    vc.extends({
        propTypes: {
           emitChooseDnsMap:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseDnsMapInfo:{
                dnsMaps:[],
                _currentDnsMapName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseDnsMap','openChooseDnsMapModel',function(_param){
                $('#chooseDnsMapModel').modal('show');
                vc.component._refreshChooseDnsMapInfo();
                vc.component._loadAllDnsMapInfo(1,10,'');
            });
        },
        methods:{
            _loadAllDnsMapInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('dnsMap.listDnsMaps',
                             param,
                             function(json){
                                var _dnsMapInfo = JSON.parse(json);
                                vc.component.chooseDnsMapInfo.dnsMaps = _dnsMapInfo.dnsMaps;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseDnsMap:function(_dnsMap){
                if(_dnsMap.hasOwnProperty('name')){
                     _dnsMap.dnsMapName = _dnsMap.name;
                }
                vc.emit($props.emitChooseDnsMap,'chooseDnsMap',_dnsMap);
                vc.emit($props.emitLoadData,'listDnsMapData',{
                    dnsMapId:_dnsMap.dnsMapId
                });
                $('#chooseDnsMapModel').modal('hide');
            },
            queryDnsMaps:function(){
                vc.component._loadAllDnsMapInfo(1,10,vc.component.chooseDnsMapInfo._currentDnsMapName);
            },
            _refreshChooseDnsMapInfo:function(){
                vc.component.chooseDnsMapInfo._currentDnsMapName = "";
            }
        }

    });
})(window.vc);
