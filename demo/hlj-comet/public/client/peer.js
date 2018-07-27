class Peer {
    constructor(host, token, option) {
        this.pings = 0;
        this.pongs = 0;
        this.onOpen = function () {
        };
        this.onMessage = function (msg) {
        };
        this.onError = function () {
        };
        this.onClose = function () {
        };

        let that = this;
        option['token'] = token;
        this.conn = new WebSocket("ws://" + host + "/api/ws?" + buildQuery(option));
        this.conn.onopen = function () {
            that.onOpen()
        };
        this.conn.onmessage = function (evt) {
            let data = JSON.parse(evt.data);
            if (data.service == "ws" && data.subject == "ping") {
                that.pong(data.payload);
            } else if (data.service == "ws" && data.subject == "pong") {
                that.pongs++;
            } else {
                that.onMessage(data)
            }
        };
        this.conn.onerror = function () {
            that.onError()
        };
        this.conn.onclose = function () {
            that.onClose()
        };
    }

    run() {
        let that = this;
        let id = self.setInterval(function () {
            if (Math.abs(that.pings - that.pongs) < 5) {
                that.ping(Date.now());
            } else {
                that.conn.close();
                self.clearInterval(id);
            }
        }, 3000);
    }

    send(service, subject, payload) {
        let msg = {
            service: service,
            subject: subject,
            payload: payload,
        };
        this.conn.send(JSON.stringify(msg))
    }

    ping(seq) {
        let msg = {
            service: "ws",
            subject: "ping",
            payload: seq,
        };
        this.conn.send(JSON.stringify(msg));
        this.pings++;
    }

    pong(seq) {
        let msg = {
            service: "ws",
            subject: "pong",
            payload: seq,
        };
        this.conn.send(JSON.stringify(msg));
    }
}
