package main

import (
	"context"
	"flag"
	"time"

	internal_authz "github.com/caos/zitadel/internal/api/authz"

	sd "github.com/caos/zitadel/internal/config/systemdefaults"
	"github.com/caos/zitadel/internal/config/types"
	"github.com/caos/zitadel/internal/domain"

	"github.com/caos/logging"
	"github.com/caos/zitadel/internal/command"
	"github.com/caos/zitadel/internal/config"
	"github.com/caos/zitadel/internal/eventstore"
)

type E2EConfig struct {
	Org                            string
	MachineKeyPath                 string
	ZitadelProjectResourceID       string
	APIURL                         string
	IssuerURL                      string
	OrgOwnerPassword               string
	OrgOwnerViewerPassword         string
	OrgProjectCreatorPassword      string
	PasswordComplexityUserPassword string
	LoginPolicyUserPassword        string
}

type setupConfig struct {
	E2E E2EConfig

	Log logging.Config

	Eventstore     types.SQL
	SystemDefaults sd.SystemDefaults
	InternalAuthZ  internal_authz.Config
}

type user struct {
	desc, role, pw string
}

var (
	e2eSetupPaths = config.NewArrayFlags("authz.yaml", "system-defaults.yaml", "setup.yaml", "e2e.yaml")
)

func main() {
	flag.Var(e2eSetupPaths, "setup-files", "paths to the setup files")
	flag.Parse()
	startE2ESetup(e2eSetupPaths.Values())
}

func startE2ESetup(configPaths []string) {

	conf := new(setupConfig)
	err := config.Read(conf, configPaths...)
	logging.Log("MAIN-EAWlt").OnError(err).Fatal("cannot read config")

	ctx := context.Background()

	es, err := eventstore.Start(conf.Eventstore)
	logging.Log("MAIN-wjQ8G").OnError(err).Fatal("cannot start eventstore")

	commands, err := command.StartCommands(
		es,
		conf.SystemDefaults,
		conf.InternalAuthZ,
		nil,
		command.OrgFeatureCheckerFunc(func(_ context.Context, _ string, _ ...string) error { return nil }),
	)
	logging.Log("MAIN-54MLq").OnError(err).Fatal("cannot start command side")

	users := []user{{
		desc: "org_owner",
		pw:   conf.E2E.OrgOwnerPassword,
		role: domain.RoleOrgOwner,
	}, {
		desc: "org_owner_viewer",
		pw:   conf.E2E.OrgOwnerViewerPassword,
		role: domain.RoleOrgOwner,
	}, {
		desc: "org_project_creator",
		pw:   conf.E2E.OrgProjectCreatorPassword,
		role: domain.RoleOrgProjectCreator,
	}, {
		desc: "login_policy_user",
		pw:   conf.E2E.LoginPolicyUserPassword,
	}, {
		desc: "password_complexity_user",
		pw:   conf.E2E.PasswordComplexityUserPassword,
	}}

	err = execute(ctx, commands, conf.E2E, users)
	logging.Log("MAIN-cgZ3p").OnError(err).Errorf("failed to execute commands steps")

	eventualConsistencyCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	err = awaitConsistency(
		eventualConsistencyCtx,
		conf.E2E,
		users,
	)
	logging.Log("MAIN-cgZ3p").OnError(err).Errorf("failed to await consistency")
}
