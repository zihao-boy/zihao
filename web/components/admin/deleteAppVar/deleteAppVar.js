(function(vc,vm){

    vc.extends({
        data:{
            deleteAppVarInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteAppVar','openDeleteAppVarModal',function(_params){

                vc.component.deleteAppVarInfo = _params;
                $('#deleteAppVarModel').modal('show');

            });
        },
        methods:{
            deleteAppVar:function(){
                vc.http.apiPost(
                    '/appService/deleteAppVar',
                    JSON.stringify(vc.component.deleteAppVarInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteAppVarModel').modal('hide');
                            vc.emit('appVarManage','listAppVar',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteAppVarModel:function(){
                $('#deleteAppVarModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
