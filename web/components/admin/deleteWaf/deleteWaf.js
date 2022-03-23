(function(vc,vm){

    vc.extends({
        data:{
            deleteWafInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteWaf','openDeleteWafModal',function(_params){

                vc.component.deleteWafInfo = _params;
                $('#deleteWafModel').modal('show');

            });
        },
        methods:{
            deleteWaf:function(){
                vc.component.deleteWafInfo.communityId=vc.getCurrentCommunity().communityId;
                vc.http.apiPost(
                    'waf.deleteWaf',
                    JSON.stringify(vc.component.deleteWafInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteWafModel').modal('hide');
                            vc.emit('wafManage','listWaf',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteWafModel:function(){
                $('#deleteWafModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
