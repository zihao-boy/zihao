(function(vc) {
    vc.extends({

        data: {
            viewRemoteBusinessImagesVersInfo: {
                vers: [],
                extImagesId: '',
                imagesName:'',
                imagesId:''
            }
        },
        _initMethod: function() {},
        _initEvent: function() {
            vc.on('viewRemoteBusinessImagesVers', 'open', function(_param) {
                $('#viewRemoteBusinessImagesVersModel').modal('show');
                $that.viewRemoteBusinessImagesVersInfo.extImagesId = _param.extImagesId;
                $that.viewRemoteBusinessImagesVersInfo.imagesName = _param.name;
                $that.viewRemoteBusinessImagesVersInfo.imagesId = _param.id;
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
                        imagesId: $that.viewRemoteBusinessImagesVersInfo.extImagesId
                    }
                };

                //发送get请求
                vc.http.apiGet('/soft/getRemoteBusinessImagesVer',
                    param,
                    function(json) {
                        var _verInfo = JSON.parse(json);
                        vc.component.viewRemoteBusinessImagesVersInfo.vers = _verInfo.data;
                        vc.emit('viewRemoteBusinessImagesVers', 'paginationPlus', 'init', {
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
            _downloadImagesVersion:function(_version){

                let _data = {
                    imagesId:$that.viewRemoteBusinessImagesVersInfo.imagesId,
                    version:_version.version,
                    typeUrl:_version.url
                }

                vc.http.apiPost(
                    '/soft/saveBusinessImagesVer',
                    JSON.stringify(_data), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#viewRemoteBusinessImagesVersModel').modal('hide');
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