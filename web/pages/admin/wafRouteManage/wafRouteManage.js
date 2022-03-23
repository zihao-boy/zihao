/**
    入驻小区
**/
(function(vc){
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data:{
            wafRouteManageInfo:{
                wafRoutes:[],
                total:0,
                records:1,
                moreCondition:false,
                routeId:'',
                conditions:{
                    routeId:'',
hostname:'',
ip:'',
port:'',

                }
            }
        },
        _initMethod:function(){
            vc.component._listWafRoutes(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent:function(){
            
            vc.on('wafRouteManage','listWafRoute',function(_param){
                  vc.component._listWafRoutes(DEFAULT_PAGE, DEFAULT_ROWS);
            });
             vc.on('pagination','page_event',function(_currentPage){
                vc.component._listWafRoutes(_currentPage,DEFAULT_ROWS);
            });
        },
        methods:{
            _listWafRoutes:function(_page, _rows){

                vc.component.wafRouteManageInfo.conditions.page = _page;
                vc.component.wafRouteManageInfo.conditions.row = _rows;
                var param = {
                    params:vc.component.wafRouteManageInfo.conditions
               };

               //发送get请求
               vc.http.apiGet('wafRoute.listWafRoutes',
                             param,
                             function(json,res){
                                var _wafRouteManageInfo=JSON.parse(json);
                                vc.component.wafRouteManageInfo.total = _wafRouteManageInfo.total;
                                vc.component.wafRouteManageInfo.records = _wafRouteManageInfo.records;
                                vc.component.wafRouteManageInfo.wafRoutes = _wafRouteManageInfo.data;
                                vc.emit('pagination','init',{
                                     total:vc.component.wafRouteManageInfo.records,
                                     currentPage:_page
                                 });
                             },function(errInfo,error){
                                console.log('请求失败处理');
                             }
                           );
            },
            _openAddWafRouteModal:function(){
                vc.emit('addWafRoute','openAddWafRouteModal',{});
            },
            _openEditWafRouteModel:function(_wafRoute){
                vc.emit('editWafRoute','openEditWafRouteModal',_wafRoute);
            },
            _openDeleteWafRouteModel:function(_wafRoute){
                vc.emit('deleteWafRoute','openDeleteWafRouteModal',_wafRoute);
            },
            _queryWafRouteMethod:function(){
                vc.component._listWafRoutes(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition:function(){
                if(vc.component.wafRouteManageInfo.moreCondition){
                    vc.component.wafRouteManageInfo.moreCondition = false;
                }else{
                    vc.component.wafRouteManageInfo.moreCondition = true;
                }
            }

             
        }
    });
})(window.vc);
