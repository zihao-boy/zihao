/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appServiceControlPortInfo: {
                appServiceControlPorts: [],
                total: 0,
                records: 1,
                asId: '',
            }
        },
        _initMethod: function () {
        },
        _initEvent: function () {

            vc.on('appServiceControlPort', 'switch', function (_param) {
                if (_param.asId == '') {
                    return;
                }
                vc.copyObject(_param, $that.appServiceControlPortInfo)
                vc.component._listappServiceControlPorts(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('appServiceControlPort', 'paginationPlus', 'page_event', function (_currentPage) {
                vc.component._listappServiceControlPorts(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listappServiceControlPorts: function (_page, _rows) {

                var param = {
                    params: {
                        page: _page,
                        row: _rows,
                        asId: $that.appServiceControlPortInfo.asId
                    }
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppServicePort',
                    param,
                    function (json, res) {
                        var _appServiceControlPortsInfo = JSON.parse(json);
                        vc.component.appServiceControlPortInfo.total = _appServiceControlPortsInfo.total;
                        vc.component.appServiceControlPortInfo.records = _appServiceControlPortsInfo.records;
                        vc.component.appServiceControlPortInfo.appServiceControlPorts = _appServiceControlPortsInfo.data;
                        vc.emit('appServiceControlPort', 'paginationPlus', 'init', {
                            total: vc.component.appServiceControlPortInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddPortModal:function(){
                vc.emit('addAppServicePort', 'openAddAppServicePortModal',$that.appServiceControlPortInfo);
            },
            _openDeletePortModal:function(_port){
                vc.emit('deleteAppServicePort','openDeleteAppServicePortModal',_port); 
            }
        }
    });
})(window.vc);
