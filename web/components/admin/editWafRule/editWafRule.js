(function(vc,vm){

    vc.extends({
        data:{
            editWafRuleInfo:{
                ruleId:'',
groupId:'',
ruleName:'',
scope:'',
objId:'',
objType:'',
seq:'',
state:'',

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('editWafRule','openEditWafRuleModal',function(_params){
                vc.component.refreshEditWafRuleInfo();
                $('#editWafRuleModel').modal('show');
                vc.copyObject(_params, vc.component.editWafRuleInfo );
                vc.component.editWafRuleInfo.communityId = vc.getCurrentCommunity().communityId;
            });
        },
        methods:{
            editWafRuleValidate:function(){
                        return vc.validate.validate({
                            editWafRuleInfo:vc.component.editWafRuleInfo
                        },{
                            'editWafRuleInfo.groupId':[
{
                            limit:"required",
                            param:"",
                            errInfo:"规则组不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"规则组不能超过64"
                        },
                    ],
'editWafRuleInfo.ruleName':[
{
                            limit:"required",
                            param:"",
                            errInfo:"规则名称不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"规则名称不能超过64"
                        },
                    ],
'editWafRuleInfo.scope':[
{
                            limit:"required",
                            param:"",
                            errInfo:"规则范围不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"规则范围不能超过64"
                        },
                    ],
'editWafRuleInfo.objId':[
{
                            limit:"required",
                            param:"",
                            errInfo:"规则对象不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"规则对象不能超过64"
                        },
                    ],
'editWafRuleInfo.objType':[
{
                            limit:"required",
                            param:"",
                            errInfo:"规则类型不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"规则类型不能超过64"
                        },
                    ],
'editWafRuleInfo.seq':[
{
                            limit:"required",
                            param:"",
                            errInfo:"优先级不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"12",
                            errInfo:"优先级不能超过12"
                        },
                    ],
'editWafRuleInfo.state':[
{
                            limit:"required",
                            param:"",
                            errInfo:"状态不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"状态不能超过64"
                        },
                    ],
'editWafRuleInfo.ruleId':[
{
                            limit:"required",
                            param:"",
                            errInfo:"编号不能为空"
                        }]

                        });
             },
            editWafRule:function(){
                if(!vc.component.editWafRuleValidate()){
                    vc.toast(vc.validate.errInfo);
                    return ;
                }

                vc.http.apiPost(
                    'wafRule.updateWafRule',
                    JSON.stringify(vc.component.editWafRuleInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editWafRuleModel').modal('hide');
                             vc.emit('wafRuleManage','listWafRule',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');

                        vc.message(errInfo);
                     });
            },
            refreshEditWafRuleInfo:function(){
                vc.component.editWafRuleInfo= {
                  ruleId:'',
groupId:'',
ruleName:'',
scope:'',
objId:'',
objType:'',
seq:'',
state:'',

                }
            }
        }
    });

})(window.vc,window.vc.component);
