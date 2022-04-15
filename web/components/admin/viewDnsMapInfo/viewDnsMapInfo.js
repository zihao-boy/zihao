/**
    hosts 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewDnsMapInfo:{
                index:0,
                flowComponent:'viewDnsMapInfo',
                host:'',
type:'',
value:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadDnsMapInfoData();
        },
        _initEvent:function(){
            vc.on('viewDnsMapInfo','chooseDnsMap',function(_app){
                vc.copyObject(_app, vc.component.viewDnsMapInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewDnsMapInfo);
            });

            vc.on('viewDnsMapInfo', 'onIndex', function(_index){
                vc.component.viewDnsMapInfo.index = _index;
            });

        },
        methods:{

            _openSelectDnsMapInfoModel(){
                vc.emit('chooseDnsMap','openChooseDnsMapModel',{});
            },
            _openAddDnsMapInfoModel(){
                vc.emit('addDnsMap','openAddDnsMapModal',{});
            },
            _loadDnsMapInfoData:function(){

            }
        }
    });

})(window.vc);
