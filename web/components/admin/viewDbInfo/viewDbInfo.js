/**
    db 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewDbInfo:{
                index:0,
                flowComponent:'viewDbInfo',
                name:'',
ip:'',
port:'',
username:'',
password:'',
dbName:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadDbInfoData();
        },
        _initEvent:function(){
            vc.on('viewDbInfo','chooseDb',function(_app){
                vc.copyObject(_app, vc.component.viewDbInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewDbInfo);
            });

            vc.on('viewDbInfo', 'onIndex', function(_index){
                vc.component.viewDbInfo.index = _index;
            });

        },
        methods:{

            _openSelectDbInfoModel(){
                vc.emit('chooseDb','openChooseDbModel',{});
            },
            _openAddDbInfoModel(){
                vc.emit('addDb','openAddDbModal',{});
            },
            _loadDbInfoData:function(){

            }
        }
    });

})(window.vc);
