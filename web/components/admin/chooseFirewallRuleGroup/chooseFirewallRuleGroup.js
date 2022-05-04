(function(vc){
    vc.extends({
        propTypes: {
           emitChooseFirewallRuleGroup:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseFirewallRuleGroupInfo:{
                firewallRuleGroups:[],
                _currentFirewallRuleGroupName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseFirewallRuleGroup','openChooseFirewallRuleGroupModel',function(_param){
                $('#chooseFirewallRuleGroupModel').modal('show');
                vc.component._refreshChooseFirewallRuleGroupInfo();
                vc.component._loadAllFirewallRuleGroupInfo(1,10,'');
            });
        },
        methods:{
            _loadAllFirewallRuleGroupInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('firewallRuleGroup.listFirewallRuleGroups',
                             param,
                             function(json){
                                var _firewallRuleGroupInfo = JSON.parse(json);
                                vc.component.chooseFirewallRuleGroupInfo.firewallRuleGroups = _firewallRuleGroupInfo.firewallRuleGroups;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseFirewallRuleGroup:function(_firewallRuleGroup){
                if(_firewallRuleGroup.hasOwnProperty('name')){
                     _firewallRuleGroup.firewallRuleGroupName = _firewallRuleGroup.name;
                }
                vc.emit($props.emitChooseFirewallRuleGroup,'chooseFirewallRuleGroup',_firewallRuleGroup);
                vc.emit($props.emitLoadData,'listFirewallRuleGroupData',{
                    firewallRuleGroupId:_firewallRuleGroup.firewallRuleGroupId
                });
                $('#chooseFirewallRuleGroupModel').modal('hide');
            },
            queryFirewallRuleGroups:function(){
                vc.component._loadAllFirewallRuleGroupInfo(1,10,vc.component.chooseFirewallRuleGroupInfo._currentFirewallRuleGroupName);
            },
            _refreshChooseFirewallRuleGroupInfo:function(){
                vc.component.chooseFirewallRuleGroupInfo._currentFirewallRuleGroupName = "";
            }
        }

    });
})(window.vc);
