package zi

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"math/rand"
	"os"
	"sync"
	"time"
)

var (
	// UsersMagnitude defines the number of users after each init() called
	UsersMagnitude = 20
)

// User defines new User struct.
type User struct {
	Id               dna.Int
	Email            dna.String
	Username         dna.String
	DisplayName      dna.String
	Birthday         dna.String
	BirthdayType     dna.Int
	Gender           dna.Int
	Mobile           dna.StringArray
	FriendTotal      dna.Int
	FriendBlock      dna.String
	FriendAvatar     dna.String
	ProfilePoint     dna.Int
	StatusWall       dna.String
	GoogleId         dna.String
	YahooId          dna.String
	Point            dna.Int
	Avatar           dna.String
	VipTotal         dna.Int
	VipBlock         dna.String
	VipAvatar        dna.String
	IsNull           dna.Bool
	Status           dna.Int
	FeedWriteWallAll dna.Bool
	FeedViewWallAll  dna.Bool
	CoverUrl         dna.String
	Checktime        time.Time
}

// NewUser return new User instance.
func NewUser() *User {
	user := new(User)
	user.Id = 0
	user.Email = ""
	user.Mobile = dna.StringArray{}
	user.ProfilePoint = 0
	user.StatusWall = ""
	user.DisplayName = ""
	user.Point = 0
	user.Username = ""
	user.Avatar = ""
	user.BirthdayType = 0
	user.VipTotal = 0
	user.VipBlock = ""
	user.VipAvatar = ""
	user.IsNull = false
	user.Gender = 0
	user.Status = 0
	user.Birthday = ""
	user.FeedWriteWallAll = false
	user.FeedViewWallAll = false
	user.CoverUrl = ""
	user.GoogleId = ""
	user.YahooId = ""
	user.FriendTotal = 0
	user.FriendBlock = ""
	user.FriendAvatar = ""
	user.Checktime = time.Now()
	return user
}

// Fill writes all fields found APIUser instance to a User variable.
func (user *User) Fill(apiUser *APIUser) {
	user.Id = apiUser.Id
	user.Email = apiUser.Email
	user.Mobile = dna.StringArray(apiUser.Mobile)
	user.ProfilePoint = apiUser.ProfilePoint
	user.StatusWall = apiUser.StatusWall
	user.DisplayName = apiUser.DisplayName.Trim()
	user.Point = apiUser.Point
	user.Username = apiUser.Username
	user.Avatar = apiUser.Avatar
	user.BirthdayType = apiUser.BirthdayType
	user.VipTotal = apiUser.Vip.Total
	user.VipBlock = apiUser.Vip.Block
	user.VipAvatar = apiUser.Vip.AvatarVip
	user.IsNull = apiUser.Benull
	user.Gender = apiUser.Gender
	user.Status = apiUser.Status
	user.Birthday = apiUser.Birthday
	user.FeedWriteWallAll = apiUser.Feed.WriteWallAll
	user.FeedViewWallAll = apiUser.Feed.ViewWallAll
	user.CoverUrl = apiUser.CoverUrl
	user.GoogleId = apiUser.GoogleId
	user.YahooId = apiUser.YahooId
	user.FriendTotal = apiUser.Friend.Total
	user.FriendBlock = apiUser.Friend.Block
	user.FriendAvatar = apiUser.Avatar
}

// IsValid checks if a user is legitimate. Currently, its result is based on the Email and Username field.
func (user *User) IsValid() dna.Bool {
	if user.Email == "" && user.Username == "" {
		return false
	} else {
		return true
	}
}

type Users struct {
	InitialId dna.Int
	List      []*User
}

func NewUsers() *Users {
	return &Users{0, nil}
}

// GetUsers returs list of users or an error.
func GetUsers(ids dna.IntArray) (*Users, error) {
	apiUsers, err := GetAPIUsers(ids)
	if err != nil {
		return nil, err
	} else {
		var users []*User
		for _, apiUser := range apiUsers {
			user := NewUser()
			user.Fill(&apiUser)
			users = append(users, user)
		}
		return &Users{ids[0], users}, nil
	}
}

func GetUsersWithInitialId(initialId dna.Int) (*Users, error) {
	rids := rand.Perm(UsersMagnitude)
	ids := dna.NewIntArray(rids)
	ids = dna.IntArray(ids.Map(func(val dna.Int, idx dna.Int) dna.Int {
		return val + initialId*dna.Int(UsersMagnitude)
	}).([]dna.Int))
	// dna.Log(ids)
	apiUsers, err := GetAPIUsers(ids)
	if err != nil {
		return nil, err
	} else {
		var users []*User
		for _, apiUser := range apiUsers {
			user := NewUser()
			user.Fill(&apiUser)
			users = append(users, user)
		}
		return &Users{initialId, users}, nil
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (users *Users) Fetch() error {
	_users, err := GetUsersWithInitialId(users.InitialId)
	if err != nil {
		return err
	} else {
		*users = *_users
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (users *Users) GetId() dna.Int {
	return users.InitialId
}

// New implements item.Item interface
// Returns new item.Item interface
func (users *Users) New() item.Item {
	return item.Item(NewUsers())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (users *Users) Init(v interface{}) {
	switch v.(type) {
	case int:
		users.InitialId = dna.Int(v.(int))
	case dna.Int:
		users.InitialId = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

var (
	mutex      = &sync.Mutex{}
	SongFile   *os.File
	TotalBytes dna.Int
)

func (users *Users) Save(db *sqlpg.DB) error {
	var queries dna.String = ""
	// var err error
	// for _, user := range users.List {
	// 	if user.IsValid() == true {
	// 		err = db.InsertIgnore(user)
	// 	}
	// }
	// return err

	for _, user := range users.List {
		if user.IsValid() == true {
			queries += sqlpg.GetInsertStatement("ziusers", user, false) + "\n"
		}
	}
	mutex.Lock()
	n, err := SongFile.WriteString(queries.String())
	if err != nil {
		dna.Log("Cannot write to file while getting song", err.Error(), "\n\n")
	}
	TotalBytes += dna.Int(n)
	mutex.Unlock()
	return nil
}
