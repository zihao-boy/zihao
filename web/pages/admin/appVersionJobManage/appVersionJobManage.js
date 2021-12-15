/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appVersionJobManageInfo: {
                appVersionJobs: [],
                total: 0,
                records: 1,
                moreCondition: false,
                jobId: '',
                component: 'appVersionJobManage',
                conditions: {
                    jobName: '',
                    state: '',
                    jobId: '',

                }
            }
        },
        _initMethod: function() {
            vc.component._listAppVersionJobs(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('appVersionJobManage', 'listAppVersionJob', function(_param) {
                $that.appVersionJobManageInfo.component = 'appVersionJobManage';
                vc.component._listAppVersionJobs(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listAppVersionJobs(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listAppVersionJobs: function(_page, _rows) {

                vc.component.appVersionJobManageInfo.conditions.page = _page;
                vc.component.appVersionJobManageInfo.conditions.row = _rows;
                let param = {
                    params: vc.component.appVersionJobManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/appVersion/getAppVersionJob',
                    param,
                    function(json, res) {
                        var _appVersionJobManageInfo = JSON.parse(json);
                        vc.component.appVersionJobManageInfo.total = _appVersionJobManageInfo.total;
                        vc.component.appVersionJobManageInfo.records = _appVersionJobManageInfo.records;
                        vc.component.appVersionJobManageInfo.appVersionJobs = _appVersionJobManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.appVersionJobManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddAppVersionJobModal: function() {
                $that.appVersionJobManageInfo.component = 'addAppVersionJob';
                vc.emit('addAppVersionJob', 'openAddAppVersionJobModal', {});
            },
            _openEditAppVersionJobModel: function(_appVersionJob) {
                vc.emit('editAppVersionJob', 'openEditAppVersionJobModal', _appVersionJob);
                $that.appVersionJobManageInfo.component = 'editAppVersionJob';
            },
            _openDeleteAppVersionJobModel: function(_appVersionJob) {
                vc.emit('deleteAppVersionJob', 'openDeleteAppVersionJobModal', _appVersionJob);
            },
            _queryAppVersionJobMethod: function() {
                vc.component._listAppVersionJobs(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.appVersionJobManageInfo.moreCondition) {
                    vc.component.appVersionJobManageInfo.moreCondition = false;
                } else {
                    vc.component.appVersionJobManageInfo.moreCondition = true;
                }
            },
            _toJob: function(_appVersionJob) {
                let data = {
                        jobId: _appVersionJob.jobId
                    }
                    //发送get请求
                vc.http.apiPost('/appVersion/doJob',
                    JSON.stringify(data), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        var _appVersionJobManageInfo = JSON.parse(json);
                        vc.toast(_appVersionJobManageInfo.msg);
                        vc.emit('appVersionJobManage', 'listAppVersionJob', {})
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }


        }
    });
})(window.vc);