package trip_service

type User struct {
	Id      int
	friends []User
	trips   []Trip
}

func (u User) GetFriends() []User {
	return u.friends
}

func (u *User) AddFriend(friend User) {
	u.friends = append(u.friends, friend)
}

func (u *User) AddTrip(trip Trip) {
	u.trips = append(u.trips, trip)
}

func (u User) Trips() []Trip {
	return u.trips
}

type UserSession struct {
}

func (UserSession) GetInstance() Session {
	return Session{}
}

type Session struct{}

func (Session) GetLoggedUser() (User, error) {
	panic(CollaboratorCallError{"UserSession.getLoggedUser() should not be called in an unit test"})
}
