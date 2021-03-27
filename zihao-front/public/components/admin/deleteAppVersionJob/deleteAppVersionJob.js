(function(vc,vm){

    vc.extends({
        data:{
            deleteAppVersionJobInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteAppVersionJob','openDeleteAppVersionJobModal',function(_params){

                vc.component.deleteAppVersionJobInfo = _params;
                $('#deleteAppVersionJobModel').modal('show');

            });
        },
        methods:{
            deleteAppVersionJob:function(){
                vc.http.apiPost(
                    '/appVersion/deleteAppVersionJob',
                    JSON.stringify(vc.component.deleteAppVersionJobInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteAppVersionJobModel').modal('hide');
                            vc.emit('appVersionJobManage','listAppVersionJob',{});
                            return ;
                        }
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.toast(json);

                     });
            },
            closeDeleteAppVersionJobModel:function(){
                $('#deleteAppVersionJobModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
