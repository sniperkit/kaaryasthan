package item

import (
	"testing"

	"github.com/kaaryasthan/kaaryasthan/db"
	"github.com/kaaryasthan/kaaryasthan/project"
	"github.com/kaaryasthan/kaaryasthan/user"
)

func TestItemCreate(t *testing.T) {
	defer db.DB.Exec("DELETE FROM users")
	defer db.DB.Exec("DELETE FROM projects")
	defer db.DB.Exec("DELETE FROM items")
	defer db.DB.Exec("DELETE FROM item_discussion_comment_search")

	usr := user.User{Username: "jack", Name: "Jack Wilber", Email: "jack@example.com", Password: "Secret@123"}
	if err := usr.Create(); err != nil {
		t.Error(err)
	}

	prj := project.Project{Name: "somename", Description: "Some description"}
	if err := prj.Create(usr); err != nil {
		t.Error(err)
	}

	itm := Item{Title: "sometitle", Description: "Some description", ProjectID: prj.ID}
	err := itm.Create(usr)
	if err != nil {
		t.Error(err)
	}
	if itm.ID <= 0 {
		t.Errorf("Data not inserted. ID: %#v", itm.ID)
	}
	if itm.Num != 1 {
		t.Errorf("Data not inserted. Num: %#v", itm.Num)
	}
}