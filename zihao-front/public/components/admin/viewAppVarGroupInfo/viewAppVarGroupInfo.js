/**
    环境组 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewAppVarGroupInfo:{
                index:0,
                flowComponent:'viewAppVarGroupInfo',
                avgName:'',
avgDesc:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadAppVarGroupInfoData();
        },
        _initEvent:function(){
            vc.on('viewAppVarGroupInfo','chooseAppVarGroup',function(_app){
                vc.copyObject(_app, vc.component.viewAppVarGroupInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewAppVarGroupInfo);
            });

            vc.on('viewAppVarGroupInfo', 'onIndex', function(_index){
                vc.component.viewAppVarGroupInfo.index = _index;
            });

        },
        methods:{

            _openSelectAppVarGroupInfoModel(){
                vc.emit('chooseAppVarGroup','openChooseAppVarGroupModel',{});
            },
            _openAddAppVarGroupInfoModel(){
                vc.emit('addAppVarGroup','openAddAppVarGroupModal',{});
            },
            _loadAppVarGroupInfoData:function(){

            }
        }
    });

})(window.vc);
