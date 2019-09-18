
.PHONY : clean,check,build
build: 
	$(MAKE) -C ./caser build
check:
	$(MAKE) -C ./caser check
clean:
	- $(MAKE) -C ./caser clean
	
