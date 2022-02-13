(function(vc){
    vc.extends({
        propTypes: {
           emitChooseDb:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseDbInfo:{
                dbs:[],
                _currentDbName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseDb','openChooseDbModel',function(_param){
                $('#chooseDbModel').modal('show');
                vc.component._refreshChooseDbInfo();
                vc.component._loadAllDbInfo(1,10,'');
            });
        },
        methods:{
            _loadAllDbInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('db.listDbs',
                             param,
                             function(json){
                                var _dbInfo = JSON.parse(json);
                                vc.component.chooseDbInfo.dbs = _dbInfo.dbs;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseDb:function(_db){
                if(_db.hasOwnProperty('name')){
                     _db.dbName = _db.name;
                }
                vc.emit($props.emitChooseDb,'chooseDb',_db);
                vc.emit($props.emitLoadData,'listDbData',{
                    dbId:_db.dbId
                });
                $('#chooseDbModel').modal('hide');
            },
            queryDbs:function(){
                vc.component._loadAllDbInfo(1,10,vc.component.chooseDbInfo._currentDbName);
            },
            _refreshChooseDbInfo:function(){
                vc.component.chooseDbInfo._currentDbName = "";
            }
        }

    });
})(window.vc);
