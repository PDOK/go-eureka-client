package eureka

import "strings"

func (c *Client) SendHeartbeat(instanceInfo *InstanceInfo) bool {
	values := []string{"apps", instanceInfo.App, instanceInfo.InstanceId}
	path := strings.Join(values, "/")
	resp, _ := c.Put(path, nil)
	if resp != nil {
		return resp.StatusCode == 200
	}
	return false
}
