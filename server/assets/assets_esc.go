// Code generated by "esc -o assets/assets_esc.go -pkg assets -include (info.json|stellar.toml) ."; DO NOT EDIT.

package assets

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/test/info.json": {
		name:    "info.json",
		local:   "test/info.json",
		size:    620,
		modtime: 1571343018,
		compressed: `
H4sIAAAAAAAC/7SQzWqEQBCE7z5FM+cgbSAXzwnknuQsHe3BAR3N/KAQfPdl1NX9cRdc2GNXFUV/9R8B
iILbxionUggngPj5el8OAMGafisuRArOeH45ypI5k6ofjbcztWWTsw6FyaLXSmdUN36UMT4xqF+NBBFH
Y5h88fH9ufMVjBFfb7wzd0dzv+iUKwtD3V508q5k7VROTjU6M/znlXlsIXzaQpIqyxfAknllvY5OIWdI
W8oDmt1KB8at8N3maDgEAAD//yOD/wxsAgAA
`,
	},

	"/test/stellar.toml": {
		name:    "stellar.toml",
		local:   "test/stellar.toml",
		size:    3152,
		modtime: 1571269440,
		compressed: `
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
`,
	},

	"/": {
		name:  "/",
		local: `.`,
		isDir: true,
	},

	"/.idea": {
		name:  ".idea",
		local: `.idea`,
		isDir: true,
	},

	"/.idea/dictionaries": {
		name:  "dictionaries",
		local: `.idea/dictionaries`,
		isDir: true,
	},

	"/account": {
		name:  "account",
		local: `account`,
		isDir: true,
	},

	"/anchor": {
		name:  "anchor",
		local: `anchor`,
		isDir: true,
	},

	"/assets": {
		name:  "assets",
		local: `assets`,
		isDir: true,
	},

	"/assets/files": {
		name:  "files",
		local: `assets/files`,
		isDir: true,
	},

	"/mp-common": {
		name:  "mp-common",
		local: `mp-common`,
		isDir: true,
	},

	"/local_server_volume": {
		name:  "local_server_volume",
		local: `local_server_volume`,
		isDir: true,
	},

	"/local_server_volume/core": {
		name:  "core",
		local: `local_server_volume/core`,
		isDir: true,
	},

	"/local_server_volume/core/bin": {
		name:  "bin",
		local: `local_server_volume/core/bin`,
		isDir: true,
	},

	"/local_server_volume/core/buckets": {
		name:  "buckets",
		local: `local_server_volume/core/buckets`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp": {
		name:  "tmp",
		local: `local_server_volume/core/buckets/tmp`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/bucket-a6f2dfb73392da28": {
		name:  "bucket-a6f2dfb73392da28",
		local: `local_server_volume/core/buckets/tmp/bucket-a6f2dfb73392da28`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85": {
		name:  "catchup-921c5aeaca0d4f85",
		local: `local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger": {
		name:  "ledger",
		local: `local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger/00": {
		name:  "00",
		local: `local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger/00`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger/00/12": {
		name:  "12",
		local: `local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger/00/12`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger/00/12/b3": {
		name:  "b3",
		local: `local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger/00/12/b3`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions": {
		name:  "transactions",
		local: `local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions/00": {
		name:  "00",
		local: `local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions/00`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions/00/12": {
		name:  "12",
		local: `local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions/00/12`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions/00/12/b3": {
		name:  "b3",
		local: `local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions/00/12/b3`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/history-cd97be9828f1ddcb": {
		name:  "history-cd97be9828f1ddcb",
		local: `local_server_volume/core/buckets/tmp/history-cd97be9828f1ddcb`,
		isDir: true,
	},

	"/local_server_volume/core/buckets/tmp/process-1fef21de2d57544b": {
		name:  "process-1fef21de2d57544b",
		local: `local_server_volume/core/buckets/tmp/process-1fef21de2d57544b`,
		isDir: true,
	},

	"/local_server_volume/core/etc": {
		name:  "etc",
		local: `local_server_volume/core/etc`,
		isDir: true,
	},

	"/local_server_volume/horizon": {
		name:  "horizon",
		local: `local_server_volume/horizon`,
		isDir: true,
	},

	"/local_server_volume/horizon/bin": {
		name:  "bin",
		local: `local_server_volume/horizon/bin`,
		isDir: true,
	},

	"/local_server_volume/horizon/etc": {
		name:  "etc",
		local: `local_server_volume/horizon/etc`,
		isDir: true,
	},

	"/local_server_volume/postgresql": {
		name:  "postgresql",
		local: `local_server_volume/postgresql`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data": {
		name:  "data",
		local: `local_server_volume/postgresql/data`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/base": {
		name:  "base",
		local: `local_server_volume/postgresql/data/base`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/base/1": {
		name:  "1",
		local: `local_server_volume/postgresql/data/base/1`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/base/12404": {
		name:  "12404",
		local: `local_server_volume/postgresql/data/base/12404`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/base/12405": {
		name:  "12405",
		local: `local_server_volume/postgresql/data/base/12405`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/base/16384": {
		name:  "16384",
		local: `local_server_volume/postgresql/data/base/16384`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/base/16385": {
		name:  "16385",
		local: `local_server_volume/postgresql/data/base/16385`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/global": {
		name:  "global",
		local: `local_server_volume/postgresql/data/global`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_clog": {
		name:  "pg_clog",
		local: `local_server_volume/postgresql/data/pg_clog`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_commit_ts": {
		name:  "pg_commit_ts",
		local: `local_server_volume/postgresql/data/pg_commit_ts`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_dynshmem": {
		name:  "pg_dynshmem",
		local: `local_server_volume/postgresql/data/pg_dynshmem`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_logical": {
		name:  "pg_logical",
		local: `local_server_volume/postgresql/data/pg_logical`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_logical/mappings": {
		name:  "mappings",
		local: `local_server_volume/postgresql/data/pg_logical/mappings`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_logical/snapshots": {
		name:  "snapshots",
		local: `local_server_volume/postgresql/data/pg_logical/snapshots`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_multixact": {
		name:  "pg_multixact",
		local: `local_server_volume/postgresql/data/pg_multixact`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_multixact/members": {
		name:  "members",
		local: `local_server_volume/postgresql/data/pg_multixact/members`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_multixact/offsets": {
		name:  "offsets",
		local: `local_server_volume/postgresql/data/pg_multixact/offsets`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_notify": {
		name:  "pg_notify",
		local: `local_server_volume/postgresql/data/pg_notify`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_replslot": {
		name:  "pg_replslot",
		local: `local_server_volume/postgresql/data/pg_replslot`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_serial": {
		name:  "pg_serial",
		local: `local_server_volume/postgresql/data/pg_serial`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_snapshots": {
		name:  "pg_snapshots",
		local: `local_server_volume/postgresql/data/pg_snapshots`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_stat": {
		name:  "pg_stat",
		local: `local_server_volume/postgresql/data/pg_stat`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_stat_tmp": {
		name:  "pg_stat_tmp",
		local: `local_server_volume/postgresql/data/pg_stat_tmp`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_subtrans": {
		name:  "pg_subtrans",
		local: `local_server_volume/postgresql/data/pg_subtrans`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_tblspc": {
		name:  "pg_tblspc",
		local: `local_server_volume/postgresql/data/pg_tblspc`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_twophase": {
		name:  "pg_twophase",
		local: `local_server_volume/postgresql/data/pg_twophase`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_xlog": {
		name:  "pg_xlog",
		local: `local_server_volume/postgresql/data/pg_xlog`,
		isDir: true,
	},

	"/local_server_volume/postgresql/data/pg_xlog/archive_status": {
		name:  "archive_status",
		local: `local_server_volume/postgresql/data/pg_xlog/archive_status`,
		isDir: true,
	},

	"/local_server_volume/postgresql/etc": {
		name:  "etc",
		local: `local_server_volume/postgresql/etc`,
		isDir: true,
	},

	"/local_server_volume/supervisor": {
		name:  "supervisor",
		local: `local_server_volume/supervisor`,
		isDir: true,
	},

	"/local_server_volume/supervisor/etc": {
		name:  "etc",
		local: `local_server_volume/supervisor/etc`,
		isDir: true,
	},

	"/scripts": {
		name:  "scripts",
		local: `scripts`,
		isDir: true,
	},

	"/server": {
		name:  "server",
		local: `server`,
		isDir: true,
	},

	"/server/middleware": {
		name:  "middleware",
		local: `server/middleware`,
		isDir: true,
	},

	"/test": {
		name:  "test",
		local: `test`,
		isDir: true,
	},

	"/transfer_server": {
		name:  "transfer_server",
		local: `transfer_server`,
		isDir: true,
	},

	"/web_auth_server": {
		name:  "web_auth_server",
		local: `web_auth_server`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	".": {},

	".idea": {},

	".idea/dictionaries": {},

	"account": {},

	"anchor": {},

	"assets": {},

	"assets/files": {},

	"mp-common": {},

	"local_server_volume": {},

	"local_server_volume/core": {},

	"local_server_volume/core/bin": {},

	"local_server_volume/core/buckets": {},

	"local_server_volume/core/buckets/tmp": {},

	"local_server_volume/core/buckets/tmp/bucket-a6f2dfb73392da28": {},

	"local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85": {},

	"local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger": {},

	"local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger/00": {},

	"local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger/00/12": {},

	"local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/ledger/00/12/b3": {},

	"local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions": {},

	"local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions/00": {},

	"local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions/00/12": {},

	"local_server_volume/core/buckets/tmp/catchup-921c5aeaca0d4f85/transactions/00/12/b3": {},

	"local_server_volume/core/buckets/tmp/history-cd97be9828f1ddcb": {},

	"local_server_volume/core/buckets/tmp/process-1fef21de2d57544b": {},

	"local_server_volume/core/etc": {},

	"local_server_volume/horizon": {},

	"local_server_volume/horizon/bin": {},

	"local_server_volume/horizon/etc": {},

	"local_server_volume/postgresql": {},

	"local_server_volume/postgresql/data": {},

	"local_server_volume/postgresql/data/base": {},

	"local_server_volume/postgresql/data/base/1": {},

	"local_server_volume/postgresql/data/base/12404": {},

	"local_server_volume/postgresql/data/base/12405": {},

	"local_server_volume/postgresql/data/base/16384": {},

	"local_server_volume/postgresql/data/base/16385": {},

	"local_server_volume/postgresql/data/global": {},

	"local_server_volume/postgresql/data/pg_clog": {},

	"local_server_volume/postgresql/data/pg_commit_ts": {},

	"local_server_volume/postgresql/data/pg_dynshmem": {},

	"local_server_volume/postgresql/data/pg_logical": {},

	"local_server_volume/postgresql/data/pg_logical/mappings": {},

	"local_server_volume/postgresql/data/pg_logical/snapshots": {},

	"local_server_volume/postgresql/data/pg_multixact": {},

	"local_server_volume/postgresql/data/pg_multixact/members": {},

	"local_server_volume/postgresql/data/pg_multixact/offsets": {},

	"local_server_volume/postgresql/data/pg_notify": {},

	"local_server_volume/postgresql/data/pg_replslot": {},

	"local_server_volume/postgresql/data/pg_serial": {},

	"local_server_volume/postgresql/data/pg_snapshots": {},

	"local_server_volume/postgresql/data/pg_stat": {},

	"local_server_volume/postgresql/data/pg_stat_tmp": {},

	"local_server_volume/postgresql/data/pg_subtrans": {},

	"local_server_volume/postgresql/data/pg_tblspc": {},

	"local_server_volume/postgresql/data/pg_twophase": {},

	"local_server_volume/postgresql/data/pg_xlog": {},

	"local_server_volume/postgresql/data/pg_xlog/archive_status": {},

	"local_server_volume/postgresql/etc": {},

	"local_server_volume/supervisor": {},

	"local_server_volume/supervisor/etc": {},

	"scripts": {},

	"server": {},

	"server/middleware": {},

	"test": {
		_escData["/test/info.json"],
		_escData["/test/stellar.toml"],
	},

	"transfer_server": {},

	"web_auth_server": {},
}