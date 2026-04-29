


type Mongo struct {
	Client &mongo.Client
	DB     *mongo.Database
}

funv Connect(ctx context.Context, cfg config.Config) (*Mongo, error) {
	
}