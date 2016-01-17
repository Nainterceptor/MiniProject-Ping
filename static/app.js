(function() {
    var API_BASE_PATH = '/api/1';
    var PINGS_BASE_PATH = API_BASE_PATH + '/pings';

    var graph = Morris.Line({
        element: 'origin-by-hour',
        xkey: 'date',
        dateFormat: function(date) {
            return new Date(date).toISOString().replace(/[TZ]/g, ' ').slice(0, 13) + 'h UTC';
        },
        yLabelFormat: function(label) {
            label = label.toFixed(2);
            return label.toString() + 'ms';
        },
        ykeys: ['averageTransferTimeMs'],
        labels: ['Average transfer time (ms)']
    });

    fetchOriginList();

    function normalize(data) {
        data.forEach(function(row, index, array) {
            var id = row['_id'];
            var year = id.year;
            var month = id.month;
            var day = id.day;
            var hour = id.hour;
            var utc_date = new Date(Date.UTC(year, month, day, hour));

            array[index] = {
                date: utc_date.toISOString(),
                averageTransferTimeMs: row['average_transfer_time_ms']

            };
        });
    }

    function fetchData(origin) {
        origin = origin.replace(/[^a-zA-Z\-]/g, '');
        jQuery.get(PINGS_BASE_PATH + '/' + origin + '/hours', {}, function(data) {
            normalize(data);
            graph.setData(data);
        });
    }

    function fetchOriginList() {
        var select = jQuery('select#origins');
        select.off();
        select.on('change', function () {
            fetchData(jQuery(this).val());
        });
        jQuery.get(PINGS_BASE_PATH + '/origins', {}, function(data) {
            select.html('<option>Please select an origin...</option>');
            data.forEach(function(origin) {
                select.append('<option>' + origin + '</option>')
            });
        });
    }
})();