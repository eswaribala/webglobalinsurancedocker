package stores

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/webglobalinsurancedocker/domain"
)

var db *gorm.DB

func ConnectionHelper() {

	var err error
	dataSourceName := "root:password@tcp(mysql:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	//db.Exec("CREATE DATABASE amexdb")
	db.Exec("USE amexdb")

	// Migration to create tables for Order and Item schema
	//db.AutoMigrate(&domain.PolicyHolder{}, &domain.Policy{}, &domain.Vehicle{})

	db.AutoMigrate(&domain.PolicyHolder{}, &domain.Policy{})
}

func CreatePolicyHolder(policyHolder domain.PolicyHolder) domain.PolicyHolder {

	db.Create(policyHolder)
	return policyHolder
}

//Read all the PolicyHolders

func ReadAllPolicyHolders() ([]*domain.PolicyHolder, error) {
	var policyHolders []*domain.PolicyHolder
	db.Preload("Policies").Find(&policyHolders)
	return policyHolders, nil
}

//PolicyHolder By Id

func GetPolicyHolder(aadharCardNo string) (*domain.PolicyHolder, error) {

	var policyHolder domain.PolicyHolder
	//db.First(&policyHolder, aadharCardNo)
	if result := db.Where("aadhaar_card_no = ?", aadharCardNo).Preload("Policies").First(&policyHolder); result.Error != nil {
		fmt.Println(result.Error)
	}

	return &policyHolder, nil
}

//Update PolicyHolder
func UpdatePolicyHolder(updatedPolicyHolder domain.PolicyHolder) domain.PolicyHolder {

	db.Save(&updatedPolicyHolder)
	return updatedPolicyHolder

}

//Delete PolicyHolder
func DeletePolicyHolder(aadharCardNo string) {

	db.Where("aadhaar_card_no = ?", aadharCardNo).Delete(&domain.PolicyHolder{})

}
