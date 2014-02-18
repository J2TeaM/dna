/*
nhac.vui.vn.

To get a song or a video, basically we have to fetch 2 link.Url formats of media link is the same. For example:

	http://hcm.nhac.vui.vn/google-bot-m472092c2p1a1.html
	http://hcm.nhac.vui.vn/asx2.php?type=1&id=472092

We assume link formats described above are only applied to a song.
Then we get all fields relating to Song type. If a field Type of the song is "video",
then the song's id is only for video and is added to FoundVideos.

We get all videos through FoundVideos variable.

The steps above is slightly different from the steps described in csn package.
But for the sake of simplicity, we ignore a general type such as SongVideo to
represent both Song and Video type.

*/
package nv
