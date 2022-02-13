/**
    oss 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewOssInfo:{
                index:0,
                flowComponent:'viewOssInfo',
                name:'',
ossType:'',
bucket:'',
accessKeySecret:'',
accessKeyId:'',
endpoint:'',
path:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadOssInfoData();
        },
        _initEvent:function(){
            vc.on('viewOssInfo','chooseOss',function(_app){
                vc.copyObject(_app, vc.component.viewOssInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewOssInfo);
            });

            vc.on('viewOssInfo', 'onIndex', function(_index){
                vc.component.viewOssInfo.index = _index;
            });

        },
        methods:{

            _openSelectOssInfoModel(){
                vc.emit('chooseOss','openChooseOssModel',{});
            },
            _openAddOssInfoModel(){
                vc.emit('addOss','openAddOssModal',{});
            },
            _loadOssInfoData:function(){

            }
        }
    });

})(window.vc);
