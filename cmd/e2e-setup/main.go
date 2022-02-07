package main

import (
	"context"
	"flag"
	"time"

	"github.com/caos/zitadel/internal/domain"

	"github.com/caos/logging"
	"github.com/caos/zitadel/internal/command"
	"github.com/caos/zitadel/internal/config"
	"github.com/caos/zitadel/internal/eventstore"
)

type user struct {
	desc, role, pw string
}

func main() {
	flag.Var(e2eSetupPaths, "setup-files", "paths to the setup files")
	flag.Parse()
	startE2ESetup(e2eSetupPaths.Values())
}

func startE2ESetup(configPaths []string) {

	conf := new(setupConfig)
	err := config.Read(conf, configPaths...)
	logging.Log("MAIN-EAWlt").OnError(err).Fatal("cannot read config")

	err = conf.E2E.validate()
	logging.Log("MAIN-NoZIV").OnError(err).Fatal("validating e2e config failed")

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

	eventualConsistencyCtx, cancel := context.WithTimeout(ctx, 15*time.Minute)
	defer cancel()
	err = awaitConsistency(
		eventualConsistencyCtx,
		conf.E2E,
		users,
	)
	logging.Log("MAIN-cgZ3p").OnError(err).Errorf("failed to await consistency")
}
