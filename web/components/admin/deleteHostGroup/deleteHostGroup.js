(function(vc,vm){

    vc.extends({
        data:{
            deleteHostGroupInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteHostGroup','openDeleteHostGroupModal',function(_params){

                vc.component.deleteHostGroupInfo = _params;
                $('#deleteHostGroupModel').modal('show');

            });
        },
        methods:{
            deleteHostGroup:function(){
               
                vc.http.apiPost(
                    '/host/deleteHostGroup',
                    JSON.stringify(vc.component.deleteHostGroupInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteHostGroupModel').modal('hide');
                            vc.emit('hostGroupManage','listHostGroup',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteHostGroupModel:function(){
                $('#deleteHostGroupModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
