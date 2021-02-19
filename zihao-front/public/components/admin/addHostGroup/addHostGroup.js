(function(vc){

    vc.extends({
        propTypes: {
               callBackListener:vc.propTypes.string, //父组件名称
               callBackFunction:vc.propTypes.string //父组件监听方法
        },
        data:{
            addHostGroupInfo:{
                groupId:'',
                name:'',
description:'',

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
            vc.on('addHostGroup','openAddHostGroupModal',function(){
                $('#addHostGroupModel').modal('show');
            });
        },
        methods:{
            addHostGroupValidate(){
                return vc.validate.validate({
                    addHostGroupInfo:vc.component.addHostGroupInfo
                },{
                    'addHostGroupInfo.name':[
{
                            limit:"required",
                            param:"",
                            errInfo:"组名称不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"组名称格式错误"
                        },
                    ],
'addHostGroupInfo.description':[
 {
                            limit:"maxLength",
                            param:"512",
                            errInfo:"描述太长"
                        },
                    ],




                });
            },
            saveHostGroupInfo:function(){
                if(!vc.component.addHostGroupValidate()){
                    vc.toast(vc.validate.errInfo);

                    return ;
                }

                vc.component.addHostGroupInfo.communityId = vc.getCurrentCommunity().communityId;
                //不提交数据将数据 回调给侦听处理
                if(vc.notNull($props.callBackListener)){
                    vc.emit($props.callBackListener,$props.callBackFunction,vc.component.addHostGroupInfo);
                    $('#addHostGroupModel').modal('hide');
                    return ;
                }

                vc.http.apiPost(
                    'hostGroup.saveHostGroup',
                    JSON.stringify(vc.component.addHostGroupInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addHostGroupModel').modal('hide');
                            vc.component.clearAddHostGroupInfo();
                            vc.emit('hostGroupManage','listHostGroup',{});

                            return ;
                        }
                        vc.message(_json.msg);

                     },
                     function(errInfo,error){
                        console.log('请求失败处理');

                        vc.message(errInfo);

                     });
            },
            clearAddHostGroupInfo:function(){
                vc.component.addHostGroupInfo = {
                                            name:'',
description:'',

                                        };
            }
        }
    });

})(window.vc);
