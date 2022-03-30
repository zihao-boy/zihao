(function(vc){
    vc.extends({
        propTypes: {
           emitChooseWafRule:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseWafRuleInfo:{
                wafRules:[],
                _currentWafRuleName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseWafRule','openChooseWafRuleModel',function(_param){
                $('#chooseWafRuleModel').modal('show');
                vc.component._refreshChooseWafRuleInfo();
                vc.component._loadAllWafRuleInfo(1,10,'');
            });
        },
        methods:{
            _loadAllWafRuleInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('wafRule.listWafRules',
                             param,
                             function(json){
                                var _wafRuleInfo = JSON.parse(json);
                                vc.component.chooseWafRuleInfo.wafRules = _wafRuleInfo.wafRules;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseWafRule:function(_wafRule){
                if(_wafRule.hasOwnProperty('name')){
                     _wafRule.wafRuleName = _wafRule.name;
                }
                vc.emit($props.emitChooseWafRule,'chooseWafRule',_wafRule);
                vc.emit($props.emitLoadData,'listWafRuleData',{
                    wafRuleId:_wafRule.wafRuleId
                });
                $('#chooseWafRuleModel').modal('hide');
            },
            queryWafRules:function(){
                vc.component._loadAllWafRuleInfo(1,10,vc.component.chooseWafRuleInfo._currentWafRuleName);
            },
            _refreshChooseWafRuleInfo:function(){
                vc.component.chooseWafRuleInfo._currentWafRuleName = "";
            }
        }

    });
})(window.vc);
