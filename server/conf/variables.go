// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// To initialize the config data.

package conf

type config struct {
	Global struct {
        Mode        string
        Port        string
        Host        string
        TimeZoneOffset  int64
        PublicSalt  string
	}

    Staticcluster []string

	Mysql struct {
		Host        string
		Port        int
		User        string
		Pass        string
		Db          string
	}

/*
	Beanstalk struct {
		Addr  []string
		Proxy string
	}


	Services map[string]struct {
		User   string
		Secret string
		Uri    []string
	}
*/
}

var Config *config = new(config)

var ConfigFile *string
