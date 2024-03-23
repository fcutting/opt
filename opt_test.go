package opt_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/fcutting/opt"
	"github.com/gkampitakis/go-snaps/snaps"
)

func TestMain(m *testing.M) {
	r := m.Run()
	snaps.Clean(m, snaps.CleanOpts{Sort: true})
	os.Exit(r)
}

type testPayload struct {
	Primitive *opt.Option[string]         `json:"primitive,omitemty"`
	Map       *opt.Option[map[string]any] `json:"map,omitempty"`
	Struct    *opt.Option[struct {
		Make  string
		Model string
	}] `json:"struct,omitempty"`
	Slice *opt.Option[[]int] `json:"slice,omitempty"`
}

type testCase struct {
	data []byte
}

var testCases = map[string]testCase{
	"Empty": {
		data: []byte(`
			{
				"empty": true
			}
		`),
	},
	"Primitive empty": {
		data: []byte(`
			{
				"primitive": ""
			}
		`),
	},
	"Primitive": {
		data: []byte(`
			{
				"primitive": "hello world"
			}
		`),
	},
	"Primitive null": {
		data: []byte(`
			{
				"primitive": null
			}
		`),
	},
	"Map empty": {
		data: []byte(`
			{
				"map": {}
			}
		`),
	},
	"Map full": {
		data: []byte(`
			{
				"map": {
					"make": "Toyota",
					"model": "Hilux"
				}
			}
		`),
	},
	"Map null": {
		data: []byte(`
			{
				"map": {}
			}
		`),
	},
	"Struct empty": {
		data: []byte(`
			{
				"struct": {}
			}
		`),
	},
	"Struct full": {
		data: []byte(`
			{
				"struct": {
					"make": "Toyota",
					"model": "Hilux"
				}
			}
		`),
	},
	"Struct null": {
		data: []byte(`
			{
				"struct": null
			}
		`),
	},
	"Slice empty": {
		data: []byte(`
			{
				"slice": []
			}
		`),
	},
	"Slice full": {
		data: []byte(`
			{
				"slice": [1, 2, 3]
			}
		`),
	},
	"Slice null": {
		data: []byte(`
			{
				"slice": null
			}
		`),
	},
}

func Test_Option(t *testing.T) {
	for n, c := range testCases {
		t.Run(n, func(t *testing.T) {
			var payload testPayload

			if err := json.Unmarshal(c.data, &payload); err != nil {
				t.Fatalf("Unexpected unmarshal error: %s", err)
			}

			t.Run("MarshalJSON", func(t *testing.T) {
				test_Marshal(t, payload)
			})

			t.Run("Exists", func(t *testing.T) {
				test_Exists(t, payload)
			})
		})
	}
}

func test_Marshal(t *testing.T, payload testPayload) {
	result, err := json.Marshal(payload)

	if err != nil {
		t.Fatalf("Unexpected marshal error: %s", err)
	}

	snaps.MatchSnapshot(t, string(result))
}

func test_Exists(t *testing.T, payload testPayload) {
	t.Run("Primitive", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Primitive.Exists())
	})

	t.Run("Map", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Map.Exists())
	})

	t.Run("Struct", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Struct.Exists())
	})

	t.Run("Slice", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Slice.Exists())
	})
}
