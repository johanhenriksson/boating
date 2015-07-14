var map = angular.module('map', []);

map.controller('MapView', function($scope) {
    var snap = Snap('.map-display');
    var map  = $('#map_render');
    var map_size = { x: 1009.0, y: 665.0 }
    var mouse = { drag: false }; 

    var half_height = parseInt(map.height()) / 2.0;
    var half_width  = parseInt(map.width()) / 2.0;
    console.log('viewport center: ' + half_width + ',' + half_height);
    var map_x = 0;
    var map_y = 0;

    var center = snap.circle(half_width, half_height, 50);

    function refresh() {
        //half_width = (map.width() - map_size.x) / 2.0;
        //half_height = (map.height() - map_size.y) / 2.0;
        var zoom = 0.5 * Math.pow($scope.zoom, 1.6) + 0.5;
        var x = map_x * zoom; 
        var y = map_y * zoom;
        var offset_x = -(half_width  - zoom * map_size.x / 2);
        var offset_y = -(half_height - zoom * map_size.y / 2);

        var transform = new Snap.Matrix();
        transform.translate(-offset_x, -offset_y)
        transform.scale(zoom);
        transform.translate(map_x, map_y);

        $scope.x = x;
        $scope.y = y;
        $scope.magnification = zoom

        g = snap.select('g');
        if (!g) return;
        g.transform(transform);
        g.attr({
            strokeWidth: 1 / zoom,
        });
    };

    $scope.init = function() {
        $scope.zoom = 1;

        Snap.load("img/map.svg", function(map) {
            var g = map.select('g');
            g.attr({
                fill: '#66CC99',
                stroke: '#555555',
                strokeWidth: 1,
            });
            snap.append(map);
            refresh();
        });

        $scope.$watch('zoom', refresh);
    };
    $scope.mouseDown = function($event) {
        mouse = { x: $event.pageX, y: $event.pageY, drag: true };
    };
    $scope.mouseUp = function() {
        mouse.drag = false;
    };
    $scope.mouseMove = function($event) {
        if (!mouse.drag)
            return;

        var delta = { x: $event.pageX - mouse.x, y: $event.pageY - mouse.y };

        var zoom = 0.5 * Math.pow($scope.zoom, 1.6) + 0.5;

        map_x += delta.x / zoom;
        map_y += delta.y / zoom;
        refresh();

        mouse = { x: $event.pageX, y: $event.pageY, drag: true };
    };
});

map.controller('MapDisplay', function($scope) {
});
