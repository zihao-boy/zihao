(function(vc,vm){

    vc.extends({
        data:{
            deleteMonitorHostGroupInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteMonitorHostGroup','openDeleteMonitorHostGroupModal',function(_params){

                vc.component.deleteMonitorHostGroupInfo = _params;
                $('#deleteMonitorHostGroupModel').modal('show');

            });
        },
        methods:{
            deleteMonitorHostGroup:function(){
                vc.http.apiPost(
                    '/monitor/deleteMonitorHostGroup',
                    JSON.stringify(vc.component.deleteMonitorHostGroupInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteMonitorHostGroupModel').modal('hide');
                            vc.emit('monitorHostGroupManage','listMonitorHostGroup',{});
                            return ;
                        }
                        vc.toast(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.toast(json);

                     });
            },
            closeDeleteMonitorHostGroupModel:function(){
                $('#deleteMonitorHostGroupModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
