package cmds

import (
	"fmt"
	"github.com/beatoz/beatoz-go/libs"
	"github.com/beatoz/sfeeder/common"
	"github.com/spf13/cobra"
)

var targetAddr []byte
var fullText bool

func NewCmd_GetSecret() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"read", "show"},
		Short:   "get secret",
		RunE:    getSecret,
	}
	addGetSecretFlags(cmd)
	return cmd
}

func addGetSecretFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BytesHexVarP(&targetAddr, "addr", "a", targetAddr, "-")
	cmd.PersistentFlags().BoolVarP(&fullText, "full", "f", fullText, "set whether the secret value is shown partially or fully.")
}

func getSecret(cmd *cobra.Command, args []string) error {
	defer func() {
		libs.ClearCredential([]byte(pass))
	}()

	secret, err := common.ReadSecret(targetAddr, []byte(pass), DataDir)
	defer libs.ClearCredential(secret)

	if err != nil {
		return err
	}

	if !fullText {
		l := len(secret)
		if l <= 8 {
			n := libs.MIN(2, l)
			secret = append(secret[:n], []byte("..")...)
		} else if l <= 16 {
			s0 := secret[:2]
			s1 := secret[l-2:]
			secret = []byte(fmt.Sprintf("%s..%s", s0, s1))
		} else {
			s0 := secret[:4]
			s1 := secret[l-4:]
			secret = []byte(fmt.Sprintf("%s..%s", s0, s1))
		}
	}

	fmt.Println("Secret:", string(secret))
	return nil
}
