package main

import (
	"context"
	"github.com/tithanayut/learnhub-ql/src/config"
	"github.com/tithanayut/learnhub-ql/src/database"
	"log"
)

func main() {
	ctx := context.Background()

	appConfig := config.LoadAppConfig()
	db := database.NewDatabaseClient(appConfig.Database)

	user, err := db.User.Create().SetUsername("test").SetPassword("hashme").SetName("Test User").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Println("user was created: ", user)

	video, err := db.Video.Create().SetVideoTitle("Kelly Clarkson - Stronger (What Doesn't Kill You) [Official Video]").
		SetVideoURL("https://www.youtube.com/watch?v=Xn676-fLq7I").
		SetVideoThumbnail("https://i.ytimg.com/vi/Xn676-fLq7I/hqdefault.jpg").
		SetCreatorName("Kelly Clarkson").
		SetCreatorURL("https://www.youtube.com/@kellyclarksonVEVO").
		SetPostedByID(user.ID).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed adding video: %v", err)
	}
	log.Println("video was added: ", video)

	video, err = db.Video.Create().SetVideoTitle("CAN'T STOP THE FEELING! (from DreamWorks Animation's \\\"TROLLS\\\") (Official Video)").
		SetVideoURL("https://www.youtube.com/watch?v=ru0K8uYEZWw").
		SetVideoThumbnail("https://i.ytimg.com/vi/ru0K8uYEZWw/hqdefault.jpg").
		SetCreatorName("Justin Timberlake").
		SetCreatorURL("https://www.youtube.com/@justintimberlakeVEVO").
		SetPostedByID(user.ID).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed adding video: %v", err)
	}
	log.Println("video was added: ", video)
}
