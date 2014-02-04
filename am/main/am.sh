START=$(date +%s)
printf "1.TRUNCATE TABLES:\n"
printf "amcredits: "
psql -c "truncate table amcredits"
printf "amdiscographies: "
psql -c "truncate table amdiscographies"
printf "amreleases: "
psql -c "truncate table amreleases"
printf "amawards: "
psql -c "truncate table amawards"
printf "amsongs: "
psql -c "truncate table amsongs"
printf "amalbums: "
psql -c "truncate table amalbums"

echo ""
echo "2.COPYING CSV FILES TO TABLES:"
printf "amcredits: "
psql -c "COPY amcredits (id,artist,albumid,job) FROM '$GOPATH/amcredits.csv' DELIMITER ',' CSV"

printf "amdiscographies: "
psql -c "COPY amdiscographies (id,title,artistid,coverart) FROM '$GOPATH/amdiscographies.csv' DELIMITER ',' CSV"

printf "amreleases: "
psql -c "COPY amreleases (id,title,albumid,format,year,label) FROM '$GOPATH/amreleases.csv' DELIMITER ',' CSV"

printf "amawards: "
psql -c "COPY amawards (id,title,albumid,section,year,chart,peak,type,prize,winnerids,winners) FROM '$GOPATH/amawards.csv' DELIMITER ',' CSV"

printf "amsongs: "
psql -c "COPY amsongs (id,title,artistids,artists,albumid,composerids,composers,duration) FROM '$GOPATH/amsongs.csv' DELIMITER ',' CSV"

printf "amalbums: "
psql -c "COPY amalbums (id,title,artistids,artists,review,coverart,duration,ratings,similars,genres,styles,moods,themes,songids,date_released,checktime) FROM '$GOPATH/amalbums.csv' DELIMITER ',' CSV"
END=$(date +%s)
DIFF=$(($END - $START))
printf "\nIt tool $DIFF seconds\n"
