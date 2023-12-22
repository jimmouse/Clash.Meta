package resource

import (
	"bytes"
	"crypto/md5"
	types "github.com/metacubex/mihomo/constant/provider"
	"github.com/samber/lo"
	"os"
	"time"
)

func (f *Fetcher[V]) SideUpdate(buf []byte) (V, bool, error) {
	now := time.Now()
	hash := md5.Sum(buf)
	if bytes.Equal(f.hash[:], hash[:]) {
		f.UpdatedAt = now
		_ = os.Chtimes(f.vehicle.Path(), now, now)
		return lo.Empty[V](), true, nil
	}

	contents, err := f.parser(buf)
	if err != nil {
		return lo.Empty[V](), false, err
	}

	if f.vehicle.Type() != types.File {
		if err := safeWrite(f.vehicle.Path(), buf); err != nil {
			return lo.Empty[V](), false, err
		}
	}

	f.UpdatedAt = now
	f.hash = hash

	return contents, false, nil
}
