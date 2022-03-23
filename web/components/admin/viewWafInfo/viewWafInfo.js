/**
    网站防火墙（waf） 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewWafInfo:{
                index:0,
                flowComponent:'viewWafInfo',
                wafName:'',
port:'',
state:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadWafInfoData();
        },
        _initEvent:function(){
            vc.on('viewWafInfo','chooseWaf',function(_app){
                vc.copyObject(_app, vc.component.viewWafInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewWafInfo);
            });

            vc.on('viewWafInfo', 'onIndex', function(_index){
                vc.component.viewWafInfo.index = _index;
            });

        },
        methods:{

            _openSelectWafInfoModel(){
                vc.emit('chooseWaf','openChooseWafModel',{});
            },
            _openAddWafInfoModel(){
                vc.emit('addWaf','openAddWafModal',{});
            },
            _loadWafInfoData:function(){

            }
        }
    });

})(window.vc);
