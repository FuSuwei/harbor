function changeStyle4() {
    var butt = document.getElementById("dark-mode-toggle");
    var obj = document.getElementById("css");
    calssName = butt.className;
    if (calssName.indexOf("fa-moon") > 0) {
        butt.className = "fas fa-sun fa-lg";
        obj.setAttribute("href", "statics/css/sunbackground.css");
    } else {
        butt.className = "fas fa-moon fa-lg";
        obj.setAttribute("href", "statics/css/moonbackground.css");
    }
}

function formatDate(time) {
    var now = new Date(time * 1000);
    var year = now.getFullYear();  //取得4位数的年份
    var month = now.getMonth() + 1;  //取得日期中的月份，其中0表示1月，11表示12月
    var date = now.getDate();      //返回日期月份中的天数（1到31）
    return year + "-" + month + "-" + date;
}
