var s;
var ui = document.querySelectorAll('#coords')[0];

var updateUi = function (z, x, y, cz, cx, cy) {
  ui.textContent = cz + ", " + cx + ", " + z + ", " + x;
};

module.exports = {
  updateUi: updateUi,
};
