(function (vc, vm) {

    vc.extends({
        data: {
            editAppServiceInfo: {
                asId: '',
                asName: '',
                asType: '',
                asDesc: '',
                asCount: '1',
                asGroupId: '',
                asGroups: [],
                hostGroups: [],
                groupId: '',
                asDeployType: '1001',
                hosts: [],
                hostId: '',
                imagesId: '',
                verId: '',
                images: [],
                vers: [],

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editAppService', 'openEditAppServiceModal', function (_params) {
                vc.component.refreshEditAppServiceInfo();
                $that._listEditAppVarGroups();
                $that._listEditHostGroups();
                $that._listEditBusinessImagess();
                $('#editAppServiceModel').modal('show');
                vc.copyObject(_params, vc.component.editAppServiceInfo);
                $that._editChangeImages();
            });
        },
        methods: {
            editAppServiceValidate: function () {
                return vc.validate.validate({
                    editAppServiceInfo: vc.component.editAppServiceInfo
                }, {
                    'editAppServiceInfo.asName': [{
                        limit: "required",
                        param: "",
                        errInfo: "应用名称不能为空"
                    },
                    {
                        limit: "maxLength",
                        param: "128",
                        errInfo: "应用名称太长"
                    },
                    ],
                    'editAppServiceInfo.asType': [{
                        limit: "required",
                        param: "",
                        errInfo: "服务类型不能为空"
                    },
                    {
                        limit: "num",
                        param: "",
                        errInfo: "服务类型格式错误"
                    },
                    ],
                    'editAppServiceInfo.asCount': [{
                        limit: "required",
                        param: "",
                        errInfo: "副本数不能为空"
                    },
                    {
                        limit: "num",
                        param: "",
                        errInfo: "副本数不是有效数字"
                    },
                    ],
                    'editAppServiceInfo.asDesc': [{
                        limit: "maxLength",
                        param: "512",
                        errInfo: "描述太长"
                    },],
                    'editAppServiceInfo.asId': [{
                        limit: "required",
                        param: "",
                        errInfo: "ID不能为空"
                    }]

                });
            },
            editAppService: function () {
                if (!vc.component.editAppServiceValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/appService/updateAppService',
                    JSON.stringify(vc.component.editAppServiceInfo), {
                    emulateJSON: true
                },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppServiceModel').modal('hide');
                            vc.emit('appServiceControlInfo', 'notify', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditAppServiceInfo: function () {
                vc.component.editAppServiceInfo = {
                    asId: '',
                    asName: '',
                    asType: '',
                    asDesc: '',
                    asCount: '1',
                    asGroupId: '',
                    asGroups: [],
                    hostGroups: [],
                    groupId: '',
                    asDeployType: '1001',
                    hosts: [],
                    hostId: '',
                    imagesId: '',
                    verId: '',
                    images: [],
                    vers: [],
                }
            },
            _listEditAppVarGroups: function (_page, _rows) {
                var param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/appService/getAppVarGroup',
                    param,
                    function (json, res) {
                        var _appVarGroupManageInfo = JSON.parse(json);
                        vc.component.editAppServiceInfo.asGroups = _appVarGroupManageInfo.data;
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _listEditHostGroups: function () {
                let param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/host/getHostGroup',
                    param,
                    function (json, res) {
                        let _hostGroupManageInfo = JSON.parse(json);
                        vc.component.editAppServiceInfo.hostGroups = _hostGroupManageInfo.data;
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _changeHostGroup: function () {
                $that._listEditHosts();
            },
            _listEditHosts: function () {


                var param = {
                    params: {
                        page: 1,
                        row: 100,
                        groupId: $that.editAppServiceInfo.groupId
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function (json, res) {
                        var _hostManageInfo = JSON.parse(json);
                        vc.component.editAppServiceInfo.hosts = _hostManageInfo.data;

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _listEditBusinessImagess: function (_page, _rows) {
                var param = {
                    params: {
                        page: 1,
                        row: 100
                    }
                };

                //发送get请求
                vc.http.apiGet('/soft/getBusinessImages',
                    param,
                    function (json, res) {
                        var _businessImagesManageInfo = JSON.parse(json);
                        vc.component.editAppServiceInfo.images = _businessImagesManageInfo.data;
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _editChangeImages:function(){
                var param = {
                    params: {
                        page: 1,
                        row: 50,
                        imagesId:$that.editAppServiceInfo.imagesId
                    }
                };

                //发送get请求
                vc.http.apiGet('/soft/getBusinessImagesVer',
                    param,
                    function (json) {
                        var _verInfo = JSON.parse(json);
                        vc.component.editAppServiceInfo.vers = _verInfo.data;
                        
                    }, function () {
                        console.log('请求失败处理');
                    }
                );
            }
        }
    });

})(window.vc, window.vc.component);