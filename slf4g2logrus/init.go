package slf4g2logrus

import (
	log "github.com/echocat/slf4g"
)

func init() {
	log.RegisterProvider(DefaultProvider)
}
