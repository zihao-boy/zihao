(function(vc, vm) {

    vc.extends({
        data: {
            installImagesInfo: {
                asType: '',
                asGroupId: '',
                asGroups: [],
                hostGroups: [],
                groupId: '',
                hosts: [],
                asDeployType: '1001',
                hostId: '',
                asDeployId: '',
                excelTemplate: '',
                appId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('installImages', 'openInstallImagesModal', function(_params) {

                //vc.component.installImagesInfo = _params;
                vc.copyObject(_params, $that.installImagesInfo)
                $('#installImagesModel').modal('show');
                $that._listImportHostGroups();
                $that._listImportAppVarGroups();
            });
            vc.on('installImages', 'openInstallImagesModal', function(_params) {
                $('#installImagesModel').modal('show');
            });
        },
        methods: {
            installImagesValidate: function() {
                return vc.validate.validate({
                    installImagesInfo: vc.component.installImagesInfo
                }, {
                    'installImagesInfo.asType': [{
                        limit: "required",
                        param: "",
                        errInfo: "服务类型不能为空"
                    }],
                    'installImagesInfo.asGroupId': [{
                        limit: "required",
                        param: "",
                        errInfo: "应用组不能为空"
                    }]

                });
            },
            installImages: function() {
                if (!vc.component.installImagesValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                if ($that.installImagesInfo.asDeployType == '1001') {
                    $that.installImagesInfo.asDeployId = $that.installImagesInfo.groupId
                } else {
                    $that.installImagesInfo.asDeployId = $that.installImagesInfo.hostId
                }

                let _data = {
                    asType: vc.component.installImagesInfo.asType,
                    asGroupId: vc.component.installImagesInfo.asGroupId,
                    groupId: vc.component.installImagesInfo.groupId,
                    asDeployType: vc.component.installImagesInfo.asDeployType,
                    asDeployId: vc.component.installImagesInfo.asDeployId,
                    appId: vc.component.installImagesInfo.appId
                }

                vc.http.apiPost(
                    '/soft/installImages',
                    JSON.stringify(_data), {
                        emulateJSON: true,
                    },
                    function(json, res) {
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#installImagesModel').modal('hide');
                            vc.component.refreshInstallImagesInfo();
                            vc.emit('imagesPoolManage', 'listBusinessImages', {});
                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            refreshInstallImagesInfo: function() {
                vc.component.installImagesInfo = {
                    asType: '',
                    asGroupId: '',
                    asGroups: [],
                    hostGroups: [],
                    groupId: '',
                    hosts: [],
                    asDeployType: '1001',
                    hostId: '',
                    asDeployId: '',
                    excelTemplate: ''
                }
            },
            _listImportHostGroups: function() {
                let param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/host/getHostGroup',
                    param,
                    function(json, res) {
                        let _hostGroupManageInfo = JSON.parse(json);
                        vc.component.installImagesInfo.hostGroups = _hostGroupManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _changeImportHostGroup: function() {
                $that._listImportHosts();
            },
            _listImportHosts: function() {


                let param = {
                    params: {
                        page: 1,
                        row: 100,
                        groupId: $that.installImagesInfo.groupId
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function(json, res) {
                        var _hostManageInfo = JSON.parse(json);
                        vc.component.installImagesInfo.hosts = _hostManageInfo.data;

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _listImportAppVarGroups: function(_page, _rows) {
                let param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/appService/getAppVarGroup',
                    param,
                    function(json, res) {
                        let _appVarGroupManageInfo = JSON.parse(json);
                        vc.component.installImagesInfo.asGroups = _appVarGroupManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            getExcelTemplate: function(e) {
                //console.log("getExcelTemplate 开始调用")
                vc.component.installImagesInfo.excelTemplate = e.target.files[0];
            },
        }
    });

})(window.vc, window.vc.component);