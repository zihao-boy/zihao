(function(vc){

    vc.extends({
        propTypes: {
               callBackListener:vc.propTypes.string, //父组件名称
               callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            addFirewallRuleInfo:{
                ruleId:'',
                groupId:'',
inOut:'',
allowLimit:'',
seq:'',
protocol:'',
srcObj:'',
dstObj:'',
remark:'',

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
            vc.on('addFirewallRule','openAddFirewallRuleModal',function(){
                $('#addFirewallRuleModel').modal('show');
            });
        },
        methods:{
            addFirewallRuleValidate(){
                return vc.validate.validate({
                    addFirewallRuleInfo:vc.component.addFirewallRuleInfo
                },{
                    'addFirewallRuleInfo.groupId':[
{
                            limit:"required",
                            param:"",
                            errInfo:"组编号不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"组编号不能超过64"
                        },
                    ],
'addFirewallRuleInfo.inOut':[
{
                            limit:"required",
                            param:"",
                            errInfo:"入站出站不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"入站出站不能超过64"
                        },
                    ],
'addFirewallRuleInfo.allowLimit':[
{
                            limit:"required",
                            param:"",
                            errInfo:"授权策略不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"授权策略不能超过64"
                        },
                    ],
'addFirewallRuleInfo.seq':[
{
                            limit:"required",
                            param:"",
                            errInfo:"顺序不能为空"
                        },
 {
                            limit:"num",
                            param:"11",
                            errInfo:"顺序不能超过11"
                        },
                    ],
'addFirewallRuleInfo.protocol':[
{
                            limit:"required",
                            param:"",
                            errInfo:"协议不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"协议不能超过64"
                        },
                    ],
'addFirewallRuleInfo.srcObj':[
{
                            limit:"required",
                            param:"",
                            errInfo:"授权对象不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"授权对象不能超过64"
                        },
                    ],
'addFirewallRuleInfo.dstObj':[
{
                            limit:"required",
                            param:"",
                            errInfo:"端口不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"端口不能超过64"
                        },
                    ],
'addFirewallRuleInfo.remark':[
{
                            limit:"required",
                            param:"",
                            errInfo:"备注不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"512",
                            errInfo:"备注不能超过512"
                        },
                    ],




                });
            },
            saveFirewallRuleInfo:function(){
                if(!vc.component.addFirewallRuleValidate()){
                    vc.toast(vc.validate.errInfo);

                    return ;
                }

                vc.component.addFirewallRuleInfo.communityId = vc.getCurrentCommunity().communityId;
                //不提交数据将数据 回调给侦听处理
                if(vc.notNull($props.callBackListener)){
                    vc.emit($props.callBackListener,$props.callBackFunction,vc.component.addFirewallRuleInfo);
                    $('#addFirewallRuleModel').modal('hide');
                    return ;
                }

                vc.http.apiPost(
                    'firewallRule.saveFirewallRule',
                    JSON.stringify(vc.component.addFirewallRuleInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addFirewallRuleModel').modal('hide');
                            vc.component.clearAddFirewallRuleInfo();
                            vc.emit('firewallRuleManage','listFirewallRule',{});

                            return ;
                        }
                        vc.message(_json.msg);

                     },
                     function(errInfo,error){
                        console.log('请求失败处理');

                        vc.message(errInfo);

                     });
            },
            clearAddFirewallRuleInfo:function(){
                vc.component.addFirewallRuleInfo = {
                                            groupId:'',
inOut:'',
allowLimit:'',
seq:'',
protocol:'',
srcObj:'',
dstObj:'',
remark:'',

                                        };
            }
        }
    });

})(window.vc);
