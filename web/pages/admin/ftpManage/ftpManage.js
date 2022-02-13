/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            ftpManageInfo: {
                ftps: [],
                total: 0,
                records: 1,
                moreCondition: false,
                ftpId: '',
                conditions: {
                    ftpId: '',
                    name: '',
                    ip: '',

                }
            }
        },
        _initMethod: function() {
            vc.component._listFtps(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('ftpManage', 'listFtp', function(_param) {
                vc.component._listFtps(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listFtps(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listFtps: function(_page, _rows) {

                vc.component.ftpManageInfo.conditions.page = _page;
                vc.component.ftpManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.ftpManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/resources/getFtp',
                    param,
                    function(json, res) {
                        var _ftpManageInfo = JSON.parse(json);
                        vc.component.ftpManageInfo.total = _ftpManageInfo.total;
                        vc.component.ftpManageInfo.records = _ftpManageInfo.records;
                        vc.component.ftpManageInfo.ftps = _ftpManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.ftpManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddFtpModal: function() {
                vc.emit('addFtp', 'openAddFtpModal', {});
            },
            _openEditFtpModel: function(_ftp) {
                vc.emit('editFtp', 'openEditFtpModal', _ftp);
            },
            _openDeleteFtpModel: function(_ftp) {
                vc.emit('deleteFtp', 'openDeleteFtpModal', _ftp);
            },
            _queryFtpMethod: function() {
                vc.component._listFtps(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.ftpManageInfo.moreCondition) {
                    vc.component.ftpManageInfo.moreCondition = false;
                } else {
                    vc.component.ftpManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);