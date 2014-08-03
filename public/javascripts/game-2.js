var scene = new THREE.Scene();
var viewSize = 900;
//var aspectRatio = window.innerWidth / window.innerHeight;
var aspect = window.innerWidth / window.innerHeight;
var d = 20;
camera = new THREE.OrthographicCamera( - d * aspect, d * aspect, d, - d, 1, 1000 );

camera.position.set( 20, 20, 20 ); // all components equal
camera.lookAt( scene.position ); // or the origin
//var camera = new THREE.PerspectiveCamera(75, window.innerWidth/window.innerHeight, 0.1, 1000);
//var camera = new THREE.OrthographicCamera(-aspectRatio * viewSize / 2, aspectRatio * viewSize / 2,
//                                          viewSize / 2, -viewSize / 2, 1, 1000);

var renderer = new THREE.WebGLRenderer();
renderer.setSize(window.innerWidth, window.innerHeight);
document.body.appendChild(renderer.domElement);

var geometry = new THREE.BoxGeometry(1, 0, 1);
var material = new THREE.MeshBasicMaterial({color: 0x80CF5A});
var cube = new THREE.Mesh(geometry, material);
scene.add(cube);

//camera.position.z = 1000;

var render = function () {
  requestAnimationFrame(render);

  //cube.rotation.x += 0.1;
  //cube.rotation.y += 0.1;

  renderer.render(scene, camera);
};

render();
