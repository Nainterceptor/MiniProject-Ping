# MiniProject-Ping

## Import

I use MongoDB for data, please find some informations here

Import command:

    mongoimport --drop --db test -c pings --type json --jsonArray pings.json


Convert date string into Date object (to be mongoDB compliant, not really useful in the project):

    # May be run in the console or in a script.
    function convert_date(date) {
        t = date.split(/[- :]/);
        return new Date(Date.UTC(t[0], t[1]-1, t[2], t[3], t[4], t[5]));
    }

    db.pings.find().forEach(function(ping){
        ping.created_at = convert_date(ping.created_at); 
        db.pings.save(ping);
    })