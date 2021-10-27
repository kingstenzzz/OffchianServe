var UID
var ChannelID
var URL ="http://localhost:8080"
function player_GetInfo() {
    document.getElementById("content").innerHTML="address:123"

}
function JoinChannel() {
    document.getElementById("content").innerHTML="joinChannel"
    //document.getElementById("btn_join").disabled = true;
    //document.getElementById("btn_create").disabled = true;
    document.getElementById("playerContent").removeAttribute("hidden")
    let playerAddr = document.getElementById("addr").value
    let deposit =document.getElementById("deposit").value
    ChannelID=document.getElementById("channel").value
    let url=URL+"/player?action=join&"+"addr="+playerAddr+"&deposit="+deposit+"&channel="+ChannelID
    let response = POST(url)
    let obj =JSON.parse(response)
    if (obj.code !=200){
        info=obj.err
        alert(info)
    }
    else {
        UID = obj.uid
        ShowInfo()
    }
}

function CreateChannel() {
    document.getElementById("content").innerHTML="createChannel"
    //document.getElementById("btn_create").disabled = true;
    //document.getElementById("btn_join").disabled = true;
    ChannelID=document.getElementById("channel").value
    let url =URL+"?channelId="+ChannelID

    let response =POST(url)
    let obj =JSON.parse(response)
    ChannelID = obj.channel
    let info = 'create channel'+ChannelID

    if (obj.code !=200)
    {
        info=obj.err
        alert(info)
    }
    else{
        info = 'create channel '+ChannelID+' successfully'
    }

}
function SendTo() {
    let toAddr =document.getElementById("text_to").value
    let amount =document.getElementById("text_amount").value
    let url=URL+"/player?action=send&uid="+UID+"&to="+toAddr+"&amount="+amount+"&channel="+ChannelID
    let response = POST(url)
    let obj =JSON.parse(response)
    if (obj.code !=200)
    {
        info=obj.err
        alert(info)

    }
    else{
        ShowInfo()
    }
}
function ExitChannel() {
    ChannelID = document.getElementById("channel").value
    let url = URL + "/player?action=exit&uid=" + UID + "&channel=" + ChannelID
    let response = POST(url)
    let obj = JSON.parse(response)
    if (obj.code != 200) {
        info = obj.err
        alert(info)


    } else {
        document.getElementById("content").innerHTML = "Delete successfully"
    }

}
function Dispute() {
    document.getElementById("content").innerHTML="dispute"
}
function CloseChannel(){
    document.getElementById("content").innerHTML="Close"

}
function ShowInfo(){
    ChannelID=document.getElementById("channel").value
    document.getElementById("playerContent").removeAttribute("hidden")
    let playerAddr = document.getElementById("addr").value
    let response = GET(URL+"/player?playerAddr="+playerAddr+"&channel="+ChannelID)
    let playerobj =JSON.parse(response)
    //document.getElementById("content").innerHTML=response
    if (playerobj.code == 404) {
        info = playerobj.err
        alert(info)
    } else {

        var info = " Channel:" + ChannelID + "<br/><br/>" +
            "UserID:" + playerobj.Uid + "<br/>" +
            "Address:" + playerobj.Addr + "<br/>" +
            "Credit:" + playerobj.Credit + "<br/>" +
            "Withdrawn:" + playerobj.Withdrawn + "<br/>" +
            "Deposit:" + playerobj.Deposit
        document.getElementById("content").innerHTML = info
        UID = playerobj.Uid
    }


}


function GET(url){
    let httpRequest = new XMLHttpRequest();//第一步：创建需要的对象

    httpRequest.open('GET', url ,false);
    httpRequest.send();//发送请求 将json写入send中

    //httpRequest.onreadystatechange = function ()
    {//请求后的回调接口，可将请求成功后要执行的程序写在其中
        if (httpRequest.readyState === 4 && httpRequest.status === 200) {//验证请求是否发送成功
            var json = httpRequest.responseText;//获取到服务端返回的数据
            return json
        }
    };
    return null

}


function POST(url){
    let httpRequest = new XMLHttpRequest();//第一步：创建需要的对象

    httpRequest.open('POST', url ,false);
    httpRequest.setRequestHeader("Content-type", "application/json");//设置请求头 注：post方式必须设置请求头（在建立连接后设置请求头）var obj = { name: 'zhansgan', age: 18 };

    httpRequest.send();//发送请求 将json写入send中
    /**
     * 获取数据后的处理程序
     */
   // httpRequest.onreadystatechange = function ()
    {//请求后的回调接口，可将请求成功后要执行的程序写在其中
        if (httpRequest.readyState === 4 && httpRequest.status === 200) {//验证请求是否发送成功
            var json = httpRequest.responseText;//获取到服务端返回的数据
            return json
        }
    };
}

function getJson(json,key){
    var jsonObj={"name":"傅红雪","age":"24","profession":"刺客"};
    //1、使用eval方法
    var eValue=eval('jsonObj.'+key);
    alert(eValue);
    //2、遍历Json串获取其属性
    for(var item in jsonObj){
        if(item==key){	//item 表示Json串中的属性，如'name'
            var jValue=jsonObj[item];//key所对应的value
            alert(jValue);
        }
    }
    //3、直接获取
    alert(jsonObj[''+key+'']);
}
