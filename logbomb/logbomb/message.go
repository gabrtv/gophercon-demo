package logbomb

import (
	"strings"

	lorem "github.com/drhodes/golorem"
)

var messageTemplate = `{"log": "#placeholder#\n", "stream": "stderr", ` +
	`"docker": {"container_id": "6a9069435788a05531ee2b9afbcdc73a22018af595f3203cb67e06f50103bf5f"}, ` +
	`"kubernetes": {"namespace_name": "foo", "pod_id": "34ebc234-2423-11e6-94aa-42010a800021", ` +
	`"pod_name": "foo-v2-web-2ggow", "container_name": "foo-web", "labels": {"app": "foo", ` +
	`"heritage": "deis", "type": "web", "version": "v2"}, "host":"gke-jchauncey-default-pool-7ae1c279-10ye"}}`

func (lb *LogBomb) getMessage() string {
	randomMessage := lorem.Sentence(lb.config.MinMessageWords, lb.config.MaxMessageWords)
	return strings.Replace(messageTemplate, "#placeholder#", randomMessage, 1)
}
