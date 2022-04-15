(function(vc, vm) {

    vc.extends({
        data: {
            editDnsInfo: {
                dnsId: '',
                dnsPort: '',
                dnsIp:'',
                state: '',
                hosts: [],
                hIds: [],
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editDns', 'openEditDnsModal', function(_params) {
                vc.component.refreshEditDnsInfo();
                $('#editDnsModel').modal('show');
                vc.copyObject(_params, vc.component.editDnsInfo);
                $that._listHosts();
                let _dnsHosts = _params.dnsHosts;

                if (!_dnsHosts) {
                    return;
                }

                _dnsHosts.forEach(item => {
                    $that.editDnsInfo.hIds.push(item.hostId);
                });



            });
        },
        methods: {
            editDnsValidate: function() {
                return vc.validate.validate({
                    editDnsInfo: vc.component.editDnsInfo
                }, {
                    'editDnsInfo.dnsPort': [{
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
                    'editDnsInfo.state': [{
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
                    'editDnsInfo.dnsId': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editDns: function() {
                if (!vc.component.editDnsValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                $that.editDnsInfo.hostIds = $that.editDnsInfo.hIds.join(',')

                vc.http.apiPost(
                    '/dns/updateDns',
                    JSON.stringify(vc.component.editDnsInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editDnsModel').modal('hide');
                            vc.emit('dnsManage', 'listDns', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditDnsInfo: function() {
                vc.component.editDnsInfo = {
                    dnsId: '',
                    dnsPort: '',
                    dnsIp:'',
                    state: '',
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
                        vc.component.editDnsInfo.hosts = _hostManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc, window.vc.component);