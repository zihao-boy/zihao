(function(vc,vm){

    vc.extends({
        data:{
            deleteFirewallRuleInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteFirewallRule','openDeleteFirewallRuleModal',function(_params){

                vc.component.deleteFirewallRuleInfo = _params;
                $('#deleteFirewallRuleModel').modal('show');

            });
        },
        methods:{
            deleteFirewallRule:function(){
                vc.component.deleteFirewallRuleInfo.communityId=vc.getCurrentCommunity().communityId;
                vc.http.apiPost(
                    'firewallRule.deleteFirewallRule',
                    JSON.stringify(vc.component.deleteFirewallRuleInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteFirewallRuleModel').modal('hide');
                            vc.emit('firewallRuleManage','listFirewallRule',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteFirewallRuleModel:function(){
                $('#deleteFirewallRuleModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
