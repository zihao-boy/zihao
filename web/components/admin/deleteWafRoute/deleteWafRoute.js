(function(vc,vm){

    vc.extends({
        data:{
            deleteWafRouteInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteWafRoute','openDeleteWafRouteModal',function(_params){

                vc.component.deleteWafRouteInfo = _params;
                $('#deleteWafRouteModel').modal('show');

            });
        },
        methods:{
            deleteWafRoute:function(){
                vc.component.deleteWafRouteInfo.communityId=vc.getCurrentCommunity().communityId;
                vc.http.apiPost(
                    'wafRoute.deleteWafRoute',
                    JSON.stringify(vc.component.deleteWafRouteInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteWafRouteModel').modal('hide');
                            vc.emit('wafRouteManage','listWafRoute',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteWafRouteModel:function(){
                $('#deleteWafRouteModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
