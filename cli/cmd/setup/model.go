package setup

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/monorkin/twice/cli/cmd"
	api "github.com/monorkin/twice/cli/internal/api"
)

type View int

const (
	DockerView       View = 0
	ProductView      View = 1
	InstallationView View = 2
)

type model struct {
	apiClient       *api.Client
	keys            keyMap
	help            help.Model
	licenseKey      string
	currentView     View
	checkingDocker  bool
	dockerInstalled bool
	dockerRunning   bool
	productsLoaded  bool
	products        []api.Product
	selectedProduct *Product
	productList     list.Model
	spinner         spinner.Model
	err             error
	quitting        bool
	width           int
	height          int
}

func initialModel(licenseKey string) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = cmd.SpinnerStyle

	productListItems := make([]list.Item, 0)
	productList := list.New(productListItems, list.NewDefaultDelegate(), 0, 0)
	productList.Title = "Available Products"
	productList.SetShowStatusBar(true)
	productList.SetFilteringEnabled(true)

	var apiClient *api.Client
	licenseParts := strings.SplitN(licenseKey, "@", 2)

	if len(licenseParts) == 2 {
		apiClient = api.NewClient(licenseParts[1], licenseParts[0])
	} else {
		apiClient = api.NewClient("localhost:3000", licenseParts[0])
	}

	return model{
		keys:           keys,
		help:           help.New(),
		apiClient:      apiClient,
		licenseKey:     licenseKey,
		currentView:    DockerView,
		checkingDocker: true,
		productList:    productList,
		spinner:        s,
		width:          80,
		height:         24,
	}
}
