package cmd

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"

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
		headers, _ := cmd.Flags().GetStringArray("header")
		requestFlag, _ := cmd.Flags().GetString("request")
		dataFlag, _ := cmd.Flags().GetString("data")

		u, err := url.Parse(args[0])

		if err != nil {
			fmt.Println(err)
			return
		}

		port := u.Port()
		hostname := u.Hostname()

		if port == "" {
			port = "80"
		}

		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port))
		if err != nil {
			fmt.Println(err)
			return
		}

		method := strings.ToUpper(requestFlag)

		request := fmt.Sprintf("%s %s HTTP/1.1\n", method, u.Path)
		request += fmt.Sprintf("Host: %s\n", u.Host)
		request += "Accept: */*\n"
		request += "Connection: close\n"

		for _, header := range headers {
			request += header + "\n"
		}

		if dataFlag != "" {
			request += fmt.Sprintf("Content-Length: %d\n", len(dataFlag))
			request += "\n"
			request += dataFlag
		}

		request += "\n"

		if verbose {
			printRequestLines(request)
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

		header, body := splitResponse(string(buffer))

		if verbose {
			printHeaderLines(header)
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

func printRequestLines(request string) {
	lines := strings.Split(request, "\n")
	for _, line := range lines[:len(lines)-1] {
		fmt.Println(">" + line)
	}
}

func splitResponse(response string) (header, body string) {
	parts := strings.Split(response, "\r\n\r\n")
	return parts[0], parts[1]
}

func printHeaderLines(header string) {
	lines := strings.Split(header, "\n")
	for _, line := range lines {
		fmt.Println("<" + line)
	}
}
