(function(vc,vm){

    vc.extends({
        data:{
            editBusinessPackageInfo:{
                id:'',
name:'',
varsion:'',

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('editBusinessPackage','openEditBusinessPackageModal',function(_params){
                vc.component.refreshEditBusinessPackageInfo();
                $('#editBusinessPackageModel').modal('show');
                vc.copyObject(_params, vc.component.editBusinessPackageInfo );
                vc.component.editBusinessPackageInfo.communityId = vc.getCurrentCommunity().communityId;
            });
        },
        methods:{
            editBusinessPackageValidate:function(){
                        return vc.validate.validate({
                            editBusinessPackageInfo:vc.component.editBusinessPackageInfo
                        },{
                            'editBusinessPackageInfo.name':[
{
                            limit:"required",
                            param:"",
                            errInfo:"名称不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"名称不能超过64"
                        },
                    ],
'editBusinessPackageInfo.varsion':[
{
                            limit:"required",
                            param:"",
                            errInfo:"版本不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"32",
                            errInfo:"版本不能超过32"
                        },
                    ],
'editBusinessPackageInfo.id':[
{
                            limit:"required",
                            param:"",
                            errInfo:"ID不能为空"
                        }]

                        });
             },
            editBusinessPackage:function(){
                if(!vc.component.editBusinessPackageValidate()){
                    vc.toast(vc.validate.errInfo);
                    return ;
                }

                vc.http.apiPost(
                    'businessPackage.updateBusinessPackage',
                    JSON.stringify(vc.component.editBusinessPackageInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editBusinessPackageModel').modal('hide');
                             vc.emit('businessPackageManage','listBusinessPackage',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');

                        vc.message(errInfo);
                     });
            },
            refreshEditBusinessPackageInfo:function(){
                vc.component.editBusinessPackageInfo= {
                  id:'',
name:'',
varsion:'',

                }
            }
        }
    });

})(window.vc,window.vc.component);
