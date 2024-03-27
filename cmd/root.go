package cmd

import (
	"fmt"
	"gurl/internal/formatter"
	"gurl/internal/httpclient"
	"gurl/internal/parser"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gurl <url>",
	Short: "A Golang curl command line tool",
	Long:  `A Golang curl command line tool that can be used to make HTTP requests to a server.`,
	Args:  validateArgs,
	Example: `
gurl http://eu.httpbin.org/get
gurl http://eu.httpbin.org/bearer -H 'Authorization: Bearer guineapig'
gurl http://eu.httpbin.org/post -X POST -d '{"name": "Robert J. Oppenheimer"}' -H "Content-Type: application/json"
gurl http://eu.httpbin.org/put -X PUT -d '{"name": "Ludwig Wittgenstein"}' -H "Content-Type: application/json"`,

	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		requestFlag, _ := cmd.Flags().GetString("request")
		dataFlag, _ := cmd.Flags().GetString("data")
		headers, _ := cmd.Flags().GetStringArray("header")

		host, port, path, err := parser.ParseURL(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		conn, err := httpclient.InitiateSocketConnection(host, port)
		if err != nil {
			fmt.Println(err)
			return
		}

		request, _ := httpclient.PrepareRequest(requestFlag, headers, dataFlag, host, path)

		if verbose {
			formatter.PrintRequestLines(request)
		}

		_, err = conn.Write([]byte(request))
		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)

		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		header, body := formatter.SplitResponse(string(buffer))

		if verbose {
			formatter.PrintHeaderLines(header)
		}

		fmt.Println(body)

		defer conn.Close()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var Verbose bool
	var Request string
	var Headers []string
	var Data string

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&Request, "request", "X", "GET", "HTTP request method")
	rootCmd.PersistentFlags().StringVarP(&Data, "data", "d", "", "HTTP request data")
	rootCmd.Flags().StringArrayVarP(&Headers, "header", "H", []string{}, "HTTP request headers")

}

func validateArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}

	requestFlag, _ := cmd.Flags().GetString("request")

	switch requestFlag {
	case "GET", "POST", "PUT", "DELETE":
	default:
		return fmt.Errorf("invalid request method: %s", requestFlag)
	}

	return nil
}
