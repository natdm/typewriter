package template

import (
	"io"

	"sort"

	log "github.com/Sirupsen/logrus"
)

func Draw(t map[string]*PackageType, out io.Writer, lang Language, verbose bool) error {
	if err := Header(out, lang); err != nil {
		return err
	}

	keys := make([]string, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := t[k]
		if err := v.Template(out, lang); err != nil {
			return err
		}
		if err := Raw(out, "\n"); err != nil && verbose {
			log.Warn("unable to create new line")
		}
	}
	return nil
}
