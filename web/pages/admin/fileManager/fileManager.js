/**
    入驻小区
**/
(function(vc) {
    vc.extends({
        data: {
            fileManagerInfo: {
                curPath:'/',
                hostId:'',
                files:[]
            }
        },
        _initMethod: function() {
            $that.fileManagerInfo.hostId = vc.getParam('hostId');
            
        },
        _initEvent: function() {
            
        },
        methods: {
            


        }
    });
})(window.vc);