/**
    开放端口 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewFirewallRuleGroupInfo:{
                index:0,
                flowComponent:'viewFirewallRuleGroupInfo',
                groupName:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadFirewallRuleGroupInfoData();
        },
        _initEvent:function(){
            vc.on('viewFirewallRuleGroupInfo','chooseFirewallRuleGroup',function(_app){
                vc.copyObject(_app, vc.component.viewFirewallRuleGroupInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewFirewallRuleGroupInfo);
            });

            vc.on('viewFirewallRuleGroupInfo', 'onIndex', function(_index){
                vc.component.viewFirewallRuleGroupInfo.index = _index;
            });

        },
        methods:{

            _openSelectFirewallRuleGroupInfoModel(){
                vc.emit('chooseFirewallRuleGroup','openChooseFirewallRuleGroupModel',{});
            },
            _openAddFirewallRuleGroupInfoModel(){
                vc.emit('addFirewallRuleGroup','openAddFirewallRuleGroupModal',{});
            },
            _loadFirewallRuleGroupInfoData:function(){

            }
        }
    });

})(window.vc);
