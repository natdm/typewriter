package template

import (
	"io"

	"strings"

	"sort"

	log "github.com/Sirupsen/logrus"
)

func Draw(t map[string]*PackageType, out io.Writer, lang Language, verbose bool) error {
	if err := Header(out, lang); err != nil {
		return err
	}

	for _, v := range t {
		if s, ok := v.Type.(*Struct); ok {

			// If any of the structs have embedded types, transfer them to the struct
			if len(s.Embedded) > 0 {
			EMBEDLOOP:
				for _, v := range s.Embedded {
					_v := strings.TrimSpace(v)
					if _, ok := t[_v]; ok {
						if s2, ok := t[_v].Type.(*Struct); ok {
							s.Fields = append(s.Fields, s2.Fields...)
						} else if verbose {
							log.Errorf("error while embedding type %s", _v)
						}
						continue EMBEDLOOP
					}
					if verbose {
						log.Errorf("unable to find type %v -- unable to embed", _v)
					}
					continue EMBEDLOOP

				}
			}
		}
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
