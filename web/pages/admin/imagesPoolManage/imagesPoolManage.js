/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            imagesPoolManageInfo: {
                businessImagess: [],
                total: 0,
                records: 1,
                moreCondition: false,
                id: '',
                conditions: {
                    id: '',
                    name: '',
                    imagesType: '',
                    imagesFlag: '',
                }
            }
        },
        _initMethod: function() {
            vc.component._listBusinessImagess(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('imagesPoolManage', 'listBusinessImages', function(_param) {
                vc.component._listBusinessImagess(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listBusinessImagess(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listBusinessImagess: function(_page, _rows) {

                vc.component.imagesPoolManageInfo.conditions.page = _page;
                vc.component.imagesPoolManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.imagesPoolManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/soft/getImagesPool',
                    param,
                    function(json, res) {
                        var _imagesPoolManageInfo = JSON.parse(json);
                        vc.component.imagesPoolManageInfo.total = _imagesPoolManageInfo.total;
                        vc.component.imagesPoolManageInfo.records = _imagesPoolManageInfo.records;
                        vc.component.imagesPoolManageInfo.businessImagess = _imagesPoolManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.imagesPoolManageInfo.records,
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
                if (vc.component.imagesPoolManageInfo.moreCondition) {
                    vc.component.imagesPoolManageInfo.moreCondition = false;
                } else {
                    vc.component.imagesPoolManageInfo.moreCondition = true;
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
            _myPublishSoft: function() {
                vc.jumpToPage('/index.html#/pages/admin/myAppPoolManage')
            },
            _publishSoft: function() {
                vc.jumpToPage('/index.html#/pages/admin/applyPublishApp');
            }

        }
    });
})(window.vc);