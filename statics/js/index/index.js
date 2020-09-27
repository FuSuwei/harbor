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

function addArticle(dataList) {
    for (var i = 0; i < dataList.length; i++) {
        var article = $(getArticle());
        article.find(".post-title").text(dataList[i]["title"]);
        article.find(".summary").text(dataList[i]["summary"]);
        article.find(".time").text(formatDate(parseInt(dataList[i]["created_on"])));
        for (var j=0; j<dataList[i]["categories"].length;j++){
            var a = $("<a href=\"\"></a>");
            a.attr("href", "www.baidu.com");
            a.html(dataList[i]["categories"][j]["name"]);
            article.find(".meta-post").append(a)
        }

        $(".posts-list").append(article);
    }
}

function formatDate(time) {
    var now = new Date(time*1000);
    var year=now.getFullYear();  //取得4位数的年份
    var month=now.getMonth()+1;  //取得日期中的月份，其中0表示1月，11表示12月
    var date=now.getDate();      //返回日期月份中的天数（1到31）
    return year+"-"+month+"-"+date;
}

function getArticle() {
    return `
<article class="post-preview">
        <a href="">
            <h2 class="post-title"></h2>
        </a>
        <div class="post-entry">

            <p></p>
            <p class="summary"></p>
            <p></p>
            <a href="" class="post-read-more">Read More</a>

        </div>

        <div class="postmeta">
    <span class="meta-post">
        <i class="fa fa-calendar-alt" style="padding-right: 5px"></i>
        <i style="padding-right: 10px" class="time"></i>
        <i class="fa fa-folder-open" style="padding-right: 10px"></i>
        <a href=""></a>
        <a href=""></a>
    </span>

        </div>
</article>`
}
$(document).ready(function(){
    $(".nextPage").click(function () {
        $.ajax({
            type: "GET",
            url: "api/v1/getArticleList",
            success: function (result) {
                if (result.code == 200) {
                    var dataList = $.parseJSON(result["data"]);
                    $(".post-preview").remove();
                    addArticle(dataList);
                } else {
                    alert(result["msg"] + "\n" + result["errors"])
                }
            }
        })
    });

})


