(function(vc) {
    vc.extends({
        data: {
            chooseHostContainerInfo: {
                containers: [],
                asId:''
            }
        },
        _initMethod: function() {},
        _initEvent: function() {
            vc.on('chooseHostContainer', 'openChooseHostContainerModel', function(_param) {
                $('#chooseHostContainerModel').modal('show');
                vc.copyObject(_param, $that.chooseHostContainerInfo);
                vc.component._loadAllHostContainerInfo();
            });
        },
        methods: {
            _loadAllHostContainerInfo: function() {
                var param = {
                    params: {
                        page: 1,
                        row: 100,
                        asId: $that.chooseHostContainerInfo.asId
                    }
                };
                //发送get请求
                vc.http.apiGet('/appService/getAppServiceContainer',
                    param,
                    function(json, res) {
                        let _appServiceControlContainersInfo = JSON.parse(json);
                        $that.chooseHostContainerInfo.containers = _appServiceControlContainersInfo.data;
                        
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            chooseHostContainer: function(_host) {
                $('#chooseHostContainerModel').modal('hide');

                 //获取主机访问token
                 let param = {
                    params: {
                        hostId: _host.hostId
                    }
                };
                //发送get请求
                vc.http.apiGet('/host/getHostToken',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);
                        let _zihaoToken = _hostManageInfo.data;
                        window.open("/webshell/console.html?hostId=" +
                            _container.hostId +
                            "&zihaoToken=" + _zihaoToken +
                            "&val=log&command=" + _container.dockerContainerId,
                            '_blank')
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            queryHosts: function() {
                vc.component._loadAllHostContainerInfo();
            },
            _refreshChooseHostContainerInfo: function() {
            }
        }

    });
})(window.vc);