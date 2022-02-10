(function(vc) {
    vc.extends({

        data: {
            viewBusinessImagesVersInfo: {
                vers: [],
                imagesId: '',
                publisherId: '',
                extImagesId:''
            }
        },
        _initMethod: function() {},
        _initEvent: function() {
            vc.on('viewBusinessImagesVers', 'open', function(_param) {
                $('#viewBusinessImagesVersModel').modal('show');
                $that.viewBusinessImagesVersInfo.imagesId = _param.id;
                $that.viewBusinessImagesVersInfo.publisherId = _param.publisherId;
                $that.viewBusinessImagesVersInfo.extImagesId = _param.extImagesId;

                vc.component._loadAllBusinessImagesInfo(1, 10);
            });
            vc.on('viewBusinessImagesVers', 'paginationPlus', 'page_event', function(_currentPage) {
                vc.component._loadAllBusinessImagesInfo(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _loadAllBusinessImagesInfo: function(_page, _row) {
                var param = {
                    params: {
                        page: _page,
                        row: _row,
                        imagesId: $that.viewBusinessImagesVersInfo.imagesId
                    }
                };

                //发送get请求
                vc.http.apiGet('/soft/getBusinessImagesVer',
                    param,
                    function(json) {
                        var _verInfo = JSON.parse(json);
                        vc.component.viewBusinessImagesVersInfo.vers = _verInfo.data;
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
            },
            _publishAppVersion:function(_version){

                let _data = {
                    imagesId:$that.viewBusinessImagesVersInfo.extImagesId,
                    version:_version.version,
                    url:_version.typeUrl,
                    publisherId:$that.viewBusinessImagesVersInfo.publisherId,
                }

                vc.http.apiPost(
                    '/soft/saveRemoteBusinessImagesVer',
                    JSON.stringify(_data), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#viewBusinessImagesVersModel').modal('hide');
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });

            }
        
        }

    });
})(window.vc);