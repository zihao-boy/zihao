/**
    mydockerfile 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewBusinessDockerfileInfo:{
                index:0,
                flowComponent:'viewBusinessDockerfileInfo',
                name:'',
dockerfile:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadBusinessDockerfileInfoData();
        },
        _initEvent:function(){
            vc.on('viewBusinessDockerfileInfo','chooseBusinessDockerfile',function(_app){
                vc.copyObject(_app, vc.component.viewBusinessDockerfileInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewBusinessDockerfileInfo);
            });

            vc.on('viewBusinessDockerfileInfo', 'onIndex', function(_index){
                vc.component.viewBusinessDockerfileInfo.index = _index;
            });

        },
        methods:{

            _openSelectBusinessDockerfileInfoModel(){
                vc.emit('chooseBusinessDockerfile','openChooseBusinessDockerfileModel',{});
            },
            _openAddBusinessDockerfileInfoModel(){
                vc.emit('addBusinessDockerfile','openAddBusinessDockerfileModal',{});
            },
            _loadBusinessDockerfileInfoData:function(){

            }
        }
    });

})(window.vc);
