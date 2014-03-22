/*
lyricfind.com.

Get song including lyric:
	https://api.lyricfind.com/lyric.do?apikey=ccabb2c8bf7302e1d8c9b87be793bfb0&reqtype=default&trackid=amg:2033&output=json

	Using "internal" of reqtype
	https://api.lyricfind.com/lyric.do?apikey=ccabb2c8bf7302e1d8c9b87be793bfb0&reqtype=internal&trackid=amg:2033&output=json&&appname=android&version=1.2&territory=us

Get metadata of a song (without lyric)
	http://api.lyricfind.com/metadata.do?apikey=4b59d60b5b74512a662b89dfb1b28680&reqtype=metadata&displaykey=ccabb2c8bf7302e1d8c9b87be793bfb0&trackid=amg:2033&output=json

Getting metadata of all tracks having lyrics from an artist
	http://api.lyricfind.com/metadata.do?apikey=4b59d60b5b74512a662b89dfb1b28680&reqtype=availablelyrics&artistid=amg:7362&offset=0&limit=10&displaykey=ccabb2c8bf7302e1d8c9b87be793bfb0&output=json

	Notice: listingtype=all . Default is "main"
	http://api.lyricfind.com/metadata.do?apikey=4b59d60b5b74512a662b89dfb1b28680&reqtype=availablelyrics&artistid=amg:7362&offset=0&limit=10&displaykey=ccabb2c8bf7302e1d8c9b87be793bfb0&output=json&listingtype=all

Top 25 tracks
	http://api.lyricfind.com/charts.do?apikey=445fe174a2b3bb659ce088797f2290df&reqtype=trackcharts&displaykey=ccabb2c8bf7302e1d8c9b87be793bfb0&output=json


Search tracks
	All tracks contains "I kissed a girl and i liked it"
	http://api.lyricfind.com/search.do?apikey=780634cb4d1a9277187f7431e2fa0139&reqtype=default&searchtype=track&lyrics=I+kissed+a+girl+and+i+liked+it&displaykey=ccabb2c8bf7302e1d8c9b87be793bfb0&output=json


*/
package lf
