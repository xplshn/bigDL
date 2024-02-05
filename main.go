// main.go // This is the main entrypoint, which calls all the different functions //

package main

import (
	"fmt"
	"os"
)

// Repositories contains the URLs for fetching metadata.
var Repositories = []string{
	"https://bin.ajam.dev/x86_64_Linux/",
	"https://bin.ajam.dev/x86_64_Linux/Baseutils/",
	"https://raw.githubusercontent.com/xplshn/Handyscripts/master/",
}

// MetadataURLs contains the URLs for fetching metadata.
var MetadataURLs = []string{
	"https://bin.ajam.dev/x86_64_Linux/METADATA.json",
	"https://bin.ajam.dev/x86_64_Linux/Baseutils/METADATA.json",
	"https://api.github.com/repos/xplshn/Handyscripts/contents",
}

const RMetadataURL = "https://raw.githubusercontent.com/metis-os/hysp-pkgs/main/data/metadata.json"

// TMPDIR is the directory for storing temporary files.
const TEMP_DIR = "/tmp/bigdl_cached"

// CACHE_FILE is the file path for caching installation information.
const CACHE_FILE = TEMP_DIR + "/bigdl_cache.log"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bigdl {list|install|remove|run|info|search} [args...]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "find_url":
		if len(os.Args) != 3 {
			fmt.Println("Usage: bigdl find_url <binary>")
			os.Exit(1)
		}
		findURLCommand(os.Args[2])
	case "install":
		if len(os.Args) < 3 {
			fmt.Println("Usage: bigdl install <binary> [install_dir] [install_message]")
			os.Exit(1)
		}
		binaryName := os.Args[2]
		var installDir, installMessage string
		if len(os.Args) > 3 {
			installDir = os.Args[3]
		}
		if len(os.Args) > 4 {
			installMessage = os.Args[4]
		}
		installCommand(binaryName, []string{installDir, installMessage})
	case "list":
		listBinaries()
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Usage: bigdl run <binary> [args...]")
			os.Exit(1)
		}
		RunFromCache(os.Args[2], os.Args[3:])
	case "info":
		if len(os.Args) != 3 {
			fmt.Println("Usage: bigdl info <package-name>")
			os.Exit(1)
		}
		packageName := os.Args[2]
		showPackageInfo(packageName)
	case "search":
		if len(os.Args) != 3 {
			fmt.Println("Usage: bigdl search <search-term>")
			os.Exit(1)
		}
		searchTerm := os.Args[2]
		fSearch(searchTerm)
	case "remove":
		if len(os.Args) != 3 {
			fmt.Println("Usage: bigdl remove <binary>")
			os.Exit(1)
		}
		binaryToRemove := os.Args[2]
		remove(binaryToRemove)
	default:
		fmt.Printf("bigdl: Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
