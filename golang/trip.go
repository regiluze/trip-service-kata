package trip_service

import ()

type Trip struct{}

type TripDao struct{}

func (TripDao) FindTripsByUser(user User) []Trip {
	panic(CollaboratorCallError{"TripDAO should not be invoked on an unit test."})
}

type TripService struct {
}

func (TripService) getTripsByUser(user User) ([]Trip, error) {
	tripList := []Trip{}
	loggedUser, err := UserSession{}.GetInstance().GetLoggedUser()
	isFriend := false
	if err != nil {
		for _, friend := range user.GetFriends() {
			if friend.Id == loggedUser.Id {
				isFriend = true
				break
			}
		}
		if isFriend {
			tripList = TripDao{}.FindTripsByUser(user)
		}
		return tripList, nil
	} else {
		return []Trip{}, UserNotLoggedInError{"User not logged"}
	}
}
