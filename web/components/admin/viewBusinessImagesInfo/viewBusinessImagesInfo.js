/**
    镜像 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewBusinessImagesInfo:{
                index:0,
                flowComponent:'viewBusinessImagesInfo',
                name:'',
imagesType:'',
typeUrl:'',
imagesFlag:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadBusinessImagesInfoData();
        },
        _initEvent:function(){
            vc.on('viewBusinessImagesInfo','chooseBusinessImages',function(_app){
                vc.copyObject(_app, vc.component.viewBusinessImagesInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewBusinessImagesInfo);
            });

            vc.on('viewBusinessImagesInfo', 'onIndex', function(_index){
                vc.component.viewBusinessImagesInfo.index = _index;
            });

        },
        methods:{

            _openSelectBusinessImagesInfoModel(){
                vc.emit('chooseBusinessImages','openChooseBusinessImagesModel',{});
            },
            _openAddBusinessImagesInfoModel(){
                vc.emit('addBusinessImages','openAddBusinessImagesModal',{});
            },
            _loadBusinessImagesInfoData:function(){

            }
        }
    });

})(window.vc);
