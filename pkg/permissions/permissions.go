package permissions

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

type permissionSet struct {
	Nodes       []*PermissionNode
	rawNodes    []string `toml:"nodes"`
	Groups      []*permissionGroup
	Name        string
	Description string
}

type permissionGroup struct {
	Sets        []*permissionSet
	rawSets     []string `toml:"sets"`
	Name        string
	Description string
	RoleID      string
}

type permissionTree struct {
	Sets   []permissionSet
	Groups []permissionGroup
}

//PermissionNode 
type PermissionNode struct {
	Identifier string
}

func (group *permissionGroup) UnmarshalTOML(i interface{}) error {
	iMap, ok := i.(map[string]interface{})

	if !ok {
		return fmt.Errorf("type assertion error: wants %T, have %T", map[string]interface{}{}, i)
	}

	name, _ := iMap["name"]
	group.Name = name.(string)

	description, _ := iMap["description"]
	group.Description = description.(string)

	role, _ := iMap["role"]
	group.RoleID = role.(string)

	rawSets, _ := iMap["sets"].([]interface{})

	for _, rawSet := range rawSets {
		setString, _ := rawSet.(string)
		group.rawSets = append(group.rawSets, setString)
	}

	return nil
}

func (set *permissionSet) UnmarshalTOML(i interface{}) error {
	iMap, ok := i.(map[string]interface{})

	if !ok {
		return fmt.Errorf("type assertion error: wants %T, have %T", map[string]interface{}{}, i)
	}

	name, _ := iMap["name"]
	set.Name = name.(string)

	description, _ := iMap["description"]
	set.Description = description.(string)

	rawNodes, _ := iMap["nodes"].([]interface{})

	for _, rawNode := range rawNodes {
		nodeString, _ := rawNode.(string)
		set.rawNodes = append(set.rawNodes, nodeString)
	}

	return nil
}

func parsePermissionConfig(tomlBytes []byte) (permissionTree, error) {
	permTree := permissionTree{}

	err := toml.Unmarshal(tomlBytes, &permTree)

	return permTree, err
}
