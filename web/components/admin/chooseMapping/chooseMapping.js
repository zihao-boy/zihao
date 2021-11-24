(function(vc){
    vc.extends({
        propTypes: {
           emitChooseMapping:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseMappingInfo:{
                mappings:[],
                _currentMappingName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseMapping','openChooseMappingModel',function(_param){
                $('#chooseMappingModel').modal('show');
                vc.component._refreshChooseMappingInfo();
                vc.component._loadAllMappingInfo(1,10,'');
            });
        },
        methods:{
            _loadAllMappingInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('mapping.listMappings',
                             param,
                             function(json){
                                var _mappingInfo = JSON.parse(json);
                                vc.component.chooseMappingInfo.mappings = _mappingInfo.mappings;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseMapping:function(_mapping){
                if(_mapping.hasOwnProperty('name')){
                     _mapping.mappingName = _mapping.name;
                }
                vc.emit($props.emitChooseMapping,'chooseMapping',_mapping);
                vc.emit($props.emitLoadData,'listMappingData',{
                    mappingId:_mapping.mappingId
                });
                $('#chooseMappingModel').modal('hide');
            },
            queryMappings:function(){
                vc.component._loadAllMappingInfo(1,10,vc.component.chooseMappingInfo._currentMappingName);
            },
            _refreshChooseMappingInfo:function(){
                vc.component.chooseMappingInfo._currentMappingName = "";
            }
        }

    });
})(window.vc);
