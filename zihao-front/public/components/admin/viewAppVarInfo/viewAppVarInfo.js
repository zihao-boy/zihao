/**
    环境变量 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewAppVarInfo:{
                index:0,
                flowComponent:'viewAppVarInfo',
                varName:'',
avgId:'',
varType:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadAppVarInfoData();
        },
        _initEvent:function(){
            vc.on('viewAppVarInfo','chooseAppVar',function(_app){
                vc.copyObject(_app, vc.component.viewAppVarInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewAppVarInfo);
            });

            vc.on('viewAppVarInfo', 'onIndex', function(_index){
                vc.component.viewAppVarInfo.index = _index;
            });

        },
        methods:{

            _openSelectAppVarInfoModel(){
                vc.emit('chooseAppVar','openChooseAppVarModel',{});
            },
            _openAddAppVarInfoModel(){
                vc.emit('addAppVar','openAddAppVarModal',{});
            },
            _loadAppVarInfoData:function(){

            }
        }
    });

})(window.vc);
