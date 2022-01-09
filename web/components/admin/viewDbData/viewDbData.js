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
                let _sql = " insert into  (</br>";

                let _values = " </br>) values (</br>";

                $that.viewDbDataInfo.dataCols.forEach((item)=>{

                    _sql += (item.name+",</br>")
                    _values += ("'"+item.value +"',</br>")
                });

                if(_sql.endsWith(",</br>")){
                    _sql = _sql.substring(0,_sql.length -6);
                    _values = _values.substring(0,_values.length -6);
                }

                _sql += (_values +"</br>);</br>");

                $that.viewDbDataInfo.sql = _sql;
            },
            _copyUpdateSql:function(){
                $that.viewDbDataInfo.view = "sql";
                let _sql = " update xx set </br>";

                let _where = " </br> where 1=1 </br>";

                $that.viewDbDataInfo.dataCols.forEach((item)=>{
                    _sql += (item.name+"='"+item.value+"',</br>")
                    _where += (" and "+item.name+"='"+item.value+"'</br>")
                });

                if(_sql.endsWith(",</br>")){
                    _sql = _sql.substring(0,_sql.length -6);
                }

                _sql += (_where +";");

                $that.viewDbDataInfo.sql = _sql;
            },
            _copyDeleteSql:function(){
                $that.viewDbDataInfo.view = "sql";
                let _sql = " delete from xx </br>";
                let _where = " </br> where 1=1 </br>";
                $that.viewDbDataInfo.dataCols.forEach((item)=>{
                    _where += (" and "+item.name+"='"+item.value+"'</br>")
                });
                _sql += (_where +";");

                $that.viewDbDataInfo.sql = _sql;
            },
        }
    });

})(window.vc, window.vc.component);