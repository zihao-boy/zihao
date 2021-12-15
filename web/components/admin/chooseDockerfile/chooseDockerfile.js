(function(vc) {
    vc.extends({
        propTypes: {
            emitChooseDockerfile: vc.propTypes.string
        },
        data: {
            chooseDockerfileInfo: {
                dockerfiles: [],
                _currentDockerfileName: '',
                data: {}
            }
        },
        _initMethod: function() {},
        _initEvent: function() {
            vc.on('chooseDockerfile', 'openChooseDockerfileModel', function(_param) {
                $('#chooseDockerfileModel').modal('show');
                vc.component._refreshChooseDockerfileInfo();
                vc.component._loadAllDockerfileInfo(1, 10, '');
                $that.chooseDockerfileInfo.data = _param;
            });

            vc.on('chooseDockerfile', 'paginationPlus', 'page_event', function(_currentPage) {
                vc.component._loadAllDockerfileInfo(_currentPage, 10);
            });
        },
        methods: {
            _loadAllDockerfileInfo: function(_page, _row, _name) {
                var param = {
                    params: {
                        page: _page,
                        row: _row,
                        name: _name
                    }
                };
                //发送get请求
                vc.http.apiGet('/soft/getBusinessDockerfile',
                    param,
                    function(json) {
                        var _dockerfileInfo = JSON.parse(json);
                        vc.component.chooseDockerfileInfo.dockerfiles = _dockerfileInfo.data;
                        vc.emit('chooseDockerfile', 'paginationPlus', 'init', {
                            total: _dockerfileInfo.records,
                            currentPage: _page
                        });
                    },
                    function() {
                        console.log('请求失败处理');
                    }
                );
            },
            chooseDockerfile: function(_dockerfile) {
                if (_dockerfile.hasOwnProperty('name')) {
                    _dockerfile.dockerfileName = _dockerfile.name;
                }
                vc.emit($props.emitChooseDockerfile, 'chooseDockerfile', _dockerfile);
                vc.copyObject(_dockerfile, $that.chooseDockerfileInfo.data);
                $that.chooseDockerfileInfo.data.businessDockerfileId = _dockerfile.id;
                $that.chooseDockerfileInfo.data.businessDockerfileName = _dockerfile.name;

                $('#chooseDockerfileModel').modal('hide');
            },
            queryDockerfiles: function() {
                vc.component._loadAllDockerfileInfo(1, 10, vc.component.chooseDockerfileInfo._currentDockerfileName);
            },
            _refreshChooseDockerfileInfo: function() {
                vc.component.chooseDockerfileInfo._currentDockerfileName = "";
            }
        }

    });
})(window.vc);