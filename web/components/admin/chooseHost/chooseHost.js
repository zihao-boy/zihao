(function(vc) {
    vc.extends({
        propTypes: {
            emitChooseHost: vc.propTypes.string,
            emitLoadData: vc.propTypes.string
        },
        data: {
            chooseHostInfo: {
                hosts: [],
                _currentHostName: '',
                condition: {
                    name: '',
                    hostId: '',
                    groupId: ''
                }
            }
        },
        _initMethod: function() {},
        _initEvent: function() {
            vc.on('chooseHost', 'openChooseHostModel', function(_param) {
                $('#chooseHostModel').modal('show');
                vc.component._refreshChooseHostInfo();
                vc.copyObject(_param, $that.chooseHostInfo.condition);
                vc.component._loadAllHostInfo(1, 10, '');
            });
        },
        methods: {
            _loadAllHostInfo: function(_page, _row, _name) {
                let _condition = $that.chooseHostInfo.condition;
                _condition.page = _page;
                _condition.row = _row;
                let param = {
                    params: _condition
                };

                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function(json) {
                        var _hostInfo = JSON.parse(json);
                        vc.component.chooseHostInfo.hosts = _hostInfo.data;
                    },
                    function() {
                        console.log('请求失败处理');
                    }
                );
            },
            chooseHost: function(_host) {
                if (_host.hasOwnProperty('name')) {
                    _host.hostName = _host.name;
                }
                vc.emit($props.emitChooseHost, 'chooseHost', _host);
                vc.emit($props.emitLoadData, 'listHostData', {
                    hostId: _host.hostId
                });
                $('#chooseHostModel').modal('hide');
            },
            queryHosts: function() {
                vc.component._loadAllHostInfo(1, 10, vc.component.chooseHostInfo._currentHostName);
            },
            _refreshChooseHostInfo: function() {
                vc.component.chooseHostInfo._currentHostName = "";
            }
        }

    });
})(window.vc);