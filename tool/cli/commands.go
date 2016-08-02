package cli

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/jaredwarren/easypostTester/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"strings"
)

type (
	// CareateCarrierAccountCommand is the command line data structure for the careate action of carrier_account
	CareateCarrierAccountCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteCarrierAccountCommand is the command line data structure for the delete action of carrier_account
	DeleteCarrierAccountCommand struct {
		ID          string
		PrettyPrint bool
	}

	// ListCarrierAccountCommand is the command line data structure for the list action of carrier_account
	ListCarrierAccountCommand struct {
		PrettyPrint bool
	}

	// ShowCarrierAccountCommand is the command line data structure for the show action of carrier_account
	ShowCarrierAccountCommand struct {
		// Carrier ID
		ID          string
		PrettyPrint bool
	}

	// UpdateCarrierAccountCommand is the command line data structure for the update action of carrier_account
	UpdateCarrierAccountCommand struct {
		Payload     string
		ContentType string
		ID          string
		PrettyPrint bool
	}

	// ShowCarrierTypesCommand is the command line data structure for the show action of carrier_types
	ShowCarrierTypesCommand struct {
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "careate",
		Short: `CarrierAccount objects may be managed through the EasyPost API using the Production mode API key only. Multiple accounts can be added for a single carrier (with the exception of the USPS).`,
	}
	tmp1 := new(CareateCarrierAccountCommand)
	sub = &cobra.Command{
		Use:   `carrier_account ["/v2/carrier_accounts"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `CarrierAccount objects may be removed from your account when they become out of date or no longer useful.`,
	}
	tmp2 := new(DeleteCarrierAccountCommand)
	sub = &cobra.Command{
		Use:   `carrier_account ["/v2/carrier_accounts/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `Retrieve an unpaginated list of all CarrierAccounts available to the authenticated account. Only Production API keys may be used to retrieve this list, as there is no test mode equivalent.`,
	}
	tmp3 := new(ListCarrierAccountCommand)
	sub = &cobra.Command{
		Use:   `carrier_account ["/v2/carrier_accounts"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp4 := new(ShowCarrierAccountCommand)
	sub = &cobra.Command{
		Use:   `carrier_account ["/v2/carrier_accounts/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(ShowCarrierTypesCommand)
	sub = &cobra.Command{
		Use:   `carrier_types ["/v2/carrier_types"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `Updates can be made to description, reference, and any fields in credentials or test_credentials.`,
	}
	tmp6 := new(UpdateCarrierAccountCommand)
	sub = &cobra.Command{
		Use:   `carrier_account ["/v2/carrier_accounts/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

// Run makes the HTTP request corresponding to the CareateCarrierAccountCommand command.
func (cmd *CareateCarrierAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v2/carrier_accounts"
	}
	var payload client.CareateCarrierAccountPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CareateCarrierAccount(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CareateCarrierAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteCarrierAccountCommand command.
func (cmd *DeleteCarrierAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v2/carrier_accounts/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteCarrierAccount(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteCarrierAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the ListCarrierAccountCommand command.
func (cmd *ListCarrierAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v2/carrier_accounts"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListCarrierAccount(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListCarrierAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowCarrierAccountCommand command.
func (cmd *ShowCarrierAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v2/carrier_accounts/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowCarrierAccount(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowCarrierAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `Carrier ID`)
}

// Run makes the HTTP request corresponding to the UpdateCarrierAccountCommand command.
func (cmd *UpdateCarrierAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v2/carrier_accounts/%v", cmd.ID)
	}
	var payload client.UpdateCarrierAccountPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateCarrierAccount(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateCarrierAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the ShowCarrierTypesCommand command.
func (cmd *ShowCarrierTypesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v2/carrier_types"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowCarrierTypes(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowCarrierTypesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}
