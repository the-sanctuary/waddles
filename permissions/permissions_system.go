package permissions

import (
	"fmt"

	"github.com/the-sanctuary/waddles/command"
)

//PermissionSystem manages all permission checking and storage
type PermissionSystem struct {
	Nodes []*permissionNode
	Tree  permissionTree
}

func CurrentPermissionSystem() *PermissionSystem {
	return &permSystem
}

func NewPermissionSystem(router *command.Router, permissionFile string) *PermissionSystem {
	permSystem = PermissionSystem{
		Tree: permissionTree{
			Sets:   make([]permissionSet, 1),
			Groups: make([]permissionGroup, 1),
		},
	}

	permSystem.generateNodes(router)
	permSystem.loadPermissionConfig(permissionFile)

	return &permSystem
}

func (pm *PermissionSystem) generateNodes(router *command.Router) {
	for _, cmd := range router.Commands {
		pm.generateNodesFromCommand(cmd, "")
	}
}

func (pm *PermissionSystem) generateNodesFromCommand(cmd *command.Command, baseNode string) {
	newBaseNode := baseNode + cmd.Name

	pm.addPermissionNode(newBaseNode)

	if cmd.HasSubcommands() {
		for _, subCmd := range cmd.SubCommands {
			pm.generateNodesFromCommand(subCmd, newBaseNode+".")
		}
	}
}

func (pm *PermissionSystem) addPermissionNode(nodeIdentifier string) {
	node := &permissionNode{Identifier: nodeIdentifier}
	pm.Nodes = append(pm.Nodes, node)
}

func (pm *PermissionSystem) GetNodeFromIdentifier(nodeIdentifier string) *permissionNode {
	for _, node := range pm.Nodes {
		if node.Identifier == nodeIdentifier {
			return node
		}
	}
	return nil
}

func (pm *PermissionSystem) GetSetFromName(setName string) *permissionSet {
	for _, set := range pm.Tree.Sets {
		if set.Name == setName {
			return &set
		}
	}
	fmt.Printf("Couldn't find set: %v in: %+v", setName, pm.Tree.Sets)
	return nil
}

func (pm *PermissionSystem) loadPermissionConfig(permissionFile string) {
	// tree, _ := toml.LoadFile(permissionFile)
	// tree.Unmarshal()
}
