package postgres

type ImageModel struct {
	ID			string	`gorm:"type:uuid;primaryKey"`
	ProfileID	string	`gorm:"type:uuid;not null"`
	AvatarUrl	string	
	Bio			string	
	Rating		float32
	Tags		[]string
	Reviews		[]ReviewModel	`gorm:"foreignKey:ImageID;references:ID"`
}

type ReviewModel struct {
	ID		string	`gorm:"type:uuid;primaryKey"`
	ImageID	string	`gorm:"type:uuid;not null"`
	UserID	string	`gorm:"type:uuid;not null"`
	Text	string
}