package opt_test

import (
	"encoding/json"
	"fmt"
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
	Primitive opt.Option[string]         `json:"primitive"`
	Map       opt.Option[map[string]any] `json:"map"`
	Struct    opt.Option[testStruct]     `json:"struct"`
	Slice     opt.Option[[]int]          `json:"slice"`
	Pointer   opt.Option[*bool]          `json:"pointer"`
}

type testStruct struct {
	Make  string
	Model string
}

type testCase struct {
	data []byte
}

var (
	primitiveDefault = "default"
	mapDefault       = map[string]any{
		"make":  "Ford",
		"model": "Focus",
	}
	structDefault = testStruct{
		Make:  "Audi",
		Model: "A5",
	}
	sliceDefault   = []int{9, 8, 7}
	pointerDefault = true

	testCases = map[string]testCase{
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
				"map": null
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
		"Pointer full": {
			data: []byte(`
			{
				"pointer": true
			}
		`),
		},
		"Pointer null": {
			data: []byte(`
			{
				"pointer": null
			}
		`),
		},
	}
)

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

			t.Run("String", func(t *testing.T) {
				test_String(t, payload)
			})

			t.Run("Exists", func(t *testing.T) {
				test_Exists(t, payload)
			})

			t.Run("Unwrap", func(t *testing.T) {
				test_Unwrap(t, payload)
			})

			t.Run("MustUnwrap", func(t *testing.T) {
				test_Unwrap(t, payload)
			})

			t.Run("UnwrapDefault", func(t *testing.T) {
				test_UnwrapDefault(t, payload)
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

	t.Run("Pointer", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Pointer.Exists())
	})
}

func test_Unwrap(t *testing.T, payload testPayload) {
	t.Run("Primitive", func(t *testing.T) {
		snaps.MatchSnapshot(t, payload.Primitive.Unwrap())
	})

	t.Run("Map", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Map.Unwrap())
	})

	t.Run("Struct", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Struct.Unwrap())
	})

	t.Run("Slice", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Slice.Unwrap())
	})

	t.Run("Pointer", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Pointer.Unwrap())
	})
}

func test_MustUnwrap(t *testing.T, payload testPayload) {
	t.Run("Primitive", func(t *testing.T) {
		snaps.MatchSnapshot(t, payload.Primitive.MustUnwrap())
	})

	t.Run("Map", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Map.MustUnwrap())
	})

	t.Run("Struct", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Struct.MustUnwrap())
	})

	t.Run("Slice", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Slice.MustUnwrap())
	})

	t.Run("Pointer", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Pointer.MustUnwrap())
	})
}

func test_UnwrapDefault(t *testing.T, payload testPayload) {
	t.Run("Primitive", func(t *testing.T) {
		snaps.MatchSnapshot(t, payload.Primitive.UnwrapDefault(primitiveDefault))
	})

	t.Run("Map", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Map.UnwrapDefault(mapDefault))
	})

	t.Run("Struct", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Struct.UnwrapDefault(structDefault))
	})

	t.Run("Slice", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Slice.UnwrapDefault(sliceDefault))
	})

	t.Run("Pointer", func(t *testing.T) {
		snaps.MatchJSON(t, payload.Pointer.UnwrapDefault(&pointerDefault))
	})
}

func test_String(t *testing.T, payload testPayload) {
	t.Run("Primitive", func(t *testing.T) {
		snaps.MatchSnapshot(t, fmt.Sprint(payload.Primitive))
	})

	t.Run("Map", func(t *testing.T) {
		snaps.MatchSnapshot(t, fmt.Sprint(payload.Map))
	})

	t.Run("Struct", func(t *testing.T) {
		snaps.MatchSnapshot(t, fmt.Sprint(payload.Struct))
	})

	t.Run("Slice", func(t *testing.T) {
		snaps.MatchSnapshot(t, fmt.Sprint(payload.Slice))
	})
}
