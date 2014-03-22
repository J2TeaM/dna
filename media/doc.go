/*
This package supports accumulation of artists, songs, albums and videos.


CUSTOME PL/PGSQL FUNCTIONS:

	1.dna_hash
	2.get_siteid
	3.get_short_form
	4.upsert_hashid
	5.upsert_hashids
	6.get_lasted_checktime
	7.get_sources

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


*/
package media
