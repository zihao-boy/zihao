(function (vc, vm) {

    vc.extends({
        data: {
            viewDbDataInfo: {
                dataCols: [],
                view:'form',
                sql:''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('viewDbData', 'openViewDbDataModal', function (_params) {
                vc.component.refreshViewDbDataInfo();
                $('#viewDbDataModel').modal('show');
                _params.dataCols.forEach(item => {
                    $that.viewDbDataInfo.dataCols.push(
                        {
                            name:item,
                            value:_params.data[item]
                        }
                    )
                });
            });
        },
        methods: {
            refreshViewDbDataInfo: function () {
                vc.component.viewDbDataInfo = {
                    dataCols: [],
                    view:'form',
                    sql:''
                }
            },
            _copyInsertSql:function(){
                $that.viewDbDataInfo.view = "sql";
                let _sql = " insert into xx(\n";

                let _values = " \n) values (\n";

                $that.viewDbDataInfo.dataCols.forEach((item)=>{

                    _sql += (item.name+",\n")
                    _values += ("'"+item.value +"',\n")
                });

                if(_sql.endsWith(",\n")){
                    _sql = _sql.substring(0,_sql.length -2);
                    _values = _values.substring(0,_values.length -2);
                }

                _sql += (_values +"\n);\n");

                $that.viewDbDataInfo.sql = _sql;
            },
            _copyUpdateSql:function(){
                $that.viewDbDataInfo.view = "sql";
                let _sql = " update xx set \n";

                let _where = " \n where 1=1 \n";

                $that.viewDbDataInfo.dataCols.forEach((item)=>{
                    _sql += (item.name+"='"+item.value+"',\n")
                    _where += (" and "+item.name+"='"+item.value+"'\n")
                });

                if(_sql.endsWith(",\n")){
                    _sql = _sql.substring(0,_sql.length -2);
                }

                _sql += (_where +";");

                $that.viewDbDataInfo.sql = _sql;
            },
            _copyDeleteSql:function(){
                $that.viewDbDataInfo.view = "sql";
                let _sql = " delete from xx \n";
                let _where = " \n where 1=1 \n";
                $that.viewDbDataInfo.dataCols.forEach((item)=>{
                    _where += (" and "+item.name+"='"+item.value+"'\n")
                });
                _sql += (_where +";");

                $that.viewDbDataInfo.sql = _sql;
            },
            _viewDbDataModal:function(){
                $that.viewDbDataInfo.view = "form";

            }
        }
    });

})(window.vc, window.vc.component);