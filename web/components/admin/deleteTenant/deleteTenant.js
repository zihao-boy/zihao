(function(vc,vm){

    vc.extends({
        data:{
            deleteTenantInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteTenant','openDeleteTenantModal',function(_params){

                vc.component.deleteTenantInfo = _params;
                $('#deleteTenantModel').modal('show');

            });
        },
        methods:{
            deleteTenant:function(){
                vc.http.apiPost(
                    '/tenant/deleteTenant',
                    JSON.stringify(vc.component.deleteTenantInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteTenantModel').modal('hide');
                            vc.emit('tenantManage','listTenant',{});
                            return ;
                        }
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.toast(json);

                     });
            },
            closeDeleteTenantModel:function(){
                $('#deleteTenantModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
