var connection = require('./connection'),
    messages = require('./messages');

function moveAvatar(y, x) {
  connection.send(JSON.stringify(messages.buildMove(y, x)));
}

function moveUp(e) {
  moveAvatar(-1, 0);
}
function moveDown(e) {
  moveAvatar(1, 0);
}
function moveLeft(e) {
  moveAvatar(0, -1);
}
function moveRight(e) {
  moveAvatar(0, 1);
}

// game loop optimized keyboard handling
kd.UP.down(moveUp);
kd.DOWN.down(moveDown);
kd.LEFT.down(moveLeft);
kd.RIGHT.down(moveRight);

module.exports = kd;
