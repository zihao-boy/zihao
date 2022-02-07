(function(vc){
    vc.extends({
        propTypes: {
           emitChooseAppPublisher:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseAppPublisherInfo:{
                appPublishers:[],
                _currentAppPublisherName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseAppPublisher','openChooseAppPublisherModel',function(_param){
                $('#chooseAppPublisherModel').modal('show');
                vc.component._refreshChooseAppPublisherInfo();
                vc.component._loadAllAppPublisherInfo(1,10,'');
            });
        },
        methods:{
            _loadAllAppPublisherInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('appPublisher.listAppPublishers',
                             param,
                             function(json){
                                var _appPublisherInfo = JSON.parse(json);
                                vc.component.chooseAppPublisherInfo.appPublishers = _appPublisherInfo.appPublishers;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseAppPublisher:function(_appPublisher){
                if(_appPublisher.hasOwnProperty('name')){
                     _appPublisher.appPublisherName = _appPublisher.name;
                }
                vc.emit($props.emitChooseAppPublisher,'chooseAppPublisher',_appPublisher);
                vc.emit($props.emitLoadData,'listAppPublisherData',{
                    appPublisherId:_appPublisher.appPublisherId
                });
                $('#chooseAppPublisherModel').modal('hide');
            },
            queryAppPublishers:function(){
                vc.component._loadAllAppPublisherInfo(1,10,vc.component.chooseAppPublisherInfo._currentAppPublisherName);
            },
            _refreshChooseAppPublisherInfo:function(){
                vc.component.chooseAppPublisherInfo._currentAppPublisherName = "";
            }
        }

    });
})(window.vc);
