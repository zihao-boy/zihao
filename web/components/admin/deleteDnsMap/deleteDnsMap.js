(function(vc,vm){

    vc.extends({
        data:{
            deleteDnsMapInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteDnsMap','openDeleteDnsMapModal',function(_params){

                vc.component.deleteDnsMapInfo = _params;
                $('#deleteDnsMapModel').modal('show');

            });
        },
        methods:{
            deleteDnsMap:function(){
                vc.component.deleteDnsMapInfo.communityId=vc.getCurrentCommunity().communityId;
                vc.http.apiPost(
                    'dnsMap.deleteDnsMap',
                    JSON.stringify(vc.component.deleteDnsMapInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteDnsMapModel').modal('hide');
                            vc.emit('dnsMapManage','listDnsMap',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteDnsMapModel:function(){
                $('#deleteDnsMapModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
