package model

import "github.com/kamva/mgm/v3"

type Prover struct {
	mgm.IDField       `json:",inline" bson:",inline"`
	Account           string `json:"account" bson:"account"`
	ProverName        string `json:"proverName" bson:"proverName"`
	ApiKey            string `json:"apiKey" bson:"apiKey"`
	Description       string `json:"description" bson:"description"`
	BinaryPath        string `json:"binaryPath" bson:"binaryPath"`
	DataDirectoryPath string `json:"dataDirectoryPath" bson:"dataDirectoryPath"`
	PkSize            int64  `json:"pkSize" bson:"pkSize"`
	VkSize            int64  `json:"vkSize" bson:"vkSize"`
	R1csSize          int64  `json:"r1csSize" bson:"r1csSize"`
	Visibility        int    `json:"visibility" bson:"visibility"` //0 is public ; 1 is private
	IsActive          int    `json:"isActive" bson:"isActive"`     //0 is no active ;1 is active
	CreateTime        int64  `json:"createTime" bson:"createTime"`
}

func (p *Prover) CollectionName() string {
	return "provers"
}

func (p *Prover) Validate() error {
	return nil
}
