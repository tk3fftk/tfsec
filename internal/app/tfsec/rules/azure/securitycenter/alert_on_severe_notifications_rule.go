package securitycenter

import (
	"github.com/aquasecurity/defsec/rules/azure/securitycenter"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
		resource "azurerm_security_center_contact" "bad_example" {
		email = "bad_example@example.com"
		phone = "+1-555-555-5555"

		alert_notifications = false
		alerts_to_admins = false
		}
		`},
		GoodExample: []string{`
		resource "azurerm_security_center_contact" "good_example" {
		email = "good_example@example.com"
		phone = "+1-555-555-5555"

		alert_notifications = true
		alerts_to_admins = true
		}
	`},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/security_center_contact#alert_notifications",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"azurerm_security_center_contact"},
		Base:           securitycenter.CheckAlertOnSevereNotifications,
	})
}