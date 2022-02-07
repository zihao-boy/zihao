/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appPublisherManageInfo: {
                appPublishers: [],
                total: 0,
                records: 1,
                moreCondition: false,
                publisherId: '',
                conditions: {
                    publisherId: '',
                    username: '',
                    email: '',
                    phone: '',
                }
            }
        },
        _initMethod: function() {
            vc.component._listAppPublishers(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('appPublisherManage', 'listAppPublisher', function(_param) {
                vc.component._listAppPublishers(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listAppPublishers(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listAppPublishers: function(_page, _rows) {

                vc.component.appPublisherManageInfo.conditions.page = _page;
                vc.component.appPublisherManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.appPublisherManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/soft/getAppPublisher',
                    param,
                    function(json, res) {
                        var _appPublisherManageInfo = JSON.parse(json);
                        vc.component.appPublisherManageInfo.total = _appPublisherManageInfo.total;
                        vc.component.appPublisherManageInfo.records = _appPublisherManageInfo.records;
                        vc.component.appPublisherManageInfo.appPublishers = _appPublisherManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.appPublisherManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddAppPublisherModal: function() {
                vc.emit('addAppPublisher', 'openAddAppPublisherModal', {});
            },
            _openEditAppPublisherModel: function(_appPublisher) {
                vc.emit('editAppPublisher', 'openEditAppPublisherModal', _appPublisher);
            },
            _openDeleteAppPublisherModel: function(_appPublisher) {
                vc.emit('deleteAppPublisher', 'openDeleteAppPublisherModal', _appPublisher);
            },
            _queryAppPublisherMethod: function() {
                vc.component._listAppPublishers(DEFAULT_PAGE, DEFAULT_ROWS);
            },
            _moreCondition: function() {
                if (vc.component.appPublisherManageInfo.moreCondition) {
                    vc.component.appPublisherManageInfo.moreCondition = false;
                } else {
                    vc.component.appPublisherManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);