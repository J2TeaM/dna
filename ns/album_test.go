package ns

import (
	. "dna"
	"testing"
)

//Testing album with fail link
func TestGetAlbum(t *testing.T) {
	_, err := GetAlbum(818243)
	if err == nil {
		t.Error("The album has to have an error")
	} else {
		if err.Error() != "Album 818243: Songids and Nsongs do not match" {
			t.Error("Wrong error message.")
		}
	}
	album, err := GetAlbum(818231)
	if err != nil {
		t.Error("The album has to have no error")
	} else {
		if album.Title != "Salute" {
			t.Error("The album 818231 has to have tittle called Salute")
		}
		if album.Topics.Length() != 1 && album.Topics[0] != "Nhạc Âu Mỹ" {
			t.Error("The album 818231 has wrong topics")
		}
		if !album.Description.Match(`Label: Sony Music Entertainment`) {
			t.Error("The album 818231 has wrong desc")
		}
		if album.Genres.Length() != 1 && album.Genres[0] != "Pop" {
			t.Error("The album 818231 has wrong genres")
		}
		if album.Songids.Length() != 16 && album.Nsongs != 16 {
			t.Error("The album 818231 has Songids and Nsongs not matching")
		}
		if album.Plays < 50 {
			t.Error("The album 818231 has wrong total of plays")
		}
		if album.Coverart != "http://st.nhacso.net/images/album/2013/11/08/1072071738/13838971304_346_120x120.jpg" {
			t.Error("The album 818231 has wrong coverart")
		}
		if album.Artists.Length() != 1 && album.Artists[0] != "Little Mix" && album.Artistid != 15617 {
			t.Error("The album 818231 has wrong artist")
		}
		if album.Label != "Sony Music Entertainment UK Limited" {
			t.Errorf("The album 818231 has wrong label")
		}
		if album.DateReleased != "2013" {
			t.Error("The album 818231 has wrong date released")
		}
	}

	al, err := GetAlbum(815400)
	if err != nil {
		t.Error("The album has to have no error")
	} else {
		if al.Title != "Người Tình Quê" {
			t.Error("The album 818231 has to have tittle called Người Tình Quê")
		}
		if al.Topics.Length() != 1 && al.Topics[0] != "Nhạc Trữ Tình" {
			t.Error("The album 818231 has wrong topics")
		}
		if !al.Description.Match(`Nhân dịp thực hiện liveshow 20`) {
			t.Error("The album 818231 has wrong desc")
		}
		if al.Genres.Length() != 0 {
			t.Error("The album 818231 has wrong genres")
		}
		if al.Songids.Length() != 10 && al.Nsongs != 10 {
			t.Error("The album 818231 has Songids and Nsongs not matching")
		}
		if al.Plays < 18500 {
			t.Error("The album 818231 has wrong total of plays")
		}
		if al.Coverart != "http://st.nhacso.net/images/album/2013/11/04/1236052905/13835294504_3208_120x120.jpg" {
			t.Error("The album 818231 has wrong coverart")
		}
		if al.Artists.Length() != 2 && al.Artists[0] != "Cẩm Ly" && al.Artists[1] != "Quốc Đại" && al.Artistid != 17 {
			t.Error("The album 818231 has wrong artist")
		}
		if al.Label != "" {
			t.Errorf("The album 818231 has wrong label")
		}
		if al.DateReleased != "2013" {
			t.Error("The album 818231 has wrong date released")
		}
	}
}

func ExampleGetAlbum() {
	album, err := GetAlbum(818448)
	if err != nil {
		Log(err.Error())
	} else {
		Log("Id:", album.Id)
		Log("Title:", album.Title)
		Log("Artists:", album.Artists)
		Log("Artistid:", album.Artistid)
		Log("Topics:", album.Topics)
		Log("Genres:", album.Genres)
		Log("Category:", album.Category)
		Log("Coverart:", album.Coverart)
		Log("Nsongs:", album.Nsongs)
		Log("Label:", album.Label)
		// Log("Plays:", album.Plays)
		Log("DateReleased:", album.DateReleased)
		Log("Songids:", album.Songids)
		if album.Plays > 0 {
			Log("Plays > 0")
		}
		Log("Description:")
		Logv(album.Description)

	}
	// Output:
	// Id: 818448
	// Title: France And French Cafe Accordion Music
	// Artists: dna.StringArray{"Bon Appétit Musique"}
	// Artistid: 99574
	// Topics: dna.StringArray{"Nhạc Không Lời"}
	// Genres: dna.StringArray{"French Pop", "Instrumental"}
	// Category: dna.StringArray{}
	// Coverart: http://st.nhacso.net/images/album/2013/11/08/1486073233/138392264619_3426_120x120.jpg
	// Nsongs: 23
	// Label: One Media Publishing
	// DateReleased: 2009
	// Songids: dna.IntArray{1313103, 1313104, 1313105, 1313106, 1313107, 1313108, 1313109, 1313110, 1313111, 1313112, 1313113, 1313114, 1313115, 1313116, 1313117, 1313118, 1313119, 1313120, 1313121, 1313122, 1313123, 1313124, 1313125}
	// Plays > 0
	// Description:
	// "Genres: French Pop, Instrumental \r\nReleased: Aug 17, 2009/ ℗ 2009 One Media Publishing"

}
