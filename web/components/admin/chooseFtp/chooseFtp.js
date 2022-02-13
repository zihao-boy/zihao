(function(vc){
    vc.extends({
        propTypes: {
           emitChooseFtp:vc.propTypes.string,
           emitLoadData:vc.propTypes.string
        },
        data:{
            chooseFtpInfo:{
                ftps:[],
                _currentFtpName:'',
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('chooseFtp','openChooseFtpModel',function(_param){
                $('#chooseFtpModel').modal('show');
                vc.component._refreshChooseFtpInfo();
                vc.component._loadAllFtpInfo(1,10,'');
            });
        },
        methods:{
            _loadAllFtpInfo:function(_page,_row,_name){
                var param = {
                    params:{
                        page:_page,
                        row:_row,
                        communityId:vc.getCurrentCommunity().communityId,
                        name:_name
                    }
                };

                //发送get请求
               vc.http.apiGet('ftp.listFtps',
                             param,
                             function(json){
                                var _ftpInfo = JSON.parse(json);
                                vc.component.chooseFtpInfo.ftps = _ftpInfo.ftps;
                             },function(){
                                console.log('请求失败处理');
                             }
                           );
            },
            chooseFtp:function(_ftp){
                if(_ftp.hasOwnProperty('name')){
                     _ftp.ftpName = _ftp.name;
                }
                vc.emit($props.emitChooseFtp,'chooseFtp',_ftp);
                vc.emit($props.emitLoadData,'listFtpData',{
                    ftpId:_ftp.ftpId
                });
                $('#chooseFtpModel').modal('hide');
            },
            queryFtps:function(){
                vc.component._loadAllFtpInfo(1,10,vc.component.chooseFtpInfo._currentFtpName);
            },
            _refreshChooseFtpInfo:function(){
                vc.component.chooseFtpInfo._currentFtpName = "";
            }
        }

    });
})(window.vc);
