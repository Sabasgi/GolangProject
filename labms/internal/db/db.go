package db

type DATABASE interface {
	creteConnection()
}

// type MongoRepo struct {
// 	Client *mongo.Client
// 	DBInfo models.MongoModel
// }

// func NewMongoDBRepo(mm models.MongoModel) *MongoRepo {
// 	clientOptions := options.Client().ApplyURI(mm.DSN)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, Cerror := mongo.Connect(ctx, clientOptions)
// 	if Cerror != nil {
// 		fmt.Println("ERROR : newMongoDB", Cerror)
// 	}
// 	return &MongoRepo{
// 		Client: client,
// 		DBInfo: mm,
// 	}
// }
// func (m *MongoRepo) CloseConnection() error {
// 	return m.Client.Disconnect(context.Background())
// }

//service not needed
// type MongoService struct {
// 	repo MongoRepo
// }

// func NewMongoService(mgo MongoRepo) *MongoService {
// 	return &MongoService{
// 		repo: mgo,
// 	}
// }

// func newMongoDB(db DBInterface) *DBStruct {
// 	return &DBStruct{
// 		DB: db,
// 	}
// }
