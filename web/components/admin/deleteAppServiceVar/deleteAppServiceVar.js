(function(vc,vm){

    vc.extends({
        data:{
            deleteAppServiceVarInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteAppServiceVar','openDeleteAppServiceVarModal',function(_params){

                vc.component.deleteAppServiceVarInfo = _params;
                $('#deleteAppServiceVarModel').modal('show');

            });
        },
        methods:{
            deleteAppServiceVar:function(){
                vc.http.apiPost(
                    '/appService/deleteAppServiceVar',
                    JSON.stringify(vc.component.deleteAppServiceVarInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteAppServiceVarModel').modal('hide');
                            vc.emit('appServiceControlVar', 'switch', vc.component.deleteAppServiceVarInfo);
                            return ;
                        }
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteAppServiceVarModel:function(){
                $('#deleteAppServiceVarModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
