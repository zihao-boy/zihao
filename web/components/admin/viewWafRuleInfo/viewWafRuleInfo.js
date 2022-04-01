/**
    规则 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewWafRuleInfo:{
                index:0,
                flowComponent:'viewWafRuleInfo',
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
            //根据请求参数查询 查询 业主信息
            vc.component._loadWafRuleInfoData();
        },
        _initEvent:function(){
            vc.on('viewWafRuleInfo','chooseWafRule',function(_app){
                vc.copyObject(_app, vc.component.viewWafRuleInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewWafRuleInfo);
            });

            vc.on('viewWafRuleInfo', 'onIndex', function(_index){
                vc.component.viewWafRuleInfo.index = _index;
            });

        },
        methods:{

            _openSelectWafRuleInfoModel(){
                vc.emit('chooseWafRule','openChooseWafRuleModel',{});
            },
            _openAddWafRuleInfoModel(){
                vc.emit('addWafRule','openAddWafRuleModal',{});
            },
            _loadWafRuleInfoData:function(){

            }
        }
    });

})(window.vc);
