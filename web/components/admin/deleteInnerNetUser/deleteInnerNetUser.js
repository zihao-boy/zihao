(function(vc,vm){

    vc.extends({
        data:{
            deleteInnerNetUserInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteInnerNetUser','openDeleteInnerNetUserModal',function(_params){

                vc.component.deleteInnerNetUserInfo = _params;
                $('#deleteInnerNetUserModel').modal('show');

            });
        },
        methods:{
            deleteInnerNetUser:function(){
                vc.http.apiPost(
                    '/innerNet/deleteInnerNetUser',
                    JSON.stringify(vc.component.deleteInnerNetUserInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteInnerNetUserModel').modal('hide');
                            vc.emit('innerNetUserManage','listInnerNetUser',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteInnerNetUserModel:function(){
                $('#deleteInnerNetUserModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
