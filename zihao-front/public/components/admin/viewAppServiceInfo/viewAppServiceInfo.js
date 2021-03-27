/**
    服务信息 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewAppServiceInfo:{
                index:0,
                flowComponent:'viewAppServiceInfo',
                asName:'',
asType:'',
asDesc:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadAppServiceInfoData();
        },
        _initEvent:function(){
            vc.on('viewAppServiceInfo','chooseAppService',function(_app){
                vc.copyObject(_app, vc.component.viewAppServiceInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewAppServiceInfo);
            });

            vc.on('viewAppServiceInfo', 'onIndex', function(_index){
                vc.component.viewAppServiceInfo.index = _index;
            });

        },
        methods:{

            _openSelectAppServiceInfoModel(){
                vc.emit('chooseAppService','openChooseAppServiceModel',{});
            },
            _openAddAppServiceInfoModel(){
                vc.emit('addAppService','openAddAppServiceModal',{});
            },
            _loadAppServiceInfoData:function(){

            }
        }
    });

})(window.vc);
