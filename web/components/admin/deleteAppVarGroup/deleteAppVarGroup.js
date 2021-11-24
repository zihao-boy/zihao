(function(vc,vm){

    vc.extends({
        data:{
            deleteAppVarGroupInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteAppVarGroup','openDeleteAppVarGroupModal',function(_params){

                vc.component.deleteAppVarGroupInfo = _params;
                $('#deleteAppVarGroupModel').modal('show');

            });
        },
        methods:{
            deleteAppVarGroup:function(){
                vc.http.apiPost(
                    '/appService/deleteAppVarGroup',
                    JSON.stringify(vc.component.deleteAppVarGroupInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteAppVarGroupModel').modal('hide');
                            vc.emit('appVarGroupManage','listAppVarGroup',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteAppVarGroupModel:function(){
                $('#deleteAppVarGroupModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
