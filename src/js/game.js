var camera = require('./camera'),
    connection = require('./connection'),
    kd = require('./keyboard'),
    scene = require('./scene'),
    lighting = require('./lighting'),
    renderer = require('./renderer'),
    updateSun = require('./ui').updateSun;

var aspect = window.innerWidth / window.innerHeight;

stats = new Stats();
stats.domElement.style.position = 'absolute';
stats.domElement.style.top = '0px';
stats.domElement.style.zIndex = 100;

document.querySelectorAll('#stats')[0].appendChild(stats.domElement);

var axis = new THREE.AxisHelper(128);
axis.position.set(0, 0, 0);
scene.add(axis);

//scene.fog = new THREE.FogExp2(0xcccccc, 0.002);
//renderer.setClearColor(scene.fog.color, 1);

var gridHelper = new THREE.GridHelper(1640, 32);
scene.add(gridHelper);

//var ch = new THREE.CameraHelper(camera);
//scene.add(ch);

scene.add(lighting);

function start () {
  function animate() {
    kd.tick();
    requestAnimationFrame(animate);
    var timer = Date.now() * 0.0001;

    //lighting.position.z += 0.00005;
    //lighting.position.x += 0.0001;
    updateSun(lighting);
    //camera.position.z = Math.sin(timer) * 200;
    //camera.position.y = Math.tan(timer) * 200;
    //camera.lookAt(scene.position);
    renderer.render(scene, camera);
    stats.update();
  }
  requestAnimationFrame(animate);
}

start();
