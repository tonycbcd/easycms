// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// The Api Controllers.

package controllers

import (
    "github.com/tonycbcd/easycms/server/components/egin"
)

func ApiGet(c *egin.Context) {
    c.JSON(200, egin.H{"Hello Get": "Api", "test":"123"})
}

func ApiSet(c *egin.Context) {
    c.JSON(200, egin.H{"Hello Set": "Api", "test":"123"})
}

func ApiDelete(c *egin.Context) {
    c.JSON(200, egin.H{"Hello Delete": "Api", "test":"123"})
}

func ApiOption(c *egin.Context) {
    c.JSON(200, egin.H{"Hello Option": "Api", "test":"123"})
}


