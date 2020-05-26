// This file is generated by generate-std.joke script. Do not edit manually!

package yaml

import (
	"fmt"
	. "github.com/candid82/joker/core"
	"os"
)

func InternsOrThunks() {
	if VerbosityLevel > 0 {
		fmt.Fprintln(os.Stderr, "Lazily running slow version of yaml.InternsOrThunks().")
	}
	yamlNamespace.ResetMeta(MakeMeta(nil, `Implements encoding and decoding of YAML.`, "1.0"))

	yamlNamespace.InternVar("read-string", read_string_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Parses the YAML-encoded data and return the result as a Joker value.`, "1.0"))

	yamlNamespace.InternVar("write-string", write_string_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("v"))),
			`Returns the YAML encoding of v.`, "1.0"))

}
