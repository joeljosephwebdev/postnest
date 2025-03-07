package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("missing login credentials")
	}
	username := cmd.Args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User switched successfully!")

	return nil
}
