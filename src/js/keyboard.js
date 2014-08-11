var connection = require('./connection'),
    messages = require('./messages');

function moveAvatar(z, x) {
  connection.send(JSON.stringify(messages.buildMove(z, x)));
}

function moveUp() {
  moveAvatar(-1, 0);
}
function moveDown() {
  moveAvatar(1, 0);
}
function moveLeft() {
  moveAvatar(0, -1);
}
function moveRight() {
  moveAvatar(0, 1);
}

// game loop optimized keyboard handling
kd.UP.down(moveUp);
kd.DOWN.down(moveDown);
kd.LEFT.down(moveLeft);
kd.RIGHT.down(moveRight);

module.exports = kd;
