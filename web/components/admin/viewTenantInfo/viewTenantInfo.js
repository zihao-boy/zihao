/**
    租户 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewTenantInfo:{
                index:0,
                flowComponent:'viewTenantInfo',
                tenantName:'',
tenantType:'',
address:'',
personName:'',
phone:'',
username:'',
passwd:'',
remark:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadTenantInfoData();
        },
        _initEvent:function(){
            vc.on('viewTenantInfo','chooseTenant',function(_app){
                vc.copyObject(_app, vc.component.viewTenantInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewTenantInfo);
            });

            vc.on('viewTenantInfo', 'onIndex', function(_index){
                vc.component.viewTenantInfo.index = _index;
            });

        },
        methods:{

            _openSelectTenantInfoModel(){
                vc.emit('chooseTenant','openChooseTenantModel',{});
            },
            _openAddTenantInfoModel(){
                vc.emit('addTenant','openAddTenantModal',{});
            },
            _loadTenantInfoData:function(){

            }
        }
    });

})(window.vc);
