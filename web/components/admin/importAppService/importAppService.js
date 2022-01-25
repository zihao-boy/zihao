(function(vc, vm) {

    vc.extends({
        data: {
            importAppServiceInfo: {
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
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('importAppService', 'openImportAppServiceModal', function(_params) {
                $('#importAppServiceModel').modal('show');
                $that._listImportHostGroups();
                $that._listImportAppVarGroups();
            });
        },
        methods: {
            importAppServiceValidate: function() {
                return vc.validate.validate({
                    importAppServiceInfo: vc.component.importAppServiceInfo
                }, {
                    'importAppServiceInfo.asType': [{
                        limit: "required",
                        param: "",
                        errInfo: "服务类型不能为空"
                    }],
                    'importAppServiceInfo.asGroupId': [{
                        limit: "required",
                        param: "",
                        errInfo: "应用组不能为空"
                    }]

                });
            },
            importAppService: function() {
                if (!vc.component.importAppServiceValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                if ($that.importAppServiceInfo.asDeployType == '1001') {
                    $that.importAppServiceInfo.asDeployId = $that.importAppServiceInfo.groupId
                } else {
                    $that.importAppServiceInfo.asDeployId = $that.importAppServiceInfo.hostId
                }

                let param = new FormData();
                param.append("uploadFile", vc.component.importAppServiceInfo.excelTemplate);
                param.append('asType', vc.component.importAppServiceInfo.asType);
                param.append('asGroupId', vc.component.importAppServiceInfo.asGroupId);
                param.append('groupId', vc.component.importAppServiceInfo.groupId);
                param.append('asDeployType', vc.component.importAppServiceInfo.asDeployType);
                param.append('asDeployId', vc.component.importAppServiceInfo.asDeployId);

                vc.http.apiPost(
                    '/appService/importAppService',
                    param, {
                        emulateJSON: true,
                        //添加请求头
                        headers: {
                            "Content-Type": "multipart/form-data"
                        }
                    },
                    function(json, res) {
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#importAppServiceModel').modal('hide');
                            vc.component.refreshImportAppServiceInfo();
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
            refreshImportAppServiceInfo: function() {
                vc.component.importAppServiceInfo = {
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
                        vc.component.importAppServiceInfo.hostGroups = _hostGroupManageInfo.data;
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
                        groupId: $that.importAppServiceInfo.groupId
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function(json, res) {
                        var _hostManageInfo = JSON.parse(json);
                        vc.component.importAppServiceInfo.hosts = _hostManageInfo.data;

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
                        vc.component.importAppServiceInfo.asGroups = _appVarGroupManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            getExcelTemplate: function(e) {
                //console.log("getExcelTemplate 开始调用")
                vc.component.importAppServiceInfo.excelTemplate = e.target.files[0];
            },
        }
    });

})(window.vc, window.vc.component);