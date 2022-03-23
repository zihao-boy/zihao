/**
    waf路由 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewWafRouteInfo:{
                index:0,
                flowComponent:'viewWafRouteInfo',
                wafId:'',
hostname:'',
ip:'',
port:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadWafRouteInfoData();
        },
        _initEvent:function(){
            vc.on('viewWafRouteInfo','chooseWafRoute',function(_app){
                vc.copyObject(_app, vc.component.viewWafRouteInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewWafRouteInfo);
            });

            vc.on('viewWafRouteInfo', 'onIndex', function(_index){
                vc.component.viewWafRouteInfo.index = _index;
            });

        },
        methods:{

            _openSelectWafRouteInfoModel(){
                vc.emit('chooseWafRoute','openChooseWafRouteModel',{});
            },
            _openAddWafRouteInfoModel(){
                vc.emit('addWafRoute','openAddWafRouteModal',{});
            },
            _loadWafRouteInfoData:function(){

            }
        }
    });

})(window.vc);
