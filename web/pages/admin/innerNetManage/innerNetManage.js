/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            innerNetManageInfo: {
                innerNets: [],
                total: 0,
                records: 1,
                moreCondition: false,
                innerNetId: '',
                conditions: {
                    innerNetId: '',
                    innerNetName: '',
                    state: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listInnerNets(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('innerNetManage', 'listInnerNet', function (_param) {
                vc.component._listInnerNets(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listInnerNets(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listInnerNets: function (_page, _rows) {

                vc.component.innerNetManageInfo.conditions.page = _page;
                vc.component.innerNetManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.innerNetManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/innerNet/getInnerNet',
                    param,
                    function (json, res) {
                        var _innerNetManageInfo = JSON.parse(json);
                        vc.component.innerNetManageInfo.total = _innerNetManageInfo.total;
                        vc.component.innerNetManageInfo.records = _innerNetManageInfo.records;
                        vc.component.innerNetManageInfo.innerNets = _innerNetManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.innerNetManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openEditInnerNetModel: function (_innerNet) {
                vc.emit('editInnerNet', 'openEditInnerNetModal', _innerNet);
            },
            _openStartInnerNetModel: function (_innerNet) {
                vc.http.apiPost(
                    '/innerNet/startInnerNet',
                    JSON.stringify(_innerNet),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('innerNetManage', 'listInnerNet', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            _openStopInnerNetModel: function (_innerNet) {
                vc.http.apiPost(
                    '/innerNet/stopInnerNet',
                    JSON.stringify(_innerNet),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('innerNetManage', 'listInnerNet', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            _queryInnerNetMethod: function () {
                vc.component._listInnerNets(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.innerNetManageInfo.moreCondition) {
                    vc.component.innerNetManageInfo.moreCondition = false;
                } else {
                    vc.component.innerNetManageInfo.moreCondition = true;
                }
            },
            _loadInnerNetConfig: function (_innerNet) {
                vc.http.apiPost(
                    '/innerNet/refreshInnerNetConfig',
                    JSON.stringify(_innerNet),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('innerNetManage', 'listInnerNet', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
           


        }
    });
})(window.vc);
