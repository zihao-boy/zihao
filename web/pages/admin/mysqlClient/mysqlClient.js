(function(vc) {

    vc.extends({
        data: {
            mysqlClientInfo: {

            }
        },
        _initMethod: function() {

            $("#text").setTextareaCount({
                width: "30px",
                bgColor: "#FFF",
                color: "#000",
                display: "inline-block"
            });
        },
        _initEvent: function() {

        },
        methods: {

            _customKeypress: function() {
                let typeSql = window.getSelection().toString();

                console.log(typeSql);
            }
        }
    });

})(window.vc);