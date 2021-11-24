/**
    租户设置 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewTenantSettingInfo:{
                index:0,
                flowComponent:'viewTenantSettingInfo',
                specCd:'',
value:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadTenantSettingInfoData();
        },
        _initEvent:function(){
            vc.on('viewTenantSettingInfo','chooseTenantSetting',function(_app){
                vc.copyObject(_app, vc.component.viewTenantSettingInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewTenantSettingInfo);
            });

            vc.on('viewTenantSettingInfo', 'onIndex', function(_index){
                vc.component.viewTenantSettingInfo.index = _index;
            });

        },
        methods:{

            _openSelectTenantSettingInfoModel(){
                vc.emit('chooseTenantSetting','openChooseTenantSettingModel',{});
            },
            _openAddTenantSettingInfoModel(){
                vc.emit('addTenantSetting','openAddTenantSettingModal',{});
            },
            _loadTenantSettingInfoData:function(){

            }
        }
    });

})(window.vc);
