(function(vc,vm){

    vc.extends({
        data:{
            deleteMonitorHostInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteMonitorHost','openDeleteMonitorHostModal',function(_params){

                vc.component.deleteMonitorHostInfo = _params;
                $('#deleteMonitorHostModel').modal('show');

            });
        },
        methods:{
            deleteMonitorHost:function(){
                
                vc.http.apiPost(
                    '/monitor/deleteMonitorHost',
                    JSON.stringify(vc.component.deleteMonitorHostInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteMonitorHostModel').modal('hide');
                            vc.emit('monitorHostManage','listMonitorHost',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteMonitorHostModel:function(){
                $('#deleteMonitorHostModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
