package eureka

import (
	"encoding/json"
	"strings"
)

func (c *Client) RegisterInstance(instanceInfo *InstanceInfo) error {
	values := []string{"apps", instanceInfo.App}
	path := strings.Join(values, "/")
	instance := &Instance{
		Instance: instanceInfo,
	}
	body, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	_, err = c.Post(path, body)
	return err
}
