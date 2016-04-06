package cuckle

// Option stores optional query parameters.
type Option map[option]interface{}

// Options with no parameters.
var (
	OptionAllowFiltering  Option = Option{optionAllowFiltering: nil}
	OptionCalled          Option = Option{optionCalled: nil}
	OptionClusteringOrder Option = Option{optionClusteringOrder: nil}
	OptionCompactStorage  Option = Option{optionCompactStorage: nil}
	OptionDistinct        Option = Option{optionDistinct: nil}
	OptionIfExists        Option = Option{optionIfExists: nil}
	OptionIfNotExists     Option = Option{optionIfNotExists: nil}
	OptionIndexKeys       Option = Option{optionIndexKeys: nil}
	OptionJSON            Option = Option{optionJSON: nil}
	OptionReplace         Option = Option{optionReplace: nil}
)

// OptionAliases returns an Option for column aliases.
func OptionAliases(aliases map[Identifier]Identifier) Option {
	return Option{optionAliases: aliases}
}

// OptionAssignments returns an Option for assignments.
func OptionAssignments(r ...Relation) Option {
	return Option{optionAssignments: r}
}

// OptionFinalFunc returns an Option for a final function.
func OptionFinalFunc(finalFunc Identifier) Option {
	return Option{optionFinalFunc: finalFunc}
}

// OptionIf returns an Option for lightweight transaction conditions.
func OptionIf(r ...Relation) Option {
	return Option{optionIf: r}
}

// OptionIndexIdentifier returns an Option for an index identifier.
func OptionIndexIdentifier(index Identifier) Option {
	return Option{optionIndexIdentifier: index}
}

// OptionInitCond returns an Option for a function initial condition.
func OptionInitCond(initCond Term) Option {
	return Option{optionInitCond: initCond}
}

// OptionLimit returns an Option for a result row limit.
func OptionLimit(limit int) Option {
	return Option{optionLimit: limit}
}

// OptionOptions returns an Option for arbitrary Term key-value pairs.
func OptionOptions(options map[Term]Term) Option {
	return Option{optionOptions: options}
}

// OptionOrder returns an Option for ordering result rows by columns and directions.
func OptionOrder(i []Identifier, o []Order) Option {
	return Option{optionOrderByColumns: i, optionOrderByDirections: o}
}

// OptionSelectors returns an Option for selecting column values.
func OptionSelectors(s ...Selector) Option {
	return Option{optionSelectors: s}
}

// OptionTTL returns an Option for insert and update time-to-lives.
func OptionTTL(ttl int64) Option {
	return Option{optionTTL: ttl}
}

// OptionTimestamp returns an Option for insert and update timestamps.
func OptionTimestamp(timestamp int64) Option {
	return Option{optionTimestamp: timestamp}
}

// OptionTriggerIdentifier returns an Option for a trigger identifier.
func OptionTriggerIdentifier(trigger Identifier) Option {
	return Option{optionTriggerIdentifier: trigger}
}

// OptionUsing returns an Option for a custom index class.
func OptionUsing(class string) Option {
	return Option{optionUsing: class}
}

// OptionWhere returns an Option for criteria relations.
func OptionWhere(r ...Relation) Option {
	return Option{optionWhere: r}
}

// OptionWith is key-value pairs.
func OptionWith(options map[Identifier]Term) Option {
	return Option{optionWith: options}
}

func combine(os []Option) Option {
	var combined = Option{}

	for _, o := range os {
		for k, v := range o {
			combined[k] = v
		}
	}

	return combined
}

type option int

const (
	optionAliases option = iota
	optionAllowFiltering
	optionAssignments
	optionCalled
	optionClusteringOrder
	optionCompactStorage
	optionDistinct
	optionFinalFunc
	optionIf
	optionIfExists
	optionIfNotExists
	optionIndexIdentifier
	optionIndexKeys
	optionInitCond
	optionJSON
	optionLimit
	optionOptions
	optionOrderByColumns
	optionOrderByDirections
	optionReplace
	optionSelectors
	optionTTL
	optionTimestamp
	optionTriggerIdentifier
	optionUsing
	optionWhere
	optionWith
)
