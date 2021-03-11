(function (vc) {

    vc.extends({
        data: {
            addJobInfo: {
                taskName: '',
                templateId: '',
                taskCron: '',
                templates: [],
                templateSpecs: [],
                hosts:[],
                hostId:'',
                noticeType:''

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addJob', 'openAddJobModal', function () {
                $that.queryTempalte();
                $that._listHosts();
                $('#addJobModel').modal('show');
            });
        },
        methods: {
            addJobValidate() {
                return vc.validate.validate({
                    addJobInfo: vc.component.addJobInfo
                }, {
                    'addJobInfo.taskName': [
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
                    'addJobInfo.templateId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "模板不能为空"
                        },
                    ],
                    'addJobInfo.taskCron': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "定时不能为空"
                        },
                    ]
                });
            },
            saveJobInfo: function () {
                if (!vc.component.addJobValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                vc.http.apiPost(
                    '/monitor/saveMonitorTask',
                    JSON.stringify(vc.component.addJobInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        let _json = JSON.parse(json);
                        let data = res.data;
                        if (_json.code == 0) {
                            //关闭model
                            $('#addJobModel').modal('hide');
                            vc.component.clearAddJobInfo();
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
            queryTempalte: function () {
                var _param = {
                    params: {
                        page: 1,
                        row: 30
                    }
                };
                //获取模板信息
                vc.http.apiGet('/monitor/listTaskTemplate',
                    _param,
                    function (json, res) {
                        console.log('task.listTaskTemplate',json);
                        let _json = JSON.parse(json);
                        let data = _json.data;
                        if (_json.code == 0) {
                            $that.addJobInfo.templates = data;
                            return;
                        }
                    }, function (errInfo, error) {
                        console.log('请求失败处理', errInfo, error);
                        vc.toast("查询地区失败");
                    });
            },
            chooseTemplate: function () {
                //根据当前 template 查询属性 渲染页面
                $that.addJobInfo.templates.forEach(item => {
                    if(item.templateId == $that.addJobInfo.templateId){
                        item.specs.forEach(specItem =>{
                            specItem.value="";
                        })
                        $that.addJobInfo.templateSpecs = item.specs;
                    }
                });             
            },
            clearAddJobInfo: function () {
                vc.component.addJobInfo = {
                    taskName: '',
                    templateId: '',
                    taskCron: '',
                    templates: [],
                    templateSpecs: [],
                    hosts:[],
                    hostId:'',
                    noticeType:''
                };
            },
            _listHosts: function (_page, _rows) {
                var param = {
                    params: {
                        page:1,
                        row:100
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function (json, res) {
                        var _hostManageInfo = JSON.parse(json);
                        vc.component.addJobInfo.hosts = _hostManageInfo.data;
                       
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);
