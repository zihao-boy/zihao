/**
    入驻小区
**/
(function(vc){
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data:{
            firewallRuleManageInfo:{
                firewallRules:[],
                total:0,
                records:1,
                moreCondition:false,
                ruleId:'',
                conditions:{
                    inOut:'',
allowLimit:'',
seq:'',
protocol:'',
srcObj:'',
dstObj:'',

                }
            }
        },
        _initMethod:function(){
            vc.component._listFirewallRules(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent:function(){
            
            vc.on('firewallRuleManage','listFirewallRule',function(_param){
                  vc.component._listFirewallRules(DEFAULT_PAGE, DEFAULT_ROWS);
            });
             vc.on('pagination','page_event',function(_currentPage){
                vc.component._listFirewallRules(_currentPage,DEFAULT_ROWS);
            });
        },
        methods:{
            _listFirewallRules:function(_page, _rows){

                vc.component.firewallRuleManageInfo.conditions.page = _page;
                vc.component.firewallRuleManageInfo.conditions.row = _rows;
                var param = {
                    params:vc.component.firewallRuleManageInfo.conditions
               };

               //发送get请求
               vc.http.apiGet('firewallRule.listFirewallRules',
                             param,
                             function(json,res){
                                var _firewallRuleManageInfo=JSON.parse(json);
                                vc.component.firewallRuleManageInfo.total = _firewallRuleManageInfo.total;
                                vc.component.firewallRuleManageInfo.records = _firewallRuleManageInfo.records;
                                vc.component.firewallRuleManageInfo.firewallRules = _firewallRuleManageInfo.data;
                                vc.emit('pagination','init',{
                                     total:vc.component.firewallRuleManageInfo.records,
                                     currentPage:_page
                                 });
                             },function(errInfo,error){
                                console.log('请求失败处理');
                             }
                           );
            },
            _openAddFirewallRuleModal:function(){
                vc.emit('addFirewallRule','openAddFirewallRuleModal',{});
            },
            _openEditFirewallRuleModel:function(_firewallRule){
                vc.emit('editFirewallRule','openEditFirewallRuleModal',_firewallRule);
            },
            _openDeleteFirewallRuleModel:function(_firewallRule){
                vc.emit('deleteFirewallRule','openDeleteFirewallRuleModal',_firewallRule);
            },
            _queryFirewallRuleMethod:function(){
                vc.component._listFirewallRules(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition:function(){
                if(vc.component.firewallRuleManageInfo.moreCondition){
                    vc.component.firewallRuleManageInfo.moreCondition = false;
                }else{
                    vc.component.firewallRuleManageInfo.moreCondition = true;
                }
            }

             
        }
    });
})(window.vc);
