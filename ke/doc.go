/*
This package implements some functionalities regarding keeng.vn

This package uses 2 methods to get info.

* METHOD 1: Fetching info directly from http website:

It just finds new albums and their songs and store them into DB.

SongFinder to find new albums. AlbumSong is a special type that can
fetch a new album and its songs. Then save the found songs and album.

SongAlbum gets a new album and songs from 2 example following links:

	http://www.keeng.vn/album/google-bot-M4U/CLJH7SVS.html
	http://www.keeng.vn/album/get-album-xml?album_identify=CLJH7SVS

* METHOD 2: Using API:

Get songs or albums or videos of an artist.

	<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
	  <soap:Header/>
	  <soap:Body>
	    <tem:getSinger_Detail_moi xmlns=" http://tempuri.org/ ">
	      <tem:token></tem:token>
	      <tem:singerid>1394</tem:singerid>
	      <tem:page>1</tem:page>
	      <tem:num>10</tem:num>
	      <tem:type>2</tem:type>
	    </tem:getSinger_Detail_moi>
	  </soap:Body>
	</soap:Envelope>

The params:

	*singerid : an artist id
	*type : [1|2|3] <=> ["song","album","video"]
	*page : current page for pagination
	*num : the number of items per page
	*token : empty

*/
package ke
