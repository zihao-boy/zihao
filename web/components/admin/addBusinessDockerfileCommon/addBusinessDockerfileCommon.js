(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addBusinessDockerfileCommonInfo: {
                name: '',
                shellContext: '',
                deployType: '',
                path: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addBusinessDockerfileCommon', 'openAddBusinessDockerfileCommonModal', function() {

            });
        },
        methods: {
            addBusinessDockerfileCommonValidate() {
                return vc.validate.validate({
                    addBusinessDockerfileCommonInfo: vc.component.addBusinessDockerfileCommonInfo
                }, {
                    'addBusinessDockerfileCommonInfo.name': [{
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
                    'addBusinessDockerfileCommonInfo.shellContext': [{
                        limit: "required",
                        param: "",
                        errInfo: "内容不能为空"
                    }],
                });
            },
            saveBusinessDockerfileCommonInfo: function() {
                if (!vc.component.addBusinessDockerfileCommonValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addBusinessDockerfileCommonInfo);
                    $('#addBusinessDockerfileCommonModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/soft/saveBusinessDockerfileCommon',
                    JSON.stringify(vc.component.addBusinessDockerfileCommonInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addBusinessDockerfileCommonModel').modal('hide');
                            vc.component.clearAddBusinessDockerfileCommonInfo();
                            vc.emit('businessDockerfileManage', 'listBusinessDockerfile', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddBusinessDockerfileCommonInfo: function() {
                vc.component.addBusinessDockerfileCommonInfo = {
                    name: '',
                    shellContext: '',
                    deployType: '',
                    path: ''
                };
            },
            _goBack: function() {
                vc.emit('businessDockerfileManage', 'listBusinessDockerfile', {});
            },
            _selectCommonBusinessPackages: function() {
                vc.emit('chooseBusinessPackage', 'openChooseBusinessPackageModel', $that.addBusinessDockerfileCommonInfo);
            },
            _changeShellContext: function() {
                $that.addBusinessDockerfileCommonInfo.shellContext = "";
                if ($that.addBusinessDockerfileCommonInfo.deployType != 'java') {
                    return;
                }
                $that.addBusinessDockerfileCommonInfo.shellContext = "" +
                    "#!/bin/bash\n" +
                    "# 最小内存\n" +
                    "min_mem=$MIN_MEM\n" +
                    "# 最大内存\n" +
                    "max_mem=$MAX_MEM\n" +
                    "# spring boot 配置文件\n" +
                    "active=$ACTIVE\n" +
                    "# java启动脚本\n" +
                    "java -jar -Dspring.profiles.active=$active $min_mem $max_mem /root/这里jar包文件名.jar";
            }
        }
    });

})(window.vc);