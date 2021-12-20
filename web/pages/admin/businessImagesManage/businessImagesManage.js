/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            businessImagesManageInfo: {
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
                    version: ''
                }
            }
        },
        _initMethod: function() {
            vc.component._listBusinessImagess(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('businessImagesManage', 'listBusinessImages', function(_param) {
                vc.component._listBusinessImagess(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listBusinessImagess(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listBusinessImagess: function(_page, _rows) {

                vc.component.businessImagesManageInfo.conditions.page = _page;
                vc.component.businessImagesManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.businessImagesManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/soft/getBusinessImages',
                    param,
                    function(json, res) {
                        var _businessImagesManageInfo = JSON.parse(json);
                        vc.component.businessImagesManageInfo.total = _businessImagesManageInfo.total;
                        vc.component.businessImagesManageInfo.records = _businessImagesManageInfo.records;
                        vc.component.businessImagesManageInfo.businessImagess = _businessImagesManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.businessImagesManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddBusinessImagesModal: function() {
                vc.emit('addBusinessImages', 'openAddBusinessImagesModal', {});
            },
            _openEditBusinessImagesModel: function(_businessImages) {
                vc.emit('editBusinessImages', 'openEditBusinessImagesModal', _businessImages);
            },
            _openDeleteBusinessImagesModel: function(_businessImages) {
                vc.emit('deleteBusinessImages', 'openDeleteBusinessImagesModal', _businessImages);
            },
            _queryBusinessImagesMethod: function() {
                vc.component._listBusinessImagess(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.businessImagesManageInfo.moreCondition) {
                    vc.component.businessImagesManageInfo.moreCondition = false;
                } else {
                    vc.component.businessImagesManageInfo.moreCondition = true;
                }
            },
            _getImagesType: function(_imagesType) {
                if (_imagesType == "1") {
                    return "远程镜像";
                } else if (_imagesType == "2") {
                    return "dockfile生成镜像"
                }

                return "导入镜像";
            }


        }
    });
})(window.vc);