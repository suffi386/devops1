package orb

import (
	"github.com/caos/orbos/mntr"
	"github.com/caos/orbos/pkg/tree"
	"github.com/caos/zitadel/operator/kinds/iam/zitadel"
	"github.com/pkg/errors"
)

func BackupListFunc() func(monitor mntr.Monitor, desiredTree *tree.Tree) (strings []string, err error) {
	return func(monitor mntr.Monitor, desiredTree *tree.Tree) (strings []string, err error) {
		desiredKind, err := parseDesiredV0(desiredTree)
		if err != nil {
			return nil, errors.Wrap(err, "parsing desired state failed")
		}
		desiredTree.Parsed = desiredKind

		if desiredKind.Spec.Verbose && !monitor.IsVerbose() {
			monitor = monitor.Verbose()
		}

		return zitadel.BackupListFunc()(monitor, desiredKind.IAM)
	}
}
