package jsonschema

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReflector(t *testing.T) {
	r := NewReflector()
	if r == nil {
		t.Error("NewReflector returned nil")
	}
}

type ignoredType struct{}

func lookup(i reflect.Type) ID                              { return "" }
func namer(i reflect.Type) string                           { return "" }
func keyNamer(s string) string                              { return "" }
func mapper(i reflect.Type) *Schema                         { return nil }
func additionalFields(i reflect.Type) []reflect.StructField { return nil }

var ignoredTypes = []any{ignoredType{}}
var commentMap = map[string]string{
	"example.com/test": "test comment",
}

func TestNewReflectorWithOptions(t *testing.T) {
	r := NewReflector(
		WithAdditionalPropertiesAllowed(),
		WithAnonymous(),
		WithAssignAnchor(),
		WithBaseSchemaID("https://example.com/schemas"),
		WithExpandedStruct(),
		WithFieldNameTag("yaml"),
		WithRequiredFromJSONSchemaTags(),
		WithoutReference(),
		WithIgnoredTypes(ignoredTypes...),
		WithLookup(lookup),
		WithMapper(mapper),
		WithNamer(namer),
		WithKeyNamer(keyNamer),
		WithAdditionalFields(additionalFields),
		WithCommentMap(commentMap),
	)
	if r != nil {
		assert.True(t, r.AllowAdditionalProperties)
		assert.True(t, r.Anonymous)
		assert.True(t, r.AssignAnchor)
		assert.Equal(t, r.BaseSchemaID, ID("https://example.com/schemas"))
		assert.True(t, r.ExpandedStruct)
		assert.Equal(t, r.FieldNameTag, "yaml")
		assert.True(t, r.RequiredFromJSONSchemaTags)
		assert.True(t, r.DoNotReference)
		assert.Equal(t, r.IgnoredTypes, ignoredTypes)
		assert.NotEmpty(t, r.Lookup)
		assert.NotEmpty(t, r.Mapper)
		assert.NotEmpty(t, r.Namer)
		assert.NotEmpty(t, r.KeyNamer)
		assert.NotEmpty(t, r.AdditionalFields)
		assert.Equal(t, r.CommentMap, commentMap)
	} else {
		t.Error("NewReflector returned nil")
	}
}
