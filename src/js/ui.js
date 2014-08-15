var s;
var ui = document.querySelectorAll('#coords')[0];

var updateUi = function (z, x, cz, cx) {
  ui.textContent = cz + ", " + cx + ", " + z + ", " + x;
};

module.exports = {
  updateUi: updateUi,
};
