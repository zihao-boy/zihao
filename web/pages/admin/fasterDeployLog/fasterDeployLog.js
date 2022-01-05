/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            fasterDeployLogInfo: {
                fasterDeploys: [],
                total: 0,
                records: 1,
                moreCondition: false,
                hostId: '',
                conditions: {
                    appName: '',
                    deployId: ''
                }
            }
        },
        _initMethod: function() {
            vc.component._listFasterDeployLogs(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {
            vc.on('fasterDeployLog', 'listFasterDeployLog', function(_param) {
                vc.component._listFasterDeployLogs(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listFasterDeployLogs(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listFasterDeployLogs: function(_page, _rows) {

                vc.component.fasterDeployLogInfo.conditions.page = _page;
                vc.component.fasterDeployLogInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.fasterDeployLogInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/appService/getFasterDeploy',
                    param,
                    function(json, res) {
                        var _fasterDeployLogInfo = JSON.parse(json);
                        vc.component.fasterDeployLogInfo.total = _fasterDeployLogInfo.total;
                        vc.component.fasterDeployLogInfo.records = _fasterDeployLogInfo.records;
                        vc.component.fasterDeployLogInfo.fasterDeploys = _fasterDeployLogInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.fasterDeployLogInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddFasterDeployLogModal: function() {
                vc.emit('addFasterDeployLog', 'openAddFasterDeployLogModal', {});
            },
            _openEditFasterDeployLogModel: function(_fasterDeploy) {
                vc.emit('editFasterDeply', 'openEditFasterDeplyModal', _fasterDeploy);
            },
            _openDeleteFasterDeployLogModel: function(_fasterDeploy) {
                vc.emit('deleteFasterDeployLog', 'openDeleteFasterDeployLogModal', _fasterDeploy);
            },
            _queryFasterDeployLogMethod: function() {
                vc.component._listFasterDeployLogs(DEFAULT_PAGE, DEFAULT_ROWS);
            },
            _moreCondition: function() {
                if (vc.component.fasterDeployLogInfo.moreCondition) {
                    vc.component.fasterDeployLogInfo.moreCondition = false;
                } else {
                    vc.component.fasterDeployLogInfo.moreCondition = true;
                }
            },
            _goBack: function() {
                vc.goBack();
            },
            _toDockerfile: function(_fasterDeploy) {
                vc.jumpToPage('/index.html#/pages/admin/businessDockerfileManage')
            }
        }
    });
})(window.vc);