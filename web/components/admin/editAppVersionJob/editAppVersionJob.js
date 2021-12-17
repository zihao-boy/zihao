(function(vc) {
    vc.extends({
        data: {
            editAppVersionJobInfo: {
                jobId: '',
                jobName: '',
                jobShell: '',
                gitUrl: '',
                gitUsername: '',
                gitPasswd: '',
                workDir: '',
                appVersionJobImages: []
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editAppVersionJob', 'openEditAppVersionJobModal', function(_param) {
                $that.clearEditAppVersionJobInfo();
                $('#editAppVersionJobModel').modal('show');
                vc.copyObject(_param, $that.editAppVersionJobInfo);
                $that._loadImages();
            });
        },
        methods: {
            editAppVersionJobValidate() {
                return vc.validate.validate({
                    editAppVersionJobInfo: vc.component.editAppVersionJobInfo
                }, {
                    'editAppVersionJobInfo.jobName': [{
                            limit: "required",
                            param: "",
                            errInfo: "名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "名称太长"
                        },
                    ],
                    'editAppVersionJobInfo.jobShell': [{
                            limit: "required",
                            param: "",
                            errInfo: "脚本不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "1024",
                            errInfo: "脚本错误"
                        },
                    ],
                    'editAppVersionJobInfo.gitUrl': [{
                            limit: "required",
                            param: "",
                            errInfo: "git地址不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "512",
                            errInfo: "git地址太长"
                        },
                    ],
                    'editAppVersionJobInfo.workDir': [{
                            limit: "required",
                            param: "",
                            errInfo: "工作目录不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "工作目录太长"
                        },
                    ],
                });
            },
            _loadImages: function() {
                let param = {
                    params: {
                        page: 1,
                        row: 100,
                        jobId: $that.editAppVersionJobInfo.jobId
                    }
                };
                //发送get请求
                vc.http.apiGet('/appVersion/getAppVersionJobImages',
                    param,
                    function(json, res) {
                        let _appVersionJobManageInfo = JSON.parse(json);
                        vc.component.editAppVersionJobInfo.appVersionJobImages = _appVersionJobManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _editAppVersionJobInfo: function() {
                if (!vc.component.editAppVersionJobValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/appVersion/updateAppVersionJob',
                    JSON.stringify(vc.component.editAppVersionJobInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppVersionJobModel').modal('hide');
                            vc.component.clearEditAppVersionJobInfo();
                            vc.emit('appVersionJobManage', 'listAppVersionJob', {});

                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            clearEditAppVersionJobInfo: function() {
                vc.component.editAppVersionJobInfo = {
                    jobId: '',
                    jobName: '',
                    jobShell: '',
                    gitUrl: '',
                    gitUsername: '',
                    gitPasswd: '',
                    workDir: '',
                    appVersionJobImages: []
                };
            },
            _addEditJobPlan: function() {
                if (!$that.editAppVersionJobInfo.appVersionJobImages) {
                    $that.editAppVersionJobInfo.appVersionJobImages = []
                }
                $that.editAppVersionJobInfo.appVersionJobImages.push({
                    packageUrl: '',
                    businessPackageId: '',
                    businessPackageName: '',
                    businessDockerfileId: '',
                    businessDockerfileName: ''
                })
            },
            _deleteEditJobPlan: function(_index) {
                $that.editAppVersionJobInfo.appVersionJobImages.splice(_index, 1);
            },
            _appEditChooseBusinessPackage: function(_data) {
                vc.emit('chooseBusinessPackage', 'openChooseBusinessPackageModel', _data);
            },
            _appEditChooseDockerfile: function(_data) {
                vc.emit('chooseDockerfile', 'openChooseDockerfileModel', _data);
            },
            _goEditBack: function() {
                vc.emit('appVersionJobManage', 'listAppVersionJob', {});
            }
        }
    });

})(window.vc);