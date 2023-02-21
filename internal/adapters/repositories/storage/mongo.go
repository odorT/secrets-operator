package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"secrets-operator/config"
	"secrets-operator/internal/core/domain"
	"secrets-operator/internal/errors"
	_ "secrets-operator/internal/errors"
	"strconv"
	"time"
)

type mongoDB struct {
	cfg    *config.Config
	l      *zap.SugaredLogger
	client *mongo.Client
}

func NewMongoDb(cfg *config.Config, l *zap.SugaredLogger) *mongoDB {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		l.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		l.Fatalln(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		l.Fatalln(err)
	}

	l.Infoln("Connected to MongoDB at", cfg.MongoHost)
	return &mongoDB{
		cfg:    cfg,
		l:      l,
		client: client,
	}
}

func (db *mongoDB) SaveFindingsReport(findingsReport domain.FindingsReport, collectionName string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.client.Database(db.cfg.MongoDBName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, findingsReport)
	if err != nil {
		return err
	}

	return nil
}

func (db *mongoDB) GetRepoFindingsById(repoId int, collectionName string) (domain.RepoFindings, error) {

	repoFindings := domain.RepoFindings{}

	filter := bson.D{{
		"repoid", repoId,
	}}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.client.Database(db.cfg.MongoDBName).Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(&repoFindings)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return domain.RepoFindings{}, errors.ErrRepositoryNotFound
		default:
			return domain.RepoFindings{}, err
		}
	}

	return repoFindings, nil
}

// SaveAndUpdateRepoFindingsById TODO: should be optimized, unnecessary variable definitions should be moved
func (db *mongoDB) SaveAndUpdateRepoFindingsById(repoFindings domain.RepoFindings, repoId int, collectionName string) error {

	repoFindingsOld := domain.RepoFindings{}

	// updating with update operators. Please check following link for more details
	// https://www.mongodb.com/docs/manual/reference/operator/update/addToSet/#-each-modifier
	repoFindingsUpdate := bson.D{{
		"$addToSet",
		bson.D{{
			"findings",
			bson.D{{
				"$each",
				repoFindings.Findings,
			}},
		}},
	}}

	filter := bson.D{{
		Key:   "repoid",
		Value: repoId,
	}}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.client.Database(db.cfg.MongoDBName).Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(&repoFindingsOld)
	if err != nil {
		switch err {
		// if this repository does not have any findings
		case mongo.ErrNoDocuments:
			_, insertErr := collection.InsertOne(ctx, repoFindings)
			if insertErr != nil {
				return err
			}
		default:
			return err
		}
	}

	_, insertErr := collection.UpdateOne(ctx, filter, repoFindingsUpdate)
	if insertErr != nil {
		return err
	}

	return nil
}

func (db *mongoDB) GetRepositoriesByName(repoName string, collectionName string) ([]map[string]string, error) {

	filter := bson.M{
		"reponame": bson.M{
			"$regex":   repoName,
			"$options": "im",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.client.Database(db.cfg.MongoDBName).Collection(collectionName)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var documents []domain.RepoFindings
	if err = cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	var response []map[string]string

	for _, document := range documents {
		id := strconv.Itoa(document.RepoID)
		response = append(response, map[string]string{"name": document.RepoName, "id": id})
	}

	return response, nil
}
