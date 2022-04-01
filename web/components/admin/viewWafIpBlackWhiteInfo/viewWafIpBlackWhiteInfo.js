/**
    黑白名单 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewWafIpBlackWhiteInfo:{
                index:0,
                flowComponent:'viewWafIpBlackWhiteInfo',
                typeCd:'',
ip:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadWafIpBlackWhiteInfoData();
        },
        _initEvent:function(){
            vc.on('viewWafIpBlackWhiteInfo','chooseWafIpBlackWhite',function(_app){
                vc.copyObject(_app, vc.component.viewWafIpBlackWhiteInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewWafIpBlackWhiteInfo);
            });

            vc.on('viewWafIpBlackWhiteInfo', 'onIndex', function(_index){
                vc.component.viewWafIpBlackWhiteInfo.index = _index;
            });

        },
        methods:{

            _openSelectWafIpBlackWhiteInfoModel(){
                vc.emit('chooseWafIpBlackWhite','openChooseWafIpBlackWhiteModel',{});
            },
            _openAddWafIpBlackWhiteInfoModel(){
                vc.emit('addWafIpBlackWhite','openAddWafIpBlackWhiteModal',{});
            },
            _loadWafIpBlackWhiteInfoData:function(){

            }
        }
    });

})(window.vc);
