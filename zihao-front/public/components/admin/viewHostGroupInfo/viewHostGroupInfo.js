/**
    主机组 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewHostGroupInfo:{
                index:0,
                flowComponent:'viewHostGroupInfo',
                name:'',
description:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadHostGroupInfoData();
        },
        _initEvent:function(){
            vc.on('viewHostGroupInfo','chooseHostGroup',function(_app){
                vc.copyObject(_app, vc.component.viewHostGroupInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewHostGroupInfo);
            });

            vc.on('viewHostGroupInfo', 'onIndex', function(_index){
                vc.component.viewHostGroupInfo.index = _index;
            });

        },
        methods:{

            _openSelectHostGroupInfoModel(){
                vc.emit('chooseHostGroup','openChooseHostGroupModel',{});
            },
            _openAddHostGroupInfoModel(){
                vc.emit('addHostGroup','openAddHostGroupModal',{});
            },
            _loadHostGroupInfoData:function(){

            }
        }
    });

})(window.vc);
