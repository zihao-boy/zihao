(function(vc,vm){

    vc.extends({
        data:{
            deleteVpnUserInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteVpnUser','openDeleteVpnUserModal',function(_params){

                vc.component.deleteVpnUserInfo = _params;
                $('#deleteVpnUserModel').modal('show');

            });
        },
        methods:{
            deleteVpnUser:function(){
                vc.http.apiPost(
                    '/vpn/deleteVpnUser',
                    JSON.stringify(vc.component.deleteVpnUserInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteVpnUserModel').modal('hide');
                            vc.emit('vpnUserManage','listVpnUser',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteVpnUserModel:function(){
                $('#deleteVpnUserModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
