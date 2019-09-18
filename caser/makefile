build: caesar.cpp
	g++ -o caesar caesar.cpp

check: build
	./caesar -e -i "meet me after the toga party" | diff ciphertext.txt -
	./caesar -d -i "PHHW PH DIWHU WKH WRJD SDUWB" | diff plantext.txt -
	./caesar -e -f "plantext.txt" | diff ciphertext.txt -
	./caesar -d -f "ciphertext.txt" | diff plantext.txt -

.PHONY : clean
clean:
	- rm caesar