(function(vc,vm){

    vc.extends({
        data:{
            deleteHostInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteHost','openDeleteHostModal',function(_params){

                vc.component.deleteHostInfo = _params;
                $('#deleteHostModel').modal('show');

            });
        },
        methods:{
            deleteHost:function(){
                vc.http.apiPost(
                    '/host/deleteHost',
                    JSON.stringify(vc.component.deleteHostInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteHostModel').modal('hide');
                            vc.emit('hostManage','listHost',{});
                            return ;
                        }
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.toast(json);

                     });
            },
            closeDeleteHostModel:function(){
                $('#deleteHostModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
