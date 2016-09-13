package eureka

import (
	"./../util"
)

var uuid string

func Register() {

	uuid = util.GetUuid()
}
