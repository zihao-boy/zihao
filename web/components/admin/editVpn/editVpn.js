(function(vc, vm) {

    vc.extends({
        data: {
            editVpnInfo: {
                vpnId: '',
                vpnPort: '',
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
            vc.on('editVpn', 'openEditVpnModal', function(_params) {
                vc.component.refreshEditVpnInfo();
                $('#editVpnModel').modal('show');
                vc.copyObject(_params, vc.component.editVpnInfo);
                $that._listHosts();
                let _vpnHosts = _params.vpnHosts;

                if (!_vpnHosts) {
                    return;
                }

                _vpnHosts.forEach(item => {
                    $that.editVpnInfo.hIds.push(item.hostId);
                });



            });
        },
        methods: {
            editVpnValidate: function() {
                return vc.validate.validate({
                    editVpnInfo: vc.component.editVpnInfo
                }, {
                    'editVpnInfo.vpnPort': [{
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
                    'editVpnInfo.tun': [{
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
                    'editVpnInfo.state': [{
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
                    'editVpnInfo.vpnId': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editVpn: function() {
                if (!vc.component.editVpnValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                $that.editVpnInfo.hostIds = $that.editVpnInfo.hIds.join(',')

                vc.http.apiPost(
                    '/vpn/updateVpn',
                    JSON.stringify(vc.component.editVpnInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editVpnModel').modal('hide');
                            vc.emit('vpnManage', 'listVpn', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditVpnInfo: function() {
                vc.component.editVpnInfo = {
                    vpnId: '',
                    vpnPort: '',
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
                        vc.component.editVpnInfo.hosts = _hostManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc, window.vc.component);