var currentPage = 1;

function addArticle(dataList) {

    for (var i = 0; i < dataList.length; i++) {
        var article = $(getArticle());
        article.find(".post-title").text(dataList[i]["title"]);
        article.find(".summary").text(dataList[i]["summary"]);
        article.find(".time").text(formatDate(parseInt(dataList[i]["created_on"])));
        article.find(".info").attr("href", "/article/" + dataList[i]["uuid"]);
        for (var j = 0; j < dataList[i]["categories"].length; j++) {
            var a = $("<a href=\"\"></a>");
            a.attr("href", "www.baidu.com");
            a.html(dataList[i]["categories"][j]["name"]);
            article.find(".meta-post").append(a)
        }

        $(".posts-list").append(article);
    }
}

function getArticle() {
    return `
<article class="post-preview">
        <a href="" class="info">
            <h2 class="post-title"></h2>
        </a>
        <div class="post-entry">

            <p ></p>
            <p class="summary"></p>
            <p></p>
            <a href="" class="post-read-more info">Read More</a>

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

$(document).ready(function () {
    var BaseUrl = $("#baseUrl").attr("baseUrl");
    $(".nextPage").click(function () {
        $.ajax({
            type: "GET",
            url: BaseUrl,
            data: {page: currentPage + 1},
            success: function (result) {
                if (result.code === 200) {
                    if (result["isNextPage"] !== true) {
                        setNextPageAttr(true)
                    }
                    setPrevPageAttr(false);
                    var dataList = $.parseJSON(result["data"]);
                    $(".post-preview").remove();
                    addArticle(dataList);
                    currentPage += 1
                } else {
                    alert(result["msg"] + "\n" + result["errors"])
                }
            }
        })
    });
    $(".prevPage").click(function () {
        $.ajax({
            type: "GET",
            url: BaseUrl,
            data: {page: currentPage - 1},
            success: function (result) {
                if (currentPage === 2) {
                    setPrevPageAttr(true)
                }
                setNextPageAttr(false);
                if (result.code === 200) {
                    var dataList = $.parseJSON(result["data"]);
                    $(".post-preview").remove();
                    addArticle(dataList);
                    currentPage -= 1
                } else {
                    alert(result["msg"] + "\n" + result["errors"])
                }
            }
        })
    });
    $.ajax({
        type: "GET",
        url: BaseUrl,
        data: {page: currentPage},
        success: function (result) {
            if (result.code === 200) {
                if (result["isNextPage"] !== true) {
                    setNextPageAttr(true)
                }
                var dataList = $.parseJSON(result["data"]);
                addArticle(dataList);
            } else {
                alert(result["msg"] + "\n" + result["errors"])
            }
        }
    });
});


function setNextPageAttr(flag) {
    if (flag) {
        $(".nextPage").attr("style", "pointer-events: none; opacity: 0.5")
    } else {
        $(".nextPage").removeAttr("style")
    }
}


function setPrevPageAttr(flag) {
    if (flag) {
        $(".prevPage").attr("style", "pointer-events: none; opacity: 0.5")
    } else {
        $(".prevPage").removeAttr("style")
    }
}