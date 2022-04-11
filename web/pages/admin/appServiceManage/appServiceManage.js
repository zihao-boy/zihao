/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appServiceManageInfo: {
                appServices: [],
                total: 0,
                records: 1,
                moreCondition: false,
                component: 'appServiceManage',
                asId: '',
                asGroups: [],
                conditions: {
                    asName: '',
                    asType: '',
                    asCount: '',
                    asGroupId: '',
                    state: ''
                }
            }
        },
        _initMethod: function() {

            $that._listListAppVarGroups();
        },
        _initEvent: function() {
            vc.on('appServiceManage', 'listAppService', function(_param) {
                $that.appServiceManageInfo.component = 'appServiceManage';
                vc.component._listAppServices(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                DEFAULT_PAGE = _currentPage;
                vc.component._listAppServices(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listAppServices: function(_page, _rows) {

                vc.component.appServiceManageInfo.conditions.page = _page;
                vc.component.appServiceManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.appServiceManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppService',
                    param,
                    function(json, res) {
                        var _appServiceManageInfo = JSON.parse(json);
                        vc.component.appServiceManageInfo.total = _appServiceManageInfo.total;
                        vc.component.appServiceManageInfo.records = _appServiceManageInfo.records;
                        vc.component.appServiceManageInfo.appServices = _appServiceManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.appServiceManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
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
                        vc.component.appServiceManageInfo.asGroups = _appVarGroupManageInfo.data;
                        if (_appVarGroupManageInfo.data.length < 1) {
                            return;
                        }
                        $that.appServiceManageInfo.conditions.asGroupId = _appVarGroupManageInfo.data[0].avgId;
                        vc.component._listAppServices(DEFAULT_PAGE, DEFAULT_ROWS);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddAppServiceModal: function() {
                $that.appServiceManageInfo.component = 'addAppService';
                vc.emit('addAppService', 'openAddAppServiceModal', {});
            },
            _openControl: function(_appService) {
                vc.jumpToPage('/index.html#/pages/admin/appServiceControl?asId=' + _appService.asId)
            },
            _copyAppService: function(_appService) {
                vc.emit('copyAppService', 'openCopyAppServiceModal', _appService);
            },
            _upgradeAppService: function(_appService) {
                vc.emit('upgradeAppService', 'openUpgradeAppServiceModal', _appService);
            },
            _openDeleteAppServiceModel: function(_appService) {
                vc.emit('deleteAppService', 'openDeleteAppServiceModal', _appService);
            },
            _queryAppServiceMethod: function() {
                vc.component._listAppServices(DEFAULT_PAGE, DEFAULT_ROWS);
            },
            _moreCondition: function() {
                if (vc.component.appServiceManageInfo.moreCondition) {
                    vc.component.appServiceManageInfo.moreCondition = false;
                } else {
                    vc.component.appServiceManageInfo.moreCondition = true;
                }
            },
            _startAppService: function(_appService) {
                vc.http.apiPost(
                    '/appService/startAppService',
                    JSON.stringify(_appService), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            vc.emit('appServiceManage', 'listAppService', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });
            },
            _stopAppService: function(_appService) {
                vc.http.apiPost(
                    '/appService/stopAppService',
                    JSON.stringify(_appService), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            vc.emit('appServiceManage', 'listAppService', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });
            },
            _openFasterDeploy: function() {
                vc.jumpToPage('/index.html#/pages/admin/fasterDeploy')
            },
            _exportYaml: function() {
                //;
                vc.emit('selectAppService', 'openSelectAppService', {
                    asGroupId: $that.appServiceManageInfo.conditions.asGroupId,
                    callBack: $that._doExportYaml
                })
            },
            _doExportYaml: function(_asIds) {
                vc.jumpToPage('/app/appService/exportAppService?asIds=' + _asIds);
            },
            _importYaml: function() {
                vc.emit('importAppService', 'openImportAppServiceModal', {});
            },
            _restartMoreApps: function() {
                vc.emit('selectAppService', 'openSelectAppService', {
                    asGroupId: $that.appServiceManageInfo.conditions.asGroupId,
                    callBack: $that._doRestartMoreApps
                })
            },
            _doRestartMoreApps: function(_asIds) {
                let _data = {
                    asIds: _asIds
                }
                vc.http.apiPost(
                    '/appService/restartAppServices',
                    JSON.stringify(_data), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            vc.emit('appServiceManage', 'listAppService', {});
                            vc.toast('重启请求已经发送，请点击查询查看应用状态');
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });
            },
            _openDockerLog: function(_appService) {
                vc.emit('chooseHostContainer', 'openChooseHostContainerModel', _appService);
            },

            swatch: function(item) {
                $that.appServiceManageInfo.conditions.asGroupId = item.avgId;
                $that._listAppServices(1, DEFAULT_ROWS);
            }
        }
    });
})(window.vc);