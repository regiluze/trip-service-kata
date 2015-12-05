package trip_service

func getTripsByUser(user User) ([]Trip, error) {
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
