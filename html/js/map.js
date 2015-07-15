var map = angular.module('map', []);

map.controller('MapView', function($scope) {
    var snap = Snap('.map-display');
    var map  = $('#map_render');
    var map_size = { x: 1009.0, y: 665.0 }
    var mouse = { drag: false }; 
    var panSpeed = 2.0;

    var half_height = parseInt(map.height()) / 2.0;
    var half_width  = parseInt(map.width()) / 2.0;
    var map_x = 0, map_y = 0;
    var offset_x = 0, offset_y = 0;

    var world;
    var text = snap.group();

    function refresh() {
        if (!world)
            return;
        var zoom = zoomFactor($scope.zoom);
        offset_x = half_width  - zoom * map_size.x / 2;
        offset_y = half_height - zoom * map_size.y / 2;

        var transform = new Snap.Matrix();
        transform.translate(offset_x, offset_y)
        transform.scale(zoom);
        transform.translate(map_x, map_y);

        $scope.x = map_x;
        $scope.y = map_y;
        $scope.magnification = zoomFactor($scope.zoom);

        world.transform(transform);
        world.select('.world-gfx').attr({
            strokeWidth: Math.min(2 / zoom, 1),
        });

        var visible = zoom > 5;
        text.attr({ 
            'visibility': visible ? 'visible' : 'hidden',
            'font-size': 12.0 / $scope.zoom, 
            'font-family': 'sans-serif',
        });
    };

    function zoomFactor(level) {
        return 0.5 * Math.pow(level, 2.0) + 0.5;
    };

    $scope.init = function() {
        $scope.zoom = 1;

        Snap.load("img/map.svg", function(map) {
            var world_gfx = map.select('g');
            world_gfx.attr({
                fill: '#C8F7C5',
                stroke: '#444444',
            });
            world = snap.group(world_gfx, text); 
            snap.append(world);

            var c = snap.paper.text(map_size.x / 2 + 4.2, map_size.y / 2, "Rome");
            c.attr({
                fill: "#222222",
                'transform-origin': '50% 50%',
            });
            text.append(c);

            snap.circle(half_width, half_height, 5);
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

        var zoom = zoomFactor($scope.zoom);

        map_x += panSpeed * delta.x / zoom;
        map_y += panSpeed * delta.y / zoom;
        refresh();

        mouse = { x: $event.pageX, y: $event.pageY, drag: true };
    };
});
