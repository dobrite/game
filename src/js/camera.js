var config = require('./config');

var aspect = window.innerWidth / window.innerHeight;
var d = 2028;

//var camera = new THREE.OrthographicCamera( - d * aspect, d * aspect, d, - d, 1, 1000 );
//var camera = new THREE.OrthographicCamera(
//  window.innerWidth / -2,
//  window.innerWidth / 2,
//  window.innerHeight / 2,
//  window.innerHeight / -2,
//  -2000,
//  2000
//);

var camera = new THREE.OrthographicCamera(
  d / -2 * aspect,
  d / 2 * aspect,
  d / 2,
  d / -2,
  -2000,
  2000
);

camera.position.set(200, 200, 200);
//camera.rotation.order = 'YXZ';
//camera.rotation.y = - Math.PI / 4;
//camera.rotation.x = Math.atan( - 1 / Math.sqrt( 2 ) );
camera.up = new THREE.Vector3(0, 1, 0);
camera.lookAt(new THREE.Vector3(0, 0, 0));
//camera.left = -1000;
//camera.right = 1000;
//camera.top = 1000;
//camera.bottom = -1000;
//
//camera.near = 0;
//camera.far = 1000;

camera.updateProjectionMatrix();
module.exports = camera;
