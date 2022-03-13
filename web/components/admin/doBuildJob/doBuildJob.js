(function (vc, vm) {

    vc.extends({
        data: {
            doBuildJobInfo: {
                jobId: '',
                appVersionJobImages: [],
                images: [],
                action:'build',
                isAll: true,
                asGroups:[],
                avgIds:[]
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('doBuildJob', 'openDoBuildJobModal', function (_params) {
                vc.copyObject(_params, $that.doBuildJobInfo);
                //load plan
                $that._doBuildLoadImages();
                $that._listListAppVarGroups();
                $('#doBuildJobModel').modal('show');

            });
        },
        methods: {
            _build: function () {

                if($that.doBuildJobInfo.action == 'buildAndStart' && $that.doBuildJobInfo.avgIds.length < 1){
                    vc.toast('请选择应用组');
                    return;
                }
                
                let data = {
                    jobId: $that.doBuildJobInfo.jobId,
                    action: $that.doBuildJobInfo.action,
                    avgIds: $that.doBuildJobInfo.avgIds.join(','),
                    images: $that.doBuildJobInfo.images.join(',')
                }
                //发送get请求
                vc.http.apiPost('/appVersion/doJob',
                    JSON.stringify(data), {
                    emulateJSON: true
                },
                    function (json, res) {
                        $('#doBuildJobModel').modal('hide');
                        var _appVersionJobManageInfo = JSON.parse(json);
                        vc.toast(_appVersionJobManageInfo.msg);
                        vc.emit('appVersionJobManage', 'listAppVersionJob', {})
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            closeDoBuildJobModel: function () {
                $('#doBuildJobModel').modal('hide');
                $that.doBuildJobInfo = {
                    jobId: '',
                    appVersionJobImages: [],
                    images: [],
                    action:'build',
                    isAll: true,
                    asGroups:[],
                    avgIds:[]
                }
            },
            _doBuildLoadImages: function () {
                let param = {
                    params: {
                        page: 1,
                        row: 100,
                        jobId: $that.doBuildJobInfo.jobId
                    }
                };
                //发送get请求
                vc.http.apiGet('/appVersion/getAppVersionJobImages',
                    param,
                    function (json, res) {
                        let _appVersionJobManageInfo = JSON.parse(json);
                        vc.component.doBuildJobInfo.appVersionJobImages = _appVersionJobManageInfo.data;
                        $that.changeAll();
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            changeAll: function () {
                $that.doBuildJobInfo.images = [];
                if (!$that.doBuildJobInfo.isAll) {
                    return;
                }

                $that.doBuildJobInfo.appVersionJobImages.forEach(item => {
                    $that.doBuildJobInfo.images.push(item.jobImagesId);
                });
            },
            changeItem: function () {
                if ($that.doBuildJobInfo.images.length < $that.doBuildJobInfo.appVersionJobImages.length) {
                    $that.doBuildJobInfo.isAll = false;
                    return;
                }
                $that.doBuildJobInfo.isAll = true;

            },
            _listListAppVarGroups: function(_page, _rows) {
                var param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/appService/getAppVarGroup',
                    param,
                    function(json, res) {
                        var _appVarGroupManageInfo = JSON.parse(json);
                        vc.component.doBuildJobInfo.asGroups = _appVarGroupManageInfo.data;
                        
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc, window.vc.component);
