package zi

import (
	"dna"
	"time"
)

func ExampleGetUsers() {
	users, err := GetUsers(dna.IntArray{438})
	if err != nil {
		dna.PanicError(err)
	} else {
		user := users.List[0]
		user.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
		if user.FriendTotal < 694 {
			panic("Wrong user's FriendTotal")
		}
		if user.Point < 22932 {
			panic("Wrong user's Point")
		}
		if user.VipTotal < 310 {
			panic("Wrong user's VipTotal")
		}

		if user.FriendAvatar == "" {
			panic("Wrong user's FriendAvatar")
		}

		if user.Avatar == "" {
			panic("Wrong user's Avatar")
		}

		if user.VipAvatar == "" {
			panic("Wrong user's VipAvatar")
		}

		if user.CoverUrl == "" {
			panic("Wrong user's CoverUrl")
		}

		user.FriendTotal = 694
		user.Point = 22932
		user.VipTotal = 310
		user.FriendAvatar = "http://s120.avatar.zdn.vn/6/7/8/a/dzobanbe_100_131.jpg"
		user.Avatar = "http://s120.avatar.zdn.vn/6/7/8/a/dzobanbe_100_131.jpg"
		user.VipAvatar = "http://s120.avatar.zdn.vn/9/4/3/c/macma2005_100_262.jpg"
		user.CoverUrl = "http://stc.thongdiephay.zdn.vn/bd54328839-9927.jpg"

		dna.LogStruct(user)
	}
	// Output:
	// Id : 438
	// Email : "dzobanbe@yahoo.vn"
	// Username : "dzobanbe"
	// DisplayName : "Nguyễn Văn Đức Trọng"
	// Birthday : "1981-07-28"
	// BirthdayType : 1
	// Gender : 0
	// Mobile : dna.StringArray{"84937388269"}
	// FriendTotal : 694
	// FriendBlock : "false"
	// FriendAvatar : "http://s120.avatar.zdn.vn/6/7/8/a/dzobanbe_100_131.jpg"
	// ProfilePoint : 0
	// StatusWall : "Tình yêu của anh đối với em giống như giá xăng vậy càng lúc càng tăng."
	// GoogleId : "dzobanbe"
	// YahooId : "dzobanbe"
	// Point : 22932
	// Avatar : "http://s120.avatar.zdn.vn/6/7/8/a/dzobanbe_100_131.jpg"
	// VipTotal : 310
	// VipBlock : "false"
	// VipAvatar : "http://s120.avatar.zdn.vn/9/4/3/c/macma2005_100_262.jpg"
	// IsNull : false
	// Status : 0
	// FeedWriteWallAll : false
	// FeedViewWallAll : true
	// CoverUrl : "http://stc.thongdiephay.zdn.vn/bd54328839-9927.jpg"
	// Checktime : "2013-11-21 00:00:00"
}

func ExampleGetUsersWithInitialId() {
	UsersMagnitude = 10
	users, err := GetUsersWithInitialId(1)
	if err != nil {
		dna.PanicError(err)
	} else {
		if len(users.List) != 9 {
			dna.Log("Complete")
		} else {
			dna.Log("Not complete")
		}
	}
	// Output:
	// Complete

}
