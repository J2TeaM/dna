/*
This package supports accumulation of artists, songs, albums and videos.



CREATING GRAND SONGS:


	Create abbbbbb temp grand songs table by select all tables
	Create

CUSTOME TYPES:
	1. source (id int,siteid int)
	2. album_source ( sources source[], songids int[] )


CUSTOM AGGREGATE FUNCTIONS:

	1.agg_array_accum

agg_array_accum accumulates all varchar arrays into a new arrays

CUSTOM TYPES:
	TYPE media_type AS ENUM ('mp3', 'm4a', 'flac','mp4','flv');
	TYPE media_bitrate AS ENUM ('128kbps', '320kbps', '32kbps','500kbps','Lossless');
	TYPE song_format AS (link varchar(300), type media_type, size int4, bitrate media_bitrate);



CUSTOME PL/PGSQL FUNCTIONS:

	1.dna_hash
	2.get_siteid
	3.get_short_form
	4.upsert_hashid
	5.upsert_hashids
	6.get_lasted_checktime
	7.get_sources
	8.to_bitrate
	9.anyarray_sort (from https://github.com/JDBurnZ/anyarray)
	10.anyarray_uniq (from https://github.com/JDBurnZ/anyarray)
	11.transform_topics
	12.get_language
	13.array_avg
	14.anyarrray_remove (from https://github.com/JDBurnZ/anyarray)
	15.get_duration
	16.get_genres
	17.anyarray_remove (from https://github.com/JDBurnZ/anyarray)
	18.get_song_details
	19.get_songids
	20.to_source
	21.convert_hashids_to_songids
	22.get_agg_album // select count(*) from get_agg_album(255555);
	23.get_high_priority_unordered_hashid
	24.hash_songids
	25.anyarray_remove_null(from https://github.com/JDBurnZ/anyarray)

1.dna_hash

	FUNCTION dna_hash(title varchar, artists varchar[])

dna_hash returns a base-64 encoded md5 hash from title , artists and secret key from a song or an album or a video.

Example:

	select dna_hash('Never say never','{"Bieber"}');
	        dna_hash
	------------------------
	 spe2i2qrj+yoGqYPnjkKyw
	(1 row)


2.get_siteid

	FUNCTION get_siteid(src_table varchar) RETURNS int

get_siteid returns siteid from a source table

Example:

	select get_siteid('nssongs');
	 get_siteid
	------------
	          1
	(1 row)

3.get_short_form

	FUNCTION get_short_form(id integer) RETURNS text

get_short_form returns a short form from a siteid.

Example:

	select get_short_form(14);
	 get_short_form
	----------------
	 nv
	(1 row)

4.upsert_hashid

	FUNCTION upsert_hashid(srcid int,src_table varchar) RETURNS  TABLE (op varchar, id int, hashid varchar,source_id int,source_table varchar)

upsert_hashid gets source id (srcid) and source table name (src_table) as input params. It returns a table which has following fields;
	op : Operation name (INSERT or UPDATE).
	id : The id in destination table. New id returns if operation is INSERT. Otherwise, it returns old id.
	hashid : An unique string describes an item.
	source_id : The returned srcid.
	source_table: The returned src_table.

Example:

	select * from upsert_hashid(1299335,'nssongs');
	   op   | id |         hashid         | source_id | source_table
	--------+----+------------------------+-----------+--------------
	 UPDATE | 10 | PVuce2DCakmizGSrUyEPSw |   1299335 | nssongs
	(1 row)

5.upsert_hashids

	FUNCTION upsert_hashids(srcids int[],src_table varchar) RETURNS  TABLE (op varchar, id int, hashid varchar,source_id int,source_table varchar)

	FUNCTION upsert_hashids(checktime varchar,src_table varchar) RETURNS  TABLE (op varchar, id int, hashid varchar,source_id int,source_table varchar)

	FUNCTION upsert_hashids(checktime timestamp,src_table varchar) RETURNS  TABLE (op varchar, id int, hashid varchar,source_id int,source_table varchar)

upsert_hashids is similar to upsert_hashid but instead of taking srcid, it takes an array of srcid as an input. Returns a table with multiple rows, each row is a result from seach srcid in the array.

upsert_hashids also takes timestamp as a condition to run queries.

Example 1:

	select * from upsert_hashids('{1382381744,1382381740,1382381834,1382381761}'::int[],'zisongs';
	OR select * from upsert_hashids(ARRAY[1382381744,1382381740,1382381834,1382381761],'zisongs')
	   op   | id |         hashid         | source_id  | source_table
	--------+----+------------------------+------------+--------------
	 INSERT | 11 | 4yIojPOLhJ1lvSmqke77bA | 1382381744 | zisongs
	 INSERT | 12 | DX2g4ZAbt1I1xFXkuQVGGg | 1382381740 | zisongs
	 UPDATE |  6 | zkIV0SIaIeY7jBI/7D4Igw | 1382381834 | zisongs
	 INSERT | 13 | dGazDJNaJ5beAxlSiLG+kA | 1382381761 | zisongs
	(4 rows)

Example 2:

	SELECT * FROM upsert_hashids('2014-03-17','nssongs');
	OR SELECT * FROM upsert_hashids('2014-03-17'::timestamp,'nssongs');
	   op   | id  |         hashid         | source_id | source_table
	--------+-----+------------------------+-----------+--------------
	 UPDATE |  42 | o4d4rpQ9r3he978/x6SRzw |   1332635 | nssongs
	 UPDATE |  43 | gWm2Rlh8guS/GCXeK1IU6A |   1332637 | nssongs
	 UPDATE |  44 | dCN5f37PVci9Pe5/1hd/0A |   1332636 | nssongs
	 UPDATE |  45 | 2MVXNxM4G3+KYiBhvu6HVA |   1332652 | nssongs
	 UPDATE |  46 | krDP/d2KkRmNo9Wu5COxkg |   1332638 | nssongs
	 UPDATE |  47 | NZiBHC99Oh04IIUdVcjbdw |   1332634 | nssongs
	 UPDATE |  48 | UgOu+kHTr6+xCGxBYvKvYw |   1332651 | nssongs
	 UPDATE |  49 | w+AUzgqhtnwdtr4o5DE6NA |   1332653 | nssongs
	 UPDATE |  50 | BAOJr5jK87Vav1zfdUbH1w |   1332649 | nssongs
	 UPDATE |  51 | 2sqz2aI/L9LP0X862fMmlw |   1332648 | nssongs



6.get_lasted_checktime

	FUNCTION get_lasted_checktime() RETURNS   TABLE (site varchar, checktime timestamp)

get_lasted_checktime returns a multiple-row table from predefined sites which has lested checktime.

Example:

	select * from get_lasted_checktime();
	   site    |         checktime
	-----------+----------------------------
	 ccalbums  | 2014-03-17 11:55:10.044454
	 csnalbums | 2014-03-17 12:29:23.862303
	 kealbums  | 2014-03-17 12:07:14.404536
	 nvalbums  | 2014-03-17 11:55:26.10098
	 nctalbums | 2014-03-17 12:17:34.889705
	 nsalbums  | 2014-03-17 12:10:19.749395
	 zialbums  | 2014-03-17 11:49:49.284374
	 ccsongs   | 2014-03-15 11:13:21.120885
	 csnsongs  | 2014-03-17 12:29:24.274701
	 kesongs   | 2014-02-28 09:37:24.914208
	 nvsongs   | 2014-03-17 11:55:16.946422
	 mvsongs   | 2014-03-20 18:18:43.525121
	 nctsongs  | 2014-02-26 09:12:13.362725
	 nssongs   | 2014-03-17 12:09:37.592643
	 vgsongs   | 2013-04-03 20:10:58
	 zisongs   | 2014-03-17 11:52:25.868106
	 ccvideos  | 2014-03-17 11:55:11.53981
	 csnvideos | 2014-03-17 12:22:06.572581
	 kevideos  | 2014-03-17 12:07:24.199338
	 nvvideos  | 2014-03-17 12:19:07.313006
	 mvvideos  | 2014-03-20 18:19:10.764905
	 nctvideos | 2014-03-17 12:17:44.647824
	 nsvideos  | 2014-03-17 12:12:08.539091
	 zivideos  | 2014-02-24 18:18:18.081383
	(24 rows)

7.get_sources

	FUNCTION get_sources(itemId int,srcTable varchar) RETURNS   TABLE (id int, site varchar, title varchar, artists varchar[])

get_sources takes sources defined by an id and a table name. It returns a new table containing all information relating to the sources.

	itemId : an id from source table
	srcTable : "songs" || "albums" || "videos"

Example:

	select * from get_sources(2002204,'songs');
	   id    |   site   |      title      |      artists
	---------+----------+-----------------+-------------------
	  514470 | nssongs  | Never Say Never | {"Justin Bieber"}
	   75324 | nssongs  | Never Say Never | {"Justin Bieber"}
	 1003411 | csnsongs | Never Say Never | {"Justin Bieber"}
	 1038058 | nctsongs | Never Say Never | {"Justin Bieber"}
	 1057022 | nctsongs | Never Say Never | {"Justin Bieber"}
	 1060176 | nctsongs | Never Say Never | {"Justin Bieber"}
	 1060200 | nctsongs | Never Say Never | {"Justin Bieber"}
	 1099316 | nctsongs | Never Say Never | {"Justin Bieber"}
	 1179839 | nctsongs | Never Say Never | {"Justin Bieber"}
	 1442452 | nctsongs | Never Say Never | {"Justin Bieber"}
	 1497947 | nctsongs | Never Say Never | {"Justin Bieber"}
	 1730815 | nctsongs | Never Say Never | {"Justin Bieber"}
	 1848622 | nctsongs | Never Say Never | {"Justin Bieber"}
	 1896946 | nctsongs | Never Say Never | {"Justin Bieber"}
	 2061544 | nctsongs | Never Say Never | {"Justin Bieber"}
	  548007 | nctsongs | Never Say Never | {"Justin Bieber"}
	  549484 | nctsongs | Never Say Never | {"Justin Bieber"}
	  600994 | nctsongs | Never Say Never | {"Justin Bieber"}
	  617591 | nctsongs | Never Say Never | {"Justin Bieber"}
	  651334 | nctsongs | Never Say Never | {"Justin Bieber"}
	  667689 | nctsongs | Never Say Never | {"Justin Bieber"}
	  697217 | nctsongs | Never Say Never | {"Justin Bieber"}
	  755052 | nctsongs | Never Say Never | {"Justin Bieber"}
	  757740 | nctsongs | Never Say Never | {"Justin Bieber"}
	  804861 | nctsongs | Never Say Never | {"Justin Bieber"}
	  857985 | nctsongs | Never Say Never | {"Justin Bieber"}
	  858667 | nctsongs | Never Say Never | {"Justin Bieber"}
	  867329 | nctsongs | Never Say Never | {"Justin Bieber"}
	  869770 | nctsongs | Never Say Never | {"Justin Bieber"}
	  870741 | nctsongs | Never Say Never | {"Justin Bieber"}
	  893501 | nctsongs | Never Say Never | {"Justin Bieber"}
	  896541 | nctsongs | Never Say Never | {"Justin Bieber"}
	  920597 | nctsongs | Never Say Never | {"Justin Bieber"}
	  934672 | nctsongs | Never Say Never | {"Justin Bieber"}
	  951375 | nctsongs | Never Say Never | {"Justin Bieber"}
	  957276 | nctsongs | Never Say Never | {"Justin Bieber"}
	  977930 | nctsongs | Never Say Never | {"Justin Bieber"}
	  738766 | ccsongs  | Never Say Never | {"Justin Bieber"}
	(38 rows)


8.to_bitrate

	 FUNCTION to_bitrate(bitrate_flag int2) RETURNS int

Example:

	select to_bitrate(7::int2);

11. transform_topics

	FUNCTION transform_topics(topics varchar[]) RETURNS varchar[]

transform_topics tranforms all topics to theirs normalized form

Example:

	 select transform_topics('{"âu, mỹ","nhạc trẻ","nhạc hàn"}') as topics;;
	                              topics
	------------------------------------------------------------------
	 {"Nhạc Âu Mỹ","Nhạc Việt Nam","Nhạc Trẻ","Nhạc Hàn","Tiếng Hàn"}
	(1 row)

12. get_language

	FUNCTION get_language(topics varchar[]) RETURNS varchar[]

get_language returns a list of langueges found in a list of topics.

Example:

	select * from get_language(transform_topics('{"âu, mỹ","nhạc trẻ","nhạc hàn"}')) as languages;
	         languages
	----------------------------
	 {"Tiếng Việt","Tiếng Hàn"}
	(1 row)

15.get_duration

	FUNCTION get_duration(int []) RETURNS int

get_duration returns the most occured duration from duration list. Zero values are omitted.

	select get_duration('{333,333,333,0,0,0,0,0,0,0,0,0,0,0,0,0,333,333,333,333,333,333,333,180,333,333,299,332,33	3,333,333,335,333,333,333,333,333,284,333,335,333,333,333,333,333,332,333,333,0,333,333,333,0}');
	 get_duration
	--------------
	          333
	(1 row)

15.get_genres

	FUNCTION get_genres(topics varchar[]) RETURNS  varchar[]

get_genres returns a list of genres from topics.

	select get_genres('{"Nhạc Việt Nam","Nhạc Trữ Tình","Nhạc Quê Hương",Pop,Ballad,Rock,"Nhạc Cách Mạng","Nhạc Đỏ"}') as genres ;
	      genres
	-------------------
	 {Pop,Ballad,Rock}
	(1 row)

18.get_song_details

	FUNCTION get_song_details(itemId int) RETURNS   TABLE (mid int,id int, site varchar, title varchar, artists varchar[], album_title varchar, album_artists varchar[], video_title varchar, video_artists varchar[], authors varchar[],topics varchar[],official boolean,plays int,bitrate int,duration int,link varchar,lyric text, date_created timestamp )

Example:

	select * from get_song_details(3144484)


19.get_songids

	FUNCTION get_songids(songids int[], src_table varchar) RETURNS varchar[]

Example:

	select * from get_songids('{1321697,1321696,1321695,1321698,1321699}','nssongs');
	                                                     get_songids
	----------------------------------------------------------------------------------------------------------------------
	 {mFphzHqzoQ4G+7Wgq7YgaQ,1krSVeZ8oOQEKnhztITs8A,/Brjgj2O8o04NJQKrwzt3w,sQsKRdgGCYS6tS7NdlKPvQ,Hb96Mv0CjU4MdMlrakwV8g}
	(1 row)


20.to_source

	FUNCTION to_source(sites varchar[],ids int[]) RETURNS source[]

Example:

	select to_source('{nctalbums,ccalbums}'::varchar[],'{12026316,18058}'::int[]);
	          to_source
	------------------------------
	 {"(4,12026316)","(8,18058)"}
	(1 row)

21.convert_hashids_to_songids

	FUNCTION convert_hashids_to_songids(hashids varchar[]) RETURNS int[]

Example:

	SELECT convert_hashids_to_songids('{0lBTYsgbp8lwm+ltx1ObRw,1OU/S3gEPSUfjvv81tKxPQ}'::varchar[]);
	 convert_hashids_to_songids
	----------------------------
	 {168515,136826}
	(1 row)

22.get_high_priority_unordered_hashid

	FUNCTION get_high_priority_unordered_hashid(sources source[],hashids varchar[]) RETURNS varchar

get_high_priority_unordered_hashid determined an order of songids in which album being selected.
For example. Album A has many sources (nct, zi, ns). Because zi has highest priority. The order of the songs in that zi album will be selected. In the case that A has many zi sources, the first one will be chosen.

From the highest priority to the lowest:
	zi 2
	ns 1
	cc 8
	nct 4
	nv 16
	ke 128
	csn 32

EX:

	select 	get_high_priority_unordered_hashid('{"(4,12026316)","(8,18058)","(16,18058)","(2,18058)"}','{hash1,hash2,hash3	,hash4}');
	 get_high_priority_unordered_hashid
	------------------------------------
	 hash4
	(1 row)

	select get_high_priority_unordered_hashid('{"(4,12026316)","(8,18058)","(1,18058)","(16,18058)"}','{hash1,	hash2,hash3,hash4}');
	 get_high_priority_unordered_hashid
	------------------------------------
	 hash3
	(1 row)

24.

	FUNCTION hash_songids(songids int[]) returns int

hash_songids return 32bit integer from songids.

EX:

	select hash_songids('{3065543,622368,1116753,168515,2092682,3161373,2833357,2441939,136826,2788118}');
	 hash_songids
	--------------
	   1278073522
	(1 row)

*/
package media
