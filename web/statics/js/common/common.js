// JS中时间戳转日期格式（YYYY-MM-dd HH:mm:ss）  // 1566981707
function formatUnixtimestamp(unixtimestamp, type) {
    var stamp = unixtimestamp;
    // console.log("unixtimestamp",unixtimestamp)
    if (stamp !== 0) {
        var unixtimestamp = new Date(unixtimestamp * 1000);
        var year = 1900 + unixtimestamp.getYear();
        var month = "0" + (unixtimestamp.getMonth() + 1);
        var date = "0" + unixtimestamp.getDate();
        var hour = "0" + unixtimestamp.getHours();
        var minute = "0" + unixtimestamp.getMinutes();
        var second = "0" + unixtimestamp.getSeconds();
        if (type === 1) {
            return year + "-" + month.substring(month.length - 2, month.length) + "-" + date.substring(date.length - 2, date.length)
        } else if (type === 11) {
            return year + "-" + month.substring(month.length - 2, month.length) + "-" + date.substring(date.length - 2, date.length)
                + "_" + hour.substring(hour.length - 2, hour.length) + ":"
                + minute.substring(minute.length - 2, minute.length) + ":"
                + second.substring(second.length - 2, second.length);
        } else if (type === 22) {
            return hour.substring(hour.length - 2, hour.length) + ":"
                + minute.substring(minute.length - 2, minute.length) + ":"
                + second.substring(second.length - 2, second.length);
        } else {
            return year + "-" + month.substring(month.length - 2, month.length) + "-" + date.substring(date.length - 2, date.length)
                + " " + hour.substring(hour.length - 2, hour.length) + ":"
                + minute.substring(minute.length - 2, minute.length) + ":"
                + second.substring(second.length - 2, second.length);
        }
    } else {
        return ""
    }
}

// 日期轉換爲時間錯
function date2Stamp(pram) {
    if (pram.length <= 4) {
        return 0
    } else {
        // var date = new Date('2014-04-23 18:55:49:123');
        var date = new Date(pram);
        // 有三种方式获取
        //var time1 = date.getTime();
        //var time2 = date.valueOf();
        var time3 = Date.parse(date);
        //console.log(time1);//1398250549123
        //console.log(time2);//1398250549123
        //console.log(time3);//1398250549000
        return Number(time3 / 1000)
    }
}

// 用JS获取地址栏参数的方法（超级简单）
function GetQueryString(key) {
    var reg = new RegExp("(^|&)" + key + "=([^&]*)(&|$)");
    var result = window.location.search.substr(1).match(reg);
    return result ? decodeURIComponent(result[2]) : null;
}

//替换指定传入参数的值,paramName为参数,replaceWith为新值
function ReplaceParamVal(paramName, replaceWith) {
    var oUrl = this.location.href.toString();
    var re = eval('/(' + paramName + '=)([^&]*)/gi');
    var nUrl = oUrl.replace(re, paramName + '=' + replaceWith);
    // this.location = nUrl;
    if (!!window.Worker) {
        history.pushState(null, null, nUrl);  //将网址设置
    } else {
        location.hash = nUrl;
    }
    // window.location.href=nUrl
}

// 判断字符串是否为JSON格式
function isJSON(str) {
    if (typeof str == 'string') {
        try {
            var obj = JSON.parse(str);
            if (typeof obj == 'object' && obj) {
                return true;
            } else {
                return false;
            }

        } catch (e) {
            //console.log('error：'+str+'!!!'+e);
            return false;
        }
    }
    return false;
    //console.log('It is not a string!')
}

// 字符串(yyyymmdd)格式转换成日期格式（yyyy-mm-dd）
function formatedDate(val) {
    if (!isNaN(val) && val.length === 8) {
        var dateString = val;
        var pattern = /(\d{4})(\d{2})(\d{2})/;
        var formatedDate = dateString.replace(pattern, '$1-$2-$3');
        return formatedDate;
    }
    return ""
}

// JS解决URL传递参数中文乱码问题
function getUrlParam(name) {
    // 用该属性获取页面 URL 地址从问号 (?) 开始的 URL（查询部分）
    var url = window.location.search;
    // 正则筛选地址栏
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    // 匹配目标参数
    var result = url.substr(1).match(reg);
    //返回参数值
    return result ? decodeURIComponent(result[2]) : null;
}

// 截取字符串，多余的部分用...代替
function getString(str, len) {
    var strlen = 0;
    var s = "";
    for (var i = 0; i < str.length; i++) {
        if (str.charCodeAt(i) > 128) {
            strlen += 2;
        } else {
            strlen++;
        }
        s += str.charAt(i);
        if (strlen >= len && s.length < str.length && s.length < len) {
            return s + "...";
        }
    }
    return s;
}

// 判断是否包含特殊符号
function checkStr(str) {
    // var pattern = new RegExp("[`~!@#$^&*()=|{}':;',\\[\\].<>《》/?~！@#￥……&*（）——|{}【】‘；：”“'。，、？ ]");
    var pattern = new RegExp("[`~!@$^&*()=|{}':;',\\[\\].<>《》/?~！@#￥……&*（）——|{}【】‘；：”“'。，、？ ]");
    if (pattern.test(str)) {
        return true;
    }
    return false;
}

//只能输入汉字和英文字母
function checkVal(t) {
    var re = /^[\u4e00-\u9fa5a-z]+$/gi;//只能输入汉字和英文字母
    if (re.test(t)) {
        return true;
    } else {
        return false;
    }
}

// 号码隐藏中间N位
function getXCode(code, start, end, typeStr) {
    if (code === "") {
        return "";
    } else if (typeStr === "phone" && !(/^1[3456789]\d{9}$/.test(code))) {
        return code;
    } else if (typeStr === "cardNo" && !checkIDCard(code)) {
        return code;
    }
    return code.substring(0, start) + "****" + code.substr(code.length - end);
}

// 函数参数必须是字符串，因为二代身份证号码是十八位，而在javascript中，十八位的数值会超出计算范围，造成不精确的结果，导致最后两位和计算的值不一致，从而该函数出现错误。
// 详情查看javascript的数值范围
function checkIDCard(idcode) {
    // 加权因子
    var weight_factor = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2];
    // 校验码
    var check_code = ['1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'];
    var code = idcode + "";
    var last = idcode[17];//最后一位
    var seventeen = code.substring(0, 17);
    // ISO 7064:1983.MOD 11-2
    // 判断最后一位校验码是否正确
    var arr = seventeen.split("");
    var len = arr.length;
    var num = 0;
    for (var i = 0; i < len; i++) {
        num = num + arr[i] * weight_factor[i];
    }
    // 获取余数
    var resisue = num % 11;
    var last_no = check_code[resisue];
    // 格式的正则
    // 正则思路
    /*
    第一位不可能是0
    第二位到第六位可以是0-9
    第七位到第十位是年份，所以七八位为19或者20
    十一位和十二位是月份，这两位是01-12之间的数值
    十三位和十四位是日期，是从01-31之间的数值
    十五，十六，十七都是数字0-9
    十八位可能是数字0-9，也可能是X
    */
    var idcard_patter = /^[1-9][0-9]{5}([1][9][0-9]{2}|[2][0][0|1][0-9])([0][1-9]|[1][0|1|2])([0][1-9]|[1|2][0-9]|[3][0|1])[0-9]{3}([0-9]|[X])$/;
    // 判断格式是否正确
    var format = idcard_patter.test(idcode);
    // 返回验证结果，校验码和格式同时正确才算是合法的身份证号码
    return last === last_no && format ? true : false;
}

/**
 * 复制内容到粘贴板
 * content : 需要复制的内容
 * message : 复制完后的提示，不传则默认提示"复制成功"
 */
function copyToClip(content, message) {
    var aux = document.createElement("input");
    aux.setAttribute("value", content);
    document.body.appendChild(aux);
    aux.select();
    document.execCommand("copy");
    document.body.removeChild(aux);
    if (message == null) {
        layer.msg("复制成功");
    } else {
        layer.msg(message);
    }
}

//函数来判断一个值是数字的最正确的方法就是： true:数值型的，false：非数值型
function myIsNaN(value) {
    return typeof value === 'number' && !isNaN(value);
}

/**
 * 获取下一个月 :: 测试 alert(getNextMonth("2014-12-25",2));
 * @date 格式为yyyy-mm-dd的日期，如：2014-01-25
 */
function getNextMonth(date, tp) {
    var arr = date.split('-');
    var year = arr[0]; //获取当前日期的年份
    var month = arr[1]; //获取当前日期的月份
    var day = arr[2]; //获取当前日期的日
    var days = new Date(year, month, 0);
    days = days.getDate(); //获取当前日期中的月的天数
    var year2 = year;
    var month2 = parseInt(month) + 1;
    if (month2 == 13) {
        year2 = parseInt(year2) + 1;
        month2 = 1;
    }
    var day2 = day;
    var days2 = new Date(year2, month2, 0);
    days2 = days2.getDate();
    if (day2 > days2) {
        day2 = days2;
    }
    if (month2 < 10) {
        month2 = '0' + month2;
    }
    if (tp === 1) {
        var t2 = year2 + '-' + month2 + '-01';
        return t2;
    } else {
        var t2 = year2 + '-' + month2 + '-' + day2;
        return t2;
    }
}

/***
 * 获取指定日期上一天、下一天 ::: 测试 getNextDate('2018-02-28', 1) getNextDate('2018-03-01', -1)
 * @param date
 * @param day
 * @returns {string}
 */
function getNextDate(date, day) {
    var dd = new Date(date);
    dd.setDate(dd.getDate() + day);
    var y = dd.getFullYear();
    var m = dd.getMonth() + 1 < 10 ? "0" + (dd.getMonth() + 1) : dd.getMonth() + 1;
    var d = dd.getDate() < 10 ? "0" + dd.getDate() : dd.getDate();
    return y + "-" + m + "-" + d;
}

// 获取产品名称 -- 订单列表
function productsForName(arr) {
    var productsName = "";
    if (arr.length > 0) {
        for (var n = 0; n < arr.length; n++) {
            if (arr[n].status === 1) {
                productsName = arr[n].name;
            }
        }
    }
    return productsName;
}

// 获取产品名称Id -- 订单列表
function productsForId(arr) {
    var productsId = [];
    if (arr.length > 0) {
        for (var n = 0; n < arr.length; n++) {
            productsId.push(arr[n].id);
        }
    }
    return productsId;
}

// 直接下载，用户体验好
function downloadFileUrl(url) {
    if (url !== "") {
        var $form = $('<form method="GET"></form>');
        $form.attr('action', url);
        $form.appendTo($('body'));
        $form.submit();
    }
}

// 保留两位小数
/**
 * @return {number}
 */
function DecimalFormat2(value) {
    //console.log("DecimalFormat2",value);
    var value = Math.round(parseFloat(value) * 100) / 10000;
    var xsd = value.toString().split(".");
    if (xsd.length === 1) {
        value = value.toString() + ".00";
        return value;
    }
    if (xsd.length > 1) {
        if (xsd[1].length < 2) {
            value = value.toString() + "0";
        }
        return value;
    }
    // return Math.round(num*100)/100;
}

/**
 * 检测字符串字符包含特殊字符串，含有为false
 * 字符”+”,”/”,”?”,”%”,”#”,”&”,”=”都将被转码
 * url地址中不能出现哪些特殊字符
 * 使用微信支付的签名函数的时候，%、&、=、/、\、=、+、#、？、。、- 和空格。这些特殊字符不应该直接出现在url中。否则签名会失败的。
 */
function checkStrIsLegal(str) {
    var checked = ['+', '/', '?', '%', '#', '&', '=', ' ', '\\'];
	var tag = true;
    for (var n = 0; n < checked.length; n++) {
		if(str.indexOf(checked[n]) !== -1){ // 如果要检索的字符串值没有出现，则该方法返回 -1。
			tag = false;
			break;
		}
    }
	return tag;
}

/** 校验微信号
 1、可以使用6-20个子母、数字、下划线和减号
 2、必须以字母开头（字母不区分大小写）
 3、不能设置中文
 */
function checkWeChatIsLegal(str) {
    var re = /^[a-zA-Z]([-_a-zA-Z0-9]{5,19})+$/;
    if(str.length === 0){
        return false;
    }else if(!re.test(str)){
        return false;
    }
    return true;
}

// 简单的非正规的身份证验证
function checkIdCard(idcard) {
    const regIdCard = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/;
    if (!regIdCard.test(idcard)) {
        // errorMsg = '身份证号填写有误';
        return false;
    } else {
        return true;
    }
}

// 修改url地址参数
function changeURLArg(url,arg,arg_val){
    var pattern=arg+'=([^&]*)';
    var replaceText=arg+'='+arg_val;
    if(url.match(pattern)){
        var tmp='/('+ arg+'=)([^&]*)/gi';
        tmp=url.replace(eval(tmp),replaceText);
        return tmp;
    }else{
        if(url.match('[\?]')){
            return url+'&'+replaceText;
        }else{
            return url+'?'+replaceText;
        }
    }
}

//替换指定传入参数的值,paramName为参数,replaceWith为新值
function replaceParamVal(paramName,replaceWith) {
    var oUrl = this.location.href.toString();
    var re=eval('/('+ paramName+'=)([^&]*)/gi');
    var nUrl = oUrl.replace(re,paramName+'='+replaceWith);
    this.location = nUrl;
    window.location.href=nUrl
}

// 输出03:05:59  时分秒
function secondToDate(result) {
    var h = Math.floor(result / 3600) < 10 ? Math.floor(result / 3600) : Math.floor(result / 3600);
    var m = Math.floor((result / 60 % 60)) < 10 ?  Math.floor((result / 60 % 60)) : Math.floor((result / 60 % 60));
    var s = Math.floor((result % 60)) < 10 ?  Math.floor((result % 60)) : Math.floor((result % 60));
    if(result >= 3600){
        return h + "小时" + m + "分" + s+"秒";
    }else if(result >= 60){
        return m + "分" + s+"秒";
    }else{
        return "秒";
    }
}

// 排序函数
function objArraySort(objArr, key) {
    let result = objArr.slice(0);
    return result.sort((a, b) => a[key] - b[key]);
}

// 判断一个数组是否包含另一个数组, aa 包含 bb
function IsContained(aa, bb) {
    if(!(aa instanceof Array) || !(bb instanceof Array) || ((aa.length < bb.length))) {
        return false;
    }
    for (var i = 0; i < bb.length; i++) {
        var flag = false;
        for(var j = 0; j < aa.length; j++){
            if(aa[j] == bb[i]){
                flag = true;
                break;
            }
        }
        if(flag == false){
            return flag;
        }
    }
    return true;
}

// 去除一个数组中与另一个数组中的值相同的元素
function DelArrForAlike(a,b) {
    //var a = [1,2,3,4,5,6,6,'a','b'];
    //var b = [2,3,4];
    var c = [];
    for(var i of a){
        if(b.indexOf(i) === -1){
            c.push(i)
        }
    }
    return c;
    // console.log(c)  Array [1, 5, 6, 6, "a", "b"]
}

/***
 * 获取当前浏览器类型
 */
function getMyBrowser() {
    var userAgent = navigator.userAgent; //取得浏览器的userAgent字符串
    var isOpera = userAgent.indexOf("Opera") > -1;
    if (isOpera) { //判断是否Opera浏览器
        return "Opera"
    }
    if (userAgent.indexOf("Firefox") > -1) { //判断是否Firefox浏览器
        return "FF";
    }
    if (userAgent.indexOf("Chrome") > -1) {
        return "Chrome";
    }
    if (userAgent.indexOf("Safari") > -1) { //判断是否Safari浏览器
        return "Safari";
    }
    if (userAgent.indexOf("compatible") > -1 && userAgent.indexOf("MSIE") > -1 && !isOpera) { //判断是否IE浏览器
        return "IE";
    }
}

function checkInt(input) {
    var re = /^[0-9]+$/;//判断字符串是否为数字//判断正整数/[1−9]+[0−9]∗]∗/
    return re.test(input)
}
