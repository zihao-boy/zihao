/**
    入驻小区
**/
(function(vc){
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data:{
            innerNetPrivilegeManageInfo:{
                innerNetPrivileges:[],
                total:0,
                records:1,
                moreCondition:false,
                pId:'',
                conditions:{
                    srcUserId:'',
targetUserId:'',
targetPort:'',

                }
            }
        },
        _initMethod:function(){
            vc.component._listInnerNetPrivileges(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent:function(){
            
            vc.on('innerNetPrivilegeManage','listInnerNetPrivilege',function(_param){
                  vc.component._listInnerNetPrivileges(DEFAULT_PAGE, DEFAULT_ROWS);
            });
             vc.on('pagination','page_event',function(_currentPage){
                vc.component._listInnerNetPrivileges(_currentPage,DEFAULT_ROWS);
            });
        },
        methods:{
            _listInnerNetPrivileges:function(_page, _rows){

                vc.component.innerNetPrivilegeManageInfo.conditions.page = _page;
                vc.component.innerNetPrivilegeManageInfo.conditions.row = _rows;
                var param = {
                    params:vc.component.innerNetPrivilegeManageInfo.conditions
               };

               //发送get请求
               vc.http.apiGet('innerNetPrivilege.listInnerNetPrivileges',
                             param,
                             function(json,res){
                                var _innerNetPrivilegeManageInfo=JSON.parse(json);
                                vc.component.innerNetPrivilegeManageInfo.total = _innerNetPrivilegeManageInfo.total;
                                vc.component.innerNetPrivilegeManageInfo.records = _innerNetPrivilegeManageInfo.records;
                                vc.component.innerNetPrivilegeManageInfo.innerNetPrivileges = _innerNetPrivilegeManageInfo.data;
                                vc.emit('pagination','init',{
                                     total:vc.component.innerNetPrivilegeManageInfo.records,
                                     currentPage:_page
                                 });
                             },function(errInfo,error){
                                console.log('请求失败处理');
                             }
                           );
            },
            _openAddInnerNetPrivilegeModal:function(){
                vc.emit('addInnerNetPrivilege','openAddInnerNetPrivilegeModal',{});
            },
            _openEditInnerNetPrivilegeModel:function(_innerNetPrivilege){
                vc.emit('editInnerNetPrivilege','openEditInnerNetPrivilegeModal',_innerNetPrivilege);
            },
            _openDeleteInnerNetPrivilegeModel:function(_innerNetPrivilege){
                vc.emit('deleteInnerNetPrivilege','openDeleteInnerNetPrivilegeModal',_innerNetPrivilege);
            },
            _queryInnerNetPrivilegeMethod:function(){
                vc.component._listInnerNetPrivileges(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition:function(){
                if(vc.component.innerNetPrivilegeManageInfo.moreCondition){
                    vc.component.innerNetPrivilegeManageInfo.moreCondition = false;
                }else{
                    vc.component.innerNetPrivilegeManageInfo.moreCondition = true;
                }
            }

             
        }
    });
})(window.vc);
