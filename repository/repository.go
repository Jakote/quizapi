package repository

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type QuizQuestion struct {
	Category       		string `bson:"category" json:"category"`
	Type 				string `bson:"type" jzon:"type"`
	Difficulty			string `bson:"difficulty" json:"difficulty"`
	Question        	string `bson:"question" json:"question"`
	Correct_answer  	string `bson:"correct_answer" json:"correct_answer"`
	Incorrect_answers 	[]string `bson:"incorrect_answers" json:"incorrect_answers"`
}

type Repository struct {
	dbSession    *mgo.Session
	dbServer     string
	dbDatabase   string
	dbCollection string
}

func NewRepository(dbServer string, dbDatabase string, dbCollection string) *Repository {
	repo := new(Repository)
	repo.dbServer = dbServer
	repo.dbDatabase = dbDatabase
	repo.dbCollection = dbCollection

	dbSession, err := mgo.Dial(repo.dbServer)
	if err != nil {
		log.Fatal(err)
	}
	repo.dbSession = dbSession
	return repo
}

func (repo *Repository) Close() {
	repo.dbSession.Close()
}

func (repo *Repository) newSession() *mgo.Session {
	return repo.dbSession.Clone()
}

func (repo *Repository) FindAll() ([]QuizQuestion, error) {
	dbSession := repo.newSession()
	defer dbSession.Close()

	coll := dbSession.DB(repo.dbDatabase).C(repo.dbCollection)

	var qqzs []QuizQuestion
	err := coll.Find(bson.M{}).All(&qqzs)
	return qqzs, err
}

func (repo *Repository) FindByCategory(ctgry string) (QuizQuestion, error) {
	dbSession := repo.newSession()
	defer dbSession.Close()

	coll := dbSession.DB(repo.dbDatabase).C(repo.dbCollection)

	var qqz QuizQuestion
	err := coll.Find(bson.M{"category": ctgry}).One(&qqz)

	return qqz, err
}

func (repo *Repository) Insert(qqz QuizQuestion) error {
	dbSession := repo.newSession()
	defer dbSession.Close()

	coll := dbSession.DB(repo.dbDatabase).C(repo.dbCollection)

	err := coll.Insert(&qqz)
	return err
}

func (repo *Repository) Delete(qqz QuizQuestion) error {
	dbSession := repo.newSession()
	defer dbSession.Close()

	coll := dbSession.DB(repo.dbDatabase).C(repo.dbCollection)

	err := coll.Remove(bson.M{"category": qqz.Category})
	return err
}

func (repo *Repository) Update(ctgry string, qqz QuizQuestion) error {
	dbSession := repo.newSession()
	defer dbSession.Close()

	coll := dbSession.DB(repo.dbDatabase).C(repo.dbCollection)

	err := coll.Update(bson.M{"category": ctgry}, &qqz)
	return err
}
