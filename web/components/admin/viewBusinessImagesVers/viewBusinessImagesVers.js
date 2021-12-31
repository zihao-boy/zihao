(function (vc) {
    vc.extends({

        data: {
            viewBusinessImagesVersInfo: {
                vers: [],
                imagesId:''
            }
        },
        _initMethod: function () {
        },
        _initEvent: function () {
            vc.on('viewBusinessImagesVers', 'open', function (_param) {
                $('#viewBusinessImagesVersModel').modal('show');
                $that.viewBusinessImagesVersInfo.imagesId = _param.id;
                vc.component._loadAllBusinessImagesInfo(1, 10);
            });
            vc.on('viewBusinessImagesVers', 'paginationPlus', 'page_event', function (_currentPage) {
                vc.component._loadAllBusinessImagesInfo(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _loadAllBusinessImagesInfo: function (_page, _row) {
                var param = {
                    params: {
                        page: _page,
                        row: _row,
                        imagesId:$that.viewBusinessImagesVersInfo.imagesId
                    }
                };

                //发送get请求
                vc.http.apiGet('/soft/getBusinessImagesVer',
                    param,
                    function (json) {
                        var _verInfo = JSON.parse(json);
                        vc.component.viewBusinessImagesVersInfo.vers = _verInfo.data;
                        vc.emit('newOaWorkflowUndo', 'paginationPlus', 'init', {
                            total: _verInfo.records,
                            dataCount: _verInfo.total,
                            currentPage: _page
                        });
                    }, function () {
                        console.log('请求失败处理');
                    }
                );
            }
        }

    });
})(window.vc);
