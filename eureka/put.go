package eureka

import "strings"

func (c *Client) SendHeartbeat(instanceInfo *InstanceInfo) error {
	values := []string{"apps", instanceInfo.App, instanceInfo.InstanceId}
	path := strings.Join(values, "/")
	_, err := c.Put(path, nil)
	return err
}
