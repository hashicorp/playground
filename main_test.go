package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// update name to an empty string
	if err := DB.Model(&User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"Name": "",
	}).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// now, we will search again for the user reusing the previous result
	// variable
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// the user name from the db should be an empty string, but it is still
	// "jinzhu" from the previous search.... the empty string isn't being set in
	// the model.
	if result.Name != "" {
		t.Errorf("Failed, name should be %q but it's %q", "", result.Name)
	}

	// Let's prove that the name is actually an empty string in the db by
	// searching again with an new variable for the model
	var newResult User
	if err := DB.First(&newResult, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if user.Name != "" {
		t.Errorf("name should be set to %q but it's now: %q", "", user.Name)
	}

}
