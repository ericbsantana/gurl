package cmd

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Make a GET request to a server",
	Args:  validateArgs,

	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

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

		request := fmt.Sprintf("GET %s HTTP/1.1\n", u.Path)
		request += fmt.Sprintf("Host: %s\n", u.Host)
		request += "Accept: */*\n"
		request += "Connection: close\n"
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

func init() {
	rootCmd.AddCommand(getCmd)
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
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
