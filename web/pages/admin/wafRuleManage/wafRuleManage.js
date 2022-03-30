/**
    入驻小区
**/
(function(vc){
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data:{
            wafRuleManageInfo:{
                wafRules:[],
                total:0,
                records:1,
                moreCondition:false,
                ruleId:'',
                conditions:{
                    groupId:'',
ruleName:'',
objType:'',
state:'',

                }
            }
        },
        _initMethod:function(){
            vc.component._listWafRules(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent:function(){
            
            vc.on('wafRuleManage','listWafRule',function(_param){
                  vc.component._listWafRules(DEFAULT_PAGE, DEFAULT_ROWS);
            });
             vc.on('pagination','page_event',function(_currentPage){
                vc.component._listWafRules(_currentPage,DEFAULT_ROWS);
            });
        },
        methods:{
            _listWafRules:function(_page, _rows){

                vc.component.wafRuleManageInfo.conditions.page = _page;
                vc.component.wafRuleManageInfo.conditions.row = _rows;
                var param = {
                    params:vc.component.wafRuleManageInfo.conditions
               };

               //发送get请求
               vc.http.apiGet('wafRule.listWafRules',
                             param,
                             function(json,res){
                                var _wafRuleManageInfo=JSON.parse(json);
                                vc.component.wafRuleManageInfo.total = _wafRuleManageInfo.total;
                                vc.component.wafRuleManageInfo.records = _wafRuleManageInfo.records;
                                vc.component.wafRuleManageInfo.wafRules = _wafRuleManageInfo.data;
                                vc.emit('pagination','init',{
                                     total:vc.component.wafRuleManageInfo.records,
                                     currentPage:_page
                                 });
                             },function(errInfo,error){
                                console.log('请求失败处理');
                             }
                           );
            },
            _openAddWafRuleModal:function(){
                vc.emit('addWafRule','openAddWafRuleModal',{});
            },
            _openEditWafRuleModel:function(_wafRule){
                vc.emit('editWafRule','openEditWafRuleModal',_wafRule);
            },
            _openDeleteWafRuleModel:function(_wafRule){
                vc.emit('deleteWafRule','openDeleteWafRuleModal',_wafRule);
            },
            _queryWafRuleMethod:function(){
                vc.component._listWafRules(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition:function(){
                if(vc.component.wafRuleManageInfo.moreCondition){
                    vc.component.wafRuleManageInfo.moreCondition = false;
                }else{
                    vc.component.wafRuleManageInfo.moreCondition = true;
                }
            }

             
        }
    });
})(window.vc);
