(function(vc,vm){

    vc.extends({
        data:{
            deleteFirewallRuleGroupInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteFirewallRuleGroup','openDeleteFirewallRuleGroupModal',function(_params){
                vc.component.deleteFirewallRuleGroupInfo = _params;
                $('#deleteFirewallRuleGroupModel').modal('show');
            });
        },
        methods:{
            deleteFirewallRuleGroup:function(){
                vc.http.apiPost(
                    '/firewall/deleteFirewallRuleGroup',
                    JSON.stringify(vc.component.deleteFirewallRuleGroupInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteFirewallRuleGroupModel').modal('hide');
                            vc.emit('firewallRuleGroupManage','listFirewallRuleGroup',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteFirewallRuleGroupModel:function(){
                $('#deleteFirewallRuleGroupModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
