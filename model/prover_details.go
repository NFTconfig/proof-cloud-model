package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProverDetails struct {
	mgm.IDField `json:",inline" bson:",inline"`
	ProverId    primitive.ObjectID `json:"proverId" bson:"proverId"`
	Account     string             `json:"account" bson:"account"`
	ProverName  string             `json:"proverName" bson:"proverName"`
	Input       string             `json:"input" bson:"input"`
	OutputPath  string             `json:"outputPath" bson:"outputPath"`
	Status      int                `json:"status" bson:"status"` //0 is pending ;1 is running; 2 is completed ;3 is cancel,4 is deleted
	CreateTime  int64              `json:"createTime" bson:"createTime"`
	InputHash   string             `json:"inputHash" bson:"inputHash"`
}

func (p *ProverDetails) CollectionName() string {
	return "prover_details"
}

func (p *ProverDetails) Validate() error {
	return nil
}
