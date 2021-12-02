(function(vc){
    vc.extends({
        propTypes: {
           emitChooseBusinessImages:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseBusinessImagesInfo:{
                businessImagess:[],
                _currentBusinessImagesName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseBusinessImages','openChooseBusinessImagesModel',function(_param){
                $('#chooseBusinessImagesModel').modal('show');
                vc.component._refreshChooseBusinessImagesInfo();
                vc.component._loadAllBusinessImagesInfo(1,10,'');
            });
        },
        methods:{
            _loadAllBusinessImagesInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('businessImages.listBusinessImagess',
                             param,
                             function(json){
                                var _businessImagesInfo = JSON.parse(json);
                                vc.component.chooseBusinessImagesInfo.businessImagess = _businessImagesInfo.businessImagess;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseBusinessImages:function(_businessImages){
                if(_businessImages.hasOwnProperty('name')){
                     _businessImages.businessImagesName = _businessImages.name;
                }
                vc.emit($props.emitChooseBusinessImages,'chooseBusinessImages',_businessImages);
                vc.emit($props.emitLoadData,'listBusinessImagesData',{
                    businessImagesId:_businessImages.businessImagesId
                });
                $('#chooseBusinessImagesModel').modal('hide');
            },
            queryBusinessImagess:function(){
                vc.component._loadAllBusinessImagesInfo(1,10,vc.component.chooseBusinessImagesInfo._currentBusinessImagesName);
            },
            _refreshChooseBusinessImagesInfo:function(){
                vc.component.chooseBusinessImagesInfo._currentBusinessImagesName = "";
            }
        }

    });
})(window.vc);
