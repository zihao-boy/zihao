(function (vc, vm) {

    vc.extends({
        data: {
            viewDbDataInfo: {
                dataCols: [],
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('viewDbData', 'openViewDbDataModal', function (_params) {
                vc.component.refreshViewDbDataInfo();
                $('#viewDbDataModel').modal('show');
                _params.dataCols.forEach(item => {
                    $that.viewDbDataInfo.dataCols.push(
                        {
                            name:item,
                            value:_params.data[item]
                        }
                    )
                });
            });
        },
        methods: {
            refreshViewDbDataInfo: function () {
                vc.component.viewDbDataInfo = {
                    dataCols: [],
                }
            },
            _copyInsertSql:function(){
                $that._copy('werwerwerwer');
            },

            _copy:function (_value) {
                let transfer = document.createElement('input');
                document.body.appendChild(transfer);
                transfer.value = _value;  // 这里表示想要复制的内容
                transfer.focus();
                transfer.select();
                if (document.execCommand('copy')) {
                    document.execCommand('copy');
                }
                transfer.blur();
                console.log('复制成功');
                document.body.removeChild(transfer);
                
            }
        }
    });

})(window.vc, window.vc.component);