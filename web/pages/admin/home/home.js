/**
    入驻小区
**/
(function(vc) {
    vc.extends({
        data: {
            homeInfo: {
                hostCount: 0,
                cpuCount: 0,
                memCount: 0,
                diskCount: 0,
                appCount: 0,
                dockerCount: 0,
                pages: [{
                    name: '主机',
                    url: '/index.html#/pages/admin/hostManage',
                    icon: '/img/hosts.png'
                }, {
                    name: '应用',
                    url: '/index.html#/pages/admin/appServiceManage',
                    icon: '/img/apps.png'
                }, {
                    name: '本地镜像',
                    url: '/index.html#/pages/admin/businessImagesManage',
                    icon: '/img/images.png'
                }, {
                    name: '构建版本',
                    url: '/index.html#/pages/admin/appVersionJobManage',
                    icon: '/img/builder.png'
                }, {
                    name: '业务包',
                    url: '/index.html#/pages/admin/businessPackageManage',
                    icon: '/img/jar.png'
                }, {
                    name: 'dockerfile',
                    url: '/index.html#/pages/admin/businessDockerfileManage',
                    icon: '/img/dockerfile.png'
                }, {
                    name: 'shell',
                    url: 'action',
                    action: 'shell',
                    icon: '/img/shell.png'
                }, {
                    name: '文件',
                    url: 'action',
                    action: 'file',
                    icon: '/img/folder.png'
                }, {
                    name: '快速部署',
                    url: '/index.html#/pages/admin/fasterDeploy',
                    icon: '/img/fasterDeploy.png'
                }, {
                    name: 'mysql客户端',
                    url: '/index.html#/pages/admin/mysqlClient',
                    icon: '/img/mysql.png'
                }],
                apps: [],
                action: '',
            }
        },
        _initMethod: function() {
            $that._loadPlatfromData();
        },
        _initEvent: function() {
            vc.on('home', 'chooseHost', function(_param) {
                $that._doShellOrFile(_param);
            });
        },
        methods: {
            newFlow: function(item) {
                if (item.url == 'action') {
                    $that.homeInfo.action = item.action;
                    vc.emit('chooseHost', 'openChooseHostModel', {})
                    return;
                }
                vc.jumpToPage(item.url);
            },
            _loadPlatfromData: function() {
                //获取主机访问token
                let param = {
                    params: {
                        platform: 'zihao'
                    }
                };
                //发送get请求
                vc.http.apiGet('/home/platformData',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);
                        if (_hostManageInfo.code != 0) {
                            console.log(_hostManageInfo);
                            return;
                        }
                        vc.copyObject(_hostManageInfo.data, $that.homeInfo);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _doShellOrFile: function(_param) {
                if ($that.homeInfo.action == 'file') {
                    vc.jumpToPage('/index.html#/pages/admin/fileManager?hostId=' + _param.hostId);
                    return;
                }
                //shell
                let param = {
                    params: {
                        hostId: _param.hostId
                    }
                };
                //发送get请求
                vc.http.apiGet('/host/getHostToken',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);
                        let _zihaoToken = _hostManageInfo.data;
                        window.open("/webshell/console.html?hostId=" + _param.hostId + "&zihaoToken=" + _zihaoToken, '_blank')
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }
        }
    });
})(window.vc);