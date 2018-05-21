# updateTranslation
A utility to regenerate and update existing translation files for all english keys.

### Usage ###

To start you need to create a yandex.json file in the same directory as your binary.

 	echo -e '{
    	"key":"YourAPIKey",
    	"path":"web/pages",
    	"languages":[
      		"es",
      		"ja"
    	]
  	}' > yandex.json
  	
### Updating all existing english key values to your languages.###

Linux AMD64

	updateTranslation myPage
	
Darwin

	updateTranslationDarwin myPage
	
###  Building the Binaries ###

	bin/build
	
Output files will be located in dist.
	
###  Deploying to your project. ###

Copy the desired binary and yandex.json file to the root directory of your project.
	
### Output ###

The "myPage" argument lookg for the i18n.json file located in your path/[myPage] directory.  All English keys will be regenerated to all specified languages.
	



