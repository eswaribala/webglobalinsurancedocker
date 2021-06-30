package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/webglobalinsurancedocker/domain"
	"github.com/webglobalinsurancedocker/stores"
	"net/http"
)

func CreatePolicyHolderHandler(w http.ResponseWriter, r *http.Request) {
	var policyHolder domain.PolicyHolder
	json.NewDecoder(r.Body).Decode(&policyHolder)
	stores.CreatePolicyHolder(policyHolder)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(policyHolder)
}

func ReadAllPolicyHoldersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	policyHolders, _ := stores.ReadAllPolicyHolders()
	json.NewEncoder(w).Encode(policyHolders)
}
func GetPolicyHolderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	aadharCardNo := params["aadhaarCardNo"]
	fmt.Println(aadharCardNo)
	policyHolder, _ := stores.GetPolicyHolder(aadharCardNo)
	json.NewEncoder(w).Encode(policyHolder)
}
func UpdatePolicyHolderHandler(w http.ResponseWriter, r *http.Request) {
	var updatedPolicyHolder domain.PolicyHolder
	json.NewDecoder(r.Body).Decode(&updatedPolicyHolder)
	policyHolder := stores.UpdatePolicyHolder(updatedPolicyHolder)
	json.NewEncoder(w).Encode(policyHolder)
}
func DeletePolicyHolderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	aadharCardNo := params["aadhaarCardNo"]
	stores.DeletePolicyHolder(aadharCardNo)
	fmt.Fprintf(w, "Record deleted")
}
