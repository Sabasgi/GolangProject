package models

type ProductModel struct {
	Id             string  `json:"id,omitempty" bson:"id,omitempty" db:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Description    string  `json:"description,omitempty" bson:"description,omitempty" db:"description"`
	ImagePaath     string  `json:"imagePath,omitempty" bson:"imagePath,omitempty" db:"imagePath"`
	Category       string  `json:"category,omitempty" bson:"category,omitempty" db:"category"`
	InventoryStock string  `json:"inventoryStock,omitempty" bson:"inventoryStock,omitempty" db:"inventoryStock"`
}

type DBInfo struct {
	DSN      string
	Database string
	Name     string
}
type DataSource struct {
	MongoDSN         string
	SQLDSN           string
	Port             string
	MongoDBName      string
	SQLDBName        string
	Collections      map[string]string
	UsersEventsTopic string
	KafkaBroker      string
}

type UserModel struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty" db:"name"`
	Age      int    `json:"age,omitempty" bson:"age,omitempty" db:"age"`
	EmailId  string `json:"emailId,omitempty" bson:"emailId,omitempty" db:"emailId"`
	Username string `json:"username,omitempty" bson:"username,omitempty" db:"username"`
	Password string `json:"password,omitempty" bson:"password,omitempty" db:"password"`
	UserId   string `json:"userId,omitempty" bson:"userId,omitempty" db:"userId"`
	Id       string `json:"id,omitempty" bson:"id,omitempty" db:"id,omitempty"`
}
