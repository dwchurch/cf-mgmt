package configcommands

import (
	"errors"
	"fmt"

	"github.com/pivotalservices/cf-mgmt/config"
)

type UpdateSpaceConfigurationCommand struct {
	ConfigManager config.Manager
	BaseConfigCommand
	OrgName               string     `long:"org" description:"Org name" required:"true"`
	SpaceName             string     `long:"space" description:"Space name" required:"true"`
	AllowSSH              string     `long:"allow-ssh" description:"Enable the Space Quota in the config" choice:"true" choice:"false"`
	EnableRemoveUsers     string     `long:"enable-remove-users" description:"Enable removing users from the space" choice:"true" choice:"false"`
	EnableSecurityGroup   string     `long:"enable-security-group" description:"Enable space level security group definitions" choice:"true" choice:"false"`
	IsoSegment            string     `long:"isolation-segment" description:"Isolation segment assigned to space"`
	ClearIsolationSegment bool       `long:"clear-isolation-segment" description:"Sets the isolation segment to blank"`
	ASGs                  []string   `long:"named-asg" description:"Named asg(s) to assign to space, specify multiple times"`
	ASGsToRemove          []string   `long:"named-asg-to-remove" description:"Named asg(s) to remove, specify multiple times"`
	NamedQuota            string     `long:"named-quota" description:"Named quota to assign to space"`
	ClearNamedQuota       bool       `long:"clear-named-quota" description:"Sets the named quota to blank"`
	Quota                 SpaceQuota `group:"quota"`
	Developer             UserRole   `group:"developer" namespace:"developer"`
	Manager               UserRole   `group:"manager" namespace:"manager"`
	Auditor               UserRole   `group:"auditor" namespace:"auditor"`
}

//Execute - updates space configuration`
func (c *UpdateSpaceConfigurationCommand) Execute(args []string) error {
	c.initConfig()
	spaceConfig, err := c.ConfigManager.GetSpaceConfig(c.OrgName, c.SpaceName)
	if err != nil {
		return err
	}
	asgConfigs, err := c.ConfigManager.GetASGConfigs()
	if err != nil {
		return err
	}

	if c.Quota.EnableSpaceQuota == "true" && c.NamedQuota != "" {
		return fmt.Errorf("cannot enable space quota and use named quotas")
	}
	
	errorString := ""

	convertToBool("allow-ssh", &spaceConfig.AllowSSH, c.AllowSSH, &errorString)
	convertToBool("enable-remove-users", &spaceConfig.RemoveUsers, c.EnableRemoveUsers, &errorString)
	convertToBool("enable-security-group", &spaceConfig.EnableSecurityGroup, c.EnableSecurityGroup, &errorString)
	if c.IsoSegment != "" {
		spaceConfig.IsoSegment = c.IsoSegment
	}
	if c.ClearIsolationSegment {
		spaceConfig.IsoSegment = ""
	}

	spaceConfig.ASGs = removeFromSlice(addToSlice(spaceConfig.ASGs, c.ASGs, &errorString), c.ASGsToRemove)
	validateASGsExist(asgConfigs, spaceConfig.ASGs, &errorString)
	updateSpaceQuotaConfig(spaceConfig, c.Quota, &errorString)

	if c.NamedQuota != "" {
		spaceConfig.NamedQuota = c.NamedQuota
	}
	if c.ClearNamedQuota {
		spaceConfig.NamedQuota = ""
	}

	c.updateUsers(spaceConfig, &errorString)

	if errorString != "" {
		return errors.New(errorString)
	}

	if err := c.ConfigManager.SaveSpaceConfig(spaceConfig); err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("The org/space [%s/%s] has been updated", c.OrgName, c.SpaceName))
	return nil
}

func (c *UpdateSpaceConfigurationCommand) updateUsers(spaceConfig *config.SpaceConfig, errorString *string) {
	updateUsersBasedOnRole(&spaceConfig.Developer, spaceConfig.GetDeveloperGroups(), &c.Developer, errorString)
	updateUsersBasedOnRole(&spaceConfig.Auditor, spaceConfig.GetAuditorGroups(), &c.Auditor, errorString)
	updateUsersBasedOnRole(&spaceConfig.Manager, spaceConfig.GetManagerGroups(), &c.Manager, errorString)

	spaceConfig.DeveloperGroup = ""
	spaceConfig.ManagerGroup = ""
	spaceConfig.AuditorGroup = ""
}

func (c *UpdateSpaceConfigurationCommand) initConfig() {
	if c.ConfigManager == nil {
		c.ConfigManager = config.NewManager(c.ConfigDirectory)
	}
}
