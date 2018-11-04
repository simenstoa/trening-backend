# trening-backend

For bruk under utvikling av S33 Trening, så man slipper å logge på/hente data fra Strava.
Innstaller MongoDB, og kjøre `mongoimport --jsonArray -d trening-s33 -c activities --file mock.json` for legge inn data. (`trening-s33` er databasen, `activities` er collection).
