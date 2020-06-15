package activity_test

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/FrancescoIlario/usplay/cmd/services/bookmaster/server/conf"
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage/mongostore"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyz")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(length int) string {
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func getRepo(collection string) (storage.Repository, error) {
	c, err := conf.GetConfigurationFromEnv()
	if err != nil {
		return nil, fmt.Errorf(`configuration error: %w`, err)
	}

	mc := mongostore.Configuration{
		ConnectionString: c.ActivitytypeMongoConnstr,
		Collection:       collection,
		Database:         c.ActivitytypeMongoDatabase,
		Password:         c.ActivitytypeMongoPassword,
		Username:         c.ActivitytypeMongoUsername,
	}

	store, err := mongostore.New(&mc)
	if err != nil {
		return nil, fmt.Errorf(`error connecting to ActivityType repository: %w`, err)
	}
	return store, nil
}
