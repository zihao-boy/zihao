(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addBusinessImagesInfo: {
                name: '',
                imagesType: '',
                typeUrl: '',
                imagesFlag: '',
                excelTemplate: '',
                isFile:'Y'
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addBusinessImages', 'openAddBusinessImagesModal', function (_param) {
                vc.copyObject(_param,$that.addBusinessImagesInfo);
                $('#addBusinessImagesModel').modal('show');
            });
        },
        methods: {
            addBusinessImagesValidate() {
                return vc.validate.validate({
                    addBusinessImagesInfo: vc.component.addBusinessImagesInfo
                }, {
                    'addBusinessImagesInfo.name': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "镜像名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "镜像名称不能超过64"
                        },
                    ],
                });
            },
            saveBusinessImagesInfo: function () {
                if (!vc.component.addBusinessImagesValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addBusinessImagesInfo);
                    $('#addBusinessImagesModel').modal('hide');
                    return;
                }
                let param = new FormData();
                param.append("uploadFile", vc.component.addBusinessImagesInfo.excelTemplate);
                param.append('name', vc.component.addBusinessImagesInfo.name);
                param.append('isFile', vc.component.addBusinessImagesInfo.isFile);
                param.append('typeUrl', vc.component.addBusinessImagesInfo.typeUrl);


                vc.http.apiPost(
                    '/soft/saveBusinessImages',
                    param,
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addBusinessImagesModel').modal('hide');
                            vc.component.clearAddBusinessImagesInfo();
                            vc.emit('businessImagesManage', 'listBusinessImages', {});

                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            getExcelTemplate: function (e) {
                //console.log("getExcelTemplate 开始调用")
                vc.component.addBusinessImagesInfo.excelTemplate = e.target.files[0];
            },
            clearAddBusinessImagesInfo: function () {
                vc.component.addBusinessImagesInfo = {
                    name: '',
                    imagesType: '',
                    typeUrl: '',
                    imagesFlag: '',
                    excelTemplate: '',
                    isFile:'Y'
                };
            }
        }
    });

})(window.vc);
