package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/mmsacky/gator/internal/config"
	"github.com/mmsacky/gator/internal/database"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Critical error opening config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)

	dbQueries := database.New(db)

	myState := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	myCommands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}

	myCommands.register("login", handlerLogin)
	myCommands.register("register", handlerRegister)
	myCommands.register("reset", handlerReset)
	myCommands.register("users", handlerUsers)
	myCommands.register("agg", handlerAggregator)
	myCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	myCommands.register("feeds", handlerFeeds)
	myCommands.register("follow", middlewareLoggedIn(handlerFollow))
	myCommands.register("following", middlewareLoggedIn(handlerFollowing))

	userArgsCount := len(os.Args)

	if userArgsCount < 2 {
		log.Fatalln("Error occurred: not enough arguments")
	}

	myCommand := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = myCommands.run(&myState, myCommand)
	if err != nil {
		log.Fatalln("Error occurred:", err)
	}
}
