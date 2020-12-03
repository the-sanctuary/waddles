package permissions

import (
	"fmt"
	"io/ioutil"

	"github.com/rs/zerolog/log"

	"github.com/the-sanctuary/waddles/command"
	"github.com/the-sanctuary/waddles/util"
)

//PermissionSystem manages all permission checking and storage
type PermissionSystem struct {
	Nodes []*permissionNode
	Tree  permissionTree
}

func BuildPermissionSystem(router *command.Router, permissionFile string) *PermissionSystem {
	permTree := loadPermissionConfig(permissionFile)

	permSystem := PermissionSystem{
		Tree:  permTree,
		Nodes: make([]*permissionNode, 1),
	}

	permSystem.generateNodes(router)
	permSystem.AddReferences()

	return &permSystem
}

func (pm *PermissionSystem) AddReferences() {
	tree := &pm.Tree

	sets := tree.Sets
	groups := tree.Groups

	for i := range sets {
		nodes := 0
		for _, rawNode := range sets[i].rawNodes {
			node := pm.getNodeFromIdentifier(rawNode)
			sets[i].Nodes = append(sets[i].Nodes, node)
			nodes++
		}
		fmt.Printf("Found %d nodes for set: %s\n", nodes, sets[i].Name)
	}

	for i := range groups {
		sets := 0
		nodes := 0
		for _, rawSet := range groups[i].rawSets {
			set := pm.getSetFromName(rawSet)
			groups[i].Sets = append(groups[i].Sets, set)
			sets++
			nodes += len(set.Nodes)
		}
		fmt.Printf("Found %d sets for group: %s (with %d nodes)\n", sets, groups[i].Name, nodes)
	}
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

//getNodeFromIdentifier Returns a pointer to the permissions.permissionNode{} the setName represents
func (pm *PermissionSystem) getNodeFromIdentifier(nodeIdentifier string) *permissionNode {
	for _, node := range pm.Nodes {
		if node.Identifier == nodeIdentifier {
			return node
		}
	}
	return nil
}

//getSetFromName Returns a pointer to the permissions.permissionSet{} the setName represents
func (pm *PermissionSystem) getSetFromName(setName string) *permissionSet {
	for _, set := range pm.Tree.Sets {
		if set.Name == setName {
			return &set
		}
	}
	fmt.Printf("Couldn't find set: %v in: %+v", setName, pm.Tree.Sets)
	return nil
}

func loadPermissionConfig(permissionFile string) permissionTree {
	tomlBytes, err := ioutil.ReadFile(permissionFile)

	if util.DebugError(err) {
		log.Fatal().Err(err).Msg("An error occured while reading config file.")
		return permissionTree{}
	}

	permTree, err := parsePermissionConfig(tomlBytes)
	if util.DebugError(err) {
		log.Fatal().Err(err).Msg("An error while parsing the config file.")
		return permissionTree{}
	}

	return permTree
}
