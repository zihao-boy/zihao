/**
 入驻小区
 **/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    var TEMP_SEARCH = "simplifyAcceptanceSearch";
    vc.extends({
        data: {
            hostDetailManageInfo: {
                _currentTab: 'hostEvents',
                hostId: '',
                groupId: '',
                groupName: '',
                name: '',
                ip: '',
                username: '',
                passwd: '',
                cpu: '',
                mem: '',
                disk: '',
                useCpu: '',
                useMem: '',
                useDisk: '',
            }
        },
        _initMethod: function() {

            $that.hostDetailManageInfo.hostId = vc.getParam('hostId');

            $that._listHostDetail();

        },
        _initEvent: function() {
            vc.on('simplifyAcceptance', 'chooseRoom', function(_room) {
                vc.copyObject(_room, $that.hostDetailManageInfo);
                vc.emit('simplifyRoomFee', 'switch', $that.hostDetailManageInfo)
            });
        },
        methods: {
            changeTab: function(_tab) {
                $that.hostDetailManageInfo._currentTab = _tab;
                vc.emit(_tab, 'switch', {
                    hostId: $that.hostDetailManageInfo.hostId
                })
            },
            _clearData: function() {
                $that.hostDetailManageInfo = {
                    _currentTab: 'hostContainers',
                    hostId: '',
                    groupId: '',
                    groupName: '',
                    name: '',
                    ip: '',
                    username: '',
                    passwd: '',
                    cpu: '',
                    mem: '',
                    disk: '',
                    useCpu: '',
                    useMem: '',
                    useDisk: '',
                }
            },
            _listHostDetail: function() {

                let param = {
                    params: {
                        page: 1,
                        row: 1,
                        hostId: $that.hostDetailManageInfo.hostId
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);

                        vc.copyObject(_hostManageInfo.data[0], $that.hostDetailManageInfo)

                        $that.changeTab($that.hostDetailManageInfo._currentTab);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openSsh: function() {
                //获取主机访问token
                var param = {
                    params: {
                        hostId: $that.hostDetailManageInfo.hostId
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHostToken',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);
                        let _zihaoToken = _hostManageInfo.data;
                        window.open("/webshell/console.html?hostId=" + $that.hostDetailManageInfo.hostId + "&zihaoToken=" + _zihaoToken, '_blank')
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _goBack: function() {
                vc.goBack();
            },
            _openFileManager: function() {
                vc.jumpToPage('/index.html#/pages/admin/fileManager?hostId=' + $that.hostDetailManageInfo.hostId)
            },

        }
    });
})(window.vc);