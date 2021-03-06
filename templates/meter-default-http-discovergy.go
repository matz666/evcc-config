package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Discovergy (Grid or PV meter/ HTTP)",
		Sample: `power: # power reading
  type: http # use http plugin
  auth:
    type: basic
    user: demo@discovergy.com # Discovergy user name
    password: demo # password 
  uri: https://api.discovergy.com/public/v1/last_reading?meterId=659a3da00324400da66cef81e1cbe3c5 # append meter id
  jq: .values.power
  scale: 0.001`,
	}

	registry.Add(template)
}
