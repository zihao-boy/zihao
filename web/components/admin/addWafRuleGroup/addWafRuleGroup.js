(function(vc){

    vc.extends({
        propTypes: {
               callBackListener:vc.propTypes.string, //父组件名称
               callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            addWafRuleGroupInfo:{
                groupId:'',
                groupName:'',
state:'',

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
            vc.on('addWafRuleGroup','openAddWafRuleGroupModal',function(){
                $('#addWafRuleGroupModel').modal('show');
            });
        },
        methods:{
            addWafRuleGroupValidate(){
                return vc.validate.validate({
                    addWafRuleGroupInfo:vc.component.addWafRuleGroupInfo
                },{
                    'addWafRuleGroupInfo.groupName':[
{
                            limit:"required",
                            param:"",
                            errInfo:"组名称不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"组名称不能超过64"
                        },
                    ],
'addWafRuleGroupInfo.state':[
{
                            limit:"required",
                            param:"",
                            errInfo:"组状态'不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"组状态'不能超过64"
                        },
                    ],




                });
            },
            saveWafRuleGroupInfo:function(){
                if(!vc.component.addWafRuleGroupValidate()){
                    vc.toast(vc.validate.errInfo);

                    return ;
                }

                vc.component.addWafRuleGroupInfo.communityId = vc.getCurrentCommunity().communityId;
                //不提交数据将数据 回调给侦听处理
                if(vc.notNull($props.callBackListener)){
                    vc.emit($props.callBackListener,$props.callBackFunction,vc.component.addWafRuleGroupInfo);
                    $('#addWafRuleGroupModel').modal('hide');
                    return ;
                }

                vc.http.apiPost(
                    'wafRuleGroup.saveWafRuleGroup',
                    JSON.stringify(vc.component.addWafRuleGroupInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addWafRuleGroupModel').modal('hide');
                            vc.component.clearAddWafRuleGroupInfo();
                            vc.emit('wafRuleGroupManage','listWafRuleGroup',{});

                            return ;
                        }
                        vc.message(_json.msg);

                     },
                     function(errInfo,error){
                        console.log('请求失败处理');

                        vc.message(errInfo);

                     });
            },
            clearAddWafRuleGroupInfo:function(){
                vc.component.addWafRuleGroupInfo = {
                                            groupName:'',
state:'',

                                        };
            }
        }
    });

})(window.vc);
