package eureka

import "strings"

func (c *Client) UnregisterInstance(instanceInfo *InstanceInfo) error {
	values := []string{"apps", instanceInfo.App, instanceInfo.InstanceId}
	path := strings.Join(values, "/")
	_, err := c.Delete(path)
	return err
}
