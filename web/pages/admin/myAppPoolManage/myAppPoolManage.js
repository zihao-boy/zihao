/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            myAppPoolManageInfo: {
                businessImagess: [],
                appPublishers: [],
                total: 0,
                records: 1,
                moreCondition: false,
                id: '',
                conditions: {
                    publisherId: "",
                    name: '',
                    state: ''
                }
            }
        },
        _initMethod: function() {
            $that._listAppPublishers();
        },
        _initEvent: function() {

            vc.on('myAppPoolManage', 'listBusinessImages', function(_param) {
                vc.component._listBusinessImagess(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listBusinessImagess(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listBusinessImagess: function(_page, _rows) {

                vc.component.myAppPoolManageInfo.conditions.page = _page;
                vc.component.myAppPoolManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.myAppPoolManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/soft/getMyAppPool',
                    param,
                    function(json, res) {
                        var _myAppPoolManageInfo = JSON.parse(json);
                        vc.component.myAppPoolManageInfo.total = _myAppPoolManageInfo.total;
                        vc.component.myAppPoolManageInfo.records = _myAppPoolManageInfo.records;
                        vc.component.myAppPoolManageInfo.businessImagess = _myAppPoolManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.myAppPoolManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _queryBusinessImagesMethod: function() {
                vc.component._listBusinessImagess(DEFAULT_PAGE, DEFAULT_ROWS);
            },
            _moreCondition: function() {
                if (vc.component.myAppPoolManageInfo.moreCondition) {
                    vc.component.myAppPoolManageInfo.moreCondition = false;
                } else {
                    vc.component.myAppPoolManageInfo.moreCondition = true;
                }
            },
            _viewImages: function(_businessImages) {
                vc.emit('viewPoolsImages', 'openViewPoolsImagesModel', _businessImages.zihaoAppImagesDtos);
            },
            _installModel: function(_businessImages) {
                vc.emit('installImages', 'openInstallImagesModal', _businessImages);
            },
            _uninstallModel: function(_businessImages) {
                vc.emit('uninstallImages', 'openUninstallImagesModal', _businessImages);
            },
            _goBack: function() {
                vc.goBack();
            },
            getState: function(_app) {
                if (_app.state == '001') {
                    return "正常";
                } else if (_app.state == '002') {
                    return "待审核";
                } else if (_app.state == '003') {
                    return "审核失败";
                } else if (_app.state == '004') {
                    return "下架";
                }

                return "未知"
            },
            _listAppPublishers: function(_page, _rows) {
                var param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/soft/getAppPublisher',
                    param,
                    function(json, res) {
                        var _appPublisherManageInfo = JSON.parse(json);
                        vc.component.myAppPoolManageInfo.appPublishers = _appPublisherManageInfo.data;
                        if (_appPublisherManageInfo.data.length < 1) {
                            return;
                        }
                        $that.myAppPoolManageInfo.conditions.publisherId = _appPublisherManageInfo.data[0].publisherId;
                        vc.component._listBusinessImagess(DEFAULT_PAGE, DEFAULT_ROWS);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },

        }
    });
})(window.vc);