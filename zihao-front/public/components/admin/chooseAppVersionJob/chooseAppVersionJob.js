(function(vc){
    vc.extends({
        propTypes: {
           emitChooseAppVersionJob:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseAppVersionJobInfo:{
                appVersionJobs:[],
                _currentAppVersionJobName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseAppVersionJob','openChooseAppVersionJobModel',function(_param){
                $('#chooseAppVersionJobModel').modal('show');
                vc.component._refreshChooseAppVersionJobInfo();
                vc.component._loadAllAppVersionJobInfo(1,10,'');
            });
        },
        methods:{
            _loadAllAppVersionJobInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('appVersionJob.listAppVersionJobs',
                             param,
                             function(json){
                                var _appVersionJobInfo = JSON.parse(json);
                                vc.component.chooseAppVersionJobInfo.appVersionJobs = _appVersionJobInfo.appVersionJobs;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseAppVersionJob:function(_appVersionJob){
                if(_appVersionJob.hasOwnProperty('name')){
                     _appVersionJob.appVersionJobName = _appVersionJob.name;
                }
                vc.emit($props.emitChooseAppVersionJob,'chooseAppVersionJob',_appVersionJob);
                vc.emit($props.emitLoadData,'listAppVersionJobData',{
                    appVersionJobId:_appVersionJob.appVersionJobId
                });
                $('#chooseAppVersionJobModel').modal('hide');
            },
            queryAppVersionJobs:function(){
                vc.component._loadAllAppVersionJobInfo(1,10,vc.component.chooseAppVersionJobInfo._currentAppVersionJobName);
            },
            _refreshChooseAppVersionJobInfo:function(){
                vc.component.chooseAppVersionJobInfo._currentAppVersionJobName = "";
            }
        }

    });
})(window.vc);
