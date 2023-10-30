package gql

import (
	"os"
	"strings"

	"github.com/leliuga/cdk/types"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	schemaInputs = `
"""
The Int64 scalar type represents a signed 64‐bit numeric non‐fractional value.
Int64 can represent values in range [-(2^63),(2^63 - 1)].
"""
scalar Int64

"""
The DateTime scalar type represents date and time as a string in RFC3339 format.
For example: "1985-04-12T23:20:50.52Z" represents 20 mins 50.52 secs after the 23rd hour of Apr 12th 1985 in UTC.
"""
scalar DateTime

input IntRange{
	min: Int!
	max: Int!
}

input FloatRange{
	min: Float!
	max: Float!
}

input Int64Range{
	min: Int64!
	max: Int64!
}

input DateTimeRange{
	min: DateTime!
	max: DateTime!
}

input StringRange{
	min: String!
	max: String!
}

enum DgraphIndex {
	int
	int64
	float
	bool
	hash
	exact
	term
	fulltext
	trigram
	regexp
	year
	month
	day
	hour
	geo
}

input AuthRule {
	and: [AuthRule]
	or: [AuthRule]
	not: AuthRule
	rule: String
}

enum HTTPMethod {
	GET
	POST
	PUT
	PATCH
	DELETE
}

enum Mode {
	BATCH
	SINGLE
}

input CustomHTTP {
	url: String!
	method: HTTPMethod!
	body: String
	graphql: String
	mode: Mode
	forwardHeaders: [String!]
	secretHeaders: [String!]
	introspectionHeaders: [String!]
	skipIntrospection: Boolean
}

type Point {
	longitude: Float!
	latitude: Float!
}

input PointRef {
	longitude: Float!
	latitude: Float!
}

input NearFilter {
	distance: Float!
	coordinate: PointRef!
}

input PointGeoFilter {
	near: NearFilter
	within: WithinFilter
}

type PointList {
	points: [Point!]!
}

input PointListRef {
	points: [PointRef!]!
}

type Polygon {
	coordinates: [PointList!]!
}

input PolygonRef {
	coordinates: [PointListRef!]!
}

type MultiPolygon {
	polygons: [Polygon!]!
}

input MultiPolygonRef {
	polygons: [PolygonRef!]!
}

input WithinFilter {
	polygon: PolygonRef!
}

input ContainsFilter {
	point: PointRef
	polygon: PolygonRef
}

input IntersectsFilter {
	polygon: PolygonRef
	multiPolygon: MultiPolygonRef
}

input PolygonGeoFilter {
	near: NearFilter
	within: WithinFilter
	contains: ContainsFilter
	intersects: IntersectsFilter
}

input GenerateQueryParams {
	get: Boolean
	query: Boolean
	password: Boolean
	aggregate: Boolean
}

input GenerateMutationParams {
	add: Boolean
	update: Boolean
	delete: Boolean
}
`
	directiveDefs = `
directive @hasInverse(field: String!) on FIELD_DEFINITION
directive @search(by: [DgraphIndex!]) on FIELD_DEFINITION
directive @dgraph(type: String, pred: String) on OBJECT | INTERFACE | FIELD_DEFINITION
directive @id(interface: Boolean) on FIELD_DEFINITION
directive @withSubscription on OBJECT | INTERFACE | FIELD_DEFINITION
directive @secret(field: String!, pred: String) on OBJECT | INTERFACE
directive @auth(
	password: AuthRule
	query: AuthRule,
	add: AuthRule,
	update: AuthRule,
	delete: AuthRule) on OBJECT | INTERFACE
directive @custom(http: CustomHTTP, dql: String) on FIELD_DEFINITION
directive @remote on OBJECT | INTERFACE | UNION | INPUT_OBJECT | ENUM
directive @remoteResponse(name: String) on FIELD_DEFINITION
directive @cascade(fields: [String]) on FIELD
directive @lambda on FIELD_DEFINITION
directive @lambdaOnMutate(add: Boolean, update: Boolean, delete: Boolean) on OBJECT | INTERFACE
directive @cacheControl(maxAge: Int!) on QUERY
directive @generate(
	query: GenerateQueryParams,
	mutation: GenerateMutationParams,
	subscription: Boolean) on OBJECT | INTERFACE
`
)

// NewSchema2Go returns a new Schema2Go instance.
func NewSchema2Go(filePath string) (*Schema2Go, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	schema, err := gqlparser.LoadSchema(&ast.Source{
		Name:  filePath,
		Input: schemaInputs + directiveDefs + string(input),
	})

	if err != nil {
		return nil, err
	}

	s := &Schema2Go{Schema: schema, Items: types.Map[Item]{
		"Point": Item{
			Implements:  []byte("import (\n\t\"strconv\"\n)\n\n// NewPoint creates a new Point instance\nfunc NewPoint(longitude, latitude string) *Point {\n\tlong, _ := strconv.ParseFloat(longitude, 64)\n\tlat, _ := strconv.ParseFloat(latitude, 64)\n\n\treturn &Point{\n\t\tType: \"Point\",\n\t\tCoordinates: []float64{long, lat},\n\t}\n}"),
			Definitions: []byte("\tPoint struct {\n\t\tType string `json:\"type\"`\n\t\tCoordinates []float64 `json:\"coordinates\"`\n\t}\n\n"),
		},
		"Type": Item{
			Implements:  []byte("// NewType creates a new Type instance\nfunc NewType(dType string) *Type {\n\treturn &Type{\n\t\tType: dType,\n\t}\n}"),
			Definitions: []byte("\tType struct {\n\t\tType string `json:\"dgraph.type\"`\n\t}\n\n"),
		},
	}}
	s.convert()

	return s, nil
}

// WriteToDir writes all the converted graphql types to the given directory.
func (s *Schema2Go) WriteToDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	parts := strings.Split(dir, "/")
	packageName := []byte("package " + parts[len(parts)-1] + "\n\n")

	for name, item := range s.Items {
		if err := os.WriteFile(dir+"/"+strings.ToLower(name)+".go", append(packageName, item.Implements...), 0644); err != nil {
			return err
		}
	}

	definitions := packageName
	definitions = append(definitions, []byte("type (\n")...)
	for _, item := range s.Items {
		definitions = append(definitions, item.Definitions...)
	}
	definitions = append(definitions, []byte(")")...)

	if err := os.WriteFile(dir+"/types.go", definitions, 0644); err != nil {
		return err
	}

	return nil
}

// convert converts the graphql schema to go types.
func (s *Schema2Go) convert() {
	for _, t := range s.Schema.Types {
		name := strings.ToLower(t.Name)
		if strings.Contains(name, "__") || strings.Contains(name, "point") || strings.Contains(name, "poly") || t.Kind != ast.Object || len(t.Fields) == 0 {
			continue
		}

		s.Items[t.Name] = s.toItem(t)
	}
}

// toItem returns the item for the given graphql type.
func (s *Schema2Go) toItem(t *ast.Definition) Item {
	var implements, definitions []byte

	implements = append(implements, []byte("// New"+t.Name+" creates a new "+t.Name+" instance\n")...)
	implements = append(implements, []byte("func New"+t.Name+"() *"+t.Name+" {\n")...)
	implements = append(implements, []byte("\treturn &"+t.Name+"{\n")...)
	implements = append(implements, []byte("\t\tType: NewType(\""+strings.ToLower(t.Name)+"\"),\n")...)
	implements = append(implements, []byte("\t}\n}")...)

	definitions = append(definitions, []byte("\t// "+t.Description+"\n")...)
	definitions = append(definitions, []byte("\t"+t.Name+" struct {\n")...)
	definitions = append(definitions, []byte("\t\t*Type\n\n")...)
	definitions = append(definitions, s.toStructFields(strings.ToLower(t.Name), t.Fields)...)
	definitions = append(definitions, []byte("\t}\n\n")...)

	return Item{
		Implements:  implements,
		Definitions: definitions,
	}
}

// toStructFields returns the struct fields for the given graphql type.
func (s *Schema2Go) toStructFields(typeName string, fields ast.FieldList) []byte {
	var result []byte
	for _, field := range fields {
		result = append(result, s.toStructField(typeName, field)...)
	}

	return result
}

// toStructField returns the struct field for the given graphql field.
func (s *Schema2Go) toStructField(typeName string, field *ast.FieldDefinition) []byte {
	return []byte("\t\t" + strings.Title(field.Name) + " " + s.toType(field.Type) + " `json:\"" + typeName + "." + field.Name + "\" csv:\"" + field.Name + "\"`// " + field.Description + "\n")
}

// toType returns the go type for the given graphql type.
func (s *Schema2Go) toType(t *ast.Type) string {
	if t.NamedType != "" {
		switch t.NamedType {
		case "Int":
			return "int"
		case "Int64":
			return "int64"
		case "Float":
			return "float64"
		case "Boolean":
			return "bool"
		case "ID", "String", "DateTime":
			return "string"
		case "PointList":
			return "[]*Point"
		default:
			return "*" + t.NamedType
		}
	}

	if t.Elem != nil {
		return "[]" + s.toType(t.Elem)
	}

	return ""
}
