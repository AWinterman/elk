package spec

const (
	// OpenAPI version 3.0.x is used.
	version = "3.0.3"

	JSON MediaType = "application/json"
)

type (
	Spec struct {
		Info       *Info            `json:"info"`
		Tags       []Tag            `json:"tags,omitempty"`
		Paths      map[string]*Path `json:"paths"`
		Components Components       `json:"components"`
	}
	Tag struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}
	Info struct {
		Title          string  `json:"title"`
		Description    string  `json:"description"`
		TermsOfService string  `json:"termsOfService"`
		Contact        Contact `json:"contact"`
		License        License `json:"license"`
		Version        string  `json:"version"`
	}
	Contact struct {
		Name  string `json:"name,omitempty"`
		Url   string `json:"url,omitempty"`
		Email string `json:"email,omitempty"`
	}
	License struct {
		Name string `json:"name"`
		URL  string `json:"url,omitempty"`
	}
	Path struct {
		Get        *Operation  `json:"get,omitempty"`
		Post       *Operation  `json:"post,omitempty"`
		Delete     *Operation  `json:"delete,omitempty"`
		Patch      *Operation  `json:"patch,omitempty"`
		Parameters []Parameter `json:"parameters,omitempty"`
	}
	Parameter struct {
		Name            string         `json:"name"`
		In              ParameterPlace `json:"in"`
		Description     string         `json:"description,omitempty"`
		Required        bool           `json:"required,omitempty"`
		Deprecated      bool           `json:"deprecated,omitempty"`
		AllowEmptyValue bool           `json:"allowEmptyValue,omitempty"`
		Schema          Type           `json:"schema"`
	}
	Operation struct {
		Summary      string              `json:"summary,omitempty"`
		Description  string              `json:"description,omitempty"`
		Tags         []string            `json:"tags,omitempty"`
		ExternalDocs *ExternalDocs       `json:"externalDocs,omitempty"`
		OperationID  string              `json:"operationId"`
		Parameters   []Parameter         `json:"parameters,omitempty"`
		RequestBody  *RequestBody        `json:"requestBody,omitempty"`
		Responses    map[string]Response `json:"responses"`
		Deprecated   bool                `json:"deprecated,omitempty"`
	}
	ExternalDocs struct {
		Description string `json:"description"`
		URL         string `json:"url"`
	}
	RequestBody struct {
		Description string  `json:"description"`
		Content     Content `json:"content"`
	}
	Content         map[MediaType]MediaTypeObject
	MediaType       string
	MediaTypeObject struct {
		Unique      bool
		Schema      *Schema
		SchemaRef   *Schema
		Response    *Response
		ResponseRef *Response
	}
	Schema struct {
		Name   string
		Fields Fields
		Edges  Edges
	}
	Fields map[string]*Field
	Field  struct {
		Type
		Unique   bool        `json:"-"`
		Required bool        `json:"-"`
		Example  interface{} `json:"example,omitempty"`
	}
	Type struct {
		Type   string `json:"type"`
		Format string `json:"format,omitempty"`
	}
	Edges map[string]Edge
	Edge  struct {
		Schema Schema  `json:"schema"`
		Ref    *Schema `json:"-"`
		Unique bool    `json:"-"`
	}
	Response struct {
		Name        string               `json:"-"`
		Description string               `json:"description"`
		Headers     map[string]Parameter `json:"headers,omitempty"`
		Content     *Content             `json:"content,omitempty"`
	}
	Components struct {
		Schemas    map[string]*Schema   `json:"schemas"`
		Responses  map[string]*Response `json:"responses"`
		Parameters map[string]Parameter `json:"parameters"`
		// ... TODO
	}
)
