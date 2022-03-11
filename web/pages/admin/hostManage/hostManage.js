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
            _openFileManager:function(_host){
                vc.jumpToPage('/index.html#/pages/admin/fileManager?hostId='+_host.hostId)
            },
            _moreCondition: function () {
                if (vc.component.hostManageInfo.moreCondition) {
                    vc.component.hostManageInfo.moreCondition = false;
                } else {
                    vc.component.hostManageInfo.moreCondition = true;
                }
            },
            
            _openConsole:function(_host){
                vc.jumpToPage('/index.html#/pages/admin/hostDetailManage?hostId='+_host.hostId)
            },
            _controlHost:function(_host){
                console.log(JSON.stringify(_host))
                vc.http.apiPost(
                    '/host/controlHost',
                    JSON.stringify(_host),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
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
            _getStateName:function(_state){
                if(_state == '1001'){
                    return '未管理';
                }else if(_state == '2002'){
                    return '正常';
                }else{
                    return '处理中';
                }
            },
            _getOnline:function(_time){
                if(!_time){
                    return "离线";
                }
                _time = _time.replace('T'," ");
                _time = _time.replace('Z',"");
                let _date = new Date(_time);
                if(new Date().getTime() - _date.getTime() < 2 * 60 * 1000){
                    return "在线"
                }

                return "离线"
            },
            _openHostShell:function(_hostId){
                //获取主机访问token
                let param = {
                    params: {
                        hostId: _hostId.hostId
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHostToken',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);
                        let _zihaoToken = _hostManageInfo.data;
                        window.open("/webshell/console.html?hostId=" +_hostId.hostId + "&zihaoToken=" + _zihaoToken, '_blank')
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }


        }
    });
})(window.vc);
