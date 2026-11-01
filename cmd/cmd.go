package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"amrita_pyq/cmd/internal/configs"
	"amrita_pyq/cmd/internal/logo"
	"amrita_pyq/cmd/internal/menu"
	"amrita_pyq/cmd/internal/requestclient"
	"amrita_pyq/cmd/internal/webclient"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "pranav-pyq",
	Short: "PRANAV PYQ FETCHER AMRITA",
	Long:  `PRANAV PYQ FETCHER AMRITA - An interactive CLI application to access Amrita University's digital repository for previous year question papers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(configs.LogoStyle.Render(logo.LOGO_ASCII))

		// Check connectivity to Amrita DSpace repository
		connected := checkAmritaConnectivity()

		if connected {
			fmt.Println(configs.SuccessStyle.Render("✓ Connected to Amrita Network (dspace.amritanet.edu:8080)"))
			configs.IsDemoMode = false
		} else {
			fmt.Println(configs.WarningStyle.Render(
				"⚠️  AMRITA CAMPUS NETWORK NOT DETECTED\n\n" +
					"The question paper repository (dspace.amritanet.edu:8080) is hosted on the\n" +
					"Amrita intranet. Live fetching requires:\n" +
					"  1. Amrita Campus Wi-Fi, OR\n" +
					"  2. Amrita University VPN\n\n" +
					"You can still explore the app right now using Demo Mode!",
			))

			var choice string
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Title("Select an option").
						Options(
							huh.NewOption("🚀 Launch Demo Mode (Browse sample courses & papers offline)", "demo"),
							huh.NewOption("🔄 Try connecting live anyway", "live"),
							huh.NewOption("❌ Quit", "quit"),
						).
						Value(&choice),
				),
			)

			if err := form.Run(); err != nil {
				os.Exit(0)
			}

			switch choice {
			case "demo":
				configs.IsDemoMode = true
				fmt.Println(configs.SuccessStyle.Render("⚡ Running in Offline Demo Mode — explore courses, semesters, and question papers!"))
			case "live":
				configs.IsDemoMode = false
			case "quit":
				os.Exit(0)
			}
		}

		cs := menu.CourseSelect{
			ReqClient: requestclient.RequestClient{
				WebClient: webclient.DefaultWebClient{},
			},
		}
		cs.ChooseCourse()
	},
}

func checkAmritaConnectivity() bool {
	client := http.Client{
		Timeout: 2500 * time.Millisecond,
	}
	resp, err := client.Get(configs.BASE_URL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
