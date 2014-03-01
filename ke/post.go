package ke

import (
	"dna"
	"dna/http"
	"encoding/json"
	"errors"
	dhttp "net/http"
)

func Post(url dna.String, header dhttp.Header, bodyStr dna.String) (*http.Result, error) {
	http.DefaulHeader = header
	return http.Post(url, bodyStr)
}

// GetAPILyric returns a lyric or an error from API using POST method.
//
// The SOAP data has the following format:
//
// 	<?xml version="1.0" encoding="utf-8"?>
// 	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
// 	  <soap:Header/>
// 	  <soap:Body>
// 	    <tem:getLyric xmlns=" http://tempuri.org/ ">
// 	      <tem:token></tem:token>
// 	      <tem:id>1944090</tem:id>
// 	    </tem:getLyric>
// 	  </soap:Body>
// 	</soap:Envelope>
func GetAPILyric(id dna.Int) (*APILyric, error) {
	if id == 0 {
		return nil, errors.New("Id is zero")
	}
	var dat dna.String = dna.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/"><soap:Header/><soap:Body>
<tem:getLyric xmlns=" http://tempuri.org/ "><tem:token></tem:token><tem:id>%v</tem:id></tem:getLyric></soap:Body>
</soap:Envelope>`, id)
	header := Header
	header.Set("SOAPAction", "http://tempuri.org/getLyric")
	ret, err := Post("http://service.keeng.vn/appwebservice/Service.asmx?wsdl", header, dat)
	if err != nil {
		return nil, err
	} else {
		var lyric APILyric
		data := ret.Data.FindAllStringSubmatch(`<return>(.+)</return>`, -1)[0][1].DecodeHTML()
		err := json.Unmarshal([]byte(data), &lyric)
		if err != nil {
			return nil, err
		} else {
			lyric.Data = lyric.Data.DecodeHTML()
			return &lyric, nil
		}
	}
}

// GetAPIAlbum returns an album or an error from API using POST method.
//
// The SOAP data has the following format:
//
// 	<?xml version="1.0" encoding="utf-8"?>
// 	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
// 	  <soap:Header/>
// 	  <soap:Body>
// 	    <tem:getAlbum_v2 xmlns=" http://tempuri.org/ ">
// 	      <tem:token></tem:token>
// 	      <tem:id>86682</tem:id>
// 	      <tem:identify></tem:identify>
// 	    </tem:getAlbum_v2>
// 	  </soap:Body>
// 	</soap:Envelope>
func GetAPIAlbum(id dna.Int) (*APIAlbum, error) {
	if id == 0 {
		return nil, errors.New("Id is zero")
	}
	var dat dna.String = dna.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/"><soap:Header/><soap:Body>
<tem:getAlbum_v2 xmlns=" http://tempuri.org/ "><tem:token></tem:token><tem:id>%v</tem:id><tem:identify></tem:identify></tem:getAlbum_v2></soap:Body>
</soap:Envelope>`, id)
	header := Header
	header.Set("SOAPAction", "http://tempuri.org/getAlbum_v2")
	ret, err := Post("http://service.keeng.vn/appwebservice/Service.asmx?wsdl", header, dat)
	if err != nil {
		return nil, err
	} else {
		var apiStatusAlbum APIStatusAlbum
		dataArr := ret.Data.FindAllStringSubmatch(`<return>(.+)</return>`, -1)
		if len(dataArr) > 0 {
			// dna.Log(data)
			err := json.Unmarshal([]byte(dataArr[0][1].DecodeHTML()), &apiStatusAlbum)
			if err != nil {
				return nil, err
			} else {
				return &apiStatusAlbum.Data, nil
			}
		} else {
			return nil, errors.New("No return value")
		}

	}
}

// GetAPISongEntry returns a song and relevant songs or an error from API using POST method.
//
// The SOAP data has the following format:
//
//	<?xml version="1.0" encoding="utf-8"?>
//	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
//	  <soap:Header/>
//	  <soap:Body>
//	    <tem:getSong_v2 xmlns=" http://tempuri.org/ ">
//	      <tem:token></tem:token>
//	      <tem:id>1968535</tem:id>
//	      <tem:identify></tem:identify>
//	    </tem:getSong_v2>
//	  </soap:Body>
//	</soap:Envelope>
func GetAPISongEntry(id dna.Int) (*APISongEntry, error) {
	if id == 0 {
		return nil, errors.New("Id is zero")
	}
	var dat dna.String = dna.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/"><soap:Header/><soap:Body>
<tem:getSong_v2 xmlns=" http://tempuri.org/ "><tem:token></tem:token><tem:id>%v</tem:id><tem:identify></tem:identify></tem:getSong_v2></soap:Body>
</soap:Envelope>`, id)
	header := Header
	header.Set("SOAPAction", "http://tempuri.org/getSong_v2")
	ret, err := Post("http://service.keeng.vn/appwebservice/Service.asmx?wsdl", header, dat)
	if err != nil {
		return nil, err
	} else {
		var apiStatusSong APIStatusSong
		data := ret.Data.FindAllStringSubmatch(`<return>(.+)</return>`, -1)[0][1].DecodeHTML()
		err := json.Unmarshal([]byte(data), &apiStatusSong)
		if err != nil {

			return nil, err
		} else {
			apiStatusSong.Data.MainSong.Lyric = apiStatusSong.Data.MainSong.Lyric.DecodeHTML()
			return &apiStatusSong.Data, nil
		}
	}
}

// GetAPIArtistEntry returns a song and relevant songs or an error from API using POST method.
//
// The SOAP data has the following format:
//
// 	<?xml version="1.0" encoding="utf-8"?>
// 	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
// 	  <soap:Header/>
// 	  <soap:Body>
// 	    <tem:getUserProfile_Singer xmlns=" http://tempuri.org/ ">
// 	      <tem:singer_id>1394</tem:singer_id>
// 	    </tem:getUserProfile_Singer>
// 	  </soap:Body>
// 	</soap:Envelope>
func GetAPIArtistEntry(id dna.Int) (*APIArtist, error) {
	if id == 0 {
		return nil, errors.New("Id is zero")
	}
	var dat dna.String = dna.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/"><soap:Header/><soap:Body>
<tem:getUserProfile_Singer xmlns=" http://tempuri.org/ "><tem:singer_id>%v</tem:singer_id></tem:getUserProfile_Singer></soap:Body>
</soap:Envelope>`, id)
	header := Header
	header.Set("SOAPAction", "http://tempuri.org/getUserProfile_Singer")
	ret, err := Post("http://service.keeng.vn/appwebservice/Service.asmx?wsdl", header, dat)
	if err != nil {
		return nil, err
	} else {
		var apiStatusArtist APIStatusArtist
		data := ret.Data.FindAllStringSubmatch(`<return>(.+)</return>`, -1)[0][1].DecodeHTML()
		err := json.Unmarshal([]byte(data), &apiStatusArtist)
		if err != nil {
			return nil, err
		} else {
			return &apiStatusArtist.Data, nil
		}
	}
}

// GetAPIArtistSongs returns a list of songs of an artist or an error from API using POST method.
//
// The SOAP data has the following format:
//
// 	<?xml version="1.0" encoding="utf-8"?>
// 	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
// 	  <soap:Header/>
// 	  <soap:Body>
// 	    <tem:getSinger_Detail_moi xmlns=" http://tempuri.org/ ">
// 	      <tem:token></tem:token>
// 	      <tem:singerid>1394</tem:singerid>
// 	      <tem:page>1</tem:page>
// 	      <tem:num>10</tem:num>
// 	      <tem:type>1</tem:type>
// 	    </tem:getSinger_Detail_moi>
// 	  </soap:Body>
// 	</soap:Envelope>
//
// The params:
//
// 	*singerid : an artist id
// 	*type : 1
// 	*page : current page for pagination
// 	*num : the number of items per page
// 	*token : empty
func GetAPIArtistSongs(id, page, num dna.Int) (*APIArtistSongs, error) {
	if id == 0 {
		return nil, errors.New("Id is zero")
	}
	var dat dna.String = dna.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/"><soap:Header/><soap:Body>
<tem:getSinger_Detail_moi xmlns=" http://tempuri.org/ "><tem:token></tem:token><tem:singerid>%v</tem:singerid><tem:page>%v</tem:page><tem:num>%v</tem:num><tem:type>1</tem:type></tem:getSinger_Detail_moi></soap:Body>
</soap:Envelope>`, id, page, num)
	header := Header
	header.Set("SOAPAction", "http://tempuri.org/getSinger_Detail_moi")
	ret, err := Post("http://service.keeng.vn/appwebservice/Service.asmx?wsdl", header, dat)
	if err != nil {
		return nil, err
	} else {
		var apiArtistSongs APIArtistSongs
		data := ret.Data.FindAllStringSubmatch(`<return>(.+)</return>`, -1)[0][1].DecodeHTML()
		// dna.Log(data)
		err := json.Unmarshal([]byte(data), &apiArtistSongs)
		if err != nil {
			return nil, err
		} else {
			return &apiArtistSongs, nil
		}
	}
}

// GetAPIArtistAlbums returns a list of albums of an artist or an error from API using POST method.
//
// The SOAP data has the following format:
//
// 	<?xml version="1.0" encoding="utf-8"?>
// 	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
// 	  <soap:Header/>
// 	  <soap:Body>
// 	    <tem:getSinger_Detail_moi xmlns=" http://tempuri.org/ ">
// 	      <tem:token></tem:token>
// 	      <tem:singerid>1394</tem:singerid>
// 	      <tem:page>1</tem:page>
// 	      <tem:num>10</tem:num>
// 	      <tem:type>2</tem:type>
// 	    </tem:getSinger_Detail_moi>
// 	  </soap:Body>
// 	</soap:Envelope>
//
// The params:
//
// 	*singerid : an artist id
// 	*type : 2
// 	*page : current page for pagination
// 	*num : the number of items per page
// 	*token : empty
func GetAPIArtistAlbums(id, page, num dna.Int) (*APIArtistAlbums, error) {
	if id == 0 {
		return nil, errors.New("Id is zero")
	}
	var dat dna.String = dna.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/"><soap:Header/><soap:Body>
<tem:getSinger_Detail_moi xmlns=" http://tempuri.org/ "><tem:token></tem:token><tem:singerid>%v</tem:singerid><tem:page>%v</tem:page><tem:num>%v</tem:num><tem:type>2</tem:type></tem:getSinger_Detail_moi></soap:Body>
</soap:Envelope>`, id, page, num)
	header := Header
	header.Set("SOAPAction", "http://tempuri.org/getSinger_Detail_moi")
	ret, err := Post("http://service.keeng.vn/appwebservice/Service.asmx?wsdl", header, dat)
	if err != nil {
		return nil, err
	} else {
		var apiArtistAlbums APIArtistAlbums
		data := ret.Data.FindAllStringSubmatch(`<return>(.+)</return>`, -1)[0][1].DecodeHTML()
		// dna.Log(data)
		err := json.Unmarshal([]byte(data), &apiArtistAlbums)
		if err != nil {
			return nil, err
		} else {
			return &apiArtistAlbums, nil
		}
	}
}

// GetAPIArtistVideos returns a list of albums of an artist or an error from API using POST method.
//
// The SOAP data has the following format:
//
// 	<?xml version="1.0" encoding="utf-8"?>
// 	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
// 	  <soap:Header/>
// 	  <soap:Body>
// 	    <tem:getSinger_Detail_moi xmlns=" http://tempuri.org/ ">
// 	      <tem:token></tem:token>
// 	      <tem:singerid>1394</tem:singerid>
// 	      <tem:page>1</tem:page>
// 	      <tem:num>10</tem:num>
// 	      <tem:type>3</tem:type>
// 	    </tem:getSinger_Detail_moi>
// 	  </soap:Body>
// 	</soap:Envelope>
//
// The params:
//
// 	*singerid : an artist id
// 	*type : 3
// 	*page : current page for pagination
// 	*num : the number of items per page
// 	*token : empty
func GetAPIArtistVideos(id, page, num dna.Int) (*APIArtistVideos, error) {
	if id == 0 {
		return nil, errors.New("Id is zero")
	}
	var dat dna.String = dna.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/"><soap:Header/><soap:Body>
<tem:getSinger_Detail_moi xmlns=" http://tempuri.org/ "><tem:token></tem:token><tem:singerid>%v</tem:singerid><tem:page>%v</tem:page><tem:num>%v</tem:num><tem:type>3</tem:type></tem:getSinger_Detail_moi></soap:Body>
</soap:Envelope>`, id, page, num)
	header := Header
	header.Set("SOAPAction", "http://tempuri.org/getSinger_Detail_moi")
	ret, err := Post("http://service.keeng.vn/appwebservice/Service.asmx?wsdl", header, dat)
	if err != nil {
		return nil, err
	} else {
		var apiArtistVideos APIArtistVideos
		data := ret.Data.FindAllStringSubmatch(`<return>(.+)</return>`, -1)[0][1].DecodeHTML()
		// dna.Log(data)
		err := json.Unmarshal([]byte(data), &apiArtistVideos)
		if err != nil {
			return nil, err
		} else {
			return &apiArtistVideos, nil
		}
	}
}

// GetAPIVideo returns an album or an error from API using POST method.
//
// The SOAP data has the following format:
//
// 	<?xml version="1.0" encoding="utf-8"?>
// 	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
// 	  <soap:Header/>
// 	  <soap:Body>
// 	    <tem:getVideo_v2 xmlns=" http://tempuri.org/ ">
// 	      <tem:token></tem:token>
// 	      <tem:id>86682</tem:id>
// 	      <tem:identify></tem:identify>
// 	    </tem:getVideo_v2>
// 	  </soap:Body>
// 	</soap:Envelope>
func GetAPIVideo(id dna.Int) (*APIVideo, error) {
	if id == 0 {
		return nil, errors.New("Id is zero")
	}
	var dat dna.String = dna.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/"><soap:Header/><soap:Body>
<tem:getVideo_v2 xmlns=" http://tempuri.org/ "><tem:token></tem:token><tem:id>%v</tem:id><tem:identify></tem:identify></tem:getVideo_v2></soap:Body>
</soap:Envelope>`, id)
	header := Header
	header.Set("SOAPAction", "http://tempuri.org/getVideo_v2")
	ret, err := Post("http://service.keeng.vn/appwebservice/Service.asmx?wsdl", header, dat)
	if err != nil {
		return nil, err
	} else {
		var apiStatusVideo APIStatusVideo
		dataArr := ret.Data.FindAllStringSubmatch(`<return>(.+)</return>`, -1)
		if len(dataArr) > 0 {
			// dna.Log(data)
			err := json.Unmarshal([]byte(dataArr[0][1].DecodeHTML()), &apiStatusVideo)
			if err != nil {
				return nil, err
			} else {
				return &apiStatusVideo.Data, nil
			}
		} else {
			return nil, errors.New("No return value")
		}

	}
}
