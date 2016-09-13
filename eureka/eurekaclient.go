package eureka

import (
	"goreka/util"
)

var uuid string

func Register() {

	uuid = util.GetUuid()
}

func Unregister() {

}
