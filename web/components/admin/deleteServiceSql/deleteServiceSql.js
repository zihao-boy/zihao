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
                vc.http.apiPost(
                    '/system/deleteServiceSql',
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
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.toast(json);

                     });
            },
            closeDeleteServiceSqlModel:function(){
                $('#deleteServiceSqlModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
