(function(vc,vm){

    vc.extends({
        data:{
            editDnsMapInfo:{
                dnsMapId:'',
host:'',
type:'',
value:'',

            }
        },
         _initMethod:function(){

         },
         _initEvent:function(){
             vc.on('editDnsMap','openEditDnsMapModal',function(_params){
                vc.component.refreshEditDnsMapInfo();
                $('#editDnsMapModel').modal('show');
                vc.copyObject(_params, vc.component.editDnsMapInfo );
                vc.component.editDnsMapInfo.communityId = vc.getCurrentCommunity().communityId;
            });
        },
        methods:{
            editDnsMapValidate:function(){
                        return vc.validate.validate({
                            editDnsMapInfo:vc.component.editDnsMapInfo
                        },{
                            'editDnsMapInfo.host':[
{
                            limit:"required",
                            param:"",
                            errInfo:"域名不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"域名不能超过64"
                        },
                    ],
'editDnsMapInfo.type':[
{
                            limit:"required",
                            param:"",
                            errInfo:"类型不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"类型不能超过64"
                        },
                    ],
'editDnsMapInfo.value':[
{
                            limit:"required",
                            param:"",
                            errInfo:"ip'不能为空"
                        },
 {
                            limit:"maxLength",
                            param:"64",
                            errInfo:"ip'不能超过64"
                        },
                    ],
'editDnsMapInfo.dnsMapId':[
{
                            limit:"required",
                            param:"",
                            errInfo:"编号不能为空"
                        }]

                        });
             },
            editDnsMap:function(){
                if(!vc.component.editDnsMapValidate()){
                    vc.toast(vc.validate.errInfo);
                    return ;
                }

                vc.http.apiPost(
                    'dnsMap.updateDnsMap',
                    JSON.stringify(vc.component.editDnsMapInfo),
                    {
                        emulateJSON:true
                     },
                     function(json,res){
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editDnsMapModel').modal('hide');
                             vc.emit('dnsMapManage','listDnsMap',{});
                            return ;
                        }
                        vc.message(_json.msg);
                     },
                     function(errInfo,error){
                        console.log('请求失败处理');

                        vc.message(errInfo);
                     });
            },
            refreshEditDnsMapInfo:function(){
                vc.component.editDnsMapInfo= {
                  dnsMapId:'',
host:'',
type:'',
value:'',

                }
            }
        }
    });

})(window.vc,window.vc.component);
