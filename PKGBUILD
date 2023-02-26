# Maintainer: James Kimble <jckimble@pm.me>
pkgname=otpcli
pkgver=1.0.0
pkgrel=1
pkgdesc="A quick tool for totp tokens. Made cause other desktop authenticators bug me."
arch=('x86_64')
url="https://github.com/jckimble/otpcli"
license=('APACHE')
makedepends=('go')

source=()

build(){
	cd ..
	go build -o otpcli .
}

package(){
	cd ..
	install -Dm755 otpcli $pkgdir/usr/bin/otpcli
}

