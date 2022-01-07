(function(vc) {

    vc.extends({
        data: {
            mysqlClientInfo: {
                dbLinks: [],
                curDbId:'',
                curDbName:'无',
                _currentTab:''

            }
        },
        _initMethod: function() {

            $("#text").setTextareaCount({
                width: "30px",
                bgColor: "#FFF",
                color: "#000",
                display: "inline-block"
            });

            $that._loadDbLink();
        },
        _initEvent: function() {

            vc.on('mysqlClient', 'load', function() {
                $that._loadDbLink();
            })

        },
        methods: {
            changeTab: function(_tab) {
                $that.mysqlClientInfo._currentTab = _tab;
                
            },
            _customKeypress: function() {
                let typeSql = window.getSelection().toString();
                console.log(typeSql);
            },
            _loadDbLink: function() {
                let _param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/dbClient/getDbLink',
                    _param,
                    function(json, res) {
                        var _appVarGroupManageInfo = JSON.parse(json);
                        vc.component.mysqlClientInfo.dbLinks = _appVarGroupManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openNewDbLinkModal: function() {
                vc.emit('newDbLink', 'openNewDbLinkModal', {})
            },
            _openEditDbLinkModal:function(_dbLink){
                vc.emit('editDbLink', 'openEditDbLinkModal',_dbLink);
            },
            _openDeleteDbLinkModal:function(_dbLink){
                vc.emit( 'deleteDbLink','openDeleteDbLinkModal',_dbLink);
            },
            _chooseDb:function(_dbLink){
                $that.mysqlClientInfo.curDbId = _dbLink.curDbId;
                $that.mysqlClientInfo.curDbName = _dbLink.name;
            }

           
        }
    });

})(window.vc);