package template

import (
	"io"

	"sort"

	log "github.com/sirupsen/logrus"
)

// Draw draws all types to a writer.
func Draw(t map[string]*PackageType, out io.Writer, lang Language, verbose bool) (int, error) {
	if err := Header(out, lang); err != nil {
		return 0, err
	}

	keys := make([]string, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := t[k]
		if err := v.Template(out, lang); err != nil {
			return 0, err
		}
		if err := Raw(out, "\n"); err != nil && verbose {
			log.WithField("type", k).Warn("unable to create new line")
		}
		if verbose {
			log.Infof("created type: %s", k)
		}
	}
	return len(keys), nil
}
