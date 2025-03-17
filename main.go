package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joeljosephwebdev/postnest.git/internal/config"
	"github.com/joeljosephwebdev/postnest.git/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	dbQueries := *database.New(db)

	programState := &state{
		db:  &dbQueries,
		cfg: &cfg,
	}

	cmds := commandsList{
		commands: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", agg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFeedFollow))
	cmds.register("following", middlewareLoggedIn(showFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
