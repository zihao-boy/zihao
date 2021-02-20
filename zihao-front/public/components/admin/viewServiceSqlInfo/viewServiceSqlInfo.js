/**
    能力信息 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewServiceSqlInfo:{
                index:0,
                flowComponent:'viewServiceSqlInfo',
                sqlCode:'',
remark:'',
sqlText:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadServiceSqlInfoData();
        },
        _initEvent:function(){
            vc.on('viewServiceSqlInfo','chooseServiceSql',function(_app){
                vc.copyObject(_app, vc.component.viewServiceSqlInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewServiceSqlInfo);
            });

            vc.on('viewServiceSqlInfo', 'onIndex', function(_index){
                vc.component.viewServiceSqlInfo.index = _index;
            });

        },
        methods:{

            _openSelectServiceSqlInfoModel(){
                vc.emit('chooseServiceSql','openChooseServiceSqlModel',{});
            },
            _openAddServiceSqlInfoModel(){
                vc.emit('addServiceSql','openAddServiceSqlModal',{});
            },
            _loadServiceSqlInfoData:function(){

            }
        }
    });

})(window.vc);
