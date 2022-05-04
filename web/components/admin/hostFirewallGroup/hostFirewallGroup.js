(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            hostFirewallGroupInfo: {
                groups: [],
                hostId: ''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            //切换 至费用页面
            vc.on('hostFirewallGroup', 'switch', function (_param) {
                if (_param.hostId == '') {
                    return;
                }
                $that.clearhostFirewallGroupInfo();
                vc.copyObject(_param, $that.hostFirewallGroupInfo)
                $that._listhostFirewallGroup(DEFAULT_PAGE, DEFAULT_ROWS);
            });
        },
        methods: {
            _listhostFirewallGroup: function (_page, _row) {

                let param = {
                    params: {
                        page: 1,
                        row: 50,
                        hostId: $that.hostFirewallGroupInfo.hostId
                    }
                }
                //发送get请求
                vc.http.apiGet('/firewall/getHostFirewallGroup',
                    param,
                    function (json, res) {
                        let _json = JSON.parse(json);
                        $that.hostFirewallGroupInfo.groups = _json.data;
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            clearhostFirewallGroupInfo: function () {
                $that.hostFirewallGroupInfo = {
                    groups: [],
                    hostId: ''
                }
            },
            _openAddHostGroupModal:function(){
                vc.emit(  'addHostFirewallGroup', 'openAddHostFirewallGroupModal',{
                    hostId:$that.hostFirewallGroupInfo.hostId
                })
            },
            _openDeleteFirewallRuleGroupModel:function(_group){
                vc.emit( 'deleteHostFirewallGroup','openDeleteHostFirewallGroupModal',_group)
            }
           
          

        }

    });
})(window.vc);
