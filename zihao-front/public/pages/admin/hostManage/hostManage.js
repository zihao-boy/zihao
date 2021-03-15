/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            hostManageInfo: {
                hosts: [],
                total: 0,
                records: 1,
                moreCondition: false,
                hostId: '',
                conditions: {
                    groupId: '',
                    name: '',
                    ip: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listHosts(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('hostManage', 'listHost', function (_param) {
                vc.component._listHosts(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listHosts(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listHosts: function (_page, _rows) {

                vc.component.hostManageInfo.conditions.page = _page;
                vc.component.hostManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.hostManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function (json, res) {
                        var _hostManageInfo = JSON.parse(json);
                        vc.component.hostManageInfo.total = _hostManageInfo.total;
                        vc.component.hostManageInfo.records = _hostManageInfo.records;
                        vc.component.hostManageInfo.hosts = _hostManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.hostManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddHostModal: function () {
                vc.emit('addHost', 'openAddHostModal', {});
            },
            _openEditHostModel: function (_host) {
                vc.emit('editHost', 'openEditHostModal', _host);
            },
            _openDeleteHostModel: function (_host) {
                vc.emit('deleteHost', 'openDeleteHostModal', _host);
            },
            _queryHostMethod: function () {
                vc.component._listHosts(DEFAULT_PAGE, DEFAULT_ROWS);
            },
            _openMonitorHostAnalysis:function(_host){
                vc.jumpToPage('/index.html#/pages/admin/monitorHostGroupManageAnalysis?hostId='+_host.hostId)
            },
            _moreCondition: function () {
                if (vc.component.hostManageInfo.moreCondition) {
                    vc.component.hostManageInfo.moreCondition = false;
                } else {
                    vc.component.hostManageInfo.moreCondition = true;
                }
            },
            _openSsh:function(_host){
                //获取主机访问token
                var param = {
                    params: {
                        hostId:_host.hostId
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHostToken',
                    param,
                    function (json, res) {
                        let _hostManageInfo = JSON.parse(json);
                        let _zihaoToken = _hostManageInfo.data;
                        window.open("/webshell/console.html?hostId="+_host.hostId+"&zihaoToken="+_zihaoToken,'_blank')
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }


        }
    });
})(window.vc);
