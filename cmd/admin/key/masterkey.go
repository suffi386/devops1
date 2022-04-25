package key

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

const (
	flagMasterKey      = "masterkeyFile"
	flagMasterKeyShort = "m"
	flagMasterKeyArg   = "masterkey"
	flagMasterKeyEnv   = "masterkeyFromEnv"
	envMasterKey       = "ZITADEL_MASTERKEY"
)

var (
	ErrNotSingleFlag = errors.New("masterkey must either be provided by file path, value or environment variable")
)

func AddMasterKeyFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(flagMasterKey, flagMasterKeyShort, "", "path to the masterkey for en/decryption keys")
	cmd.PersistentFlags().String(flagMasterKeyArg, "", "masterkey as argument for en/decryption keys")
	cmd.PersistentFlags().Count(flagMasterKeyEnv, "read masterkey for en/decryption keys from environment variable (ZITADEL_MASTERKEY)")
}

func MasterKey(cmd *cobra.Command) (string, error) {
	masterKeyFile, _ := cmd.Flags().GetString(flagMasterKey)
	masterKeyFromArg, _ := cmd.Flags().GetString(flagMasterKeyArg)
	masterKeyFromEnv, _ := cmd.Flags().GetCount(flagMasterKeyEnv)
	if err := checkSingleFlag(masterKeyFile, masterKeyFromArg, masterKeyFromEnv != 0); err != nil {
		return "", err
	}
	if masterKeyFromArg != "" {
		return masterKeyFromArg, nil
	}
	if masterKeyFromEnv != 0 {
		return os.Getenv(envMasterKey), nil
	}
	data, err := ioutil.ReadFile(masterKeyFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func checkSingleFlag(masterKeyFile, masterKeyFromArg string, masterKeyFromEnv bool) error {
	var flags int
	if masterKeyFile != "" {
		flags++
	}
	if masterKeyFromArg != "" {
		flags++
	}
	if masterKeyFromEnv {
		flags++
	}
	if flags != 1 {
		return ErrNotSingleFlag
	}
	return nil
}
