(function(vc){

    vc.extends({
        propTypes: {
               callBackListener:vc.propTypes.string, //父组件名称
               callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            addFirewallRuleGroupInfo:{
                groupId:'',
                groupName:'',

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
            vc.on('addFirewallRuleGroup','openAddFirewallRuleGroupModal',function(){
                $('#addFirewallRuleGroupModel').modal('show');
            });
        },
        methods:{
            addFirewallRuleGroupValidate(){
                return vc.validate.validate({
                    addFirewallRuleGroupInfo:vc.component.addFirewallRuleGroupInfo
                },{
                    'addFirewallRuleGroupInfo.groupName':[
{
                            limit:"required",
                            param:"",
                            errInfo:"名称不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"名称不能超过64"
                        },
                    ],




                });
            },
            saveFirewallRuleGroupInfo:function(){
                if(!vc.component.addFirewallRuleGroupValidate()){
                    vc.toast(vc.validate.errInfo);

                    return ;
                }

                vc.component.addFirewallRuleGroupInfo.communityId = vc.getCurrentCommunity().communityId;
                //不提交数据将数据 回调给侦听处理
                if(vc.notNull($props.callBackListener)){
                    vc.emit($props.callBackListener,$props.callBackFunction,vc.component.addFirewallRuleGroupInfo);
                    $('#addFirewallRuleGroupModel').modal('hide');
                    return ;
                }

                vc.http.apiPost(
                    'firewallRuleGroup.saveFirewallRuleGroup',
                    JSON.stringify(vc.component.addFirewallRuleGroupInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addFirewallRuleGroupModel').modal('hide');
                            vc.component.clearAddFirewallRuleGroupInfo();
                            vc.emit('firewallRuleGroupManage','listFirewallRuleGroup',{});

                            return ;
                        }
                        vc.message(_json.msg);

                     },
                     function(errInfo,error){
                        console.log('请求失败处理');

                        vc.message(errInfo);

                     });
            },
            clearAddFirewallRuleGroupInfo:function(){
                vc.component.addFirewallRuleGroupInfo = {
                                            groupName:'',

                                        };
            }
        }
    });

})(window.vc);
