(function(vc) {
    vc.extends({
        data: {
            viewPoolsImagesInfo: {
                images: []
            }
        },
        _initMethod: function() {},
        _initEvent: function() {
            vc.on('viewPoolsImages', 'openViewPoolsImagesModel', function(_params) {
                $that.viewPoolsImagesInfo.images = _params
                $('#viewPoolsImagesModel').modal('show');
            });
        },
        methods: {

        }

    });
})(window.vc);