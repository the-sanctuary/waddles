package permissions

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/the-sanctuary/waddles/pkg/util"
)

//PermissionSystem manages all permission checking and storage
type PermissionSystem struct {
	Nodes []*permissionNode
	Tree  permissionTree
}

//BuildPermissionSystem returns a to a loaded PermissionSystem based on the given file
func BuildPermissionSystem(permissionFile string) PermissionSystem {
	permTree := loadPermissionConfig(permissionFile)

	permSystem := PermissionSystem{
		Tree:  permTree,
		Nodes: make([]*permissionNode, 0),
	}

	return permSystem
}

//AddReferences loops over all sets and groups in the PermissionSystem, correctly populating any lists of associations based on the raw* field
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
		log.Trace().Msgf("Found %d nodes for set: %s\n", nodes, sets[i].Name)
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
		log.Trace().Msgf("Found %d sets for group: %s (with %d nodes)\n", sets, groups[i].Name, nodes)
	}
}

//AddPermissionNode creates a new permissionNode{} with the given identifier and appends it to the list of Nodes in this PermissionSystem
func (pm *PermissionSystem) AddPermissionNode(nodeIdentifier string) {
	node := &permissionNode{Identifier: nodeIdentifier}
	pm.Nodes = append(pm.Nodes, node)
	log.Debug().Msgf("Added permission node: %s", node.Identifier)
}

//UserHasPermissionNode will return whether or not the given Member has the given nodeIdentifier or not
func (pm *PermissionSystem) UserHasPermissionNode(member *discordgo.Member, nodeIdentifer string) bool {
	memberNodes := pm.getMemberNodes(member)

	var nodeidents []string

	for _, mn := range memberNodes {
		nodeidents = append(nodeidents, mn)
		log.Trace().Msgf("mn:%s | node:%s", mn, nodeIdentifer)
		if matchNodes(mn, nodeIdentifer) {
			return true
		}
	}

	log.Trace().Msgf("User: `%s` has perms: %+v", member.User.Username, nodeidents)

	return false
}

func matchNodes(actualNode string, testNode string) bool {
	if actualNode == testNode {
		return true
	}

	if strings.Contains(testNode, "*") {
		wildcard := strings.TrimSuffix(testNode, ".*")
		fmt.Println(wildcard)
		if strings.HasPrefix(actualNode, wildcard) {
			return true
		}
	}

	return false
}

func (pm *PermissionSystem) getMemberNodes(member *discordgo.Member) []string {
	var nodeSet map[string]interface{} = make(map[string]interface{})

	for _, group := range pm.Tree.Groups {
		if util.SliceContains(member.Roles, group.RoleID) {
			for _, set := range group.Sets {
				for _, node := range set.Nodes {
					nodeSet[node.Identifier] = struct{}{}
				}
			}
		}
	}

	var memberNodes []string

	for node := range nodeSet {
		memberNodes = append(memberNodes, node)
	}

	return memberNodes
}

//getNodeFromIdentifier Returns a pointer to the permissions.permissionNode{} the nodeIdentifier represents
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
