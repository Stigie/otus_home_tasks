package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

const testDir = "test"

func createTestDir(t *testing.T) {
	err := os.Mkdir(testDir, os.ModePerm)
	require.NoError(t, err)
}

func createTestFile(name, value string, t *testing.T) {
	file, err := os.Create(path.Join(testDir, name))
	defer file.Close()
	require.NoError(t, err)
	_, err = file.WriteString(value)
	require.NoError(t, err)
}

func deleteTestDir(t *testing.T) {
	err := os.RemoveAll(testDir)
	require.NoError(t, err)
}

func TestReadDir(t *testing.T) {
	t.Run("read some variants", func(t *testing.T) {
		createTestDir(t)

		createTestFile("HELLO", `"hello"`, t)
		createTestFile("BAR", "bar\nPLEASE IGNORE SECOND LINE", t)
		createTestFile("FOO", "   foo\u0000with new line", t)
		createTestFile("UNSET", "", t)
		createTestFile("WITH_SPACE_IN_THE_END", "VALUE ", t)
		createTestFile("WITH_TAB_IN_THE_END", "VALUE \t", t)

		result, err := ReadDir(testDir)

		require.Nil(t, err)
		require.Equal(t, Environment{
			"HELLO":                 `"hello"`,
			"BAR":                   "bar",
			"FOO":                   "   foo\nwith new line",
			"UNSET":                 "",
			"WITH_SPACE_IN_THE_END": "VALUE",
			"WITH_TAB_IN_THE_END":   "VALUE",
		}, result)

		deleteTestDir(t)
		createTestDir(t)
		createTestFile("WITH_=_IN_NAME", "VALUE", t)
		defer deleteTestDir(t)
		result, err = ReadDir(testDir)

		require.Equal(t, ErrIncorrectEnvName, err)
	})
}
