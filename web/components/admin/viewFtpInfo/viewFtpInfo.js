/**
    ftp 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewFtpInfo:{
                index:0,
                flowComponent:'viewFtpInfo',
                name:'',
ip:'',
port:'',
username:'',
passwd:'',
path:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadFtpInfoData();
        },
        _initEvent:function(){
            vc.on('viewFtpInfo','chooseFtp',function(_app){
                vc.copyObject(_app, vc.component.viewFtpInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewFtpInfo);
            });

            vc.on('viewFtpInfo', 'onIndex', function(_index){
                vc.component.viewFtpInfo.index = _index;
            });

        },
        methods:{

            _openSelectFtpInfoModel(){
                vc.emit('chooseFtp','openChooseFtpModel',{});
            },
            _openAddFtpInfoModel(){
                vc.emit('addFtp','openAddFtpModal',{});
            },
            _loadFtpInfoData:function(){

            }
        }
    });

})(window.vc);
