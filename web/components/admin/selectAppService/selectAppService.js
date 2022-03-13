(function(vc, vm) {

    vc.extends({
        data: {
            selectAppServiceInfo: {
                appServices: [],
                apps: [],
                isAll: true,
                callBack: {},
                asGroupId:''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('selectAppService', 'openSelectAppService', function(_params) {
                $that.selectAppServiceInfo.callBack = _params.callBack;
                $that.selectAppServiceInfo.asGroupId = _params.asGroupId;
                //load plan
                $that._loadSelectAppService();
                $('#selectAppServiceModel').modal('show');

            });
        },
        methods: {
            _build: function(_param) {

                $that.selectAppServiceInfo.callBack($that.selectAppServiceInfo.apps.join(','));
                $('#selectAppServiceModel').modal('hide');
            },
            closeDoBuildJobModel: function() {
                $('#selectAppServiceModel').modal('hide');
                $that.selectAppServiceInfo = {
                    appServices: [],
                    apps: [],
                    isAll: true,
                    callBack: {}
                }
            },
            _loadSelectAppService: function() {
                let param = {
                    params: {
                        page: 1,
                        row: 100,
                        asGroupId:$that.selectAppServiceInfo.asGroupId
                    }
                };
                //发送get请求
                vc.http.apiGet('/appService/getAppService',
                    param,
                    function(json, res) {
                        let _appVersionJobManageInfo = JSON.parse(json);
                        vc.component.selectAppServiceInfo.appServices = _appVersionJobManageInfo.data;
                        $that.changeAll();
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            changeAll: function() {
                $that.selectAppServiceInfo.apps = [];
                if (!$that.selectAppServiceInfo.isAll) {
                    return;
                }

                $that.selectAppServiceInfo.appServices.forEach(item => {
                    $that.selectAppServiceInfo.apps.push(item.asId);
                });
            },
            changeItem: function() {
                if ($that.selectAppServiceInfo.apps.length < $that.selectAppServiceInfo.appServices.length) {
                    $that.selectAppServiceInfo.isAll = false;
                    return;
                }
                $that.selectAppServiceInfo.isAll = true;

            }
        }
    });

})(window.vc, window.vc.component);