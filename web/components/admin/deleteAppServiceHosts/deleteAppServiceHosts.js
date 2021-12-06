(function(vc,vm){

    vc.extends({
        data:{
            deleteAppServiceHostsInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteAppServiceHosts','openDeleteAppServiceHostsModal',function(_params){

                vc.component.deleteAppServiceHostsInfo = _params;
                $('#deleteAppServiceHostsModel').modal('show');

            });
        },
        methods:{
            deleteAppServiceHosts:function(){
                vc.http.apiPost(
                    '/appService/deleteAppServiceHosts',
                    JSON.stringify(vc.component.deleteAppServiceHostsInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteAppServiceHostsModel').modal('hide');
                            vc.emit('appServiceControlHosts', 'switch', vc.component.deleteAppServiceHostsInfo);
                            return ;
                        }
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteAppServiceHostsModel:function(){
                $('#deleteAppServiceHostsModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
