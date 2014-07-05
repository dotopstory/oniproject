var EventEmitter = require('events').EventEmitter;

function Net(url) {
	this.events = new EventEmitter();
	var websocket = new WebSocket(url);
	this.ws = websocket;
	var that = this;
	websocket.binaryType = "arraybuffer";
	websocket.onopen = function(event) {
		console.info("Net open", url);
		that.events.emit("open", event);
	};
	websocket.onerror = function(event) {
		console.info("Net error", event);
		that.events.emit("error", event);
	};
	websocket.onclose = function(event) {
		console.info("Net close");
		that.events.emit("close", event);
	};
	websocket.onmessage = function(event) {
		that.events.emit("message", CBOR.decode(event.data), event);
	};
}
Net.prototype = EventEmitter.prototype;
Net.prototype.constructor = Net;
Net.prototype.close = function() {
	this.ws.close();
}
Net.prototype.send = function(message) {
	this.ws.send(CBOR.encode(message));
}

module.exports = Net;
