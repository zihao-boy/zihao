(function(vc){

    vc.extends({
        propTypes: {
               callBackListener:vc.propTypes.string, //父组件名称
               callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            addWafRuleInfo:{
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
            vc.on('addWafRule','openAddWafRuleModal',function(){
                $('#addWafRuleModel').modal('show');
            });
        },
        methods:{
            addWafRuleValidate(){
                return vc.validate.validate({
                    addWafRuleInfo:vc.component.addWafRuleInfo
                },{
                    'addWafRuleInfo.groupId':[
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
'addWafRuleInfo.ruleName':[
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
'addWafRuleInfo.scope':[
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
'addWafRuleInfo.objId':[
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
'addWafRuleInfo.objType':[
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
'addWafRuleInfo.seq':[
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
'addWafRuleInfo.state':[
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




                });
            },
            saveWafRuleInfo:function(){
                if(!vc.component.addWafRuleValidate()){
                    vc.toast(vc.validate.errInfo);

                    return ;
                }

                vc.component.addWafRuleInfo.communityId = vc.getCurrentCommunity().communityId;
                //不提交数据将数据 回调给侦听处理
                if(vc.notNull($props.callBackListener)){
                    vc.emit($props.callBackListener,$props.callBackFunction,vc.component.addWafRuleInfo);
                    $('#addWafRuleModel').modal('hide');
                    return ;
                }

                vc.http.apiPost(
                    'wafRule.saveWafRule',
                    JSON.stringify(vc.component.addWafRuleInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addWafRuleModel').modal('hide');
                            vc.component.clearAddWafRuleInfo();
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
            clearAddWafRuleInfo:function(){
                vc.component.addWafRuleInfo = {
                                            groupId:'',
ruleName:'',
scope:'',
objId:'',
objType:'',
seq:'',
state:'',

                                        };
            }
        }
    });

})(window.vc);
