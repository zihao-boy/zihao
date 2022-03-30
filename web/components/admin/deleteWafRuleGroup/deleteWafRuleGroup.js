(function(vc,vm){

    vc.extends({
        data:{
            deleteWafRuleGroupInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteWafRuleGroup','openDeleteWafRuleGroupModal',function(_params){

                vc.component.deleteWafRuleGroupInfo = _params;
                $('#deleteWafRuleGroupModel').modal('show');

            });
        },
        methods:{
            deleteWafRuleGroup:function(){
                vc.component.deleteWafRuleGroupInfo.communityId=vc.getCurrentCommunity().communityId;
                vc.http.apiPost(
                    'wafRuleGroup.deleteWafRuleGroup',
                    JSON.stringify(vc.component.deleteWafRuleGroupInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteWafRuleGroupModel').modal('hide');
                            vc.emit('wafRuleGroupManage','listWafRuleGroup',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteWafRuleGroupModel:function(){
                $('#deleteWafRuleGroupModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
