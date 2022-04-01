/**
    安全组 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewWafRuleGroupInfo:{
                index:0,
                flowComponent:'viewWafRuleGroupInfo',
                groupName:'',
state:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadWafRuleGroupInfoData();
        },
        _initEvent:function(){
            vc.on('viewWafRuleGroupInfo','chooseWafRuleGroup',function(_app){
                vc.copyObject(_app, vc.component.viewWafRuleGroupInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewWafRuleGroupInfo);
            });

            vc.on('viewWafRuleGroupInfo', 'onIndex', function(_index){
                vc.component.viewWafRuleGroupInfo.index = _index;
            });

        },
        methods:{

            _openSelectWafRuleGroupInfoModel(){
                vc.emit('chooseWafRuleGroup','openChooseWafRuleGroupModel',{});
            },
            _openAddWafRuleGroupInfoModel(){
                vc.emit('addWafRuleGroup','openAddWafRuleGroupModal',{});
            },
            _loadWafRuleGroupInfoData:function(){

            }
        }
    });

})(window.vc);
