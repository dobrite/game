var messageToHandler = require('./handlers/handlers');

var connection = new WebSocket('ws://localhost:3000/sock/');

connection.onopen = function () {
  console.log("connected");
};

connection.onerror = function (error) {
  console.log('WebSocket Error ' + error);
};

// Log messages from the server
connection.onmessage = function (e) {
  var message = JSON.parse(e.data);
  messageToHandler[message.event](message);
};

module.exports = connection;
