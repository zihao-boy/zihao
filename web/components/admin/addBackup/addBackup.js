(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addBackupInfo: {
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
            vc.on('addBackup', 'openAddBackupModal', function() {
                $('#addBackupModel').modal('show');
            });
        },
        methods: {
            addBackupValidate() {
                return vc.validate.validate({
                    addBackupInfo: vc.component.addBackupInfo
                }, {
                    'addBackupInfo.name': [{
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
                    'addBackupInfo.execTime': [{
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
                    'addBackupInfo.typeCd': [{
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
                    'addBackupInfo.srcId': [{
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
                    'addBackupInfo.srcObject': [{
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
                    'addBackupInfo.targetId': [{
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




                });
            },
            _changeSrcType: function() {
                let _srcType = $that.addBackupInfo.typeCd;
                $that.addBackupInfo.srcId = '';
                if (!_srcType) {
                    return;
                }

                if (_srcType == '10001') {
                    $that.addBackupInfo.srcObject = "*";
                    $that._listAddDbs('src');
                } else {
                    $that.addBackupInfo.srcObject = "";
                    $that._listAddHosts();
                }

            },
            _changeTargetType: function() {
                let _targetType = $that.addBackupInfo.targetTypeCd;
                $that.addBackupInfo.targetId = '';
                if (!_targetType) {
                    return;
                }

                if (_targetType == '001') {
                    $that._listAddFtp();
                } else if (_targetType == '002') {
                    $that._listAddOss();
                } else {
                    $that._listAddDbs('target');
                }
            },
            _changeExecFlag: function() {
                let _execFlag = $that.addBackupInfo.execFlag;
                if (_execFlag != '2') {
                    $that.addBackupInfo.execFlag = '2';
                } else {
                    $that.addBackupInfo.execFlag = '1';
                }
            },
            saveBackupInfo: function() {
                if (!vc.component.addBackupValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addBackupInfo);
                    $('#addBackupModel').modal('hide');
                    return;
                }

                if ($that.addBackupInfo.srcId == $that.addBackupInfo.targetId) {
                    vc.toast('备份源和目标不能一样');
                    return;
                }

                vc.http.apiPost(
                    '/resources/saveBackUp',
                    JSON.stringify(vc.component.addBackupInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addBackupModel').modal('hide');
                            vc.component.clearAddBackupInfo();
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
            clearAddBackupInfo: function() {
                vc.component.addBackupInfo = {
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
            _listAddHosts: function() {
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
                        $that.addBackupInfo.srcs = _hostManageInfo.data;
                        $that.addBackupInfo.srcs.forEach(item => {
                            item.id = item.hostId
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _listAddDbs: function(_flag) {
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
                            $that.addBackupInfo.srcs = _dbManageInfo.data;
                            $that.addBackupInfo.srcs.forEach(item => {
                                item.id = item.dbId
                            });
                        } else {
                            $that.addBackupInfo.targets = _dbManageInfo.data;
                            $that.addBackupInfo.targets.forEach(item => {
                                item.id = item.dbId
                            });
                        }

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _listAddFtp: function() {
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
                        $that.addBackupInfo.targets = _ftpManageInfo.data;
                        $that.addBackupInfo.targets.forEach(item => {
                            item.id = item.ftpId
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    })
            },
            _listAddOss: function() {
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
                        $that.addBackupInfo.targets = _ossManageInfo.data;
                        $that.addBackupInfo.targets.forEach(item => {
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