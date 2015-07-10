$(document).ready(function() {
    var scene = new THREE.Scene();
    var camera = new THREE.PerspectiveCamera(45, window.innerWidth/window.innerHeight, 0.1, 1000 );

    var renderer = new THREE.WebGLRenderer();
    renderer.setSize( window.innerWidth, window.innerHeight );
    var renderJobs = [ ];

    $('#world-display').append(renderer.domElement);

    /* Setup earth */
    var geometry         = new THREE.SphereGeometry(0.5, 64, 64);
    var material         = new THREE.MeshPhongMaterial();
    material.map         = THREE.ImageUtils.loadTexture('img/new.png');
    material.fog = false;
    material.shininess = 0;

    var earthMesh = new THREE.Mesh(geometry, material)
    scene.add( earthMesh );

    camera.position.z = 0.85;
    var cam_angle = 0;
    var cam_distance = 1;

    var mouse = {x : 0, y : 0};
    var down = false;
    var down_pos = { x:0, y:0 };
    
    document.addEventListener('mousedown', function(event){
        down = true;
        mouse.x = (event.clientX / window.innerWidth ) 
        mouse.y = (event.clientY / window.innerHeight)
        down_pos.x = mouse.x;
        down_pos.y = mouse.y;
        
        var vector = new THREE.Vector3((event.clientX / window.innerWidth) * 2 - 1, -(event.clientY / window.innerHeight) * 2 + 1, 1, 10000);
        vector.unproject(camera);

        var raycaster = new THREE.Raycaster(camera.position, vector.sub(camera.position).normalize());

        var intersect = raycaster.intersectObject(earthMesh);
        if (intersect.length > 0) {
            console.debug(intersect);
        }
    });
    document.addEventListener('mouseup', function(event){
        down = false;
    });
    document.addEventListener('mousemove', function(event){
        mouse.x = (event.clientX / window.innerWidth )
        mouse.y = (event.clientY / window.innerHeight)
        event.preventDefault();
    }, false)
    document.addEventListener("mousewheel", function(event) {
        w = event.wheelDeltaY;
        cam_distance -= Math.sign(w) * 0.02;
        if (cam_distance < 0.60) cam_distance = 0.60;
        camera.position.y = Math.sin(cam_angle) * cam_distance;
        camera.position.z = Math.cos(cam_angle) * cam_distance;
        camera.lookAt(earthMesh.position);
        event.preventDefault();
    }, false);
    renderJobs.push(function(delta, now){
        earthMesh.rotateY(0.00005);
        if (!down) return;

        m = { x: mouse.x - down_pos.x, y: mouse.y - down_pos.y };

        r = m.x * 10 * delta
        earthMesh.rotateY(r);

        cam_angle += m.y * 0.1;
        max_angle = Math.PI / 3;
        if (cam_angle > max_angle) cam_angle = max_angle;
        if (cam_angle < -max_angle) cam_angle = -max_angle;

        camera.position.y = Math.sin(cam_angle) * cam_distance;
        camera.position.z = Math.cos(cam_angle) * cam_distance;
        camera.lookAt(earthMesh.position);
    })

    /* scene lighting */
    var ambient = new THREE.AmbientLight( 0x888888 )
    scene.add(ambient)
    var sol = new THREE.DirectionalLight( 0xcccccc, 0.65 )
    sol.position.set(3,3,25)
    scene.add(sol)
    

    var material = new THREE.LineBasicMaterial({
        color: 0x0000ff
    });

    var geometry = new THREE.Geometry();
    geometry.vertices.push(new THREE.Vector3(0, -1, 0));
    geometry.vertices.push(new THREE.Vector3(0, 1, 0));

    var line = new THREE.Line(geometry, material);
    scene.add(line);

    renderJobs.push(function(){
        renderer.render( scene, camera );       
    })
    
    var lastTimeMsec= null
    requestAnimationFrame(function animate(nowMsec){
        // keep looping
        requestAnimationFrame( animate );
        // measure time
        lastTimeMsec    = lastTimeMsec || nowMsec-1000/60
        var deltaMsec   = Math.min(200, nowMsec - lastTimeMsec)
        lastTimeMsec    = nowMsec
        // call each update function
        renderJobs.forEach(function(renderJob){
            renderJob(deltaMsec/1000, nowMsec/1000)
        })
    })
});
