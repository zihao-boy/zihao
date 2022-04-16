/**
    组网访问 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewInnerNetPrivilegeInfo:{
                index:0,
                flowComponent:'viewInnerNetPrivilegeInfo',
                srcUserId:'',
targetUserId:'',
targetPort:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadInnerNetPrivilegeInfoData();
        },
        _initEvent:function(){
            vc.on('viewInnerNetPrivilegeInfo','chooseInnerNetPrivilege',function(_app){
                vc.copyObject(_app, vc.component.viewInnerNetPrivilegeInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewInnerNetPrivilegeInfo);
            });

            vc.on('viewInnerNetPrivilegeInfo', 'onIndex', function(_index){
                vc.component.viewInnerNetPrivilegeInfo.index = _index;
            });

        },
        methods:{

            _openSelectInnerNetPrivilegeInfoModel(){
                vc.emit('chooseInnerNetPrivilege','openChooseInnerNetPrivilegeModel',{});
            },
            _openAddInnerNetPrivilegeInfoModel(){
                vc.emit('addInnerNetPrivilege','openAddInnerNetPrivilegeModal',{});
            },
            _loadInnerNetPrivilegeInfoData:function(){

            }
        }
    });

})(window.vc);
