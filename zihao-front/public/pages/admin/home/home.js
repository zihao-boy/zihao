/**
    入驻小区
**/
(function(vc){
    vc.extends({
        data:{
            communityInfo:{
            }
        },
        _initMethod:function(){
        },
        _initEvent:function(){
            vc.on('home','listMyCommunity',function(_param){
                  vc.component.listMyCommunity();
            });
        },
        methods:{
            
        }
    });
})(window.vc);