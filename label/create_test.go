package label

import (
	"testing"

	"github.com/kaaryasthan/kaaryasthan/db"
	"github.com/kaaryasthan/kaaryasthan/project"
	"github.com/kaaryasthan/kaaryasthan/user"
)

func TestLabelCreate(t *testing.T) {
	defer db.DB.Exec("DELETE FROM users")
	defer db.DB.Exec("DELETE FROM projects")
	defer db.DB.Exec("DELETE FROM labels")

	usr := user.User{Username: "jack", Name: "Jack Wilber", Email: "jack@example.com", Password: "Secret@123"}
	if err := usr.Create(); err != nil {
		t.Error(err)
	}

	prj := project.Project{Name: "somename", Description: "Some description"}
	if err := prj.Create(usr); err != nil {
		t.Error(err)
	}

	lbl := Label{Name: "somename", Color: "#ee0701", ProjectID: prj.ID}
	if err := lbl.Create(usr); err != nil {
		t.Error(err)
	}
	if lbl.ID <= 0 {
		t.Errorf("Data not inserted. ID: %#v", lbl.ID)
	}
}
