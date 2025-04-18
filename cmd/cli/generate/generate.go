package generate

import (
	"fmt"
	"log"

	"github.com/mrjxtr-dev/mr-shush/cmd/cli/root"
	"github.com/mrjxtr-dev/mr-shush/internal/config"
	"github.com/mrjxtr-dev/mr-shush/internal/generator"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Password",
	Run: func(cmd *cobra.Command, args []string) {
		// Debug line to confirm command is running
		fmt.Println("DEBUG: Generate command is running")
		
		length, strength, err := getFlagValues(cmd)
		if err != nil {
			log.Fatal(err)
		}

		gen := generator.New()
		password, err := gen.GeneratePassword(length, strength)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Password:", password)
	},
}

func init() {
	root.RootCmd.AddCommand(generateCmd)

	generateCmd.Flags().
		BoolP("strong", "s", false, "Generate a strong password")
	generateCmd.Flags().
		BoolP("good", "g", false, "Generate a good password")
	generateCmd.Flags().
		BoolP("weak", "w", false, "Generate a weak password")
	generateCmd.Flags().
		IntP("length", "l", config.DefaultLength, "Length of the password")
}

func getFlagValues(cmd *cobra.Command) (int, string, error) {
	length, err := cmd.Flags().GetInt("length")
	if err != nil {
		return 0, "", err
	}

	strong, err := cmd.Flags().GetBool("strong")
	if err != nil {
		return 0, "", err
	}
	good, err := cmd.Flags().GetBool("good")
	if err != nil {
		return 0, "", err
	}
	weak, err := cmd.Flags().GetBool("weak")
	if err != nil {
		return 0, "", err
	}

	var strength string
	if strong {
		strength = "strong"
	} else if good {
		strength = "good"
	} else if weak {
		strength = "weak"
	} else {
		strength = "good"
	}

	return length, strength, nil
}
