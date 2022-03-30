(function(vc,vm){

    vc.extends({
        data:{
            deleteWafRuleInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteWafRule','openDeleteWafRuleModal',function(_params){

                vc.component.deleteWafRuleInfo = _params;
                $('#deleteWafRuleModel').modal('show');

            });
        },
        methods:{
            deleteWafRule:function(){
                vc.http.apiPost(
                    '/firewall/deleteWafRule',
                    JSON.stringify(vc.component.deleteWafRuleInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteWafRuleModel').modal('hide');
                            vc.emit('wafRuleManage','listWafRule',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteWafRuleModel:function(){
                $('#deleteWafRuleModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
