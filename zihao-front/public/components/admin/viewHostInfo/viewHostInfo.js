/**
    主机 组件
**/
(function(vc){

    vc.extends({
        propTypes: {
           callBackListener:vc.propTypes.string, //父组件名称
           callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            viewHostInfo:{
                index:0,
                flowComponent:'viewHostInfo',
                groupId:'',
name:'',
ip:'',
username:'',
passwd:'',
cpu:'',
mem:'',
disk:'',

            }
        },
        _initMethod:function(){
            //根据请求参数查询 查询 业主信息
            vc.component._loadHostInfoData();
        },
        _initEvent:function(){
            vc.on('viewHostInfo','chooseHost',function(_app){
                vc.copyObject(_app, vc.component.viewHostInfo);
                vc.emit($props.callBackListener,$props.callBackFunction,vc.component.viewHostInfo);
            });

            vc.on('viewHostInfo', 'onIndex', function(_index){
                vc.component.viewHostInfo.index = _index;
            });

        },
        methods:{

            _openSelectHostInfoModel(){
                vc.emit('chooseHost','openChooseHostModel',{});
            },
            _openAddHostInfoModel(){
                vc.emit('addHost','openAddHostModal',{});
            },
            _loadHostInfoData:function(){

            }
        }
    });

})(window.vc);
