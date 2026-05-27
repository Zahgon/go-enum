package generator

// GeneratorConfig holds all the configuration options for the Generator
type GeneratorConfig struct {
	NoPrefix          bool              `json:"no_prefix"`
	NoIota            bool              `json:"no_iota"`
	LowercaseLookup   bool              `json:"lowercase_lookup"`
	CaseInsensitive   bool              `json:"case_insensitive"`
	Marshal           bool              `json:"marshal"`
	SQL               bool              `json:"sql"`
	SQLInt            bool              `json:"sql_int"`
	Flag              bool              `json:"flag"`
	Names             bool              `json:"names"`
	Values            bool              `json:"values"`
	LeaveSnakeCase    bool              `json:"leave_snake_case"`
	JSONPkg           string            `json:"json_pkg"`
	Prefix            string            `json:"prefix"`
	SQLNullInt        bool              `json:"sql_null_int"`
	SQLNullStr        bool              `json:"sql_null_str"`
	Ptr               bool              `json:"ptr"`
	MustParse         bool              `json:"must_parse"`
	ForceLower        bool              `json:"force_lower"`
	ForceUpper        bool              `json:"force_upper"`
	NoComments        bool              `json:"no_comments"`
	NoParse           bool              `json:"no_parse"`
	BuildTags         []string          `json:"build_tags"`
	ReplacementNames  map[string]string `json:"replacement_names"`
	TemplateFileNames []string          `json:"template_file_names"`
}

func NewGeneratorConfig() *GeneratorConfig { _ = "STUB: not implemented"; return nil }

// Option is a function that modifies a Generator
type Option func(*GeneratorConfig)

// WithNoPrefix is used to change the enum const values generated to not have the enum on them.
func WithNoPrefix() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNoIota is used to generate enum constants with explicit values instead of using iota.
func WithNoIota() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLowercaseVariant is used to change the enum const values generated to not have the enum on them.
func WithLowercaseVariant() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCaseInsensitiveParse is used to change the enum const values generated to not have the enum on them.
func WithCaseInsensitiveParse() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMarshal is used to add marshalling to the enum
func WithMarshal() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSQLDriver is used to add marshalling to the enum
func WithSQLDriver() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSQLInt is used to signal a string to be stored as an int.
func WithSQLInt() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithFlag is used to add flag methods to the enum
func WithFlag() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNames is used to add Names methods to the enum
func WithNames() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithValues is used to add Values methods to the enum
func WithValues() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithoutSnakeToCamel is used to add flag methods to the enum
func WithoutSnakeToCamel() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithJsonPkg is used to add a custom json package to the imports
func WithJsonPkg(pkg string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPrefix is used to add a custom prefix to the enum constants
func WithPrefix(prefix string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPtr adds a way to get a pointer value straight from the const value.
func WithPtr() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSQLNullInt is used to add a null int option for SQL interactions.
func WithSQLNullInt() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSQLNullStr is used to add a null string option for SQL interactions.
func WithSQLNullStr() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMustParse is used to add a method `MustParse` that will panic on failure.
func WithMustParse() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithForceLower is used to force enums names to lower case while keeping variable names the same.
func WithForceLower() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithForceUpper is used to force enums names to upper case while keeping variable names the same.
func WithForceUpper() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNoComments is used to remove auto generated comments from the enum.
func WithNoComments() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBuildTags will add build tags to the generated file.
func WithBuildTags(tags ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAliases will set up aliases for the generator.
func WithAliases(aliases map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTemplates is used to provide the filenames of additional templates.
func WithTemplates(filenames ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Note: Template processing is deferred to the generator constructor
// because we need access to the template collection and knownTemplates

// WithNoParse is used to remove the public Parse method from the enum.
func WithNoParse() Option { _ = "STUB: not implemented"; return *new(Option) }
