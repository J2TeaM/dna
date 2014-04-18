/*
 This package shows table definitions.



	DROP AGGREGATE IF EXISTS agg_lyric(varchar, text, int,int,BOOLEAN);
	DROP FUNCTION IF EXISTS tmp_agg_lyric_final( tmp_agg_lyric[]);
	DROP FUNCTION tmp_agg_lyric_state( tmp_agg_lyric[],varchar,text, int,int,boolean);
	DROP TYPE IF EXISTS tmp_agg_lyric;


	CREATE TYPE tmp_agg_lyric AS (site varchar, lyric text, status int, plays int, has_lrc boolean );
	CREATE OR REPLACE FUNCTION tmp_agg_lyric_state(prev tmp_agg_lyric[], e1 varchar, e2 text, e3 int,e4 int, e5 	BOOLEAN)
	    RETURNS tmp_agg_lyric[] AS
	$$
	    SELECT  array_append(prev, (e1,e2,e3,e4,e5)::tmp_agg_lyric);
	$$
	LANGUAGE 'sql' IMMUTABLE;


	CREATE OR REPLACE FUNCTION tmp_agg_lyric_final(last tmp_agg_lyric[])
	    RETURNS text AS
	$$
		DECLARE
			lorder int = 0;
			current_order float = 0;
			current_index int = 0;
			pl float = 0;
			current_score float = 0;
			score float = 0;
		BEGIN
			IF last IS NOT NULL AND array_length(last,1) IS NOT NULL THEN
				FOR idx IN array_lower(last, 1)..array_upper(last, 1) LOOP
					IF last[idx].lyric <> '' THEN
						IF last[idx].status = 1 THEN
							RETURN last[idx].lyric;
						ELSE
							CASE get_short_form(get_siteid(last[idx].site))
								WHEN 'zi' THEN
									lorder = 9;
									IF last[idx].plays > 45512911 THEN pl = 1;END IF;
									pl = last[idx].plays/45512911::float;
								WHEN 'ns' THEN
									lorder = 10;
									IF last[idx].plays > 1088934 THEN pl = 1;END IF;
									pl = last[idx].plays/1088934::float;
								WHEN 'cc' THEN
									lorder = 7;
									IF last[idx].plays > 638790 THEN pl = 1;END IF;
									pl = last[idx].plays/638790::float;
								WHEN 'nv' THEN
									lorder = 5;
									IF last[idx].plays > 5591248 THEN pl = 1;END IF;
									pl = last[idx].plays/5591248::float;
								WHEN 'ke' THEN
									lorder = 4;
									IF last[idx].plays > 1088934 THEN pl = 1;END IF;
									pl = last[idx].plays/1088934::float;
								WHEN 'csn' THEN
									lorder = 6;
									IF last[idx].plays > 15455812 THEN pl = 1;END IF;
									pl = last[idx].plays/15455812::float;
								WHEN 'nct' THEN
									lorder = 8;
									IF last[idx].plays > 13235612 THEN pl = 1;END IF;
									pl = last[idx].plays/13235612::float;
								ELSE
									lorder = 1;
									IF last[idx].plays > 45512911 THEN pl = 1;END IF;
									pl = last[idx].plays/45512911::float;
							END CASE;
							IF last[idx].has_lrc = true THEN score = 1;
							ELSE  score = 0.6*(lorder/10::float) + 0.4*pl; END IF;
							IF score > current_score THEN
								current_index = idx;
								current_score  = score;
							END IF;
						END IF;
					END IF;
				END LOOP;
				--RAISE WARNING '% %', current_index, last[current_index].site;
				IF current_index = 0 THEN RETURN last[1].lyric ;
				ELSE RETURN last[current_index].lyric; END IF;
			ELSE
				RETURN '';
			END IF ;

		END;
	$$
	LANGUAGE 'plpgsql' IMMUTABLE;



	CREATE AGGREGATE agg_lyric(varchar, text, int,int, BOOLEAN) (
	      SFUNC=tmp_agg_lyric_state,
	      STYPE=tmp_agg_lyric[],
	      FINALFUNC=tmp_agg_lyric_final,
	      INITCOND = '{}'
	);

ALBUMS TABLES:

	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                 CCALBUMS                 │  │                 NVALBUMS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│ccalbums     id                 int4      │  │nvalbums     id                 int4      │
	│ccalbums     title              varchar   │  │nvalbums     title              varchar   │
	│ccalbums     artists            varchar[] │  │nvalbums     artists            varchar[] │
	│ccalbums     nsongs             int8      │  │nvalbums     topics             varchar[] │
	│ccalbums     plays              int4      │  │nvalbums     nsongs             int4      │
	│ccalbums     coverart           varchar   │  │nvalbums     plays              int8      │
	│ccalbums     description        text      │  │nvalbums     coverart           varchar   │
	│ccalbums     checktime          timestamp │  │nvalbums     date_created       timestamp │
	│ccalbums     songids            int[]     │  │nvalbums     checktime          timestamp │
	│ccalbums     year_released      int4      │  │nvalbums     songids            int[]     │
	│ccalbums     topics             varchar[] │  │nvalbums     description        text      │
	└──────────────────────────────────────────┘  └──────────────────────────────────────────┘



	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                NCTALBUMS                 │  │                 NSALBUMS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│nctalbums    id                 int8      │  │nsalbums     id                 int4      │
	│nctalbums    key                varchar   │  │nsalbums     title              varchar   │
	│nctalbums    title              varchar   │  │nsalbums     topics             varchar[] │
	│nctalbums    artists            varchar[] │  │nsalbums     genres             varchar[] │
	│nctalbums    topics             varchar[] │  │nsalbums     category           varchar[] │
	│nctalbums    likes              int4      │  │nsalbums     description        text      │
	│nctalbums    plays              int4      │  │nsalbums     date_released      varchar   │
	│nctalbums    songids            int[]     │  │nsalbums     nsongs             int4      │
	│nctalbums    nsongs             int4      │  │nsalbums     plays              int8      │
	│nctalbums    description        text      │  │nsalbums     coverart           varchar   │
	│nctalbums    coverart           varchar   │  │nsalbums     artists            varchar[] │
	│nctalbums    link_key           varchar   │  │nsalbums     artistid           int4      │
	│nctalbums    link_share         varchar   │  │nsalbums     checktime          timestamp │
	│nctalbums    type               varchar   │  │nsalbums     songids            int[]     │
	│nctalbums    official           bool      │  │nsalbums     label              varchar   │
	│nctalbums    has_feature        bool      │  └──────────────────────────────────────────┘
	│nctalbums    date_created       timestamp │
	│nctalbums    checktime          timestamp │
	└──────────────────────────────────────────┘

	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                 ZIALBUMS                 │  │                 KEALBUMS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│zialbums     id                 int4      │  │kealbums     id                 int4      │
	│zialbums     key                varchar   │  │kealbums     key                varchar   │
	│zialbums     encoded_key        varchar   │  │kealbums     title              varchar   │
	│zialbums     title              varchar   │  │kealbums     artists            varchar[] │
	│zialbums     artists            varchar[] │  │kealbums     nsongs             int4      │
	│zialbums     coverart           varchar   │  │kealbums     plays              int4      │
	│zialbums     topics             varchar[] │  │kealbums     coverart           varchar   │
	│zialbums     plays              int4      │  │kealbums     description        text      │
	│zialbums     year_released      varchar   │  │kealbums     songids            int[]     │
	│zialbums     nsongs             int4      │  │kealbums     date_created       timestamp │
	│zialbums     description        text      │  │kealbums     checktime          timestamp │
	│zialbums     date_created       timestamp │  └──────────────────────────────────────────┘
	│zialbums     checktime          timestamp │
	│zialbums     songids            int[]     │
	│zialbums     is_hit             int2      │
	│zialbums     is_official        int2      │
	│zialbums     is_album           int2      │
	│zialbums     likes              int4      │
	│zialbums     comments           int4      │
	│zialbums     status_id          int2      │
	│zialbums     artist_ids         int[]     │
	└──────────────────────────────────────────┘

	┌──────────────────────────────────────────┐
	│                CSNALBUMS                 │
	├──────────────────────────────────────────┤
	│csnalbums    id                 int8      │
	│csnalbums    title              varchar   │
	│csnalbums    artists            varchar[] │
	│csnalbums    href               varchar   │
	│csnalbums    nsongs             int8      │
	│csnalbums    coverart           varchar   │
	│csnalbums    producer           varchar   │
	│csnalbums    downloads          int8      │
	│csnalbums    plays              int8      │
	│csnalbums    date_released      varchar   │
	│csnalbums    checktime          timestamp │
	│csnalbums    topics             varchar[] │
	│csnalbums    date_created       timestamp │
	│csnalbums    songids            int[]     │
	└──────────────────────────────────────────┘

SONGS TABLES:

	┌──────────────────────────────────────────┐
	│               GRANDS SONGS               │
	│──────────────────────────────────────────│
	│    id                   int              │
	│    site                 varchar          │
	│    title                varchar          │
	│    artists              varchar[]        │
	│    album_title          varchar          │
	│    album_artists        varchar[]        │
	│    video_title          varchar          │
	│    video_artists        varchar[]        │
	│    authors              varchar[]        │
	│    topics               varchar[]        │
	│    official             boolean          │
	│    plays                int              │
	│    bitrate              int              │
	│    duration             int              │
	│    link                 varchar          │
	│    lyric                text             │
	│    date_created         timestamp        │
	└──────────────────────────────────────────┘

	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                 CSNSONGS                 │  │                  CCSONGS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│csnsongs     id                 int8      │  │ccsongs      id                 int4      │
	│csnsongs     title              varchar   │  │ccsongs      title              varchar   │
	│csnsongs     artists            varchar[] │  │ccsongs      artistid           int4      │
	│csnsongs     authors            varchar[] │  │ccsongs      artists            varchar[] │
	│csnsongs     topics             varchar[] │  │ccsongs      plays              int8      │
	│csnsongs     album_title        varchar   │  │ccsongs      topics             varchar[] │
	│csnsongs     album_href         varchar   │  │ccsongs      duration           int4      │
	│csnsongs     album_coverart     varchar   │  │ccsongs      bitrate            int4      │
	│csnsongs     producer           varchar   │  │ccsongs      coverart           varchar   │
	│csnsongs     downloads          int8      │  │ccsongs      link               varchar   │
	│csnsongs     plays              int8      │  │ccsongs      checktime          timestamp │
	│csnsongs     formats            json      │  │ccsongs      lyrics             text      │
	│csnsongs     href               varchar   │  └──────────────────────────────────────────┘
	│csnsongs     is_lyric           int4      │
	│csnsongs     lyric              text      │
	│csnsongs     date_released      varchar   │
	│csnsongs     date_created       timestamp │
	│csnsongs     checktime          timestamp │
	│csnsongs     type               bool      │
	└──────────────────────────────────────────┘

	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                  KESONGS                 │  │                  NVSONGS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│kesongs      id                 int4      │  │nvsongs      id                 int4      │
	│kesongs      key                varchar   │  │nvsongs      title              varchar   │
	│kesongs      title              varchar   │  │nvsongs      artists            varchar[] │
	│kesongs      artists            varchar[] │  │nvsongs      authors            varchar[] │
	│kesongs      plays              int4      │  │nvsongs      topics             varchar[] │
	│kesongs      listen_type        int4      │  │nvsongs      plays              int8      │
	│kesongs      link               varchar   │  │nvsongs      lyric              text      │
	│kesongs      lyric              text      │  │nvsongs      link               varchar   │
	│kesongs      media_url_mono     varchar   │  │nvsongs      link320            varchar   │
	│kesongs      media_url_pre      varchar   │  │nvsongs      checktime          timestamp │
	│kesongs      download_url       varchar   │  │nvsongs      type               varchar   │
	│kesongs      is_download        varchar   │  └──────────────────────────────────────────┘
	│kesongs      ringbacktone_code  varchar   │
	│kesongs      ringbacktone_price int4      │
	│kesongs      price              int4      │
	│kesongs      copyright          int4      │
	│kesongs      crbt_id            int4      │
	│kesongs      coverart           varchar   │
	│kesongs      coverart310        varchar   │
	│kesongs      date_created       timestamp │
	│kesongs      checktime          timestamp │
	│kesongs      has_lyric          bool      │
	└──────────────────────────────────────────┘

	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                 NCTSONGS                 │  │                  NSSONGS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│nctsongs     id                 int4      │  │nssongs      id                 int4      │
	│nctsongs     key                varchar   │  │nssongs      duration           int4      │
	│nctsongs     title              varchar   │  │nssongs      title              varchar   │
	│nctsongs     artists            varchar[] │  │nssongs      link               varchar   │
	│nctsongs     topics             varchar[] │  │nssongs      artists            varchar[] │
	│nctsongs     link_key           varchar   │  │nssongs      artistid           int4      │
	│nctsongs     type               varchar   │  │nssongs      authors            varchar[] │
	│nctsongs     bitrate            int4      │  │nssongs      authorid           int4      │
	│nctsongs     official           bool      │  │nssongs      plays              int4      │
	│nctsongs     likes              int4      │  │nssongs      topics             varchar[] │
	│nctsongs     plays              int4      │  │nssongs      category           varchar[] │
	│nctsongs     link_share         varchar   │  │nssongs      bitrate            int4      │
	│nctsongs     stream_url         varchar   │  │nssongs      official           int2      │
	│nctsongs     image              varchar   │  │nssongs      islyric            int2      │
	│nctsongs     coverart           varchar   │  │nssongs      date_created       timestamp │
	│nctsongs     duration           int4      │  │nssongs      date_updated       timestamp │
	│nctsongs     linkdown           varchar   │  │nssongs      lyric              text      │
	│nctsongs     linkdown_hq        varchar   │  │nssongs      checktime          timestamp │
	│nctsongs     lyricid            int4      │  │nssongs      same_artist        int2      │
	│nctsongs     has_lyric          bool      │  └──────────────────────────────────────────┘
	│nctsongs     lyric              text      │
	│nctsongs     lyric_status       int4      │
	│nctsongs     has_lrc            bool      │
	│nctsongs     lrc                text      │
	│nctsongs     lrc_url            varchar   │
	│nctsongs     username_created   varchar   │
	│nctsongs     checktime          timestamp │
	└──────────────────────────────────────────┘


	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                  ZISONGS                 │  │                  MVSONGS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│zisongs      id                 int4      │  │mvsongs      id                 int8      │
	│zisongs      key                varchar   │  │mvsongs      title              varchar   │
	│zisongs      title              varchar   │  │mvsongs      artists            varchar[] │
	│zisongs      artists            varchar[] │  │mvsongs      topics             varchar[] │
	│zisongs      authors            varchar[] │  │mvsongs      plays              int8      │
	│zisongs      plays              int8      │  │mvsongs      link               varchar   │
	│zisongs      topics             varchar[] │  │mvsongs      checktime          timestamp │
	│zisongs      link               varchar   │  └──────────────────────────────────────────┘
	│zisongs      path               varchar   │
	│zisongs      lyric              text      │  ┌──────────────────────────────────────────┐
	│zisongs      date_created       timestamp │  │                  VGSONGS                 │
	│zisongs      checktime          timestamp │  ├──────────────────────────────────────────┤
	│zisongs      album_id           int4      │  │vgsongs      id                 int8      │
	│zisongs      is_hit             int2      │  │vgsongs      title              varchar   │
	│zisongs      is_official        int2      │  │vgsongs      artists            varchar[] │
	│zisongs      download_status    int2      │  │vgsongs      authors            varchar[] │
	│zisongs      copyright          varchar   │  │vgsongs      topics             varchar[] │
	│zisongs      bitrate_flags      int2      │  │vgsongs      plays              int4      │
	│zisongs      likes              int4      │  │vgsongs      link               varchar   │
	│zisongs      comments           int4      │  │vgsongs      checktime          timestamp │
	│zisongs      video_id           int4      │  └──────────────────────────────────────────┘
	│zisongs      artist_ids         int[]     │
	│zisongs      thumbnail          varchar   │
	└──────────────────────────────────────────┘

VIDEOS TABLE:


	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                CSNVIDEOS                 │  │                 CCVIDEOS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│csnvideos    id                 int8      │  │ccvideos     id                 int4      │
	│csnvideos    title              varchar   │  │ccvideos     title              varchar   │
	│csnvideos    artists            varchar[] │  │ccvideos     artists            varchar[] │
	│csnvideos    authors            varchar[] │  │ccvideos     topics             varchar[] │
	│csnvideos    topics             varchar[] │  │ccvideos     plays              int4      │
	│csnvideos    thumbnail          varchar   │  │ccvideos     resolution_flags   int4      │
	│csnvideos    producer           varchar   │  │ccvideos     thumbnail          varchar   │
	│csnvideos    downloads          int8      │  │ccvideos     lyric              text      │
	│csnvideos    plays              int8      │  │ccvideos     links              varchar[] │
	│csnvideos    formats            json      │  │ccvideos     year_released      int4      │
	│csnvideos    href               varchar   │  │ccvideos     checktime          timestamp │
	│csnvideos    is_lyric           int4      │  └──────────────────────────────────────────┘
	│csnvideos    lyric              text      │
	│csnvideos    date_released      varchar   │
	│csnvideos    date_created       timestamp │
	│csnvideos    checktime          timestamp │
	│csnvideos    type               bool      │
	└──────────────────────────────────────────┘

	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                 KEVIDEOS                 │  │                 NVVIDEOS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│kevideos     id                 int4      │  │nvvideos     id                 int4      │
	│kevideos     key                varchar   │  │nvvideos     title              varchar   │
	│kevideos     title              varchar   │  │nvvideos     artists            varchar[] │
	│kevideos     artists            varchar[] │  │nvvideos     authors            varchar[] │
	│kevideos     plays              int4      │  │nvvideos     topics             varchar[] │
	│kevideos     listen_type        int4      │  │nvvideos     plays              int8      │
	│kevideos     link               varchar   │  │nvvideos     lyric              text      │
	│kevideos     is_download        int4      │  │nvvideos     link               varchar   │
	│kevideos     download_url       varchar   │  │nvvideos     link320            varchar   │
	│kevideos     ringbacktone_code  varchar   │  │nvvideos     checktime          timestamp │
	│kevideos     ringbacktone_price int4      │  │nvvideos     type               varchar   │
	│kevideos     price              int4      │  │nvvideos     thumbnail          varchar   │
	│kevideos     copyright          int4      │  └──────────────────────────────────────────┘
	│kevideos     crbt_id            int4      │
	│kevideos     thumbnail          varchar   │
	│kevideos     date_created       timestamp │
	│kevideos     checktime          timestamp │
	└──────────────────────────────────────────┘

	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                NCTVIDEOS                 │  │                 NSVIDEOS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│nctvideos    id                 int4      │  │nsvideos     id                 int4      │
	│nctvideos    key                varchar   │  │nsvideos     title              varchar   │
	│nctvideos    title              varchar   │  │nsvideos     artists            varchar[] │
	│nctvideos    artists            varchar[] │  │nsvideos     topics             varchar[] │
	│nctvideos    artistid           int4      │  │nsvideos     plays              int4      │
	│nctvideos    topics             varchar[] │  │nsvideos     duration           int4      │
	│nctvideos    plays              int4      │  │nsvideos     official           int4      │
	│nctvideos    likes              int4      │  │nsvideos     producerid         int4      │
	│nctvideos    duration           int4      │  │nsvideos     link               varchar   │
	│nctvideos    thumbnail          varchar   │  │nsvideos     sublink            varchar   │
	│nctvideos    image              varchar   │  │nsvideos     thumbnail          varchar   │
	│nctvideos    type               varchar   │  │nsvideos     date_created       timestamp │
	│nctvideos    link_key           varchar   │  │nsvideos     checktime          timestamp │
	│nctvideos    link_share         varchar   │  └──────────────────────────────────────────┘
	│nctvideos    lyric              text      │
	│nctvideos    stream_url         varchar   │
	│nctvideos    date_created       timestamp │
	│nctvideos    checktime          timestamp │
	│nctvideos    relatives          varchar[] │
	└──────────────────────────────────────────┘

	┌──────────────────────────────────────────┐  ┌──────────────────────────────────────────┐
	│                 ZIVIDEOS                 │  │                 MVVIDEOS                 │
	├──────────────────────────────────────────┤  ├──────────────────────────────────────────┤
	│zivideos     id                 int8      │  │mvvideos     id                 int8      │
	│zivideos     title              varchar   │  │mvvideos     title              varchar   │
	│zivideos     artists            varchar[] │  │mvvideos     artists            varchar[] │
	│zivideos     topics             varchar[] │  │mvvideos     topic              varchar[] │
	│zivideos     plays              int4      │  │mvvideos     plays              int8      │
	│zivideos     thumbnail          varchar   │  │mvvideos     link               varchar   │
	│zivideos     link               varchar   │  │mvvideos     checktime          timestamp │
	│zivideos     lyric              text      │  └──────────────────────────────────────────┘
	│zivideos     date_created       timestamp │
	│zivideos     checktime          timestamp │
	│zivideos     artist_ids         int[]     │
	│zivideos     duration           int4      │
	│zivideos     resolution_flags   int2      │
	│zivideos     status_id          int2      │
	│zivideos     likes              int4      │
	│zivideos     comments           int4      │
	└──────────────────────────────────────────┘

	1274262	 nhạc âu mỹ => [ 'Nhạc Âu Mỹ' ]
	918678	 pop => [ 'pop' ]
	863464	 rock => [ 'rock' ]
	835756	 nhạc trẻ => [ 'Nhạc Việt Nam', 'Nhạc Trẻ' ]
	804006	 âu, mỹ => [ 'Nhạc Âu Mỹ' ]
	527417	 âu mỹ => [ 'Nhạc Âu Mỹ' ]
	338572	 ballad => [ 'ballad' ]
	293277	 việt nam => [ 'Nhạc Việt Nam' ]
	278535	 nhạc nhật => [ 'Nhạc Nhật', 'Tiếng Nhật' ]
	249791	 tui hát => []
	246816	 thể loại khác => []
	221534	 trữ tình => [ 'Nhạc Việt Nam', 'Nhạc Trữ Tình' ]
	197278	 nhạc hoa => [ 'Nhạc Hoa', 'Tiếng Hoa' ]
	191485	 rap việt => [ 'Nhạc Việt Nam', 'Rap' ]
	191410	 nhật bản => [ 'Nhạc Nhật', 'Tiếng Nhật' ]
	180467	 hàn quốc => [ 'Nhạc Hàn', 'Tiếng Hàn' ]
	180344	 nhạc trữ tình => [ 'Nhạc Việt Nam', 'Nhạc Trữ Tình' ]
	152256	 rap => [ 'rap' ]
	151714	 nhạc không lời => [ 'Nhạc Không Lời' ]
	150313	 không lời => [ 'Nhạc Không Lời' ]
	143437	 nhạc hàn => [ 'Nhạc Hàn', 'Tiếng Hàn' ]
	141210	 hòa tấu => [ 'Nhạc Hòa Tấu' ]
	133669	 dance => [ 'Dance' ]
	101816	 hoa ngữ => [ 'Nhạc Hoa', 'Tiếng Hoa' ]
	96623	 jazz => [ 'jazz' ]
	92786	 => []
	84068	 classical => [ 'classical' ]
	83516	 nhạc phim => [ 'nhạc phim' ]
	81351	 country => [ 'country' ]
	80380	 hiphop => [ 'Hip hop' ]
	76301	 new age => [ 'new age' ]
	74445	 hip hop => [ 'Hip hop' ]
	73616	 giải trí => [ 'giải trí' ]
	73263	 âu => [ 'Nhạc Âu Mỹ' ]
	73263	 mỹ => [ 'Nhạc Âu Mỹ' ]
	71512	 electronic => [ 'electronic' ]
	67637	 nhạc các nước khác => [ 'Nhạc Quốc Tế' ]
	64504	 r&b => [ 'r&b' ]
	59005	 khác => []
	54299	 nhạc vàng => [ 'Nhạc Việt Nam', 'Nhạc Vàng' ]
	54130	 playback => [ 'playback' ]
	50164	 phim => [ 'Nhạc Phim' ]
	43798	 blues => [ 'blues' ]
	41909	 world music => [ 'world music' ]
	40407	 hong kong => [ 'Hồng kông', 'Tiếng Hoa' ]
	39980	 soundtrack => [ 'soundtrack' ]
	39082	 remix => [ 'remix' ]
	37833	 nước khác => [ 'Nhạc Quốc Tế' ]
	37785	 nhạc việt nam => [ 'Nhạc Việt Nam' ]
	34407	 soul => [ 'soul' ]
	33877	 âu mỹ khác => [ 'Nhạc Âu Mỹ' ]
	32642	 nhạc pháp => [ 'Nhạc Pháp', 'Tiếng Pháp' ]
	30956	 đài loan => [ 'Nhạc Hoa', 'Đài Loan', 'Tiếng Hoa' ]
	30178	 piano => [ 'piano' ]
	29573	 nhạc hàn quốc => [ 'Nhạc Hàn', 'Tiếng Hàn' ]
	27000	 dance hot => [ 'Dance' ]
	26817	 chưa phân loại => []
	26599	 trung quốc => [ 'Nhạc Hoa', 'Tiếng Hoa' ]
	22765	 nhạc trẻ hay nhất => [ 'Nhạc Việt Nam', 'Nhạc Trẻ' ]
	21666	 video clip => [ 'video clip' ]
	21231	 folk => [ 'folk' ]
	20817	 nhạc quê hương => [ 'Nhạc Việt Nam', 'Nhạc Quê Hương' ]
	20462	 hài kịch => [ 'hài kịch' ]
	19077	 phim nước ngoài => [ 'phim nước ngoài' ]
	18692	 nhạc tổng hợp => [ 'nhạc tổng hợp' ]
	18335	 việt nam khác => [ 'Nhạc Việt Nam' ]
	18278	 nhạc hải ngoại => [ 'nhạc hải ngoại' ]
	18227	 metal => [ 'metal' ]
	16737	 nhạc trữ tình hay nhất => [ 'Nhạc Việt Nam', 'Nhạc Trữ Tình' ]
	15809	 nhạc trịnh => [ 'Nhạc Việt Nam', 'Nhạc Trịnh' ]
	13642	 instrumental => [ 'instrumental' ]
	13068	 latin => [ 'latin' ]
	13012	 nhạc hot nhất => [ 'nhạc hot nhất' ]
	12551	 alternative => [ 'alternative' ]
	12029	 nhạc vàng hay nhất => [ 'Nhạc Việt Nam', 'Nhạc Vàng' ]
	11463	 violin => [ 'violin' ]
	11461	 guitar => [ 'guitar' ]
	11242	 phim hoạt hình => [ 'Nhạc Phim Hoạt Hình' ]
	10430	 nhạc cụ khác => [ 'nhạc cụ khác' ]
	9550	 house => [ 'house' ]
	9519	 cách mạng => [ 'Nhạc Việt Nam', 'Nhạc Cách Mạng', 'Nhạc Đỏ' ]
	9440	 nhạc cụ dân tộc => [ 'Nhạc Việt Nam', 'Nhạc Cụ Dân Tộc' ]
	8809	 thiếu nhi => [ 'Nhạc Việt Nam', 'Nhạc Thiếu Nhi' ]
	8657	 giải trí khác => [ 'giải trí khác' ]
	8237	 tất cả => []
	8183	 beat => [ 'beat' ]
	7983	 hàn => [ 'Nhạc Hàn', 'Tiếng Hàn' ]
	7573	 indie => [ 'indie' ]
	6975	 trance => [ 'trance' ]
	6957	 nhạc nhật bản => [ 'Nhạc Nhật', 'Tiếng Nhật' ]
	6808	 nhạc âu - mỹ => [ 'Nhạc Âu Mỹ' ]
	6229	 hip-hop => [ 'hip-hop' ]
	6004	 nhạc cách mạng => [ 'Nhạc Việt Nam', 'Nhạc Cách Mạng', 'Nhạc Đỏ' ]
	5925	 nhạc giáng sinh => [ 'nhạc giáng sinh', 'Christmas' ]
	5870	 techno => [ 'techno' ]
	5787	 liveshow => [ 'liveshow' ]
	5624	 rock việt => [ 'Nhạc Việt Nam', 'Rock' ]
	5417	 tiền chiến => [ 'Nhạc Việt Nam', 'Nhạc Tiền Chiến' ]
	4901	 nhạc tiếng anh hay nhất => [ 'Nhạc Âu Mỹ', 'Tiếng Anh' ]
	4901	 nhạc tiếng anh => [ 'Nhạc Âu Mỹ', 'Tiếng Anh' ]
	4843	 âm nhạc => []
	4758	 acoustic/audiophile => [ 'acoustic', 'audiophile' ]
	4658	 nhạc dân tộc => [ 'Nhạc Việt Nam', 'Nhạc Dân Tộc' ]
	4327	 phim trung quốc => [ 'Phim Trung Quốc', 'Tiếng Hoa' ]
	4290	 nhạc thiếu nhi => [ 'Nhạc Việt Nam', 'Nhạc Thiếu Nhi' ]
	4278	 nhạc tiền chiến => [ 'Nhạc Việt Nam', 'Nhạc Tiền Chiến' ]
	4155	 karaoke => [ 'Karaoke' ]
	4086	 nhạc đồng quê => [ 'Nhạc Việt Nam', 'Nhạc Đồng Quê' ]
	4065	 quê hương => [ 'Nhạc Việt Nam', 'Nhạc Quê Hương' ]
	4061	 beats => [ 'beats' ]
	3765	 cello => [ 'cello' ]
	3422	 nhạc spa => [ 'nhạc spa' ]
	3398	 nhạc không lời - hòa tấu hay nhất => [ 'Nhạc Không Lời', 'Nhạc Hòa Tấu' ]
	3398	 nhạc không lời - hòa tấu => [ 'Nhạc Không Lời', 'Nhạc Hòa Tấu' ]
	3156	 nhạc xuân => [ 'Nhạc Việt Nam', 'Nhạc Xuân' ]
	3131	 r&b/hip hop/rap => [ 'r&b', 'Hip hop', 'rap' ]
	3010	 việt remix => [ 'Nhạc Việt Nam', 'Remix' ]
	2949	 saxophone => [ 'saxophone' ]
	2768	 phim âu mỹ => [ 'Phim Âu Mỹ', 'Tiếng Anh' ]
	2727	 nhạc đỏ => [ 'Nhạc Việt Nam', 'Nhạc Cách Mạng', 'Nhạc Đỏ' ]
	2718	 love songs => [ 'love songs' ]
	2642	 anime => [ 'anime' ]
	2640	 hoa => [ 'Nhạc Hoa', 'Tiếng Hoa' ]
	2542	 nhạc chế - hài hước => [ 'nhạc chế', 'hài hước' ]
	2426	 nhạc sàn => [ 'nhạc sàn' ]
	2016	 nhacvui => []
	1931	 nhạc châu á => [ 'nhạc châu á' ]
	1931	 massage => [ 'massage' ]
	1867	 nhạc quê hương hay nhất => [ 'Nhạc Việt Nam', 'Nhạc Quê Hương' ]
	1861	 bài hát yêu thích => [ 'bài hát yêu thích' ]
	1854	 singapore => [ 'Nhạc Singapore', 'Tiếng Hoa' ]
	1805	 r&b/soul => [ 'r&b', 'soul' ]
	1758	 phim thiếu nhi => [ 'phim thiếu nhi' ]
	1754	 style => [ 'style' ]
	1704	 blue => [ 'blue' ]
	1574	 audiophile => [ 'audiophile' ]
	1501	 nhạc việt => [ 'Nhạc Việt Nam' ]
	1491	 thư giãn => [ 'thư giãn' ]
	1464	 dân ca => [ 'Nhạc Việt Nam', 'Dân Ca' ]
	1440	 blue/jazz => [ 'blue', 'jazz' ]
	1433	 cải lương => [ 'Cải Lương', 'Tân Cổ' ]
	1397	 truyện - sách audio hay nhất => [ 'truyện', 'sách audio' ]
	1397	 truyện - sách audio => [ 'truyện', 'sách audio' ]
	1357	 nhạc truyền thống hay nhất => [ 'Nhạc Việt Nam', 'Nhạc Truyền Thống' ]
	1357	 nhạc truyền thống => [ 'Nhạc Việt Nam', 'Nhạc Truyền Thống' ]
	1339	 nhạc quốc tế remix => [ 'Nhạc Quốc Tế' ]
	1300	 dj => [ 'dj' ]
	1260	 nonstop => [ 'nonstop' ]
	1241	 sách nói => [ 'sách nói' ]
	1238	 phim hàn quốc => [ 'Phim Hàn Quốc', 'Tiếng Hàn' ]
	1189	 nhạc rock => [ 'Rock' ]
	1185	 nhạc phim châu á hay nhất => [ 'nhạc phim châu á' ]
	1185	 nhạc phim châu á => [ 'nhạc phim châu á' ]
	1165	 nhạc rap - hip hop hay nhất => [ 'Rap', 'Hip hop' ]
	1165	 nhạc rap - hip hop => [ 'Rap', 'Hip hop' ]
	1137	 nhạc trịnh hay nhất => [ 'Nhạc Việt Nam', 'Nhạc Trịnh' ]
	1113	 nhạc dance - dj hay nhất => [ 'Dance', 'dj' ]
	1113	 nhạc dance - dj => [ 'Dance', 'dj' ]
	1102	 nhạc rap => [ 'Rap' ]
	1096	 clip vui => [ 'clip vui' ]
	1004	 nhạc nga => [ 'Nhạc Nga', 'Tiếng Nga' ]
	999	 dance việt => [ 'Nhạc Việt Nam', 'Dance' ]
	978	 nhạc tiếng nhật - hàn hay nhất => [ 'nhạc tiếng nhật', 'Nhạc Hàn', 'Tiếng Hàn' ]
	978	 nhạc tiếng nhật - hàn => [ 'nhạc tiếng nhật', 'Nhạc Hàn', 'Tiếng Hàn' ]
	963	 lounge => [ 'lounge' ]
	943	 pháp => [ 'Nhạc Pháp', 'Tiếng Pháp' ]
	925	 chilout => [ 'chilout' ]
	877	 phim việt nam => [ 'phim việt nam' ]
	873	 nhạc rock hay nhất => [ 'Rock' ]
	852	 dj - dance - remix => [ 'dj', 'Dance', 'remix' ]
	809	 n+ show => []
	804	 nhạc dance => [ 'Dance' ]
	802	 just for laughs gags (phần 2) => [ 'just for laughs gags (phần 2)' ]
	798	 nhạc chế vui => [ 'nhạc chế vui' ]
	781	 live show => [ 'live show' ]
	767	 nhạc tiếng hoa hay nhất => [ 'Nhạc Hoa', 'Tiếng Hoa' ]
	767	 nhạc tiếng hoa => [ 'Nhạc Hoa', 'Tiếng Hoa' ]
	766	 nhạc tiền chiến hay nhất => [ 'Nhạc Việt Nam', 'Nhạc Tiền Chiến' ]
	694	 nhạc spa | thư giãn => [ 'nhạc spa', 'thư giãn' ]
	687	 nhạc bà bầu & baby => [ 'nhạc bà bầu', 'baby' ]
	671	 nhạc bà bầu => [ 'nhạc bà bầu' ]
	615	 malaysia => [ 'Nhạc Malaysia', 'Tiếng Mã Lai' ]
	615	 audition => [ 'audition' ]
	604	 unknow => []
	571	 phim thái lan => [ 'Phim Thái Lan', 'Tiếng Thái' ]
	555	 phim nhật bản => [ 'Phim Nhật Bản', 'Tiếng Nhật' ]
	520	 nhạc cải lương - tân cổ hay nhất => [ 'Cải Lương', 'Tân Cổ', 'tân cổ' ]
	520	 nhạc cải lương - tân cổ => [ 'Cải Lương', 'Tân Cổ', 'tân cổ' ]
	473	 dân tộc => [ 'Nhạc Việt Nam', 'Nhạc Dân Tộc' ]
	447	 nhạc khiêu vũ => [ 'nhạc khiêu vũ' ]
	435	 nhạc game => [ 'nhạc game' ]
	434	 buddha => [ 'buddha' ]
	421	 music => [ 'music' ]
	394	 christian & gospel => [ 'christian', 'gospel' ]
	391	 nhạc dân ca hay nhất => [ 'Nhạc Việt Nam', 'Dân Ca' ]
	391	 nhạc dân ca => [ 'Nhạc Việt Nam', 'Dân Ca' ]
	365	 hoạt hình => [ 'Nhạc Phim Hoạt Hình' ]
	361	 nct tv => []
	349	 các game khác => [ 'các game khác' ]
	320	 nhạc nước khác hay nhất => [ 'Nhạc Quốc Tế' ]
	320	 nhạc nước khác => [ 'Nhạc Quốc Tế' ]
	317	 nhạc chuông => [ 'Nhạc Chuông' ]
	316	 nhạc thành viên hát hay nhất => []
	316	 nhạc thành viên hát => []
	304	 radio online => [ 'radio online' ]
	297	 nhạc spa & massage => [ 'nhạc spa', 'massage' ]
	292	 r&b/hiphop/rap => [ 'r&b', 'Hip hop', 'rap' ]
	275	 hiphop - rap => [ 'Hip hop', 'rap' ]
	273	 hit song => [ 'hit song' ]
	270	 yêu cầu nhạc => [ 'yêu cầu nhạc' ]
	268	 nhạc baby => [ 'nhạc baby' ]
	266	 nhạc học đường hay nhất => [ 'nhạc học đường' ]
	266	 nhạc học đường => [ 'nhạc học đường' ]
	262	 semi classic => [ 'semi classic' ]
	261	 nhạc phim tiếng anh hay nhất => [ 'nhạc phim tiếng anh' ]
	261	 nhạc phim tiếng anh => [ 'nhạc phim tiếng anh' ]
	242	 nhạc đạo => [ 'nhạc đạo' ]
	238	 nhạc phim việt nam hay nhất => [ 'nhạc phim việt nam' ]
	238	 nhạc phim việt nam => [ 'nhạc phim việt nam' ]
	238	 harp => [ 'harp' ]
	235	 nhạc game hay nhất => [ 'nhạc game' ]
	232	 vòng liveshow => [ 'vòng liveshow' ]
	232	 nhạc chế - hài hay nhất => [ 'nhạc chế', 'hài' ]
	232	 nhạc chế - hài => [ 'nhạc chế', 'hài' ]
	217	 giọng hát việt 2013 => [ 'Nhạc Việt Nam', 'Giọng Hát Việt' ]
	199	 dancesport => [ 'dancesport' ]
	193	 nhạc phim hot nhất => [ 'nhạc phim hot nhất' ]
	193	 acoustic => [ 'acoustic' ]
	192	 quan họ bắc ninh => [ 'Nhạc Việt Nam', 'Nhạc Quê Hương', 'Quan Họ Bắc Ninh' ]
	192	 nhạc tiếng pháp hay nhất => [ 'Nhạc Pháp', 'Tiếng Pháp' ]
	192	 nhạc tiếng pháp => [ 'Nhạc Pháp', 'Tiếng Pháp' ]
	173	 trailer => [ 'trailer' ]
	169	 tv show => [ 'tv show' ]
	168	 opera => [ 'opera' ]
	168	 giọng hát việt nhí => [ 'Nhạc Việt Nam', 'Giọng Hát Việt Nhí' ]
	166	 radio => [ 'radio' ]
	162	 vn's got talent 2011 => [ 'vn\'s got talent 2011' ]
	162	 nhạc karaoke hay nhất => [ 'Karaoke' ]
	162	 nhạc karaoke => [ 'Karaoke' ]
	161	 korean pop => [ 'korean pop' ]
	154	 nhạc thái => [ 'Nhạc Thái', 'Tiếng Thái' ]
	154	 cảm xúc => [ 'cảm xúc' ]
	148	 shining show => [ 'shining show' ]
	148	 nhạc phim hoạt hình hay nhất => [ 'Nhạc Phim Hoạt Hình' ]
	148	 nhạc phim hoạt hình => [ 'Nhạc Phim Hoạt Hình' ]
	147	 trumpet => [ 'trumpet' ]
	145	 nhạc thiền (meditation) => [ 'nhạc thiền (meditation)' ]
	142	 giọng hát việt 2012 => [ 'Nhạc Việt Nam', 'Giọng Hát Việt' ]
	141	 children's music => [ 'children\'s music' ]
	139	 ost => [ 'ost' ]
	139	 live video => [ 'live video' ]
	133	 sao vui => [ 'sao vui' ]
	130	 breakdance => [ 'breakdance' ]
	127	 nhạc quốc tế => [ 'Nhạc Quốc Tế' ]
	121	 nhạc đám cưới hay => [ 'Nhạc Đám Cưới' ]
	119	 giọng hát việt nhí 2013 => [ 'Nhạc Việt Nam', 'Giọng Hát Việt Nhí' ]
	116	 bellydance => [ 'bellydance' ]
	110	 quảng cáo quốc tế => [ 'Nhạc Quốc Tế' ]
	109	 new age/ world music => [ 'new age', 'world music' ]
	105	 nhảy cùng âm nhạc => [ 'nhảy cùng' ]
	105	 bước nhảy => [ 'bước nhảy' ]
	99	 ngâm thơ => [ 'Nhạc Việt Nam', 'Ngâm Thơ' ]
	97	 thành viên hát => [ 'thành viên hát' ]
	96	 christmas => [ 'christmas' ]
	90	 vòng đối đầu => [ 'vòng đối đầu' ]
	89	 music box => [ 'music box' ]
	89	 dance pop => [ 'dance pop' ]
	88	 pan flute => [ 'pan flute' ]
	86	 harmonica => [ 'harmonica' ]
	82	 mừng qt phụ nữ 8/3 => [ 'mừng qt phụ nữ 8', '3' ]
	81	 발라드 (ballad) => [ '발라드 (ballad)' ]
	79	 the voice vietnam (2012) => [ 'the voice vietnam (2012)' ]
	75	 vietnam idol 2013 => [ 'vietnam idol 2013' ]
	75	 hài vui nhộn => [ 'hài vui nhộn' ]
	73	 pop rock => [ 'Pop', 'Rock' ]
	73	 nhạc quảng cáo => [ 'nhạc quảng cáo' ]
	71	 house of dreams => [ 'house of dreams' ]
	71	 drum => [ 'drum' ]
	64	 vietnam's got talent => [ 'vietnam\'s got talent' ]
	63	 vietnam idol 2012 => [ 'vietnam idol 2012' ]
	63	 easy listening => [ 'easy listening' ]
	62	 nhạc tuyển tập => [ 'nhạc tuyển tập' ]
	60	 vòng giấu mặt => [ 'vòng giấu mặt' ]
	53	 vocal jazz => [ 'vocal jazz' ]
	52	 âm nhạc và bước nhảy => [ 'và bước nhảy' ]
	52	 nhạc phim khác hay nhất => [ 'nhạc phim khác' ]
	52	 nhạc phim khác => [ 'nhạc phim khác' ]
	51	 world => [ 'world' ]
	50	 nhạc chuông hay nhất => [ 'Nhạc Chuông' ]
	50	 french pop => [ 'Nhạc Pháp', 'Pop', 'Tiếng Pháp' ]
	50	 drama ost => [ 'drama ost' ]
	48	 cặp đôi hoàn hảo => [ 'cặp đôi hoàn hảo' ]
	47	 alternative rock => [ 'alternative rock' ]
	46	 lyrics video => [ 'lyrics video' ]
	43	 giải trí, hài kịch => [ 'giải trí', 'hài kịch' ]
	41	 electronica => [ 'electronica' ]
	41	 dân ca - ca cổ - cải lương => [ 'Nhạc Việt Nam', 'Dân Ca', 'ca cổ', 'Cải Lương', 'Tân Cổ' ]
	41	 chillout/lounge => [ 'chillout', 'lounge' ]
	40	 vocal => [ 'vocal' ]
	40	 latino => [ 'latino' ]
	38	 tv drama => [ 'tv drama' ]
	38	 indie pop => [ 'indie pop' ]
	38	 hip-hop/rap => [ 'hip-hop', 'rap' ]
	36	 classical crossover => [ 'classical crossover' ]
	35	 chillout => [ 'chillout' ]
	34	 댄스 (dance) => [ '댄스 (dance)' ]
	34	 vòng đo ván => [ 'vòng đo ván' ]
	34	 nhạc flash việt => [ 'Nhạc Việt Nam' ]
	34	 idol => [ 'idol' ]
	34	 hài => [ 'hài' ]
	33	 tv => [ 'tv' ]
	32	 sea show => [ 'sea show' ]
	31	 thành viên sáng tác hay nhất => [ 'thành viên sáng tác' ]
	31	 thành viên sáng tác => [ 'thành viên sáng tác' ]
	30	 guzheng => [ 'guzheng' ]
	29	 smooth jazz => [ 'smooth jazz' ]
	29	 reggae => [ 'reggae' ]
	29	 lời yêu thương => [ 'lời yêu thương' ]
	28	 solo piano => [ 'solo piano' ]
	28	 indie rock => [ 'indie rock' ]
	28	 ambient => [ 'ambient' ]
	26	 adult alternative => [ 'adult alternative' ]
	25	 quảng cáo => [ 'quảng cáo' ]
	25	 pop latino => [ 'pop latino' ]
	25	 mừng qt phụ nữ 8 => [ 'mừng qt phụ nữ 8' ]
	25	 downtempo => [ 'downtempo' ]
	25	 3 => [ '3' ]
	22	 드라마음악 => [ '드라마음악' ]
	22	 songwriter => [ 'songwriter' ]
	22	 singer => [ 'singer' ]
	22	 hard rock => [ 'hard rock' ]
	21	 dance cover => [ 'dance cover' ]
	21	 chanson => [ 'chanson' ]
	21	 [""] => [ '[""]' ]
	20	 teen pop => [ 'teen pop' ]
	20	 pop-rock => [ 'pop-rock' ]
	20	 gospel => [ 'gospel' ]
	19	 nhạc trữ tình, việt remix => [ 'Nhạc Việt Nam', 'Nhạc Trữ Tình', 'Remix' ]
	19	 french => [ 'Nhạc Pháp', 'Tiếng Pháp' ]
	19	 female vocal => [ 'female vocal' ]
	19	 celtic => [ 'celtic' ]
	17	 랩 => [ '랩' ]
	17	 dance video => [ 'dance video' ]
	16	 댄스 팝 (dance pop) => [ '댄스 팝 (dance pop)' ]
	16	 flash vui nhộn => [ 'flash vui nhộn' ]
	16	 billboard music award 2013 => [ 'billboard music award 2013' ]
	15	 알앤비 => [ '알앤비' ]
	15	 힙합 (rap => [ '힙합 (rap' ]
	15	 vòng chung kết => [ 'vòng chung kết' ]
	15	 j-pop => [ 'j-pop' ]
	14	 punk => [ 'punk' ]
	14	 nhạc chế => [ 'nhạc chế' ]
	14	 ngẫu hứng âm nhạc => [ 'ngẫu hứng' ]
	14	 hài hước => [ 'hài hước' ]
	14	 holiday => [ 'holiday' ]
	14	 [] => [ '[]' ]
	13	 nhạc thành viên chế hay nhất => [ 'nhạc thành viên chế' ]
	13	 nhạc thành viên chế => [ 'nhạc thành viên chế' ]
	13	 electronic / dance => [ 'electronic', 'Dance' ]
	12	 world music","rock","electronic => [ 'world music', 'rock', 'electronic' ]
	12	 vmvc 2011 => [ 'vmvc 2011' ]
	12	 tech => [ 'tech' ]
	12	 soul","folk","trance => [ 'soul', 'folk', 'trance' ]
	12	 relax => [ 'relax' ]
	12	 punk rock => [ 'punk rock' ]
	12	 nhảy cùng âm nhạc & bước nhảy => [ 'nhảy cùng', 'bước nhảy' ]
	12	 jazz","r&b => [ 'jazz', 'r&b' ]
	12	 jazz pop => [ 'jazz pop' ]
	12	 hip hop","blues => [ 'Hip hop', 'blues' ]
	12	 heavy metal => [ 'heavy metal' ]
	12	 gothic rock => [ 'gothic rock' ]
	12	 funk => [ 'funk' ]
	12	 fanmade => [ 'fanmade' ]
	12	 dance","rap => [ 'Dance', 'rap' ]
	12	 christian => [ 'christian' ]
	12	 acoustic guitar => [ 'acoustic guitar' ]
	11	 nhạc vàng, việt remix => [ 'Nhạc Việt Nam', 'Nhạc Vàng', 'Remix' ]
	11	 nhạc nước ngoài => [ 'Nhạc Quốc Tế' ]
	11	 modern classical => [ 'modern classical' ]
	11	 hiphop) => [ 'hiphop)' ]
	11	 ethnic => [ 'ethnic' ]
	11	 bossa nova => [ 'bossa nova' ]
	11	 bluegrass => [ 'bluegrass' ]
	10	 락 (rock) => [ '락 (rock)' ]
	10	 truyện dài => [ 'truyện dài' ]
	10	 pop dance => [ 'pop dance' ]
	10	 original score => [ 'original score' ]
	10	 heineken => [ 'heineken' ]
	9	 인디뮤직 (indie music) => [ '인디뮤직 (indie music)' ]
	9	 어반 (r&b => [ '어반 (r&b' ]
	9	 synthpop => [ 'synthpop' ]
	9	 pop ballad => [ 'pop ballad' ]
	9	 giọng hát việt => [ 'giọng hát việt' ]
	9	 electro => [ 'electro' ]
	9	 alt. rock => [ 'alt. rock' ]
	9	 alt-country => [ 'alt-country' ]
	9	 album một tuần giận nhau => [ 'album một tuần giận nhau' ]
	8	 발라드(ballad) => [ '발라드(ballad)' ]
	8	 urban) => [ 'urban)' ]
	8	 trot => [ 'trot' ]
	8	 score => [ 'score' ]
	8	 piano solo => [ 'piano solo' ]
	8	 mtv exit => [ 'mtv exit' ]
	8	 indiepop => [ 'indiepop' ]
	8	 folk rock => [ 'folk rock' ]
	8	 flash quốc tế => [ 'Nhạc Quốc Tế' ]
	8	 contemporary jazz => [ 'contemporary jazz' ]
	8	 ["âu mỹ","country","pop","new age => [ '["âu mỹ', 'country', 'pop', 'new age' ]
	7	 드라마 => [ '드라마' ]
	7	 đời sống - tâm hồn => [ 'đời sống', 'tâm hồn' ]
	7	 the voice - giọng hát việt => [ 'the voice', 'giọng hát việt' ]
	7	 nhạc nước ngoài hay nhất => [ 'Nhạc Quốc Tế' ]
	7	 indie-pop => [ 'indie-pop' ]
	7	 gangsta rap => [ 'gangsta rap' ]
	7	 folk pop => [ 'folk pop' ]
	7	 electropop => [ 'electropop' ]
	7	 club => [ 'club' ]
	6	 west coast rap => [ 'west coast rap' ]
	6	 salsa y tropical => [ 'salsa y tropical' ]
	6	 reggaeton y hip-hop => [ 'reggaeton y hip-hop' ]
	6	 progressive rock => [ 'progressive rock' ]
	6	 pop-folk => [ 'pop-folk' ]
	6	 nhạc việt hay nhất => [ 'Nhạc Việt Nam' ]
	6	 melodic rock => [ 'melodic rock' ]
	6	 latin pop => [ 'latin pop' ]
	6	 hip-hop) => [ 'hip-hop)' ]
	6	 epic music => [ 'epic music' ]
	6	 dubstep => [ 'dubstep' ]
	6	 contemporary r&b => [ 'contemporary r&b' ]
	6	 contemporary country => [ 'contemporary country' ]
	6	 chinh phục đỉnh cao => [ 'chinh phục đỉnh cao' ]
	6	 chacha star (thành viên hát) => [ 'chacha star (thành viên hát)' ]
	6	 art rock => [ 'art rock' ]
	6	 adult contemporary => [ 'adult contemporary' ]
	5	 truyện audio => [ 'truyện audio' ]
	5	 soul) => [ 'soul)' ]
	5	 rnb => [ 'rnb' ]
	5	 prog-rock => [ 'prog-rock' ]
	5	 pop punk => [ 'pop punk' ]
	5	 pop folk => [ 'pop folk' ]
	5	 nu jazz => [ 'nu jazz' ]
	5	 musik => [ 'musik' ]
	5	 movie ost => [ 'movie ost' ]
	5	 meditative => [ 'meditative' ]
	5	 live => [ 'live' ]
	5	 jazz. => [ 'jazz.' ]
	5	 dream pop => [ 'dream pop' ]
	5	 christian rock => [ 'christian rock' ]
	5	 children’s music => [ 'children’s music' ]
	5	 britpop => [ 'britpop' ]
	5	 blues rock => [ 'blues rock' ]
	5	 alternative country => [ 'alternative country' ]
	4	 댄스(dance) => [ '댄스(dance)' ]
	4	 vocal pop => [ 'vocal pop' ]
	4	 the voice => [ 'the voice' ]
	4	 symphonic rock => [ 'symphonic rock' ]
	4	 style ballad => [ 'style ballad' ]
	4	 psychedelic => [ 'psychedelic' ]
	4	 piano. => [ 'piano.' ]
	4	 piano rock => [ 'piano rock' ]
	4	 pepsi dj bus => [ 'pepsi dj bus' ]
	4	 nhạc trẻ, nhạc trữ tình => [ 'Nhạc Việt Nam', 'Nhạc Trẻ', 'Nhạc Trữ Tình' ]
	4	 nhạc trữ tình, nhạc quê hương => [ 'Nhạc Việt Nam', 'Nhạc Trữ Tình', 'Nhạc Quê Hương' ]
	4	 nhạc chuông mobile => [ 'Nhạc Chuông' ]
	4	 new wave => [ 'new wave' ]
	4	 latin jazz => [ 'latin jazz' ]
	4	 just for laughs gags (phần 1) => [ 'just for laughs gags (phần 1)' ]
	4	 instrumental rock => [ 'instrumental rock' ]
	4	 inspirational => [ 'inspirational' ]
	4	 indie folk => [ 'indie folk' ]
	4	 hành động => [ 'hành động' ]
	4	 honky tonk => [ 'honky tonk' ]
	4	 healing => [ 'healing' ]
	4	 hardcore rap => [ 'hardcore rap' ]
	4	 gương mặt thân quen => [ 'gương mặt thân quen' ]
	4	 guzheng | instrumental => [ 'guzheng', 'instrumental' ]
	4	 fitness & workout => [ 'fitness', 'workout' ]
	4	 disco => [ 'disco' ]
	4	 classic => [ 'classic' ]
	4	 chill out => [ 'chill out' ]
	4	 black metal => [ 'black metal' ]
	4	 americana => [ 'americana' ]
	4	 ["country","âu mỹ","pop","new age => [ '["country', 'Nhạc Âu Mỹ', 'pop', 'new age' ]
	4	 .. pop => [ '.. pop' ]
	4	 . r&b => [ '. r&b' ]
	3	 일렉트로니카 (electronica) => [ '일렉트로니카 (electronica)' ]
	3	 팝락 (pop rock) => [ '팝락 (pop rock)' ]
	3	 포크 (folk) => [ '포크 (folk)' ]
	3	 락(rock) => [ '락(rock)' ]
	3	 urban => [ 'urban' ]
	3	 symphonic power metal => [ 'symphonic power metal' ]
	3	 swing => [ 'swing' ]
	3	 singer-songwriter => [ 'singer-songwriter' ]
	3	 schlager => [ 'schlager' ]
	3	 r & b => [ 'r', 'b' ]
	3	 psychill => [ 'psychill' ]
	3	 post-punk => [ 'post-punk' ]
	3	 pop latin => [ 'pop latin' ]
	3	 pipa => [ 'pipa' ]
	3	 ost tv drama => [ 'ost tv drama' ]
	3	 operatic pop => [ 'operatic pop' ]
	3	 nhạc trữ tình, nhạc vàng => [ 'Nhạc Việt Nam', 'Nhạc Trữ Tình', 'Nhạc Vàng' ]
	3	 neo-classical => [ 'neo-classical' ]
	3	 lullabies => [ 'lullabies' ]
	3	 lo-fi => [ 'lo-fi' ]
	3	 k-pop => [ 'k-pop' ]
	3	 game võ lâm => [ 'game võ lâm' ]
	3	 experimental => [ 'experimental' ]
	3	 europop => [ 'europop' ]
	3	 crossover => [ 'crossover' ]
	3	 country rock => [ 'country rock' ]
	3	 contemporary gospel => [ 'contemporary gospel' ]
	3	 classical music => [ 'classical music' ]
	3	 brazilian => [ 'brazilian' ]
	3	 arena rock => [ 'arena rock' ]
	3	 alternative rap => [ 'alternative rap' ]
	3	 acoustic rock => [ 'acoustic rock' ]
	3	 acoustic blues => [ 'acoustic blues' ]
	3	 . country => [ '. country' ]
	3	 (k-pop) \t 발라드(ballad) => [ '(k-pop) \\t 발라드(ballad)' ]
	2	 일렉트로니카(electronica) => [ '일렉트로니카(electronica)' ]
	2	 알앤비(r&b) => [ '알앤비(r&b)' ]
	2	 알앤비 (r&b) => [ '알앤비 (r&b)' ]
	2	 어반 (r & b) => [ '어반 (r', 'b)' ]
	2	 어반 (r & b => [ '어반 (r', 'b' ]
	2	 ℗ 2013 pink revolver => [ '℗ 2013 pink revolver' ]
	2	 ℗ 1997 emec records => [ '℗ 1997 emec records' ]
	2	 록 (rock) => [ '록 (rock)' ]
	2	 yangqin => [ 'yangqin' ]
	2	 xiao => [ 'xiao' ]
	2	 world music","hoa ngữ","đài loan","trung quốc","hòa tấu","guitar","nhạc cụ dân tộc","harmonica => [ 'world 	music',
	2	  'Nhạc Hoa',
	2	  'Tiếng Hoa',
	2	  'Đài Loan',
	2	  'Nhạc Hòa Tấu',
	2	  'guitar',
	2	  'Nhạc Việt Nam',
	2	  'Nhạc Cụ Dân Tộc',
	2	  'harmonica' ]
	2	 vocal trance => [ 'vocal trance' ]
	2	 underground rap => [ 'underground rap' ]
	2	 truyện thiếu nhi => [ 'truyện thiếu nhi' ]
	2	 tango => [ 'tango' ]
	2	 synth-pop => [ 'synth-pop' ]
	2	 synth pop => [ 'synth pop' ]
	2	 style jazz => [ 'style jazz' ]
	2	 standards => [ 'standards' ]
	2	 spiritua => [ 'spiritua' ]
	2	 soundtrack. => [ 'soundtrack.' ]
	2	 soul. => [ 'soul.' ]
	2	 soul pop => [ 'soul pop' ]
	2	 soul / r&b => [ 'soul', 'r&b' ]
	2	 soft rock => [ 'soft rock' ]
	2	 screamo => [ 'screamo' ]
	2	 saxophone. => [ 'saxophone.' ]
	2	 samba-jazz => [ 'samba-jazz' ]
	2	 samba => [ 'samba' ]
	2	 salsa and tropical => [ 'salsa and tropical' ]
	2	 romantic => [ 'romantic' ]
	2	 rockabilly => [ 'rockabilly' ]
	2	 relaxing => [ 'relaxing' ]
	2	 power pop => [ 'power pop' ]
	2	 power metal => [ 'power metal' ]
	2	 post-hardcore => [ 'post-hardcore' ]
	2	 pop-rap => [ 'pop-rap' ]
	2	 pop-punk => [ 'pop-punk' ]
	2	 pop rap => [ 'pop rap' ]
	2	 pop metal => [ 'pop metal' ]
	2	 pop jazz => [ 'pop jazz' ]
	2	 phim ..." /> => [ 'phim ..."' ]
	2	 pan flute","violin"," => [ 'pan flute', 'violin' ]
	2	 orcestral => [ 'orcestral' ]
	2	 nu-jazz => [ 'nu-jazz' ]
	2	 nhạc trẻ, việt remix => [ 'Nhạc Việt Nam', 'Nhạc Trẻ', 'Remix' ]
	2	 nhạc trẻ, rap việt => [ 'Nhạc Việt Nam', 'Nhạc Trẻ', 'Rap' ]
	2	 nhạc trữ tình, nhạc trịnh => [ 'Nhạc Việt Nam', 'Nhạc Trữ Tình', 'Nhạc Trịnh' ]
	2	 new age. => [ 'new age.' ]
	2	 neofolk => [ 'neofolk' ]
	2	 neoclassical => [ 'neoclassical' ]
	2	 music - => [ 'music -' ]
	2	 meditation => [ 'meditation' ]
	2	 matouqin => [ 'matouqin' ]
	2	 latin rap => [ 'latin rap' ]
	2	 l acoustic => [ 'l acoustic' ]
	2	 jazzy blues => [ 'jazzy blues' ]
	2	 jazzy => [ 'jazzy' ]
	2	 jazzl => [ 'jazzl' ]
	2	 jazz-rock => [ 'jazz-rock' ]
	2	 jazz-pop => [ 'jazz-pop' ]
	2	 italo-disco => [ 'italo-disco' ]
	2	 instrumental. => [ 'instrumental.' ]
	2	 instrumental pop => [ 'instrumental pop' ]
	2	 instrumental jazz => [ 'instrumental jazz' ]
	2	 indie pop|latin => [ 'indie pop|latin' ]
	2	 indie pop-rock => [ 'indie pop-rock' ]
	2	 hulusi => [ 'hulusi' ]
	2	 holidays => [ 'holidays' ]
	2	 harmony vocal => [ 'harmony vocal' ]
	2	 hardcore punk => [ 'hardcore punk' ]
	2	 guitar. => [ 'guitar.' ]
	2	 guitar solo => [ 'guitar solo' ]
	2	 game soundtrack => [ 'game soundtrack' ]
	2	 future pop => [ 'future pop' ]
	2	 fusion => [ 'fusion' ]
	2	 folklore => [ 'folklore' ]
	2	 flute => [ 'flute' ]
	2	 flamenco guitar => [ 'flamenco guitar' ]
	2	 film music => [ 'film music' ]
	2	 europe => [ 'europe' ]
	2	 ethiopian pop => [ 'ethiopian pop' ]
	2	 erhu => [ 'erhu' ]
	2	 electronic pop => [ 'electronic pop' ]
	2	 electro pop => [ 'electro pop' ]
	2	 dub => [ 'dub' ]
	2	 drum & bass => [ 'drum', 'bass' ]
	2	 deutsch-pop => [ 'deutsch-pop' ]
	2	 death metal => [ 'death metal' ]
	2	 dark trip-hop => [ 'dark trip-hop' ]
	2	 damtv => [ 'damtv' ]
	2	 cpop => [ 'cpop' ]
	2	 country pop => [ 'country pop' ]
	2	 contenporary jazz => [ 'contenporary jazz' ]
	2	 contemporary folk => [ 'contemporary folk' ]
	2	 classical. => [ 'classical.' ]
	2	 classical pop => [ 'classical pop' ]
	2	 classic rock => [ 'classic rock' ]
	2	 classic opera => [ 'classic opera' ]
	2	 classic christian => [ 'classic christian' ]
	2	 chillwave => [ 'chillwave' ]
	2	 chill house => [ 'chill house' ]
	2	 chill => [ 'chill' ]
	2	 british rock folk => [ 'british rock folk' ]
	2	 british folk => [ 'british folk' ]
	2	 blues-rock => [ 'blues-rock' ]
	2	 bass => [ 'bass' ]
	2	 aor => [ 'aor' ]
	2	 alternative pop => [ 'alternative pop' ]
	1	 alt.rock => [ 'alt.rock' ]
	1	 african folk => [ 'african folk' ]
	1	 ["âu mỹ","new age => [ '["âu mỹ', 'new age' ]
	1	 . soul => [ '. soul' ]
	1	 . latin => [ '. latin' ]
	1	 . jazz => [ '. jazz' ]
	1	 . indie => [ '. indie' ]
	1	 . alternative => [ '. alternative' ]
	1	 일렉트로닉 팝 (electronic pop) => [ '일렉트로닉 팝 (electronic pop)' ]
	1	 인디뮤직 (indie) => [ '인디뮤직 (indie)' ]
	1	 인디뮤직 (indie music l 캐롤 (carol) => [ '인디뮤직 (indie music l 캐롤 (carol)' ]
	1	 트로트 (trot) => [ '트로트 (trot)' ]
	1	 댄스dance => [ '댄스dance' ]
	1	 힙합(rap => [ '힙합(rap' ]
	1	 어반(r & b => [ '어반(r', 'b' ]
	1	 어반( r & b => [ '어반( r', 'b' ]
	1	 인디 락 (indie rock) => [ '인디 락 (indie rock)' ]
	1	 인디 (indie) => [ '인디 (indie)' ]
	1	 힙합 (hip-hop) => [ '힙합 (hip-hop)' ]
	1	 힙합 => [ '힙합' ]
	1	 팝락 => [ '팝락' ]
	1	 댄스 => [ '댄스' ]
	1	 게임 => [ '게임' ]
	1	 сountry => [ 'сountry' ]
	1	 ：new age => [ '：new age' ]
	1	 팝 락 (pop rock) => [ '팝 락 (pop rock)' ]
	1	 팝 (pop) => [ '팝 (pop)' ]
	1	 락 (pop rock) => [ '락 (pop rock)' ]
	1	 whistle => [ 'whistle' ]
	1	 virgin records ltd => [ 'virgin records ltd' ]
	1	 universal music operations ltd. => [ 'universal music operations ltd.' ]
	1	 ukulele pop => [ 'ukulele pop' ]
	1	 trailer music => [ 'trailer music' ]
	1	 top 20 => [ 'top 20' ]
	1	 symphonic metal => [ 'symphonic metal' ]
	1	 symphonic gothic metal => [ 'symphonic gothic metal' ]
	1	 style rock => [ 'style rock' ]
	1	 style rap => [ 'style rap' ]
	1	 style r&b => [ 'style r&b' ]
	1	 style pop rock => [ 'style pop rock' ]
	1	 style indie rock => [ 'style indie rock' ]
	1	 style indie => [ 'style indie' ]
	1	 style dance => [ 'style dance' ]
	1	 style ost => [ 'style ost' ]
	1	 spiritual => [ 'spiritual' ]
	1	 southern rock => [ 'southern rock' ]
	1	 sould => [ 'sould' ]
	1	 soul jazz => [ 'soul jazz' ]
	1	 soul country => [ 'soul country' ]
	1	 smooth soul => [ 'smooth soul' ]
	1	 ska => [ 'ska' ]
	1	 sitcom ost => [ 'sitcom ost' ]
	1	 singer/songwriter => [ 'singer', 'songwriter' ]
	1	 rock punk => [ 'rock punk' ]
	1	 rock music => [ 'rock music' ]
	1	 rock 'n' roll => [ 'rock \'n\' roll' ]
	1	 rock & roll => [ 'rock', 'roll' ]
	1	 roc k => [ 'roc k' ]
	1	 rnb. => [ 'rnb.' ]
	1	 retro => [ 'retro' ]
	1	 religious => [ 'religious' ]
	1	 reggea => [ 'reggea' ]
	1	 reggaeton => [ 'reggaeton' ]
	1	 rap. => [ 'rap.' ]
	1	 rap-pop => [ 'rap-pop' ]
	1	 rap) => [ 'rap)' ]
	1	 rap việt, nhạc việt => [ 'Nhạc Việt Nam', 'Rap' ]
	1	 r&n => [ 'r&n' ]
	1	 r&b (rap => [ 'r&b (rap' ]
	1	 quality 320 kbits => [ 'quality 320 kbits' ]
	1	 punk (pop punk) => [ 'punk (pop punk)' ]
	1	 progressive => [ 'progressive' ]
	1	 postmodern => [ 'postmodern' ]
	1	 post-rock => [ 'post-rock' ]
	1	 post-dubstep => [ 'post-dubstep' ]
	1	 post punk => [ 'post punk' ]
	1	 pop-county => [ 'pop-county' ]
	1	 pop vocal => [ 'pop vocal' ]
	1	 pop french => [ 'Nhạc Pháp', 'Pop', 'Tiếng Pháp' ]
	1	 pop disco => [ 'pop disco' ]
	1	 pop (r&b) => [ 'pop (r&b)' ]
	1	 pop & rock => [ 'pop', 'rock' ]
	1	 pop & contemporary => [ 'pop', 'contemporary' ]
	1	 piano lounge => [ 'piano lounge' ]
	1	 piano jazz => [ 'piano jazz' ]
	1	 others => [ 'others' ]
	1	 ost> tv drama => [ 'ost> tv drama' ]
	1	 ost v drama => [ 'ost v drama' ]
	1	 orchestral => [ 'orchestral' ]
	1	 oldies => [ 'oldies' ]
	1	 nu jazz vocal => [ 'nu jazz vocal' ]
	1	 nouveau flamenco => [ 'nouveau flamenco' ]
	1	 nhạc trẻ, nhạc đỏ => [ 'Nhạc Việt Nam', 'Nhạc Trẻ', 'Nhạc Cách Mạng', 'Nhạc Đỏ' ]
	1	 nhạc trẻ, nhạc âu mỹ => [ 'Nhạc Việt Nam', 'Nhạc Trẻ', 'Nhạc Âu Mỹ' ]
	1	 nhạc trữ tình, rap việt, nhạc chế - hài hước => [ 'Nhạc Việt Nam', 'Nhạc Trữ Tình', 'Rap', 'nhạc chế', 	'hài hước' ]
	1	 nhạc cover => [ 'nhạc cover' ]
	1	 newage => [ 'newage' ]
	1	 new age-chillout => [ 'new age-chillout' ]
	1	 neo classical - piano => [ 'neo classical', 'piano' ]
	1	 nct channel => [ 'nct channel' ]
	1	 nature => [ 'nature' ]
	1	 natural sound => [ 'natural sound' ]
	1	 napoletana => [ 'napoletana' ]
	1	 muziek => [ 'muziek' ]
	1	 musical => [ 'musical' ]
	1	 music released 02 march 2012 ℗ 2012 the copyright in this compilation is owned by emi records ltd => [ 	'music released 02 march 2012 ℗ 2012 the copyright in this compilation is owned by emi records ltd' ]
	1	 modern rock => [ 'modern rock' ]
	1	 modern classic => [ 'modern classic' ]
	1	 minimalism => [ 'minimalism' ]
	1	 melodic symphonic power metal => [ 'melodic symphonic power metal' ]
	1	 mainstream jazz => [ 'mainstream jazz' ]
	1	 ly kỳ => [ 'ly kỳ' ]
	1	 lo fi => [ 'lo fi' ]
	1	 llatin => [ 'llatin' ]
	1	 lite rock => [ 'lite rock' ]
	1	 liem test => [ 'liem test' ]
	1	 leftfield => [ 'leftfield' ]
	1	 latin christian => [ 'latin christian' ]
	1	 label xl recordings => [ 'label xl recordings' ]
	1	 l country => [ 'l country' ]
	1	 korean classical => [ 'korean classical' ]
	1	 jazzy pop => [ 'jazzy pop' ]
	1	 jazz-pop smooth jazz => [ 'jazz-pop smooth jazz' ]
	1	 jazz vocals => [ 'jazz vocals' ]
	1	 jazz vocal jazz => [ 'jazz vocal jazz' ]
	1	 jazz vocal => [ 'jazz vocal' ]
	1	 jazz soul => [ 'jazz soul' ]
	1	 jazz bossa nova => [ 'jazz bossa nova' ]
	1	 italian pop => [ 'italian pop' ]
	1	 instrumental hip-hop => [ 'instrumental hip-hop' ]
	1	 instrumental hip hop => [ 'instrumental hip hop' ]
	1	 instrumental guitar rock => [ 'instrumental guitar rock' ]
	1	 instrumental classical piano => [ 'instrumental classical piano' ]
	1	 instrumental classical => [ 'instrumental classical' ]
	1	 instrument => [ 'instrument' ]
	1	 instrumeltal => [ 'instrumeltal' ]
	1	 industrial gothic metal => [ 'industrial gothic metal' ]
	1	 indie (rock) => [ 'indie (rock)' ]
	1	 hình sự/hành động /tâm lí" /> => [ 'hình sự', 'hành động', 'tâm lí"' ]
	1	 hiphop-rap => [ 'hiphop-rap' ]
	1	 hip-hop r&b => [ 'hip-hop r&b' ]
	1	 hip hop. => [ 'hip hop.' ]
	1	 hardcore => [ 'hardcore' ]
	1	 gypsy jazz => [ 'gypsy jazz' ]
	1	 gypsy => [ 'gypsy' ]
	1	 guzheng (tranh) => [ 'guzheng (tranh)' ]
	1	 guqin => [ 'guqin' ]
	1	 grime => [ 'grime' ]
	1	 gothic metal => [ 'gothic metal' ]
	1	 glam hard rock => [ 'glam hard rock' ]
	1	 gangsta rap (rap) => [ 'gangsta rap (rap)' ]
	1	 game ost => [ 'game ost' ]
	1	 future garage => [ 'future garage' ]
	1	 french ins. => [ 'french ins.' ]
	1	 french chanson => [ 'french chanson' ]
	1	 france pop => [ 'france pop' ]
	1	 folk. => [ 'folk.' ]
	1	 folk-rock => [ 'folk-rock' ]
	1	 eurodance => [ 'eurodance' ]
	1	 euro pop => [ 'euro pop' ]
	1	 euro dance => [ 'euro dance' ]
	1	 emo => [ 'emo' ]
	1	 electropop / indie pop / electronic => [ 'electropop', 'indie pop', 'electronic' ]
	1	 electronic (dance-pop) => [ 'electronic (dance-pop)' ]
	1	 electric violin => [ 'electric violin' ]
	1	 dub techno => [ 'dub techno' ]
	1	 dreampop => [ 'dreampop' ]
	1	 disco released jan 22 => [ 'disco released jan 22' ]
	1	 delta blues => [ 'delta blues' ]
	1	 dance-pop => [ 'dance-pop' ]
	1	 dance (dance => [ 'dance (dance' ]
	1	 dance & dj => [ 'Dance', 'dj' ]
	1	 dance pop => [ 'dance pop' ]
	1	 cross over => [ 'cross over' ]
	1	 country & wester => [ 'country', 'wester' ]
	1	 contemporary classical => [ 'contemporary classical' ]
	1	 contemporary acoustic => [ 'contemporary acoustic' ]
	1	 contemporary => [ 'contemporary' ]
	1	 comedy => [ 'comedy' ]
	1	 classical vocal => [ 'classical vocal' ]
	1	 classical instrumental => [ 'classical instrumental' ]
	1	 classical christmas => [ 'classical christmas' ]
	1	 classic crossover => [ 'classic crossover' ]
	1	 christmas rock => [ 'christmas rock' ]
	1	 christmas pop => [ 'christmas pop' ]
	1	 choral => [ 'choral' ]
	1	 chirstmas => [ 'chirstmas' ]
	1	 chanson pop => [ 'chanson pop' ]
	1	 celtic folk => [ 'celtic folk' ]
	1	 ccm => [ 'ccm' ]
	1	 carol => [ 'carol' ]
	1	 caribbean music => [ 'caribbean music' ]
	1	 buddhist => [ 'buddhist' ]
	1	 breakbeat => [ 'breakbeat' ]
	1	 bitrate vbr => [ 'bitrate vbr' ]
	1	 beyond jazz => [ 'beyond jazz' ]
	1	 beach reggae => [ 'beach reggae' ]
	1	 bass music => [ 'bass music' ]
	1	 baroque => [ 'baroque' ]
	1	 ballads => [ 'ballads' ]
	1	 ballad"] => [ 'ballad"]' ]
	1	 ballad", "nhật bản", "pop => [ 'ballad"', '"nhật bản"', '"pop' ]
	1	 balkan => [ 'balkan' ]
	1	 authentic => [ 'authentic' ]
	1	 aucostic => [ 'aucostic' ]
	1	 arabic music => [ 'arabic music' ]
	1	 alternative metal => [ 'alternative metal' ]
	1	 alt. folk => [ 'alt. folk' ]
	1	 alt. country => [ 'alt. country' ]
	1	 album xin yêu tôi bằng cả tình... => [ 'album xin yêu tôi bằng cả tình...' ]
	1	 album vẫn yêu em như ngày xưa... => [ 'album vẫn yêu em như ngày xưa...' ]
	1	 adult alternative pop => [ 'adult alternative pop' ]
	1	 acoustic pop-punk => [ 'acoustic pop-punk' ]
	1	 acoustic folk => [ 'acoustic folk' ]
	1	 acid jazz => [ 'acid jazz' ]
	1	 acappella => [ 'acappella' ]
	1	 abstract chill => [ 'abstract chill' ]
		 ["hàn quốc", "pop => [ '["hàn quốc"', '"pop' ]
		 2010 ℗ 2010 sony music entertainment germany gmbh => [ '2010 ℗ 2010 sony music entertainment germany gmbh' ]
		 . soundtrack => [ '. soundtrack' ]
		 . pop => [ '. pop' ]
		 . hip-hop => [ '. hip-hop' ]
		 (k-pop) \t 얼터너티브 (alternative) => [ '(k-pop) \\t 얼터너티브 (alternative)' ]
		 (k-pop) \t 댄스(dance) => [ '(k-pop) \\t 댄스(dance)' ]
		 (k-pop) \t 락(rock) => [ '(k-pop) \\t 락(rock)' ]













*/
package table
