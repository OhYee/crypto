build: caesar.cpp
	g++ -o caesar caesar.cpp

check: build
	./caesar -e -i "meet me after the toga party" | diff -B ciphertext.txt -
	./caesar -d -i "PHHW PH DIWHU WKH WRJD SDUWB" | diff -B plantext.txt -
	./caesar -e -f "plantext.txt" | diff -B ciphertext.txt -
	./caesar -d -f "ciphertext.txt" | diff -B plantext.txt -

.PHONY : clean
clean:
	- rm caesar