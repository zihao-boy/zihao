/**
    入驻小区
**/
(function(vc){
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data:{
            firewallRuleGroupManageInfo:{
                firewallRuleGroups:[],
                total:0,
                records:1,
                moreCondition:false,
                groupId:'',
                conditions:{
                    groupId:'',
groupName:'',

                }
            }
        },
        _initMethod:function(){
            vc.component._listFirewallRuleGroups(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent:function(){
            
            vc.on('firewallRuleGroupManage','listFirewallRuleGroup',function(_param){
                  vc.component._listFirewallRuleGroups(DEFAULT_PAGE, DEFAULT_ROWS);
            });
             vc.on('pagination','page_event',function(_currentPage){
                vc.component._listFirewallRuleGroups(_currentPage,DEFAULT_ROWS);
            });
        },
        methods:{
            _listFirewallRuleGroups:function(_page, _rows){

                vc.component.firewallRuleGroupManageInfo.conditions.page = _page;
                vc.component.firewallRuleGroupManageInfo.conditions.row = _rows;
                var param = {
                    params:vc.component.firewallRuleGroupManageInfo.conditions
               };

               //发送get请求
               vc.http.apiGet('firewallRuleGroup.listFirewallRuleGroups',
                             param,
                             function(json,res){
                                var _firewallRuleGroupManageInfo=JSON.parse(json);
                                vc.component.firewallRuleGroupManageInfo.total = _firewallRuleGroupManageInfo.total;
                                vc.component.firewallRuleGroupManageInfo.records = _firewallRuleGroupManageInfo.records;
                                vc.component.firewallRuleGroupManageInfo.firewallRuleGroups = _firewallRuleGroupManageInfo.data;
                                vc.emit('pagination','init',{
                                     total:vc.component.firewallRuleGroupManageInfo.records,
                                     currentPage:_page
                                 });
                             },function(errInfo,error){
                                console.log('请求失败处理');
                             }
                           );
            },
            _openAddFirewallRuleGroupModal:function(){
                vc.emit('addFirewallRuleGroup','openAddFirewallRuleGroupModal',{});
            },
            _openEditFirewallRuleGroupModel:function(_firewallRuleGroup){
                vc.emit('editFirewallRuleGroup','openEditFirewallRuleGroupModal',_firewallRuleGroup);
            },
            _openDeleteFirewallRuleGroupModel:function(_firewallRuleGroup){
                vc.emit('deleteFirewallRuleGroup','openDeleteFirewallRuleGroupModal',_firewallRuleGroup);
            },
            _queryFirewallRuleGroupMethod:function(){
                vc.component._listFirewallRuleGroups(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition:function(){
                if(vc.component.firewallRuleGroupManageInfo.moreCondition){
                    vc.component.firewallRuleGroupManageInfo.moreCondition = false;
                }else{
                    vc.component.firewallRuleGroupManageInfo.moreCondition = true;
                }
            }

             
        }
    });
})(window.vc);
