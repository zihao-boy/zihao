(function(vc,vm){

    vc.extends({
        data:{
            deleteServiceSqlInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteServiceSql','openDeleteServiceSqlModal',function(_params){

                vc.component.deleteServiceSqlInfo = _params;
                $('#deleteServiceSqlModel').modal('show');

            });
        },
        methods:{
            deleteServiceSql:function(){
                vc.component.deleteServiceSqlInfo.communityId=vc.getCurrentCommunity().communityId;
                vc.http.apiPost(
                    'serviceSql.deleteServiceSql',
                    JSON.stringify(vc.component.deleteServiceSqlInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteServiceSqlModel').modal('hide');
                            vc.emit('serviceSqlManage','listServiceSql',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteServiceSqlModel:function(){
                $('#deleteServiceSqlModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
