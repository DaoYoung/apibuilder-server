function getSmartReception(host, channelId, callback) {
    let url = 'http://' + host + "/api/channels/" + channelId + "/smart_reception";

    let ajax = new XMLHttpRequest();
    ajax.open('GET', url, true);
    ajax.onreadystatechange = function () {
        if (ajax.readyState === 4 && ajax.status >= 200 && ajax.status < 300 ) {
            let json = JSON.parse(ajax.responseText);
            callback(json['data']);
        }
    };

    ajax.send();
}