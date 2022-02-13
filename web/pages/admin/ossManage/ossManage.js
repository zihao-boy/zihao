/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            ossManageInfo: {
                osss: [],
                total: 0,
                records: 1,
                moreCondition: false,
                ossId: '',
                conditions: {
                    ossId: '',
                    name: '',
                    ossType: '',

                }
            }
        },
        _initMethod: function() {
            vc.component._listOsss(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('ossManage', 'listOss', function(_param) {
                vc.component._listOsss(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listOsss(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listOsss: function(_page, _rows) {

                vc.component.ossManageInfo.conditions.page = _page;
                vc.component.ossManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.ossManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/resources/getOss',
                    param,
                    function(json, res) {
                        var _ossManageInfo = JSON.parse(json);
                        vc.component.ossManageInfo.total = _ossManageInfo.total;
                        vc.component.ossManageInfo.records = _ossManageInfo.records;
                        vc.component.ossManageInfo.osss = _ossManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.ossManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddOssModal: function() {
                vc.emit('addOss', 'openAddOssModal', {});
            },
            _openEditOssModel: function(_oss) {
                vc.emit('editOss', 'openEditOssModal', _oss);
            },
            _openDeleteOssModel: function(_oss) {
                vc.emit('deleteOss', 'openDeleteOssModal', _oss);
            },
            _queryOssMethod: function() {
                vc.component._listOsss(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.ossManageInfo.moreCondition) {
                    vc.component.ossManageInfo.moreCondition = false;
                } else {
                    vc.component.ossManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);