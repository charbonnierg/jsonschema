package jsonschema

import "reflect"

// Option is a function that can be used to configure a Reflector.
// By convention, options are named starting with "With".
// Each option function returns another function that accepts a Reflector
// pointer as an argument. This allows options to be composed.
type Option func(*Reflector)

// NewReflector is a convenience function to create a new Reflector with
// options. A Reflector reflects values into a Schema.
func NewReflector(opts ...Option) *Reflector {
	r := &Reflector{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// WithBaseSchemaID is an option to define the URI that will be used as a base to determine Schema
// IDs for models. For example, a base Schema ID of `https://invopop.com/schemas`
// when defined with a struct called `User{}`, will result in a schema with an
// ID set to `https://invopop.com/schemas/user`.
//
// If this option is provided, we'll take the type's complete package path
// and use that as a base instead. Set `Anonymous` to try if you do not want to
// include a schema ID.
func WithBaseSchemaID(id ID) Option {
	return func(r *Reflector) {
		r.BaseSchemaID = id
	}
}

// WithAnonymous is an option to hide the auto-generated Schema ID and provide what is
// known as an "anonymous schema". As a rule, this is not recommended.
func WithAnonymous() Option {
	return func(r *Reflector) {
		r.Anonymous = true
	}
}

// WithAssignAnchor is an option to use the original struct's name as an anchor inside
// every definition, including the root schema. These can be useful for having a
// reference to the original struct's name in CamelCase instead of the snake-case used
// by default for URI compatibility.
//
// Anchors do not appear to be widely used out in the wild, so at this time the
// anchors themselves will not be used inside generated schema.
func WithAssignAnchor() Option {
	return func(r *Reflector) {
		r.AssignAnchor = true
	}
}

// WithAdditionalPropertiesAllowed is an option to configure Reflector to generate a schema
// without additionalProperties set to 'false' for all struct types. This means
// the presence of additional keys in JSON objects will not cause validation
// to fail. Note said additional keys will simply be dropped when the
// validated JSON is unmarshaled.
func WithAdditionalPropertiesAllowed() Option {
	return func(r *Reflector) {
		r.AllowAdditionalProperties = true
	}
}

// WithRequiredFromJSONSchemaTags is an option to configure Reflector to generate a schema
// that requires any key tagged with `jsonschema:required`, overriding the
// default of requiring any key *not* tagged with `json:,omitempty`.
func WithRequiredFromJSONSchemaTags() Option {
	return func(r *Reflector) {
		r.RequiredFromJSONSchemaTags = true
	}
}

// WithoutReference is an option to not reference definitions. This will remove the top-level $defs map and
// instead cause the entire structure of types to be output in one tree. The
// list of type definitions (`$defs`) will not be included.
func WithoutReference() Option {
	return func(r *Reflector) {
		r.DoNotReference = true
	}
}

// WithExpandedStruct is an option to include the reflected type's definition in the
// root as opposed to a definition with a reference.
func WithExpandedStruct() Option {
	return func(r *Reflector) {
		r.ExpandedStruct = true
	}
}

// WithFieldNameTag is an option to set the tag used to get field names. json tags are used by default.
func WithFieldNameTag(tag string) Option {
	return func(r *Reflector) {
		r.FieldNameTag = tag
	}
}

// WithIgnoredTypes is an option to define a slice of types that should be ignored in the schema,
// switching to just allowing additional properties instead.
func WithIgnoredTypes(types ...any) Option {
	return func(r *Reflector) {
		r.IgnoredTypes = types
	}
}

// WithLookup is an option to define a function that will provide a custom mapping of
// types to Schema IDs. This allows existing schema documents to be referenced
// by their ID instead of being embedded into the current schema definitions.
// Reflected types will never be pointers, only underlying elements.
func WithLookup(fn func(reflect.Type) ID) Option {
	return func(r *Reflector) {
		r.Lookup = fn
	}
}

// WithMapper an option to define a function that can be used to map custom Go types to jsonschema schemas.
func WithMapper(fn func(reflect.Type) *Schema) Option {
	return func(r *Reflector) {
		r.Mapper = fn
	}
}

// WithNamer is an option to define a function that can customize type names. The default is to use the type's name
// provided by the reflect package.
func WithNamer(fn func(reflect.Type) string) Option {
	return func(r *Reflector) {
		r.Namer = fn
	}
}

// WithKeyNamer is an option to define a function that can customize key names.
// The default is to use the key's name as is, or the json tag if present.
// If a json tag is present, KeyNamer will receive the tag's name as an argument, not the original key name.
func WithKeyNamer(fn func(string) string) Option {
	return func(r *Reflector) {
		r.KeyNamer = fn
	}
}

// WithAdditionalFields is an option to define a function that can add structfields for a given type
func WithAdditionalFields(fn func(reflect.Type) []reflect.StructField) Option {
	return func(r *Reflector) {
		r.AdditionalFields = fn
	}
}

// CommentMap is a function to define a dictionary of fully qualified go types and fields to comment
// strings that will be used if a description has not already been provided in
// the tags. Types and fields are added to the package path using "." as a
// separator.
//
// Type descriptions should be defined like:
//
//	map[string]string{"github.com/invopop/jsonschema.Reflector": "A Reflector reflects values into a Schema."}
//
// And Fields defined as:
//
//	map[string]string{"github.com/invopop/jsonschema.Reflector.DoNotReference": "Do not reference definitions."}
//
// See also: AddGoComments
func WithCommentMap(m map[string]string) Option {
	return func(r *Reflector) {
		r.CommentMap = m
	}
}
