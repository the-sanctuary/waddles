package permissions

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"github.com/the-sanctuary/waddles/command"
)

func Test_generateNodesFromCommand(t *testing.T) {
	testCmdSub1 := &command.Command{
		Name:    "sub1",
		Handler: func(c *command.Context) {},
	}

	testCmdSub21 := &command.Command{
		Name:    "sub21",
		Handler: func(c *command.Context) {},
	}

	testCmdSub2 := &command.Command{
		Name:        "sub2",
		Handler:     func(c *command.Context) {},
		SubCommands: []*command.Command{testCmdSub21},
	}

	testCmd := &command.Command{
		Name:        "testcmd",
		Handler:     func(c *command.Context) {},
		SubCommands: []*command.Command{testCmdSub1, testCmdSub2},
	}

	system := &PermissionSystem{}
	system.generateNodesFromCommand(testCmd, "")

	generatedNodes := system.Nodes

	assert.Equal(t, 4, len(generatedNodes))

	assert.Equal(t, "testcmd", generatedNodes[0].Identifier)
	assert.Equal(t, "testcmd.sub1", generatedNodes[1].Identifier)
	assert.Equal(t, "testcmd.sub2", generatedNodes[2].Identifier)
	assert.Equal(t, "testcmd.sub2.sub21", generatedNodes[3].Identifier)
}

func Test_ParseSet(t *testing.T) {
	router := command.Router{
		Prefix: "~",
	}

	router.RegisterCommands(
		&command.Command{Name: "test1node1"},
		&command.Command{Name: "test2node1"},
	)

	NewPermissionSystem(&router, "")

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
	assert.Equal(t, "test1node1", set1.Nodes[0].Identifier)

	assert.Equal(t, "test2", set2.Name)
	assert.Equal(t, "test2 description", set2.Description)
	assert.Equal(t, "test1node1", set2.Nodes[0].Identifier)
	assert.Equal(t, "test2node1", set2.Nodes[1].Identifier)
}

func Test_ParseGroup(t *testing.T) {
	router := command.Router{
		Prefix: "~",
	}

	router.RegisterCommands(
		&command.Command{Name: "test1node1"},
		&command.Command{Name: "test2node1"},
	)

	ps := NewPermissionSystem(&router, "")

	permSet1 := permissionSet{
		Name:        "test1set",
		Description: "test1set description",
		Nodes:       []*permissionNode{ps.GetNodeFromIdentifier("test1node1")},
	}

	permSet2 := permissionSet{
		Name:        "test2set",
		Description: "test2set description",
		Nodes: []*permissionNode{
			ps.GetNodeFromIdentifier("test1node1"),
			ps.GetNodeFromIdentifier("test2node1"),
		},
	}

	ps.Tree.Sets = append(ps.Tree.Sets, permSet1, permSet2)

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

	assert.Equal(t, permSet1.Name, group0.Sets[0].Name)
	assert.Equal(t, permSet1.Description, group0.Sets[0].Description)
	assert.Equal(t, permSet1.Nodes[0].Identifier, group0.Sets[0].Nodes[0].Identifier)

	assert.Equal(t, "test2group", group1.Name)
	assert.Equal(t, "test2group description", group1.Description)
	assert.Equal(t, "0987654321", group1.RoleID)

	assert.Equal(t, permSet1.Name, group1.Sets[0].Name)
	assert.Equal(t, permSet1.Description, group1.Sets[0].Description)
	assert.Equal(t, permSet1.Nodes[0].Identifier, group1.Sets[0].Nodes[0].Identifier)

	assert.Equal(t, permSet2.Name, group1.Sets[1].Name)
	assert.Equal(t, permSet2.Description, group1.Sets[1].Description)
	assert.Equal(t, permSet2.Nodes[0].Identifier, group1.Sets[1].Nodes[0].Identifier)
	assert.Equal(t, permSet2.Nodes[1].Identifier, group1.Sets[1].Nodes[1].Identifier)
}

func Test_ParsePermissionConfig(t *testing.T) {
	router := command.Router{
		Prefix: "~",
	}

	router.RegisterCommands(
		&command.Command{Name: "test1node1"},
		&command.Command{Name: "test2node1"},
	)

	NewPermissionSystem(&router, "")

	tomlBytes := []byte(`
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

	actual, err := parsePermissionConfig(tomlBytes)

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
