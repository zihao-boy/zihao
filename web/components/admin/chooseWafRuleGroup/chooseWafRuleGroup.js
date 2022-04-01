(function(vc){
    vc.extends({
        propTypes: {
           emitChooseWafRuleGroup:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseWafRuleGroupInfo:{
                wafRuleGroups:[],
                _currentWafRuleGroupName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseWafRuleGroup','openChooseWafRuleGroupModel',function(_param){
                $('#chooseWafRuleGroupModel').modal('show');
                vc.component._refreshChooseWafRuleGroupInfo();
                vc.component._loadAllWafRuleGroupInfo(1,10,'');
            });
        },
        methods:{
            _loadAllWafRuleGroupInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('wafRuleGroup.listWafRuleGroups',
                             param,
                             function(json){
                                var _wafRuleGroupInfo = JSON.parse(json);
                                vc.component.chooseWafRuleGroupInfo.wafRuleGroups = _wafRuleGroupInfo.wafRuleGroups;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseWafRuleGroup:function(_wafRuleGroup){
                if(_wafRuleGroup.hasOwnProperty('name')){
                     _wafRuleGroup.wafRuleGroupName = _wafRuleGroup.name;
                }
                vc.emit($props.emitChooseWafRuleGroup,'chooseWafRuleGroup',_wafRuleGroup);
                vc.emit($props.emitLoadData,'listWafRuleGroupData',{
                    wafRuleGroupId:_wafRuleGroup.wafRuleGroupId
                });
                $('#chooseWafRuleGroupModel').modal('hide');
            },
            queryWafRuleGroups:function(){
                vc.component._loadAllWafRuleGroupInfo(1,10,vc.component.chooseWafRuleGroupInfo._currentWafRuleGroupName);
            },
            _refreshChooseWafRuleGroupInfo:function(){
                vc.component.chooseWafRuleGroupInfo._currentWafRuleGroupName = "";
            }
        }

    });
})(window.vc);
