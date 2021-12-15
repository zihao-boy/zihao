(function(vc) {
    vc.extends({
        data: {
            addAppVersionJobInfo: {
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
            vc.on('addAppVersionJob', 'openAddAppVersionJobModal', function() {
                $('#addAppVersionJobModel').modal('show');
            });
        },
        methods: {
            addAppVersionJobValidate() {
                return vc.validate.validate({
                    addAppVersionJobInfo: vc.component.addAppVersionJobInfo
                }, {
                    'addAppVersionJobInfo.jobName': [{
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
                    'addAppVersionJobInfo.jobShell': [{
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
                    'addAppVersionJobInfo.gitUrl': [{
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
                    'addAppVersionJobInfo.gitUsername': [{
                            limit: "required",
                            param: "",
                            errInfo: "git名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "git名称太长"
                        },
                    ],
                    'addAppVersionJobInfo.gitPasswd': [{
                            limit: "required",
                            param: "",
                            errInfo: "git密码不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "git密码太长"
                        },
                    ],
                    'addAppVersionJobInfo.workDir': [{
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
            saveAppVersionJobInfo: function() {
                if (!vc.component.addAppVersionJobValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/appVersion/saveAppVersionJob',
                    JSON.stringify(vc.component.addAppVersionJobInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppVersionJobModel').modal('hide');
                            vc.component.clearAddAppVersionJobInfo();
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
            clearAddAppVersionJobInfo: function() {
                vc.component.addAppVersionJobInfo = {
                    jobName: '',
                    jobShell: '',
                    gitUrl: '',
                    gitUsername: '',
                    gitPasswd: '',
                    workDir: '',
                    appVersionJobImages: []
                };
            },
            _addJobPlan: function() {
                $that.addAppVersionJobInfo.appVersionJobImages.push({
                    packageUrl: '',
                    businessPackageId: '',
                    businessPackageName: '',
                    businessDockerfileId: '',
                    businessDockerfileName: ''
                })
            },
            _deleteJobPlan: function(_index) {
                $that.addAppVersionJobInfo.appVersionJobImages.splice(_index, 1);
            },
            _appChooseBusinessPackage: function(_data) {
                vc.emit('chooseBusinessPackage', 'openChooseBusinessPackageModel', _data);
            },
            _appChooseDockerfile: function(_data) {
                vc.emit('chooseDockerfile', 'openChooseDockerfileModel', _data);
            },
            _goBack: function() {
                vc.emit('appVersionJobManage', 'listAppVersionJob', {});

            }
        }
    });

})(window.vc);