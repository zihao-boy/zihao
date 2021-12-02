(function (vc, vm) {

    vc.extends({
        data: {
            editBusinessImagesInfo: {
                id: '',
                name: '',
                imagesType: '',
                typeUrl: '',
                imagesFlag: '',
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editBusinessImages', 'openEditBusinessImagesModal', function (_params) {
                vc.component.refreshEditBusinessImagesInfo();
                $('#editBusinessImagesModel').modal('show');
                vc.copyObject(_params, vc.component.editBusinessImagesInfo);
            });
        },
        methods: {
            editBusinessImagesValidate: function () {
                return vc.validate.validate({
                    editBusinessImagesInfo: vc.component.editBusinessImagesInfo
                }, {
                    'editBusinessImagesInfo.name': [
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
                    'editBusinessImagesInfo.id': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "编号不能为空"
                        }]
                });
            },
            editBusinessImages: function () {
                if (!vc.component.editBusinessImagesValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/soft/updateBusinessImages',
                    JSON.stringify(vc.component.editBusinessImagesInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editBusinessImagesModel').modal('hide');
                            vc.emit('businessImagesManage', 'listBusinessImages', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditBusinessImagesInfo: function () {
                vc.component.editBusinessImagesInfo = {
                    id: '',
                    name: '',
                    imagesType: '',
                    typeUrl: '',
                    imagesFlag: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
