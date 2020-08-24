pkgname=repos
pkgver=0.0.1
pkgrel=1
pkgdesc="A tiny script to recursively list git repositories under a given path."
url="https://github.com/kevinschoon/repos"
arch=(x86_64 aarch64 armv7h armv7l)
license=('MIT')
md5sums=()
validpgpkeys=()

build() {
	cp ../main.go .
	go build -o ${pkgname}
}

package() {
	install -Dm755 "${pkgname}" -t "${pkgdir}"/usr/bin/
}
