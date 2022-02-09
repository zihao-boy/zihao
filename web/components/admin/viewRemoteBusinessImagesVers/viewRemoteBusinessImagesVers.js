(function(vc) {
    vc.extends({

        data: {
            viewRemoteBusinessImagesVersInfo: {
                vers: [],
                extImagesId: '',
            }
        },
        _initMethod: function() {},
        _initEvent: function() {
            vc.on('viewRemoteBusinessImagesVers', 'open', function(_param) {
                $('#viewRemoteBusinessImagesVersModel').modal('show');
                $that.viewRemoteBusinessImagesVersInfo.extImagesId = _param.extImagesId;
                vc.component._loadAllRemoteBusinessImagesInfo(1, 10);
            });
            vc.on('viewRemoteBusinessImagesVers', 'paginationPlus', 'page_event', function(_currentPage) {
                vc.component._loadAllBusinessImagesInfo(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _loadAllRemoteBusinessImagesInfo: function(_page, _row) {
                var param = {
                    params: {
                        page: _page,
                        row: _row,
                        imagesId: $that.viewRemoteBusinessImagesVersInfo.imagesId
                    }
                };

                //发送get请求
                vc.http.apiGet('/soft/getBusinessImagesVer',
                    param,
                    function(json) {
                        var _verInfo = JSON.parse(json);
                        vc.component.viewRemoteBusinessImagesVersInfo.vers = _verInfo.data;
                        vc.emit('newOaWorkflowUndo', 'paginationPlus', 'init', {
                            total: _verInfo.records,
                            dataCount: _verInfo.total,
                            currentPage: _page
                        });
                    },
                    function() {
                        console.log('请求失败处理');
                    }
                );
            }
        }

    });
})(window.vc);