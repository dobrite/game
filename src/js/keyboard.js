var connection = require('./connection'),
    messages = require('./messages');
    throttle = require('lodash.throttle');

var moveAvatar = function (z, x) {
  connection.send(JSON.stringify(messages.buildMove(z, x)));
};

var moveUp = function () {
  moveAvatar(-1, 0);
};

var moveDown = function () {
  moveAvatar(1, 0);
};

var moveLeft = function () {
  moveAvatar(0, -1);
};

var moveRight = function () {
  moveAvatar(0, 1);
};

kd.UP.down(throttle(moveUp, 50));
kd.DOWN.down(throttle(moveDown, 50));
kd.LEFT.down(throttle(moveLeft, 50));
kd.RIGHT.down(throttle(moveRight, 50));

module.exports = kd;
