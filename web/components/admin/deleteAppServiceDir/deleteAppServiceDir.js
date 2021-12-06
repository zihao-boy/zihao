(function(vc,vm){

    vc.extends({
        data:{
            deleteAppServiceDirInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteAppServiceDir','openDeleteAppServiceDirModal',function(_params){

                vc.component.deleteAppServiceDirInfo = _params;
                $('#deleteAppServiceDirModel').modal('show');

            });
        },
        methods:{
            deleteAppServiceDir:function(){
                vc.http.apiPost(
                    '/appService/deleteAppServiceDir',
                    JSON.stringify(vc.component.deleteAppServiceDirInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteAppServiceDirModel').modal('hide');
                            vc.emit('appServiceControlDir', 'switch', vc.component.deleteAppServiceDirInfo);
                            return ;
                        }
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteAppServiceDirModel:function(){
                $('#deleteAppServiceDirModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
