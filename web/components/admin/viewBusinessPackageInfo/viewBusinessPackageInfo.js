/**
    业务包 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewBusinessPackageInfo:{
                index:0,
                flowComponent:'viewBusinessPackageInfo',
                name:'',
varsion:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadBusinessPackageInfoData();
        },
        _initEvent:function(){
            vc.on('viewBusinessPackageInfo','chooseBusinessPackage',function(_app){
                vc.copyObject(_app, vc.component.viewBusinessPackageInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewBusinessPackageInfo);
            });

            vc.on('viewBusinessPackageInfo', 'onIndex', function(_index){
                vc.component.viewBusinessPackageInfo.index = _index;
            });

        },
        methods:{

            _openSelectBusinessPackageInfoModel(){
                vc.emit('chooseBusinessPackage','openChooseBusinessPackageModel',{});
            },
            _openAddBusinessPackageInfoModel(){
                vc.emit('addBusinessPackage','openAddBusinessPackageModal',{});
            },
            _loadBusinessPackageInfoData:function(){

            }
        }
    });

})(window.vc);
