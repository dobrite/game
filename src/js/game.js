var camera = require('./camera'),
    connection = require('./connection'),
    kd = require('./keyboard'),
    scene = require('./scene');


var aspect = window.innerWidth / window.innerHeight;


//var axis = new THREE.AxisHelper(128);
//axis.position.set(0, 0, 0);
//scene.add(axis);
//
//var gridHelper = new THREE.GridHelper(1024, 32);
//scene.add(gridHelper);

//var ch = new THREE.CameraHelper(camera);
//scene.add(ch);

var renderer = new THREE.WebGLRenderer();

renderer.setSize(config.SCENE_WIDTH, config.SCENE_HEIGHT);
document.body.appendChild(renderer.domElement);

var ambientLight = new THREE.AmbientLight(0x10);
scene.add(ambientLight);

var directionalLight = new THREE.DirectionalLight(0xffffff);
directionalLight.position.x = 500;
directionalLight.position.y = 1000;
directionalLight.position.z = 500;
directionalLight.position.normalize();
scene.add(directionalLight);

function start () {
  function animate() {
    kd.tick();
    requestAnimationFrame(animate);
    //var timer = Date.now() * 0.0001;

    //camera.position.x = Math.cos(timer) * 200;
    //camera.position.z = Math.sin(timer) * 200;
    //camera.position.y = Math.tan(timer) * 200;
    //camera.lookAt(scene.position);
    renderer.render(scene, camera);
  }
  requestAnimationFrame(animate);
}

start();
