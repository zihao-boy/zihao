/**
    资源备份 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewBackupInfo:{
                index:0,
                flowComponent:'viewBackupInfo',
                name:'',
execTime:'',
typeCd:'',
srcId:'',
srcObject:'',
targetId:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadBackupInfoData();
        },
        _initEvent:function(){
            vc.on('viewBackupInfo','chooseBackup',function(_app){
                vc.copyObject(_app, vc.component.viewBackupInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewBackupInfo);
            });

            vc.on('viewBackupInfo', 'onIndex', function(_index){
                vc.component.viewBackupInfo.index = _index;
            });

        },
        methods:{

            _openSelectBackupInfoModel(){
                vc.emit('chooseBackup','openChooseBackupModel',{});
            },
            _openAddBackupInfoModel(){
                vc.emit('addBackup','openAddBackupModal',{});
            },
            _loadBackupInfoData:function(){

            }
        }
    });

})(window.vc);
