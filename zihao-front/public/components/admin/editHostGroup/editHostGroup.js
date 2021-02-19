(function(vc,vm){

    vc.extends({
        data:{
            editHostGroupInfo:{
                groupId:'',
name:'',
description:'',

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('editHostGroup','openEditHostGroupModal',function(_params){
                vc.component.refreshEditHostGroupInfo();
                $('#editHostGroupModel').modal('show');
                vc.copyObject(_params, vc.component.editHostGroupInfo );
                vc.component.editHostGroupInfo.communityId = vc.getCurrentCommunity().communityId;
            });
        },
        methods:{
            editHostGroupValidate:function(){
                        return vc.validate.validate({
                            editHostGroupInfo:vc.component.editHostGroupInfo
                        },{
                            'editHostGroupInfo.name':[
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
'editHostGroupInfo.description':[
 {
                            limit:"maxLength",
                            param:"512",
                            errInfo:"描述太长"
                        },
                    ],
'editHostGroupInfo.groupId':[
{
                            limit:"required",
                            param:"",
                            errInfo:"组ID不能为空"
                        }]

                        });
             },
            editHostGroup:function(){
                if(!vc.component.editHostGroupValidate()){
                    vc.toast(vc.validate.errInfo);
                    return ;
                }

                vc.http.apiPost(
                    'hostGroup.updateHostGroup',
                    JSON.stringify(vc.component.editHostGroupInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editHostGroupModel').modal('hide');
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
            refreshEditHostGroupInfo:function(){
                vc.component.editHostGroupInfo= {
                  groupId:'',
name:'',
description:'',

                }
            }
        }
    });

})(window.vc,window.vc.component);
