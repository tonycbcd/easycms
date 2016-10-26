// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// The models.

package models

type User struct {
    BaseModel
    Name string
}

func (User) TableName() string {
    return "user_passport"
}


