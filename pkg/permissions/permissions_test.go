package permissions

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
)

func BasePermissionSystem() PermissionSystem {
	// router := cmd.Router{
	// 	Prefix: "~",
	// }

	// router.RegisterCommands(
	// 	&cmd.Command{Name: "test1node1"},
	// 	&cmd.Command{Name: "test2node1"},
	// )

	permSystem := PermissionSystem{
		Tree:  permissionTree{},
		Nodes: make([]*PermissionNode, 0),
	}

	// router.generatePermissionNodes()

	return permSystem

}

const tomlBytesAll = (`
## sets
[[sets]]
	name = "test1set"
	description = "test1set description"
	nodes = [
		"test1node1",
	]
[[sets]]
	name = "test2set"
	description = "test2set description"
	nodes = [
		"test1node1",
		"test2node1",
	]
## groups
[[groups]]
	name = "test1group"
	description = "test1group description"
	role = "1234567890"
	sets = [
		"test1set"
	]
[[groups]]
	name = "test2group"
	description = "test2group description"
	role = "0987654321"
	sets = [
		"test1set",
		"test2set"
	]
`)

func Test_matchNodes(t *testing.T) {
	assert.True(t, matchNodes("test.sub.perm1", "test.*"))
	assert.True(t, matchNodes("test.sub.perm1", "test.sub.*"))
	assert.True(t, matchNodes("test.sub.perm1", "test.sub.perm1"))

	assert.False(t, matchNodes("test.sub.perm1", "test"))
	assert.False(t, matchNodes("test.sub.perm1", "test.sub"))
	assert.False(t, matchNodes("test.sub.perm1", "test.sub.perm2"))

	assert.False(t, matchNodes("test", "test.sub.perm1"))
	assert.False(t, matchNodes("test.sub", "test.sub.perm1"))

	assert.False(t, matchNodes("test.*.test.*", "test.sub.test.sub"))
}

func Test_ParseSet(t *testing.T) {
	bytes := []byte(`
		[[sets]]
			name = "test1"
			description = "test1 description"
			nodes = [
				"test1node1",
			]
		[[sets]]
			name = "test2"
			description = "test2 description"
			nodes = [
				"test1node1",
				"test2node1",
			]
	`)

	actual := []permissionSet{}

	meta := struct {
		Sets []permissionSet
	}{
		actual,
	}

	err := toml.Unmarshal(bytes, &meta)

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, 2, len(meta.Sets))

	set1 := meta.Sets[0]
	set2 := meta.Sets[1]

	assert.Equal(t, "test1", set1.Name)
	assert.Equal(t, "test1 description", set1.Description)
	assert.Equal(t, "test1node1", set1.rawNodes[0])

	assert.Equal(t, "test2", set2.Name)
	assert.Equal(t, "test2 description", set2.Description)
	assert.Equal(t, "test1node1", set2.rawNodes[0])
	assert.Equal(t, "test2node1", set2.rawNodes[1])
}

func Test_ParseGroup(t *testing.T) {
	bytes := []byte(`
		[[groups]]
			name = "test1group"
			description = "test1group description"
			role = "1234567890"
			sets = [
				"test1set"
			]
		[[groups]]
			name = "test2group"
			description = "test2group description"
			role = "0987654321"
			sets = [
				"test1set",
				"test2set"
			]
	`)

	actual := []permissionGroup{}

	meta := struct {
		Groups []permissionGroup
	}{
		actual,
	}

	_ = toml.Unmarshal(bytes, &meta)

	assert.Equal(t, 2, len(meta.Groups))

	group0 := meta.Groups[0]
	group1 := meta.Groups[1]

	assert.Equal(t, "test1group", group0.Name)
	assert.Equal(t, "test1group description", group0.Description)
	assert.Equal(t, "1234567890", group0.RoleID)

	assert.Equal(t, "test1set", group0.rawSets[0])

	assert.Equal(t, "test2group", group1.Name)
	assert.Equal(t, "test2group description", group1.Description)
	assert.Equal(t, "0987654321", group1.RoleID)

	assert.Equal(t, "test1set", group1.rawSets[0])
	assert.Equal(t, "test2set", group1.rawSets[1])
}

func Test_ParsePermissionConfig(t *testing.T) {
	t.Skip() //TODO: Fix this test
	actual, err := parsePermissionConfig([]byte(tomlBytesAll))

	assert.NoError(t, err, "Error occured during unmarshal.")

	permSystem := BasePermissionSystem()

	permSystem.Tree = actual
	permSystem.AddReferences()

	assert.NoError(t, err, "Error occured during unmarshal.")

	assert.Equal(t, 2, len(actual.Sets))
	assert.Equal(t, 2, len(actual.Groups))

	set1 := actual.Sets[0]
	set2 := actual.Sets[1]

	group1 := actual.Groups[0]
	group2 := actual.Groups[1]

	assert.Equal(t, set1.Name, group1.Sets[0].Name)
	assert.Equal(t, set2.Name, group2.Sets[1].Name)

	assert.Equal(t, "test1group", group1.Name)
	assert.Equal(t, "test1group description", group1.Description)
	assert.Equal(t, "1234567890", group1.RoleID)

	assert.Equal(t, set1.Name, group1.Sets[0].Name)
	assert.Equal(t, set1.Description, group1.Sets[0].Description)
	assert.Equal(t, set1.Nodes[0].Identifier, group1.Sets[0].Nodes[0].Identifier)

	assert.Equal(t, "test2group", group2.Name)
	assert.Equal(t, "test2group description", group2.Description)
	assert.Equal(t, "0987654321", group2.RoleID)

	assert.Equal(t, set1.Name, group2.Sets[0].Name)
	assert.Equal(t, set1.Description, group2.Sets[0].Description)
	assert.Equal(t, set1.Nodes[0].Identifier, group2.Sets[0].Nodes[0].Identifier)

	assert.Equal(t, set2.Name, group2.Sets[1].Name)
	assert.Equal(t, set2.Description, group2.Sets[1].Description)
	assert.Equal(t, set1.Nodes[0].Identifier, group2.Sets[0].Nodes[0].Identifier)
	assert.Equal(t, set2.Nodes[0].Identifier, group2.Sets[1].Nodes[0].Identifier)
	assert.Equal(t, set2.Nodes[1].Identifier, group2.Sets[1].Nodes[1].Identifier)
}
