(function(vc,vm){

    vc.extends({
        data:{
            deleteAppServiceInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteAppService','openDeleteAppServiceModal',function(_params){

                vc.component.deleteAppServiceInfo = _params;
                $('#deleteAppServiceModel').modal('show');

            });
        },
        methods:{
            deleteAppService:function(){
                vc.http.apiPost(
                    '/appService/deleteAppService',
                    JSON.stringify(vc.component.deleteAppServiceInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteAppServiceModel').modal('hide');
                            vc.emit('appServiceManage','listAppService',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteAppServiceModel:function(){
                $('#deleteAppServiceModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
