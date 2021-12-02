(function(vc,vm){

    vc.extends({
        data:{
            deleteBusinessImagesInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteBusinessImages','openDeleteBusinessImagesModal',function(_params){

                vc.component.deleteBusinessImagesInfo = _params;
                $('#deleteBusinessImagesModel').modal('show');

            });
        },
        methods:{
            deleteBusinessImages:function(){
                vc.http.apiPost(
                    '/soft/deleteBusinessImages',
                    JSON.stringify(vc.component.deleteBusinessImagesInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteBusinessImagesModel').modal('hide');
                            vc.emit('businessImagesManage','listBusinessImages',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteBusinessImagesModel:function(){
                $('#deleteBusinessImagesModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
