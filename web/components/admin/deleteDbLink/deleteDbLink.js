(function(vc,vm){

    vc.extends({
        data:{
            deleteDbLinkInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteDbLink','openDeleteDbLinkModal',function(_params){

                vc.component.deleteDbLinkInfo = _params;
                $('#deleteDbLinkModel').modal('show');

            });
        },
        methods:{
            deleteDbLink:function(){
                vc.http.apiPost(
                    '/dbClient/deleteDbLink',
                    JSON.stringify(vc.component.deleteDbLinkInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteDbLinkModel').modal('hide');
                            vc.emit('mysqlClient', 'load', {});
                            return ;
                        }
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.toast(json);

                     });
            },
            closeDeleteDbLinkModel:function(){
                $('#deleteDbLinkModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
