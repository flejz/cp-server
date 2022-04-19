package cli

import (
	"fmt"
	"os"

	"github.com/flejz/cp-server/internal/buffer"
	"github.com/flejz/cp-server/internal/db"
	"github.com/flejz/cp-server/internal/repository"
	"github.com/flejz/cp-server/internal/user"
	"github.com/spf13/cobra"
)

var (
	usr        string
	pwd        string
	sqliteFile string
)

var rootCmd = &cobra.Command{
	Use:   "cp-cli",
	Short: "cp-cli gets or sets value to your local or remote cp-server instance",
}

var cmdGet = &cobra.Command{
	Use:   "get [key]",
	Short: "Get the value for a key",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handler := initCmd()
		if err := handler.Login(usr, pwd); err != nil {
			writeErr(err)
		}

		value, err := handler.Get(args)
		if err != nil {
			writeErr(err)
		} else {
			write(value)
		}
	},
}

var cmdSet = &cobra.Command{
	Use:   "set [key] value",
	Short: "Set the value for a key",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handler := initCmd()
		if err := handler.Login(usr, pwd); err != nil {
			writeErr(err)
		}

		err := handler.Set(args)
		if err != nil {
			writeErr(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmdGet)
	rootCmd.AddCommand(cmdSet)
	rootCmd.PersistentFlags().StringVarP(&usr, "user", "u", "", "user")
	rootCmd.PersistentFlags().StringVarP(&pwd, "pass", "p", "", "password")
	rootCmd.PersistentFlags().StringVarP(&sqliteFile, "sqlite", "s", "", "sqlite file")
	rootCmd.MarkPersistentFlagRequired("user")
	rootCmd.MarkPersistentFlagRequired("pass")
}

func initCmd() *buffer.Cmd {
	// init db
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	// init repositorys
	bufferRepository := buffer.NewBufferRepository(db)
	userRepository := user.NewUserRepository(db)

	if err := repository.Init([]repository.Repository{bufferRepository, userRepository}); err != nil {
		panic(err)
	}

	// init models
	bufferService := buffer.BufferService{bufferRepository}
	userService := user.UserService{userRepository}
	return &buffer.Cmd{
		BufferService: bufferService,
		UserService:   userService,
	}
}

func write(value string) {
	fmt.Println(value)
}

func writeErr(err error) {
	fmt.Printf("%v\n", err)
}
