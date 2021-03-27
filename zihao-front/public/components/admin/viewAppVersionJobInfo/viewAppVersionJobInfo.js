/**
    版本构建 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewAppVersionJobInfo:{
                index:0,
                flowComponent:'viewAppVersionJobInfo',
                jobName:'',
jobShell:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadAppVersionJobInfoData();
        },
        _initEvent:function(){
            vc.on('viewAppVersionJobInfo','chooseAppVersionJob',function(_app){
                vc.copyObject(_app, vc.component.viewAppVersionJobInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewAppVersionJobInfo);
            });

            vc.on('viewAppVersionJobInfo', 'onIndex', function(_index){
                vc.component.viewAppVersionJobInfo.index = _index;
            });

        },
        methods:{

            _openSelectAppVersionJobInfoModel(){
                vc.emit('chooseAppVersionJob','openChooseAppVersionJobModel',{});
            },
            _openAddAppVersionJobInfoModel(){
                vc.emit('addAppVersionJob','openAddAppVersionJobModal',{});
            },
            _loadAppVersionJobInfoData:function(){

            }
        }
    });

})(window.vc);
