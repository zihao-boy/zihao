(function(vc, vm) {

    vc.extends({
        data: {
            editInnerNetInfo: {
                innerNetId: '',
                innerNetPort: '',
                tun: '',
                tunName: '',
                dns:'',
                protocol:'tcp',
                state: '',
                hosts: [],
                hIds: [],
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editInnerNet', 'openEditInnerNetModal', function(_params) {
                vc.component.refreshEditInnerNetInfo();
                $('#editInnerNetModel').modal('show');
                vc.copyObject(_params, vc.component.editInnerNetInfo);
                $that._listHosts();
                let _innerNetHosts = _params.innerNetHosts;

                if (!_innerNetHosts) {
                    return;
                }

                _innerNetHosts.forEach(item => {
                    $that.editInnerNetInfo.hIds.push(item.hostId);
                });



            });
        },
        methods: {
            editInnerNetValidate: function() {
                return vc.validate.validate({
                    editInnerNetInfo: vc.component.editInnerNetInfo
                }, {
                    'editInnerNetInfo.innerNetPort': [{
                            limit: "required",
                            param: "",
                            errInfo: "端口不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "端口不能超过64"
                        },
                    ],
                    'editInnerNetInfo.tun': [{
                            limit: "required",
                            param: "",
                            errInfo: "网段不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "网段不能超过64"
                        },
                    ],
                    'editInnerNetInfo.state': [{
                            limit: "required",
                            param: "",
                            errInfo: "状态不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "12",
                            errInfo: "状态不能超过12"
                        },
                    ],
                    'editInnerNetInfo.innerNetId': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editInnerNet: function() {
                if (!vc.component.editInnerNetValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                $that.editInnerNetInfo.hostIds = $that.editInnerNetInfo.hIds.join(',')

                vc.http.apiPost(
                    '/innerNet/updateInnerNet',
                    JSON.stringify(vc.component.editInnerNetInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editInnerNetModel').modal('hide');
                            vc.emit('innerNetManage', 'listInnerNet', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditInnerNetInfo: function() {
                vc.component.editInnerNetInfo = {
                    innerNetId: '',
                    innerNetPort: '',
                tun: '',
                tunName: '',
                dns:'',
                protocol:'tcp',
                    state: '',
                    hosts: [],
                    hIds: [],
                }
            },
            _listHosts: function(_page, _rows) {

                let param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);
                        vc.component.editInnerNetInfo.hosts = _hostManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc, window.vc.component);