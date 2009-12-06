all:
	cd glib && make
	cd gtk && make

install:
	cd glib && make install
	cd gtk && make install

clean:
	cd glib && make clean
	cd gtk && make clean

example: install
	cd example && make
