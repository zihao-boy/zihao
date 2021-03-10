(function (vc, vm) {

    vc.extends({
        data: {
            editJobInfo: {
                taskId: '',
                taskName: '',
                templateId: '',
                taskCron: '',
                taskAttr: [],
                templates: [],
                templateSpecs: []
            }
        },
        _initMethod: function () {
           
        },
        _initEvent: function () {
            vc.on('editJob', 'openEditJobModal', function (_params) {
                vc.component.refreshEditJobInfo();
                $('#editJobModel').modal('show');
                vc.copyObject(_params, vc.component.editJobInfo);
                $that.queryEditTaskAttrs();
                //vc.component.editJobInfo.communityId = vc.getCurrentCommunity().communityId;
            });
        },
        methods: {
            editJobValidate: function () {
                return vc.validate.validate({
                    editJobInfo: vc.component.editJobInfo
                }, {
                    'editJobInfo.taskName': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "任务名称不能为空"
                        },
                        {
                            limit: "maxin",
                            param: "2,50",
                            errInfo: "任务名称必须在2至50字符之间"
                        },
                    ],
                    'editJobInfo.templateId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "模板不能为空"
                        },
                    ],
                    'editJobInfo.taskCron': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "定时不能为空"
                        },
                    ]

                });
            },
            editJob: function () {
                if (!vc.component.editJobValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/monitor/updateMonitorTask',
                    JSON.stringify(vc.component.editJobInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        let data = res.data;
                        if (_json.code == 0) {
                            //关闭model
                            $('#editJobModel').modal('hide');
                            vc.component.refreshEditJobInfo();
                            vc.emit('jobManage', 'listJob', {});
                            return;
                        }
                        vc.toast(json);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshEditJobInfo: function () {
                vc.component.editJobInfo = {
                    taskId: '',
                    taskName: '',
                    templateId: '',
                    taskCron: '',
                    taskAttr: [],
                    templates: [],
                    templateSpecs: []

                }
            },
            queryEditTaskAttrs: function () {
                var _param = {
                    params: {
                        page: 1,
                        row: 30,
                        taskId: $that.editJobInfo.taskId
                    }
                };
                //获取模板信息
                vc.http.apiGet('/monitor/listTaskAttrs',
                    _param,
                    function (json, res) {
                        let _json = JSON.parse(json);
                        let data = _json.data;
                        if (_json.code == 0) {
                            $that.editJobInfo.templateSpecs = data;
                            return;
                        }
                    }, function (errInfo, error) {
                        console.log('请求失败处理', errInfo, error);
                        vc.toast("查询地区失败");
                    });
            },
        }
    });

})(window.vc, window.vc.component);
