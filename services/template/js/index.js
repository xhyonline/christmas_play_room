//一般直接写在一个js文件中
layui.use(['layer', 'form', 'jquery'], function () {
    var layer = layui.layer
        , form = layui.form;
    var $ = layui.$
    maxWidht = 768
    if (maxWidht < document.body.clientWidth) {
        $("body").text("请使用移动端设备访问页面")
    }
    layer.msg('Hello World');
});