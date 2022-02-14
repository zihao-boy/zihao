/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            backupManageInfo: {
                backups: [],
                total: 0,
                records: 1,
                moreCondition: false,
                id: '',
                conditions: {
                    id: '',
                    name: '',
                    typeCd: '',
                }
            }
        },
        _initMethod: function() {
            vc.component._listBackups(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('backupManage', 'listBackup', function(_param) {
                vc.component._listBackups(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listBackups(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listBackups: function(_page, _rows) {

                vc.component.backupManageInfo.conditions.page = _page;
                vc.component.backupManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.backupManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/resources/getBackUp',
                    param,
                    function(json, res) {
                        var _backupManageInfo = JSON.parse(json);
                        vc.component.backupManageInfo.total = _backupManageInfo.total;
                        vc.component.backupManageInfo.records = _backupManageInfo.records;
                        vc.component.backupManageInfo.backups = _backupManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.backupManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddBackupModal: function() {
                vc.emit('addBackup', 'openAddBackupModal', {});
            },
            _openEditBackupModel: function(_backup) {
                vc.emit('editBackup', 'openEditBackupModal', _backup);
            },
            _openDeleteBackupModel: function(_backup) {
                vc.emit('deleteBackup', 'openDeleteBackupModal', _backup);
            },
            _queryBackupMethod: function() {
                vc.component._listBackups(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.backupManageInfo.moreCondition) {
                    vc.component.backupManageInfo.moreCondition = false;
                } else {
                    vc.component.backupManageInfo.moreCondition = true;
                }
            },
            _getTargetTypeCdName: function(_backup) {
                if (_backup.targetTypeCd == '001') {
                    return 'ftp';
                } else if (_backup.targetTypeCd == '002') {
                    return 'oss';
                } else {
                    return '数据库';
                }
            },
            _getExecTime: function(_time) {
                if (_time == '0 0 0 */1 * ?') {
                    return "每天0点";
                } else if (_time == '0 0 0 * * 1') {
                    return "每周周一0点";
                } else if (_time == '0 0 0 1 */1 ?') {
                    return "每月1日0点";
                } else {
                    return _time;
                }
            },
            _openStartModel: function(_backup) {
                vc.http.apiPost(
                    '/resources/startBackUp',
                    JSON.stringify(_backup), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('backupManage', 'listBackup', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });
            },
            _openStopModel: function(_backup) {
                vc.http.apiPost(
                    '/resources/stopBackUp',
                    JSON.stringify(_backup), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('backupManage', 'listBackup', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });
            }


        }
    });
})(window.vc);