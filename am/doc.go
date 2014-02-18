/*
allmusic.com.

Notice: This step is deprecated!

Import from CSV files to DB tables:

	psql -c "COPY amcredits (id,artist,albumid,job) FROM '$GOPATH/amcredits.csv' DELIMITER ',' CSV"

	psql -c "COPY amdiscographies (id,title,artistid,coverart) FROM '$GOPATH/amdiscographies.csv' DELIMITER ',' CSV"

	psql -c "COPY amreleases (id,title,albumid,format,year,label) FROM '$GOPATH/amreleases.csv' DELIMITER ',' CSV"

	psql -c "COPY amawards (id,title,albumid,section,year,chart,peak,type,prize,winnerids,winners) FROM '$GOPATH/amawards.csv' DELIMITER ',' CSV"

	psql -c "COPY amsongs (id,title,artistids,artists,albumid,composerids,composers,duration) FROM '$GOPATH/amsongs.csv' DELIMITER ',' CSV"

	psql -c "COPY amalbums (id,title,artistids,artists,review,coverart,duration,ratings,similars,genres,styles,moods,themes,songids,date_released,checktime) FROM '$GOPATH/amalbums.csv' DELIMITER ',' CSV"

*/
package am
