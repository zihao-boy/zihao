(function(vc,vm){

    vc.extends({
        data:{
            deleteInnerNetPrivilegeInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteInnerNetPrivilege','openDeleteInnerNetPrivilegeModal',function(_params){

                vc.component.deleteInnerNetPrivilegeInfo = _params;
                $('#deleteInnerNetPrivilegeModel').modal('show');

            });
        },
        methods:{
            deleteInnerNetPrivilege:function(){
                vc.http.apiPost(
                    '/innerNet/deleteInnerNetPrivilege',
                    JSON.stringify(vc.component.deleteInnerNetPrivilegeInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteInnerNetPrivilegeModel').modal('hide');
                            vc.emit('innerNetPrivilegeManage','listInnerNetPrivilege',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteInnerNetPrivilegeModel:function(){
                $('#deleteInnerNetPrivilegeModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
