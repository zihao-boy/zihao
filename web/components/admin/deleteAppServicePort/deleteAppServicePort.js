(function(vc,vm){

    vc.extends({
        data:{
            deleteAppServicePortInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteAppServicePort','openDeleteAppServicePortModal',function(_params){

                vc.component.deleteAppServicePortInfo = _params;
                $('#deleteAppServicePortModel').modal('show');

            });
        },
        methods:{
            deleteAppServicePort:function(){
                vc.http.apiPost(
                    '/appService/deleteAppServicePort',
                    JSON.stringify(vc.component.deleteAppServicePortInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteAppServicePortModel').modal('hide');
                            vc.emit('appServiceControlPort', 'switch', vc.component.deleteAppServicePortInfo);
                            return ;
                        }
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteAppServicePortModel:function(){
                $('#deleteAppServicePortModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
