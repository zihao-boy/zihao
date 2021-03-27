/**
    监控主机 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewMonitorHostInfo:{
                index:0,
                flowComponent:'viewMonitorHostInfo',
                hostId:'',
monDisk:'',
cpuThreshold:'',
memThreshold:'',
diskThreshold:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadMonitorHostInfoData();
        },
        _initEvent:function(){
            vc.on('viewMonitorHostInfo','chooseMonitorHost',function(_app){
                vc.copyObject(_app, vc.component.viewMonitorHostInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewMonitorHostInfo);
            });

            vc.on('viewMonitorHostInfo', 'onIndex', function(_index){
                vc.component.viewMonitorHostInfo.index = _index;
            });

        },
        methods:{

            _openSelectMonitorHostInfoModel(){
                vc.emit('chooseMonitorHost','openChooseMonitorHostModel',{});
            },
            _openAddMonitorHostInfoModel(){
                vc.emit('addMonitorHost','openAddMonitorHostModal',{});
            },
            _loadMonitorHostInfoData:function(){

            }
        }
    });

})(window.vc);
