/**
 入驻小区
 **/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    var TEMP_SEARCH = "simplifyAcceptanceSearch";
    vc.extends({
        data: {
            appServiceControlInfo: {
                _currentTab: 'appServiceControlPort',
                asId: '',
                asName: '',
                asCount: '',
                stateName: '',
                imagesName: '',
                imagesVersion: '',
                asType: '',
                avgName: '',
                hostGroupName: '',
                hostName: '',
                asGroupId: '',
                hostId: '',
                imagesId: '',
                verId: '',
                state: '',
                asDesc: '',
                asDeployType: '',
                asDeployId: ''
            }
        },
        _initMethod: function() {
            $that.appServiceControlInfo.asId = vc.getParam('asId');
            $that._listAppServices();
        },
        _initEvent: function() {
            vc.on('appServiceControlInfo', 'notify', function(_room) {
                $that._listAppServices();
                vc.emit('simplifyRoomFee', 'switch', $that.appServiceControlInfo)
            });
        },
        methods: {
            changeTab: function(_tab) {
                $that.appServiceControlInfo._currentTab = _tab;
                vc.emit(_tab, 'switch', {
                    asId: $that.appServiceControlInfo.asId,
                    asDeployId: $that.appServiceControlInfo.asDeployId,
                    asDeployType: $that.appServiceControlInfo.asDeployType,
                })
            },
            _clearData: function() {
                $that.appServiceControlInfo = {
                    _currentTab: 'hostContainers',
                    asId: '',
                    asName: '',
                    asCount: '',
                    stateName: '',
                    imagesName: '',
                    imagesVersion: '',
                    asType: '',
                    avgName: '',
                    hostGroupName: '',
                    hostName: '',
                    imagesId: '',
                    verId: '',
                    state: '',
                    asDesc: '',
                    asDeployType: '',
                    asDeployId: ''
                }
            },
            _listAppServices: function() {

                let param = {
                    params: {
                        page: 1,
                        row: 1,
                        asId: $that.appServiceControlInfo.asId
                    }
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppService',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);

                        vc.copyObject(_hostManageInfo.data[0], $that.appServiceControlInfo)

                        $that.changeTab($that.appServiceControlInfo._currentTab);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _startAppService: function() {
                vc.http.apiPost(
                    '/appService/startAppService',
                    JSON.stringify(vc.component.appServiceControlInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            $that._listAppServices();
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            _stopAppService: function() {
                vc.http.apiPost(
                    '/appService/stopAppService',
                    JSON.stringify(vc.component.appServiceControlInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            $that._listAppServices();
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },

            _goBack: function() {
                vc.goBack();
            },
            _getAsType: function(_asType) {
                if (_asType == '001') {
                    return '数据库';
                } else if (_asType == '002') {
                    return '缓存';
                }
                return '计算应用';

            },
            _editAppService: function() {
                vc.emit('editAppService', 'openEditAppServiceModal', $that.appServiceControlInfo);
            }

        }
    });
})(window.vc);