(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addBusinessPackageInfo: {
                id: '',
                name: '',
                varsion: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addBusinessPackage', 'openAddBusinessPackageModal', function () {
                $('#addBusinessPackageModel').modal('show');
            });
        },
        methods: {
            addBusinessPackageValidate() {
                return vc.validate.validate({
                    addBusinessPackageInfo: vc.component.addBusinessPackageInfo
                }, {
                    'addBusinessPackageInfo.name': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "名称不能超过64"
                        },
                    ],
                    'addBusinessPackageInfo.varsion': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "版本不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "32",
                            errInfo: "版本不能超过32"
                        },
                    ],




                });
            },
            saveBusinessPackageInfo: function () {
                if (!vc.component.addBusinessPackageValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                vc.component.addBusinessPackageInfo.communityId = vc.getCurrentCommunity().communityId;
                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addBusinessPackageInfo);
                    $('#addBusinessPackageModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/soft/saveBusinessPackages',
                    JSON.stringify(vc.component.addBusinessPackageInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addBusinessPackageModel').modal('hide');
                            vc.component.clearAddBusinessPackageInfo();
                            vc.emit('businessPackageManage', 'listBusinessPackage', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddBusinessPackageInfo: function () {
                vc.component.addBusinessPackageInfo = {
                    name: '',
                    varsion: '',

                };
            }
        }
    });

})(window.vc);
