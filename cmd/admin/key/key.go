package key

import (
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sigs.k8s.io/yaml"

	"github.com/caos/zitadel/cmd/helper"

	"github.com/caos/zitadel/internal/crypto"
	cryptoDB "github.com/caos/zitadel/internal/crypto/database"
	"github.com/caos/zitadel/internal/database"
)

const (
	flagMasterKey = "masterkey"
	flagKeyFile   = "file"
)

type Config struct {
	Database database.Config
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "manage encryption keys",
	}
	cmd.PersistentFlags().String(flagMasterKey, "", "masterkey for en/decryption keys")
	cmd.AddCommand(newKey())
	return cmd
}

func newKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new [keyID=key]... [-f file]",
		Short: "create new encryption key(s)",
		Long: `create new encryption key(s) (encrypted by the provided master key)
provide key(s) by YAML file and/or by argument
Requirements:
- cockroachdb`,
		Example: `new -f keys.yaml
new key1=somekey key2=anotherkey
new -f keys.yaml key2=anotherkey`,
		RunE: func(cmd *cobra.Command, args []string) error {
			keys, err := keysFromArgs(args)
			if err != nil {
				return err
			}
			filePath, _ := cmd.Flags().GetString(flagKeyFile)
			if filePath != "" {
				file, err := openFile(filePath)
				if err != nil {
					return err
				}
				yamlKeys, err := keysFromYAML(file)
				if err != nil {
					return err
				}
				keys = append(keys, yamlKeys...)
			}
			config := new(Config)
			if err := viper.Unmarshal(config); err != nil {
				return err
			}
			masterKey, _ := cmd.Flags().GetString(flagMasterKey)
			storage, err := keyStorage(config.Database, masterKey)
			if err != nil {
				return err
			}
			return storage.CreateKeys(keys...)
		},
	}
	cmd.PersistentFlags().StringP(flagKeyFile, "f", "", "path to keys file")
	return cmd
}

func keysFromArgs(args []string) ([]crypto.Key, error) {
	keys := make([]crypto.Key, len(args))
	for i, arg := range args {
		key := strings.Split(arg, "=")
		if len(key) != 2 {
			return nil, helper.NewUserError("argument is not in the valid format [keyID=key]")
		}
		keys[i] = crypto.Key{
			ID:    key[0],
			Value: key[1],
		}
	}
	return keys, nil
}

func keysFromYAML(file io.Reader) ([]crypto.Key, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, helper.NewUserErrorf("unable to extract keys from file").WithParent(err)
	}
	keysYAML := make(map[string]string)
	if err = yaml.Unmarshal(data, &keysYAML); err != nil {
		return nil, helper.NewUserError("unable to extract keys from file").WithParent(err)
	}
	keys := make([]crypto.Key, 0, len(keysYAML))
	for id, key := range keysYAML {
		keys = append(keys, crypto.Key{
			ID:    id,
			Value: key,
		})
	}
	return keys, nil
}

func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, helper.NewUserErrorf("failed to open file: %s", fileName)
	}
	return file, nil
}

func keyStorage(config database.Config, masterKey string) (crypto.KeyStorage, error) {
	db, err := database.Connect(config)
	if err != nil {
		return nil, err
	}
	return cryptoDB.NewKeyStorage(db, masterKey)
}
