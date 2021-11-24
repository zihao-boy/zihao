(function(vc){
    vc.extends({
        propTypes: {
           emitChooseServiceSql:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseServiceSqlInfo:{
                serviceSqls:[],
                _currentServiceSqlName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseServiceSql','openChooseServiceSqlModel',function(_param){
                $('#chooseServiceSqlModel').modal('show');
                vc.component._refreshChooseServiceSqlInfo();
                vc.component._loadAllServiceSqlInfo(1,10,'');
            });
        },
        methods:{
            _loadAllServiceSqlInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('serviceSql.listServiceSqls',
                             param,
                             function(json){
                                var _serviceSqlInfo = JSON.parse(json);
                                vc.component.chooseServiceSqlInfo.serviceSqls = _serviceSqlInfo.serviceSqls;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseServiceSql:function(_serviceSql){
                if(_serviceSql.hasOwnProperty('name')){
                     _serviceSql.serviceSqlName = _serviceSql.name;
                }
                vc.emit($props.emitChooseServiceSql,'chooseServiceSql',_serviceSql);
                vc.emit($props.emitLoadData,'listServiceSqlData',{
                    serviceSqlId:_serviceSql.serviceSqlId
                });
                $('#chooseServiceSqlModel').modal('hide');
            },
            queryServiceSqls:function(){
                vc.component._loadAllServiceSqlInfo(1,10,vc.component.chooseServiceSqlInfo._currentServiceSqlName);
            },
            _refreshChooseServiceSqlInfo:function(){
                vc.component.chooseServiceSqlInfo._currentServiceSqlName = "";
            }
        }

    });
})(window.vc);
