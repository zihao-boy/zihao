(function(vc){
    vc.extends({
        propTypes: {
           emitChooseFirewallRule:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseFirewallRuleInfo:{
                firewallRules:[],
                _currentFirewallRuleName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseFirewallRule','openChooseFirewallRuleModel',function(_param){
                $('#chooseFirewallRuleModel').modal('show');
                vc.component._refreshChooseFirewallRuleInfo();
                vc.component._loadAllFirewallRuleInfo(1,10,'');
            });
        },
        methods:{
            _loadAllFirewallRuleInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('firewallRule.listFirewallRules',
                             param,
                             function(json){
                                var _firewallRuleInfo = JSON.parse(json);
                                vc.component.chooseFirewallRuleInfo.firewallRules = _firewallRuleInfo.firewallRules;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseFirewallRule:function(_firewallRule){
                if(_firewallRule.hasOwnProperty('name')){
                     _firewallRule.firewallRuleName = _firewallRule.name;
                }
                vc.emit($props.emitChooseFirewallRule,'chooseFirewallRule',_firewallRule);
                vc.emit($props.emitLoadData,'listFirewallRuleData',{
                    firewallRuleId:_firewallRule.firewallRuleId
                });
                $('#chooseFirewallRuleModel').modal('hide');
            },
            queryFirewallRules:function(){
                vc.component._loadAllFirewallRuleInfo(1,10,vc.component.chooseFirewallRuleInfo._currentFirewallRuleName);
            },
            _refreshChooseFirewallRuleInfo:function(){
                vc.component.chooseFirewallRuleInfo._currentFirewallRuleName = "";
            }
        }

    });
})(window.vc);
