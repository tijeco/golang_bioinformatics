package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tijeco/cluck/modules"
)

// Create a new type for a fasta of Strings
type stringList []string

// Implement the flag.Value interface
func (s *stringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

func main() {
	// Subcommands
	fastaCommand := flag.NewFlagSet("fasta", flag.ExitOnError)
	blastaxCommand := flag.NewFlagSet("blastax", flag.ExitOnError)

	// fasta subcommand flag pointers
	fastaInputFile := fastaCommand.String("input", "", "Text to parse. (Required)")
	fastaMetricPtr := fastaCommand.String("output", "chars", "Metric <chars|words|lines>. (Required)")

	// blastax subcommand
	blastaxInputFile := blastaxCommand.String("input", "", "Input blast file. (Required)")
	blastaxOptions := blastaxCommand.String("taxon", "phylum", "Metric <kingdom|phylum|class|order|family|genus|species>. (Required)")
	blastaxDB := blastaxCommand.String("makedb", "", "Metric <id_parent|id_nodes|accession_id> ")

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("fasta  subcommand is required")
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case "fasta":
		fastaCommand.Parse(os.Args[2:])
	case "blastax":
		blastaxCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	if fastaCommand.Parsed() {
		// Required Flags
		if *fastaInputFile == "" {
			fastaCommand.PrintDefaults()
			os.Exit(1)
		}
		//Choice flag
		metricChoices := map[string]bool{"chars": true, "words": true, "lines": true}
		if _, validChoice := metricChoices[*fastaMetricPtr]; !validChoice {
			fastaCommand.PrintDefaults()
			os.Exit(1)
		}
		// Print
		fmt.Printf("textPtr: %s, metricPtr: %s\n",
			*fastaInputFile,
			*fastaMetricPtr,
			// *fastaUniquePtr,
		)
	} //end fasta command parse

	if blastaxCommand.Parsed() {
		// Required Flags
		if *blastaxInputFile == "" {
			blastaxCommand.PrintDefaults()
			os.Exit(1)
		}
		//Choice flag
		blastaxChoices := map[string]bool{"kingdom": true, "phylum": true, "class": true, "order": true, "family": true, "genus": true}
		if _, validChoice := blastaxChoices[*blastaxOptions]; !validChoice {
			blastaxCommand.PrintDefaults()
			os.Exit(1)
		}
		blastaxDBChoices := map[string]bool{"id_parent": true, "id_nodes": true, "accession_id": true}
		if _, validChoice := blastaxDBChoices[*blastaxDB]; !validChoice {
			fmt.Println("YESSSSS")
			blastaxCommand.PrintDefaults()
			os.Exit(1)
		}
		// Print
		fmt.Printf("input: %s, taxon: %s\n",
			*blastaxInputFile,
			*blastaxOptions,
		)

		modules.MakeDB(*blastaxInputFile, *blastaxDB)
		fmt.Println(*blastaxDB, *blastaxInputFile)
		// demo()

	} //end blastax command parse
}
