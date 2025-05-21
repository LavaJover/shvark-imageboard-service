package domain

type Image struct {
	ID			string	
	ProfileID	string	
	AvatarUrl	string	
	Bio			string	
	Rating		float32
	Tags		[]string
	Reviews		[]Review
}

type Review struct {
	ID		string
	ImageID	string	
	UserID	string	
	Text	string
}