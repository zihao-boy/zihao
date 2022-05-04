(function(vc,vm){

    vc.extends({
        data:{
            deleteHostFirewallGroupInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteHostFirewallGroup','openDeleteHostFirewallGroupModal',function(_params){

                vc.component.deleteHostFirewallGroupInfo = _params;
                $('#deleteHostFirewallGroupModel').modal('show');

            });
        },
        methods:{
            deleteHostFirewallGroup:function(){
                vc.http.apiPost(
                    '/firewall/deleteHostFirewallGroup',
                    JSON.stringify(vc.component.deleteHostFirewallGroupInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteHostFirewallGroupModel').modal('hide');
                            vc.emit('hostFirewallGroup', 'switch', vc.component.deleteHostFirewallGroupInfo);
                            return ;
                        }
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteHostFirewallGroupModel:function(){
                $('#deleteHostFirewallGroupModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
