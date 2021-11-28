(function(vc){
    vc.extends({
        propTypes: {
           emitChooseBusinessPackage:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseBusinessPackageInfo:{
                businessPackages:[],
                _currentBusinessPackageName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseBusinessPackage','openChooseBusinessPackageModel',function(_param){
                $('#chooseBusinessPackageModel').modal('show');
                vc.component._refreshChooseBusinessPackageInfo();
                vc.component._loadAllBusinessPackageInfo(1,10,'');
            });
        },
        methods:{
            _loadAllBusinessPackageInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('businessPackage.listBusinessPackages',
                             param,
                             function(json){
                                var _businessPackageInfo = JSON.parse(json);
                                vc.component.chooseBusinessPackageInfo.businessPackages = _businessPackageInfo.businessPackages;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseBusinessPackage:function(_businessPackage){
                if(_businessPackage.hasOwnProperty('name')){
                     _businessPackage.businessPackageName = _businessPackage.name;
                }
                vc.emit($props.emitChooseBusinessPackage,'chooseBusinessPackage',_businessPackage);
                vc.emit($props.emitLoadData,'listBusinessPackageData',{
                    businessPackageId:_businessPackage.businessPackageId
                });
                $('#chooseBusinessPackageModel').modal('hide');
            },
            queryBusinessPackages:function(){
                vc.component._loadAllBusinessPackageInfo(1,10,vc.component.chooseBusinessPackageInfo._currentBusinessPackageName);
            },
            _refreshChooseBusinessPackageInfo:function(){
                vc.component.chooseBusinessPackageInfo._currentBusinessPackageName = "";
            }
        }

    });
})(window.vc);
