package main // import "arp242.net/autofox"

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"arp242.net/sconfig"
)

type config struct {
	//Addon []string
	Set [][]string

	// Shortcuts
	DisableOnboarding bool
	ReasonablePrivacy bool
	DisablePocket     bool
}

func fatal(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "autofox: %v\n", err)
	os.Exit(1)
}

var reNum = regexp.MustCompile(`^\d+$`)

// Set config values. This works by adding a call to user.js:
//
//   user_pref("key", value);
//
// https://developer.mozilla.org/en-US/docs/Mozilla/Preferences/A_brief_guide_to_Mozilla_preferences
func main() {
	var c config
	err := sconfig.Parse(&c, "autofox.conf", sconfig.Handlers{
		"Set": func(line []string) error {
			// TODO: empty item at end of line?
			// TODO: making c.Set a map and then setting something means it
			// stops processing other values??
			//fmt.Printf("XX %v -> %#v\n", len(line), line)
			if len(line) < 2 {
				return errors.New("need value")
			}
			//c.Set[line[0]] = strings.Join(line[1:], " ")
			c.Set = append(c.Set, []string{line[0], strings.TrimSpace(strings.Join(line[1:], " "))})
			return nil
		},
	})

	if c.DisableOnboarding {
		c.Set = append(c.Set, shortcuts["disable-onboarding"]...)
	}
	if c.DisablePocket {
		c.Set = append(c.Set, shortcuts["disable-pocket"]...)
	}
	if c.ReasonablePrivacy {
		c.Set = append(c.Set, shortcuts["reasonable-privacy"]...)
	}

	var b strings.Builder
	for _, s := range c.Set {
		var v string
		switch {
		case s[1] == "true", s[1] == "false", reNum.MatchString(s[1]):
			v = s[1]
		default:
			j, err := json.Marshal(s[1])
			if err != nil {
				fatal(err)
			}
			v = string(j)
		}

		b.WriteString(fmt.Sprintf("user_pref(\"%s\", %s);\n", s[0], v))
	}

	// Write user.js
	p, err := profile()
	if err != nil {
		fatal(err)
	}
	fmt.Println(p)
	fp, err := os.Create(p + "/user.js")
	if err != nil {
		fatal(err)
	}
	_, err = fp.WriteString(b.String())
	if err != nil {
		_ = fp.Close()
		fatal(err)
	}

	err = fp.Close()
	if err != nil {
		fatal(err)
	}
}

// Get location to default profile directory.
func profile() (string, error) {
	pdir := os.Getenv("HOME") + "/.mozilla/firefox"
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}

		// TODO: profile dir.
		// TODO: macOS?
	}

	// TODO: read profiles.ini
	m, err := filepath.Glob(pdir + "/*.default")
	if err != nil {
		return "", fmt.Errorf("could not locate default profile: %v", err)
	}
	if len(m) == 0 {
		return "", fmt.Errorf("could not locate default profile")
	}
	if len(m) > 1 {
		return "", fmt.Errorf("could not locate default profile")
	}

	return m[0], nil
}
