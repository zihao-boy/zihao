(function(vc) {

    vc.extends({

        data: {
            editBackupInfo: {
                id: '',
                name: '',
                execTime: '',
                typeCd: '',
                srcId: '',
                srcObject: '',
                targetId: '',
                targetTypeCd: '',
                srcs: [],
                targets: [],
                execFlag: '1'

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editBackup', 'openEditBackupModal', function(_param) {
                vc.copyObject(_param, $that.editBackupInfo);
                $('#editBackupModel').modal('show');
                $that._changeSrcType();
                $that._changeTargetType();
                $that.editBackupInfo.srcId = _param.srcId;
                $that.editBackupInfo.srcObject = _param.srcObject;
                $that.editBackupInfo.targetId = _param.targetId;

                if (_param.execTime == '0 0 0 */1 * ?') {
                    $that.editBackupInfo.execFlag = '1';
                } else if (_param.execTime == '0 0 0 * * 1') {
                    $that.editBackupInfo.execFlag = '1';
                } else if (_param.execTime == '0 0 0 1 */1 ?') {
                    $that.editBackupInfo.execFlag = '1';
                } else {
                    $that.editBackupInfo.execFlag = '2';
                }

            });
        },
        methods: {
            editBackupValidate() {
                return vc.validate.validate({
                    editBackupInfo: vc.component.editBackupInfo
                }, {
                    'editBackupInfo.name': [{
                            limit: "required",
                            param: "",
                            errInfo: "名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "名称不能超过64"
                        },
                    ],
                    'editBackupInfo.execTime': [{
                            limit: "required",
                            param: "",
                            errInfo: "执行时间不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "执行时间不能超过64"
                        },
                    ],
                    'editBackupInfo.typeCd': [{
                            limit: "required",
                            param: "",
                            errInfo: "类型不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "12",
                            errInfo: "类型不能超过12"
                        },
                    ],
                    'editBackupInfo.srcId': [{
                            limit: "required",
                            param: "",
                            errInfo: "备份编号不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "备份编号不能超过64"
                        },
                    ],
                    'editBackupInfo.srcObject': [{
                            limit: "required",
                            param: "",
                            errInfo: "备份对象不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "1024",
                            errInfo: "备份对象不能超过1024"
                        },
                    ],
                    'editBackupInfo.targetId': [{
                            limit: "required",
                            param: "",
                            errInfo: "目标不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "目标不能超过64"
                        },
                    ],
                    'editBackupInfo.id': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }],
                });
            },
            _changeSrcType: function() {
                let _srcType = $that.editBackupInfo.typeCd;
                $that.editBackupInfo.srcId = '';
                if (!_srcType) {
                    return;
                }

                if (_srcType == '10001') {
                    $that.editBackupInfo.srcObject = "*";
                    $that._listEditDbs('src');
                } else {
                    $that.editBackupInfo.srcObject = "";
                    $that._listEditHosts();
                }

            },
            _changeTargetType: function() {
                let _targetType = $that.editBackupInfo.targetTypeCd;
                $that.editBackupInfo.targetId = '';
                if (!_targetType) {
                    return;
                }

                if (_targetType == '001') {
                    $that._listEditFtp();
                } else if (_targetType == '002') {
                    $that._listEditOss();
                } else {
                    $that._listEditDbs('target');
                }
            },
            _changeExecFlag: function() {
                let _execFlag = $that.editBackupInfo.execFlag;
                if (_execFlag != '2') {
                    $that.editBackupInfo.execFlag = '2';
                } else {
                    $that.editBackupInfo.execFlag = '1';
                }
            },
            updateBackupInfo: function() {
                if (!vc.component.editBackupValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                if ($that.editBackupInfo.srcId == $that.editBackupInfo.targetId) {
                    vc.toast('备份源和目标不能一样');
                    return;
                }

                vc.http.apiPost(
                    '/resources/updateBackUp',
                    JSON.stringify(vc.component.editBackupInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editBackupModel').modal('hide');
                            vc.component.clearEditBackupInfo();
                            vc.emit('backupManage', 'listBackup', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearEditBackupInfo: function() {
                vc.component.editBackupInfo = {
                    id: '',
                    name: '',
                    execTime: '',
                    typeCd: '',
                    srcId: '',
                    srcObject: '',
                    targetId: '',
                    targetTypeCd: '',
                    srcs: [],
                    targets: [],
                    execFlag: '1'
                };
            },
            _listEditHosts: function() {
                let param = {
                    params: {
                        page: 1,
                        row: 100
                    }
                };
                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);
                        $that.editBackupInfo.srcs = _hostManageInfo.data;
                        $that.editBackupInfo.srcs.forEach(item => {
                            item.id = item.hostId
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _listEditDbs: function(_flag) {
                let param = {
                    params: {
                        page: 1,
                        row: 100
                    }
                };
                //发送get请求
                vc.http.apiGet('/resources/getDb',
                    param,
                    function(json, res) {
                        let _dbManageInfo = JSON.parse(json);
                        if (_flag == 'src') {
                            $that.editBackupInfo.srcs = _dbManageInfo.data;
                            $that.editBackupInfo.srcs.forEach(item => {
                                item.id = item.dbId
                            });
                        } else {
                            $that.editBackupInfo.targets = _dbManageInfo.data;
                            $that.editBackupInfo.targets.forEach(item => {
                                item.id = item.dbId
                            });
                        }

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _listEditFtp: function() {
                var param = {
                    params: {
                        page: 1,
                        row: 100
                    }
                };

                //发送get请求
                vc.http.apiGet('/resources/getFtp',
                    param,
                    function(json, res) {
                        let _ftpManageInfo = JSON.parse(json);
                        $that.editBackupInfo.targets = _ftpManageInfo.data;
                        $that.editBackupInfo.targets.forEach(item => {
                            item.id = item.ftpId
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    })
            },
            _listEditOss: function() {
                let param = {
                    params: {
                        page: 1,
                        row: 100
                    }
                };
                vc.http.apiGet('/resources/getOss',
                    param,
                    function(json, res) {
                        let _ossManageInfo = JSON.parse(json);
                        $that.editBackupInfo.targets = _ossManageInfo.data;
                        $that.editBackupInfo.targets.forEach(item => {
                            item.id = item.ossId
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }
        }
    });

})(window.vc);