(function(vc){
    vc.extends({
        propTypes: {
           emitChooseBackup:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseBackupInfo:{
                backups:[],
                _currentBackupName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseBackup','openChooseBackupModel',function(_param){
                $('#chooseBackupModel').modal('show');
                vc.component._refreshChooseBackupInfo();
                vc.component._loadAllBackupInfo(1,10,'');
            });
        },
        methods:{
            _loadAllBackupInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('backup.listBackups',
                             param,
                             function(json){
                                var _backupInfo = JSON.parse(json);
                                vc.component.chooseBackupInfo.backups = _backupInfo.backups;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseBackup:function(_backup){
                if(_backup.hasOwnProperty('name')){
                     _backup.backupName = _backup.name;
                }
                vc.emit($props.emitChooseBackup,'chooseBackup',_backup);
                vc.emit($props.emitLoadData,'listBackupData',{
                    backupId:_backup.backupId
                });
                $('#chooseBackupModel').modal('hide');
            },
            queryBackups:function(){
                vc.component._loadAllBackupInfo(1,10,vc.component.chooseBackupInfo._currentBackupName);
            },
            _refreshChooseBackupInfo:function(){
                vc.component.chooseBackupInfo._currentBackupName = "";
            }
        }

    });
})(window.vc);
