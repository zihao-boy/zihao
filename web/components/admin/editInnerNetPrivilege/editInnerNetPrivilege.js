(function(vc,vm){

    vc.extends({
        data:{
            editInnerNetPrivilegeInfo:{
                pId:'',
srcUserId:'',
targetUserId:'',
targetPort:'',

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('editInnerNetPrivilege','openEditInnerNetPrivilegeModal',function(_params){
                vc.component.refreshEditInnerNetPrivilegeInfo();
                $('#editInnerNetPrivilegeModel').modal('show');
                vc.copyObject(_params, vc.component.editInnerNetPrivilegeInfo );
                vc.component.editInnerNetPrivilegeInfo.communityId = vc.getCurrentCommunity().communityId;
            });
        },
        methods:{
            editInnerNetPrivilegeValidate:function(){
                        return vc.validate.validate({
                            editInnerNetPrivilegeInfo:vc.component.editInnerNetPrivilegeInfo
                        },{
                            'editInnerNetPrivilegeInfo.srcUserId':[
{
                            limit:"required",
                            param:"",
                            errInfo:"用户不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"用户不能超过64"
                        },
                    ],
'editInnerNetPrivilegeInfo.targetUserId':[
{
                            limit:"required",
                            param:"",
                            errInfo:"主机不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"主机不能超过64"
                        },
                    ],
'editInnerNetPrivilegeInfo.targetPort':[
{
                            limit:"required",
                            param:"",
                            errInfo:"端口'不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"端口'不能超过64"
                        },
                    ],
'editInnerNetPrivilegeInfo.pId':[
{
                            limit:"required",
                            param:"",
                            errInfo:"编号不能为空"
                        }]

                        });
             },
            editInnerNetPrivilege:function(){
                if(!vc.component.editInnerNetPrivilegeValidate()){
                    vc.toast(vc.validate.errInfo);
                    return ;
                }

                vc.http.apiPost(
                    'innerNetPrivilege.updateInnerNetPrivilege',
                    JSON.stringify(vc.component.editInnerNetPrivilegeInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editInnerNetPrivilegeModel').modal('hide');
                             vc.emit('innerNetPrivilegeManage','listInnerNetPrivilege',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');

                        vc.message(errInfo);
                     });
            },
            refreshEditInnerNetPrivilegeInfo:function(){
                vc.component.editInnerNetPrivilegeInfo= {
                  pId:'',
srcUserId:'',
targetUserId:'',
targetPort:'',

                }
            }
        }
    });

})(window.vc,window.vc.component);
