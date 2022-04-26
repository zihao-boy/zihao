(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addHostInfo: {
                hostId: '',
                groupId: '',
                name: '',
                ip: '',
                username: 'root',
                passwd: '',
                cpu: '0',
                mem: '0',
                disk: '0',
                osName:'linux',
                hostGroups:[]

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addHost', 'openAddHostModal', function () {
                $('#addHostModel').modal('show');
                $that.loadAddHostHostGroup();
            });
        },
        methods: {
            addHostValidate() {
                return vc.validate.validate({
                    addHostInfo: vc.component.addHostInfo
                }, {
                    'addHostInfo.groupId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "主机组不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "主机组格式错误"
                        },
                    ],
                    'addHostInfo.name': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "主机名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "主机名称太长"
                        },
                    ],
                    'addHostInfo.ip': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "主机IP不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "主机IP太长"
                        },
                    ],
                    'addHostInfo.username': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "用户名不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "用户名太长"
                        },
                    ],
                    'addHostInfo.passwd': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "密码不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "密码太长"
                        },
                    ],
                    'addHostInfo.cpu': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "CPU核数不能为空"
                        },
                        {
                            limit: "num",
                            param: "",
                            errInfo: "cpu 不是有效数字"
                        },
                    ],
                    'addHostInfo.mem': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "内存不能为空"
                        },
                        {
                            limit: "money",
                            param: "",
                            errInfo: "内存格式错误"
                        },
                    ],
                    'addHostInfo.disk': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "磁盘大小不能为空"
                        },
                        {
                            limit: "num",
                            param: "",
                            errInfo: "磁盘大小不是有效数字"
                        },
                    ],
                });
            },
            saveHostInfo: function () {
                if (!vc.component.addHostValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addHostInfo);
                    $('#addHostModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/host/saveHost',
                    JSON.stringify(vc.component.addHostInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addHostModel').modal('hide');
                            vc.component.clearAddHostInfo();
                            vc.emit('hostManage', 'listHost', {});

                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            clearAddHostInfo: function () {
                vc.component.addHostInfo = {
                    groupId: '',
                    name: '',
                    ip: '',
                    username: 'root',
                    passwd: '',
                    cpu: '',
                    mem: '',
                    disk: '',
                    hostGroups:[],
                    osName:'linux',

                };
            },
            _changeOs:function(){
                let _osName = $that.addHostInfo.osName;
                if(_osName == 'linux'){
                    $that.addHostInfo.username = 'root';
                }else{
                    $that.addHostInfo.username = 'Administrator';
                }
            },
            loadAddHostHostGroup:function(){
                var param = {
                    params: {
                        page:1,
                        row:50
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHostGroup',
                    param,
                    function (json, res) {
                        var _hostGroupManageInfo = JSON.parse(json);
                        vc.component.addHostInfo.hostGroups = _hostGroupManageInfo.data;
                        
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }
        }
    });

})(window.vc);
