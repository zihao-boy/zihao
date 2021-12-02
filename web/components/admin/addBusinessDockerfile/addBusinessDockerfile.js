(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addBusinessDockerfileInfo: {
                id: '',
                name: '',
                dockerfile: '',
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addBusinessDockerfile', 'openAddBusinessDockerfileModal', function() {
                $that._initDockerfile();
                $('#addBusinessDockerfileModel').modal('show');
            });
        },
        methods: {
            addBusinessDockerfileValidate() {
                return vc.validate.validate({
                    addBusinessDockerfileInfo: vc.component.addBusinessDockerfileInfo
                }, {
                    'addBusinessDockerfileInfo.name': [{
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
                    'addBusinessDockerfileInfo.dockerfile': [{
                        limit: "required",
                        param: "",
                        errInfo: "内容不能为空"
                    }],
                });
            },
            saveBusinessDockerfileInfo: function() {
                if (!vc.component.addBusinessDockerfileValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addBusinessDockerfileInfo);
                    $('#addBusinessDockerfileModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/soft/saveBusinessDockerfile',
                    JSON.stringify(vc.component.addBusinessDockerfileInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addBusinessDockerfileModel').modal('hide');
                            vc.component.clearAddBusinessDockerfileInfo();
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
            clearAddBusinessDockerfileInfo: function() {
                vc.component.addBusinessDockerfileInfo = {
                    name: '',
                    dockerfile: '',

                };
                $that._initDockerfile();
            },
            _initDockerfile: function() {
                $that.addBusinessDockerfileInfo.dockerfile = "" +
                    "# 指定源于一个基础镜像\n" +
                    "FROM registry.cn-beijing.aliyuncs.com/sxd/ubuntu-java8:1.0\n" +
                    "# 维护者/拥有者\n" +
                    "MAINTAINER xxx <xxx@xx.com>\n" +
                    "# 从宿主机上传文件 ，这里上传一个脚本，\n" +
                    "# 具体脚本可以去业务包上传后复制路径\n" +
                    "ADD bin/start_api.sh /root/\n" +
                    "# 从宿主机上传文件 ，这里上传一个业务文件，\n" +
                    "# 具体脚本可以去业务包上传后复制路径\n" +
                    "ADD xxx /root/\n" +
                    "# 容器内执行相应指令\n" +
                    "RUN chmod u+x /root/start_api.sh\n" +
                    "# 运行命令\n" +
                    "# CMD <command>   or CMD [<command>]\n" +
                    "# 整个Dockerfile 中只能有一个,多个会被覆盖的\n" +
                    "CMD [\"/root/start_api.sh\", \"dev -Dcache -Xms512m -Xmx512m\"]\n";
            }
        }
    });

})(window.vc);