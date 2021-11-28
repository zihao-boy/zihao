/**
    主机监控组 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewMonitorHostGroupInfo:{
                index:0,
                flowComponent:'viewMonitorHostGroupInfo',
                name:'',
monCron:'',
noticeType:'',
remark:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadMonitorHostGroupInfoData();
        },
        _initEvent:function(){
            vc.on('viewMonitorHostGroupInfo','chooseMonitorHostGroup',function(_app){
                vc.copyObject(_app, vc.component.viewMonitorHostGroupInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewMonitorHostGroupInfo);
            });

            vc.on('viewMonitorHostGroupInfo', 'onIndex', function(_index){
                vc.component.viewMonitorHostGroupInfo.index = _index;
            });

        },
        methods:{

            _openSelectMonitorHostGroupInfoModel(){
                vc.emit('chooseMonitorHostGroup','openChooseMonitorHostGroupModel',{});
            },
            _openAddMonitorHostGroupInfoModel(){
                vc.emit('addMonitorHostGroup','openAddMonitorHostGroupModal',{});
            },
            _loadMonitorHostGroupInfoData:function(){

            }
        }
    });

})(window.vc);
