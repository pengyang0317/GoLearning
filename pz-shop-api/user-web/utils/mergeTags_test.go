package utils

import (
	"fmt"
	"testing"
)

func TestMergeTags(t *testing.T) {

	var user User
	var profile UserProfile

	MergeTags(&user, &profile)

	fmt.Println(user.ID, user.Name, user.Password) // 输出: 0  user_name

}
