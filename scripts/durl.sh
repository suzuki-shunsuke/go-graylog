find . \
  -type d -name node_modules -prune -o \
  -type d -name .git -prune -o \
  -type d -name vendor -prune -o \
  -type d -name dist -prune -o \
  -type f -print | \
  grep -v package-lock.json | \
  grep -v Gopkg.toml | \
  grep -v Gopkg.lock | \
  durl check || exit 1
