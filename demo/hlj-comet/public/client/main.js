let peer;
let cometHost = window.location.host;
let cssHost = window.location.hostname + ":" + 8010;

let tokenEle = document.getElementById("user-token");
let osEle = document.getElementById("os");
let cidEle = document.getElementById("cid");
let btnEle = document.getElementById("conn-btn");
let panelEle = document.getElementById("online-panel");

let subServiceEle = document.getElementById("sub-service");
let subSubjectEle = document.getElementById("sub-subject");
let subArgsEle = document.getElementById("sub-args");

let remoteFieldEle = document.getElementById("remote-field");
let remoteValueEle = document.getElementById("remote-value");

let msgServiceEle = document.getElementById("msg-service");
let msgSubjectEle = document.getElementById("msg-subject");
let msgPayloadEle = document.getElementById("msg-payload");

function connect() {
    let token = tokenEle.value;
    let os = osEle.value;
    let cid = parseInt(cidEle.value);
    let option = {
        city: cid,
        devicekind: os,
    };

    peer = new Peer(cometHost, token, option);
    peer.onOpen = function () {
        alert("connected");
    };
    peer.onMessage = function (msg) {
        let historyEle = document.getElementById("msg-history");
        let p = document.createElement("p");
        p.innerText = JSON.stringify(msg);
        historyEle.appendChild(p)
    };
    peer.onClose = function () {
        alert("closed");
    };
    peer.run();

    tokenEle.disabled = true;
    osEle.disabled = true;
    cidEle.disabled = true;
    btnEle.disabled = true;
    btnEle.innerText = "connected";
    panelEle.style.display = "";
}

function subscribe() {
    peer.send("comet", "subscribe", {
        service: subServiceEle.value,
        subject: subSubjectEle.value,
        args: subArgsEle.value
    })
}

function unsubscribe() {
    peer.send("comet", "unsubscribe", {
        service: subServiceEle.value,
        subject: subSubjectEle.value,
        args: subArgsEle.value
    })
}

function modifyRemote() {
    let remote = {};
    remote[remoteFieldEle.value] = remoteValueEle.value;
    peer.send("comet", "set-remote", remote)
}

function send() {
    let data = JSON.parse(msgPayloadEle.value);
    peer.send(msgServiceEle.value, msgSubjectEle.value, data)
}

function testCors() {
    let result;
    getSmartReception(cssHost, 378, function(data) {
        result = data
    })
}
