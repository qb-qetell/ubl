//-- p --
package main

//-- r --
import (
	"github.com/qeetell/songTrack"
	"github.com/qeetell/rsblFetusTrackMngmt"
)
//-- i --
func main () {
	_x5100, _x6100 := rsblFetusTrackMngmt.ManageFetusTrack ([][3]interface {} {
		/*[3]interface {} {"Error Reporter",
			songTrack.Track_Create (f1),
			"a"},*/
		[3]interface {} {"Functionalityy Manager",
			songTrack.Track_Create (fnctnlMngr),
			"10.20"},
		[3]interface {} {"HTTP Interface Manager",
			songTrack.Track_Create (httpIntfcMngr),
			"10.30"},
	}, nil)
	if _x5100 == true {
		return
	}
	rsblFetusTrackMngmt.ManageFetusTrack (nil, _x6100)
}
