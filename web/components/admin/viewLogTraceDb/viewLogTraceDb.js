(function(vc, vm) {

    vc.extends({
        data: {
            viewLogTraceDbInfo: {
                dbSql: '',
                param: '',
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('viewLogTraceDb', 'openViewLogTraceDbModal', function(_params) {
                vc.component.refreshViewLogTraceDbInfo();
                $('#viewLogTraceDbModel').modal('show');
                vc.copyObject(_params, $that.viewLogTraceDbInfo);
            });
        },
        methods: {
            refreshViewLogTraceDbInfo: function() {
                vc.component.viewLogTraceDbInfo = {
                    dbSql: '',
                    param: '',
                }
            },
        }
    });

})(window.vc, window.vc.component);