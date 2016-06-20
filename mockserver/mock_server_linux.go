package mockserver

/*
#cgo LDFLAGS: ${SRCDIR}/../libs/libpact_mock_server.dylib

// Library headers
int create_mock_server(char* pact, int port);
*/
import "C"
import (
	"fmt"
	"time"
)

func CreateMockServer() {
	res := C.create_mock_server(C.CString(`
		{
		  "consumer": {
		    "name": "billy"
		  },
		  "provider": {
		    "name": "bobby"
		  },
		  "interactions": [

		  ],
		  "metadata": {
		    "pactSpecificationVersion": "2.0.0"
		  }
		}
	`), 0)
	fmt.Println("Mock Server running on port:", res)
	<-time.After(3000 * time.Second)
}
