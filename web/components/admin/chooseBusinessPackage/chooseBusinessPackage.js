(function(vc) {
    vc.extends({
        propTypes: {
            emitChooseBusinessPackage: vc.propTypes.string
        },
        data: {
            chooseBusinessPackageInfo: {
                businessPackages: [],
                _currentBusinessPackageName: '',
                data: {}
            }
        },
        _initMethod: function() {},
        _initEvent: function() {
            vc.on('chooseBusinessPackage', 'openChooseBusinessPackageModel', function(_param) {
                $('#chooseBusinessPackageModel').modal('show');
                vc.component._refreshChooseBusinessPackageInfo();
                vc.component._loadAllBusinessPackageInfo(1, 10, '');
                $that.chooseBusinessPackageInfo.data = _param;
            });

            vc.on('chooseBusinessPackage', 'paginationPlus', 'page_event', function(_currentPage) {
                vc.component._loadAllBusinessPackageInfo(_currentPage, 10);
            });
        },
        methods: {
            _loadAllBusinessPackageInfo: function(_page, _row, _name) {
                var param = {
                    params: {
                        page: _page,
                        row: _row,
                        name: _name
                    }
                };
                //发送get请求
                vc.http.apiGet('/soft/getBusinessPackages',
                    param,
                    function(json) {
                        var _businessPackageInfo = JSON.parse(json);
                        vc.component.chooseBusinessPackageInfo.businessPackages = _businessPackageInfo.data;
                        vc.emit('chooseBusinessPackage', 'paginationPlus', 'init', {
                            total: _businessPackageInfo.records,
                            currentPage: _page
                        });
                    },
                    function() {
                        console.log('请求失败处理');
                    }
                );
            },
            chooseBusinessPackage: function(_businessPackage) {
                if (_businessPackage.hasOwnProperty('name')) {
                    _businessPackage.businessPackageName = _businessPackage.name;
                }
                vc.emit($props.emitChooseBusinessPackage, 'chooseBusinessPackage', _businessPackage);
                vc.copyObject(_businessPackage, $that.chooseBusinessPackageInfo.data);
                $that.chooseBusinessPackageInfo.data.businessPackageId = _businessPackage.id;
                $that.chooseBusinessPackageInfo.data.businessPackageName = _businessPackage.name;

                $('#chooseBusinessPackageModel').modal('hide');
            },
            queryBusinessPackages: function() {
                vc.component._loadAllBusinessPackageInfo(1, 10, vc.component.chooseBusinessPackageInfo._currentBusinessPackageName);
            },
            _refreshChooseBusinessPackageInfo: function() {
                vc.component.chooseBusinessPackageInfo._currentBusinessPackageName = "";
            }
        }

    });
})(window.vc);