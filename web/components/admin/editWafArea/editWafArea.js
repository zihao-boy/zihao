(function(vc, vm) {

    vc.extends({
        data: {
            editWafAreaInfo: {
                id: '',
                typeCd: '',
                ip: '',
                scope:'*',
                seq:'',
                state:'start',
                groupId:'',
                wafRuleGroups:[],
                china:'Y',
                areaNames:[],
                chinaAreas:[],
                foreignAreas:[]
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editWafArea', 'openEditWafAreaModal', function(_params) {
                vc.component.refreshEditWafAreaInfo();
                $that._listWafEditRuleGroups();
                $that._loadChinaEditAreas();
                $that._loadForeignEditAreas();
                $('#editWafAreaModel').modal('show');
                $that.editWafAreaInfo.areaNames = _params.areaName.split(',')
                vc.copyObject(_params, vc.component.editWafAreaInfo);
                if($that.editWafAreaInfo.areaNames.length < 1){
                    return ;
                }
                $that.editWafAreaInfo.china = 'N';
                $that.editWafAreaInfo.chinaAreas.forEach(item => {
                    if(item == $that.editWafAreaInfo.areaNames[0]){
                        $that.editWafAreaInfo.china = 'Y'
                    }
                });
            });
        },
        methods: {
            editWafAreaValidate: function() {
                return vc.validate.validate({
                    editWafAreaInfo: vc.component.editWafAreaInfo
                }, {
                    'editWafAreaInfo.typeCd': [{
                            limit: "required",
                            param: "",
                            errInfo: "类型不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "类型不能超过64"
                        },
                    ],
                    'editWafAreaInfo.id': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editWafArea: function() {
                if (!vc.component.editWafAreaValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                if($that.editWafAreaInfo.areaNames.length < 1){
                    vc.toast('请选择位置');
                    return ;
                }

                $that.editWafAreaInfo.areaName = $that.editWafAreaInfo.areaNames.join(',');

                vc.http.apiPost(
                    '/firewall/updateWafArea',
                    JSON.stringify(vc.component.editWafAreaInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editWafAreaModel').modal('hide');
                            vc.emit('wafAreaManage', 'listWafArea', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditWafAreaInfo: function() {
                vc.component.editWafAreaInfo = {
                    id: '',
                    typeCd: '',
                    ip: '',
                    scope:'*',
                    seq:'',
                    state:'start',
                    groupId:'',
                    wafRuleGroups:[],
                    china:'Y',
                    areaNames:[],
                    chinaAreas:[],
                    foreignAreas:[]

                }
            },
            _listWafEditRuleGroups: function () {

                var param = {
                    params: {
                        page:1,
                        row:100
                    }
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWafRuleGroup',
                    param,
                    function (json, res) {
                        var _wafRuleGroupManageInfo = JSON.parse(json);
                       
                        vc.component.editWafAreaInfo.wafRuleGroups = _wafRuleGroupManageInfo.data;
                       
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _loadChinaEditAreas:function(){
                $that.editWafAreaInfo.chinaAreas=[
                    '山东','河北','吉林', '黑龙江','辽宁','内蒙古','新疆','甘肃','宁夏','山西','陕西','河南','安徽','江苏','浙江','福建','广东',
                    '江西','海南','广西','贵州','湖南','湖北','四川','云南','西藏','青海','天津','上海','重庆','北京','台湾','香港','澳门',
                ]
            },
            _loadForeignEditAreas:function(){
                $that.editWafAreaInfo.foreignAreas=[
                    '阿富汗','阿尔巴尼亚','阿尔及利亚','美属萨摩亚','安道​​尔','安哥拉','安提瓜和巴布达','阿根廷','亚美尼亚','阿鲁巴','澳大利亚','奥地利',
                    '阿塞拜疆','巴哈马','巴林','孟加拉国','巴巴多斯','白俄罗斯','比利时','伯利兹','贝宁','不丹','玻利维亚','波斯尼亚和黑塞哥维那','博茨瓦纳',
                    '巴西','英属维尔京群岛','文莱','保加利亚','布基纳法索','布隆迪','柬埔寨','喀麦隆','加拿大','佛得角','开曼群岛','中非共和国','乍得','智利',
                    '哥伦比亚','科摩罗','刚果','哥斯达黎加','科特迪瓦','克罗地亚','古巴','塞浦路斯','捷克共和国','朝鲜','刚果(扎伊尔)','丹麦','吉布提','多米尼加',
                    '多米尼加共和国','东帝汶','厄瓜多尔','埃及','萨尔瓦多','赤道几内亚','厄立特里亚','爱沙尼亚','埃塞俄比亚','福克兰群岛(马尔维纳斯群岛)',
                    '法罗群岛','斐济','芬兰','法国','法属圭亚那','法属波利尼西亚','加蓬','冈比亚','格鲁吉亚','德国','加纳','希腊','格陵兰','瓜德罗普岛',
                    '危地马拉','根西岛','几内亚','几内亚比绍','圭亚那','海地','赫德岛和麦当劳群岛','洪都拉斯','匈牙利','冰岛','印度','印度尼西亚','伊朗',
                    '伊拉克','爱尔兰','以色列','意大利','牙买加','约旦','哈萨克斯坦','肯尼亚','基里巴斯','科威特','吉尔吉斯斯坦','老挝','拉脱维亚','黎巴嫩',
                    '莱索托','利比里亚','阿拉伯利比亚民众国','列支敦士登','立陶宛','卢森堡','马达加斯加','马拉维','马来西亚','马尔代夫','马里','马耳他','马提尼克岛',
                    '毛里塔尼亚','马约特岛','墨西哥','密克罗尼西亚(联邦) ','摩尔多瓦共和国','莫桑比克','缅甸','纳米比亚','尼泊尔','荷兰','荷属安的列斯','新喀里多尼亚',
                    '新西兰','尼加拉瓜','尼日尔','尼日利亚','诺福克岛','北马里亚纳群岛','挪威','阿曼','巴基斯坦','帕劳','巴拿马','巴布亚新几内亚','巴拉圭','秘鲁',
                    '菲律宾','波兰','葡萄牙','波多黎各','卡塔尔','韩国','罗马尼亚','俄罗斯','卢旺达','圣基茨和尼维斯','圣卢西亚','圣皮埃尔和密克隆','圣文森特和格林纳丁斯',
                    '萨摩亚','圣马力诺','圣多美和普林西比','沙特阿拉伯','塞内加尔','塞拉利昂','斯洛伐克','斯洛文尼亚','所罗门群岛','索马里','比勒陀利亚','西班牙','苏丹',
                    '苏里南','斯威士兰','瑞典','瑞士','阿拉伯叙利亚共和国','塔吉克斯坦','泰国','马其顿','多哥','汤加','突尼斯','土耳其','土库曼斯坦','图瓦卢','乌干达',
                    '乌克兰','阿联酋','英国','坦桑尼亚','美国','美属维尔京群岛','乌拉圭','乌兹别克斯坦','瓦努阿图','委内瑞拉','越南','南斯拉夫','赞比亚','津巴布韦','日本'
                ]
            }
        }
    });

})(window.vc, window.vc.component);