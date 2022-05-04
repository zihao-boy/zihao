/**
    端口 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewFirewallRuleInfo:{
                index:0,
                flowComponent:'viewFirewallRuleInfo',
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
            //根据请求参数查询 查询 业主信息
            vc.component._loadFirewallRuleInfoData();
        },
        _initEvent:function(){
            vc.on('viewFirewallRuleInfo','chooseFirewallRule',function(_app){
                vc.copyObject(_app, vc.component.viewFirewallRuleInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewFirewallRuleInfo);
            });

            vc.on('viewFirewallRuleInfo', 'onIndex', function(_index){
                vc.component.viewFirewallRuleInfo.index = _index;
            });

        },
        methods:{

            _openSelectFirewallRuleInfoModel(){
                vc.emit('chooseFirewallRule','openChooseFirewallRuleModel',{});
            },
            _openAddFirewallRuleInfoModel(){
                vc.emit('addFirewallRule','openAddFirewallRuleModal',{});
            },
            _loadFirewallRuleInfoData:function(){

            }
        }
    });

})(window.vc);
