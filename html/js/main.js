$(document).ready(function() {
    var CityTemplate = $('script#city-template').html().trim();
    var VehicleTemplate = $('script#vehicle-template').html().trim();
    Mustache.parse(CityTemplate);
    Mustache.parse(VehicleTemplate);

    var City = Backbone.Model.extend({
    });

    var Cities = Backbone.Collection.extend({
        model: City,
        url: 'http://localhost:8000/v1/123/city/',
    });

    var CityView = Backbone.View.extend({
        className: 'city',
        render: function() {
            this.$el.html(Mustache.render(CityTemplate, this.model.toJSON()));
            return this;
        },
    });

    var CitiesView = Backbone.View.extend({
        cityViews: [],
        initialize: function() {
            this.cities = new Cities();
            this.cities.on('sync', function() {
                this.cityViews = [ ];
                this.cities.each(function(city) {
                    this.cityViews.push(new CityView({ model: city }));
                }, this);
                this.render();
            }, this);

            setInterval(_.bind(function() {
                this.cities.fetch();
            }, this), 1000);
        },
        render: function() {
            this.$el.html('');
            _.each(this.cityViews, function(cityView) {
                this.$el.append(cityView.render().el);
            }, this);
            return this;
        },
    });

    var Vehicle = Backbone.Model.extend({
        Move: function(from, to) {
            var order = {
                type: 1,
                from: from.get('id'),
                to:   to.get('id'),
            };

        },
    });

    var Vehicles = Backbone.Collection.extend({
        model: Vehicle,
        url: '/v1/123/vehicle/',
    });

    var VehicleView = Backbone.View.extend({
        className: 'vehicle',
        render: function() {
            this.$el.html(Mustache.render(VehicleTemplate, this.model.toJSON()));
            return this;
        },
    });

    var VehiclesView = Backbone.View.extend({
        vehicleViews: [ ],
        initialize: function() {
            this.vehicles = new Vehicles();
            this.vehicles.on('sync', function() {
                this.vehicleViews = [ ];
                this.vehicles.each(function(vehicle) {
                    this.vehicleViews.push(new VehicleView({ model: vehicle }));
                }, this);
                this.render();
            }, this);

            setInterval(_.bind(function() {
                this.vehicles.fetch();
            }, this), 1000);
        },
        render: function() {
            this.$el.html('');
            _.each(this.vehicleViews, function(vehicleView) {
                this.$el.append(vehicleView.render().el);
            }, this);
            return this;
        },
    });


    new CitiesView({ el: $('#cities') });
    new VehiclesView({ el: $('#vehicles') });
});
