package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Sonnenbatterie Eco/10 (Battery/ HTTP)",
		Sample: `power: # power reading
  type: http # use http plugin
  uri: http://192.0.2.2:8080/api/v1/status
  jq: .Pac_total_W
  scale: -1 # reverse direction
soc:
  type: http
  uri: http://192.0.2.2:8080/api/v1/status
  jq: .USOC`,
	}

	registry.Add(template)
}
