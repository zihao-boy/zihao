(function(vc,vm){

    vc.extends({
        data:{
            deleteBusinessPackageInfo:{

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('deleteBusinessPackage','openDeleteBusinessPackageModal',function(_params){

                vc.component.deleteBusinessPackageInfo = _params;
                $('#deleteBusinessPackageModel').modal('show');

            });
        },
        methods:{
            deleteBusinessPackage:function(){
                vc.component.deleteBusinessPackageInfo.communityId=vc.getCurrentCommunity().communityId;
                vc.http.apiPost(
                    '/soft/deleteBusinessPackages',
                    JSON.stringify(vc.component.deleteBusinessPackageInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteBusinessPackageModel').modal('hide');
                            vc.emit('businessPackageManage','listBusinessPackage',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');
                        vc.message(json);

                     });
            },
            closeDeleteBusinessPackageModel:function(){
                $('#deleteBusinessPackageModel').modal('hide');
            }
        }
    });

})(window.vc,window.vc.component);
