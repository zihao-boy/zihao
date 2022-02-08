(function(vc) {
    vc.extends({
        data: {
            applyPublishAppInfo: {
                name: '',
                version: '',
                publisherId: '',
                appPublishers: [],
                appShell: '',
                apps: [],
            }
        },
        _initMethod: function() {
            $that._listAppPublishers();

        },
        _initEvent: function() {
            vc.on("applyPublishApp", "chooseAppService", function(_param) {
                $that.applyPublishAppInfo.apps.push(_param);
            })
        },
        methods: {
            applyPublishAppValidate() {
                return vc.validate.validate({
                    applyPublishAppInfo: vc.component.applyPublishAppInfo
                }, {
                    'applyPublishAppInfo.name': [{
                            limit: "required",
                            param: "",
                            errInfo: "软件名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "软件名称太长"
                        },
                    ],
                    'applyPublishAppInfo.version': [{
                            limit: "required",
                            param: "",
                            errInfo: "软件版本不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "12",
                            errInfo: "软件版本错误"
                        },
                    ],
                    'applyPublishAppInfo.publisherId': [{
                        limit: "required",
                        param: "",
                        errInfo: "发布者不能为空"
                    }],
                });
            },
            _doApplyPublishApp: function() {
                if (!vc.component.applyPublishAppValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                if ($that.applyPublishAppInfo.apps.length < 1) {
                    vc.toast('请选择应用');
                    return;
                }
                vc.http.apiPost(
                    '/soft/applyPublishApp',
                    JSON.stringify(vc.component.applyPublishAppInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.goBack();
                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            _addJobPlan: function() {
                vc.emit('chooseAppService', 'openChooseAppServiceModel', {})

            },
            _deleteJobPlan: function(_index) {
                $that.applyPublishAppInfo.apps.splice(_index, 1);
            },
            _appChooseBusinessPackage: function(_data) {
                vc.emit('chooseBusinessPackage', 'openChooseBusinessPackageModel', _data);
            },
            _appChooseDockerfile: function(_data) {
                vc.emit('chooseDockerfile', 'openChooseDockerfileModel', _data);
            },
            _goBack: function() {
                vc.goBack();

            },
            _listAppPublishers: function(_page, _rows) {
                var param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/soft/getAppPublisher',
                    param,
                    function(json, res) {
                        var _appPublisherManageInfo = JSON.parse(json);
                        vc.component.applyPublishAppInfo.appPublishers = _appPublisherManageInfo.data;

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);