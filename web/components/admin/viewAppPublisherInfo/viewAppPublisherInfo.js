/**
    发布者 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewAppPublisherInfo:{
                index:0,
                flowComponent:'viewAppPublisherInfo',
                username:'',
email:'',
phone:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadAppPublisherInfoData();
        },
        _initEvent:function(){
            vc.on('viewAppPublisherInfo','chooseAppPublisher',function(_app){
                vc.copyObject(_app, vc.component.viewAppPublisherInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewAppPublisherInfo);
            });

            vc.on('viewAppPublisherInfo', 'onIndex', function(_index){
                vc.component.viewAppPublisherInfo.index = _index;
            });

        },
        methods:{

            _openSelectAppPublisherInfoModel(){
                vc.emit('chooseAppPublisher','openChooseAppPublisherModel',{});
            },
            _openAddAppPublisherInfoModel(){
                vc.emit('addAppPublisher','openAddAppPublisherModal',{});
            },
            _loadAppPublisherInfoData:function(){

            }
        }
    });

})(window.vc);
