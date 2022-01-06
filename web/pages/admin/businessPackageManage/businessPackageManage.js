/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            businessPackageManageInfo: {
                businessPackages: [],
                total: 0,
                records: 1,
                moreCondition: false,
                id: '',
                conditions: {
                    id: '',
                    name: '',
                    varsion: '',

                }
            }
        },
        _initMethod: function() {
            vc.component._listBusinessPackages(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {
            vc.on('businessPackageManage', 'listBusinessPackage', function(_param) {
                vc.component._listBusinessPackages(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listBusinessPackages(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listBusinessPackages: function(_page, _rows) {

                vc.component.businessPackageManageInfo.conditions.page = _page;
                vc.component.businessPackageManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.businessPackageManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/soft/getBusinessPackages',
                    param,
                    function(json, res) {
                        var _businessPackageManageInfo = JSON.parse(json);
                        vc.component.businessPackageManageInfo.total = _businessPackageManageInfo.total;
                        vc.component.businessPackageManageInfo.records = _businessPackageManageInfo.records;
                        vc.component.businessPackageManageInfo.businessPackages = _businessPackageManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.businessPackageManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddBusinessPackageModal: function() {
                vc.emit('addBusinessPackage', 'openAddBusinessPackageModal', {});
            },
            _openEditBusinessPackageModel: function(_businessPackage) {
                vc.emit('editBusinessPackage', 'openEditBusinessPackageModal', _businessPackage);
            },
            _openDeleteBusinessPackageModel: function(_businessPackage) {
                vc.emit('deleteBusinessPackage', 'openDeleteBusinessPackageModal', _businessPackage);
            },
            _queryBusinessPackageMethod: function() {
                vc.component._listBusinessPackages(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.businessPackageManageInfo.moreCondition) {
                    vc.component.businessPackageManageInfo.moreCondition = false;
                } else {
                    vc.component.businessPackageManageInfo.moreCondition = true;
                }
            },
            _toFileSystem: function(_businessPackage) {
                vc.saveData('5f3761ed-008a-475b-ad37-8cc35c88402c_curPath', _businessPackage.basePath + "/" + _businessPackage.id);
                vc.jumpToPage('/index.html#/pages/admin/fileManager?hostId=5f3761ed-008a-475b-ad37-8cc35c88402c');
            }


        }
    });
})(window.vc);