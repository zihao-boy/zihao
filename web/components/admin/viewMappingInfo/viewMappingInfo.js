/**
    编码映射 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewMappingInfo:{
                index:0,
                flowComponent:'viewMappingInfo',
                domain:'',
name:'',
keys:'',
value:'',
remark:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadMappingInfoData();
        },
        _initEvent:function(){
            vc.on('viewMappingInfo','chooseMapping',function(_app){
                vc.copyObject(_app, vc.component.viewMappingInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewMappingInfo);
            });

            vc.on('viewMappingInfo', 'onIndex', function(_index){
                vc.component.viewMappingInfo.index = _index;
            });

        },
        methods:{

            _openSelectMappingInfoModel(){
                vc.emit('chooseMapping','openChooseMappingModel',{});
            },
            _openAddMappingInfoModel(){
                vc.emit('addMapping','openAddMappingModal',{});
            },
            _loadMappingInfoData:function(){

            }
        }
    });

})(window.vc);
