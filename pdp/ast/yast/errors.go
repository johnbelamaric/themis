package yast

/* AUTOMATICALLY GENERATED FROM errors.yaml - DO NOT EDIT */

import (
	"fmt"
	"github.com/infobloxopen/themis/pdp"
	"strings"
)

const (
	externalErrorID                       = 0
	rootKeysErrorID                       = 1
	stringErrorID                         = 2
	missingStringErrorID                  = 3
	mapErrorID                            = 4
	missingMapErrorID                     = 5
	listErrorID                           = 6
	missingListErrorID                    = 7
	attributeTypeErrorID                  = 8
	policyAmbiguityErrorID                = 9
	policyMissingKeyErrorID               = 10
	unknownRCAErrorID                     = 11
	missingRCAErrorID                     = 12
	invalidRCAErrorID                     = 13
	missingMapRCAParamErrorID             = 14
	missingDefaultRuleRCAErrorID          = 15
	missingErrorRuleRCAErrorID            = 16
	unknownMapperPCAOrderID               = 17
	unknownMapperRCAOrderID               = 18
	notImplementedRCAErrorID              = 19
	unknownPCAErrorID                     = 20
	missingPCAErrorID                     = 21
	invalidPCAErrorID                     = 22
	missingMapPCAParamErrorID             = 23
	missingDefaultPolicyPCAErrorID        = 24
	missingErrorPolicyPCAErrorID          = 25
	notImplementedPCAErrorID              = 26
	mapperArgumentTypeErrorID             = 27
	conditionTypeErrorID                  = 28
	unknownEffectErrorID                  = 29
	noSMPItemsErrorID                     = 30
	tooManySMPItemsErrorID                = 31
	unknownMatchFunctionErrorID           = 32
	matchFunctionCastErrorID              = 33
	matchFunctionArgsNumberErrorID        = 34
	invalidMatchFunctionArgErrorID        = 35
	matchFunctionBothValuesErrorID        = 36
	matchFunctionBothAttrsErrorID         = 37
	unknownFunctionErrorID                = 38
	functionCastErrorID                   = 39
	unknownAttributeErrorID               = 40
	unknownTypeErrorID                    = 41
	invalidTypeErrorID                    = 42
	missingContentErrorID                 = 43
	notImplementedValueTypeErrorID        = 44
	invalidAddressErrorID                 = 45
	invalidNetworkErrorID                 = 46
	invalidDomainErrorID                  = 47
	selectorURIErrorID                    = 48
	selectorLocationErrorID               = 49
	unsupportedSelectorSchemeErrorID      = 50
	entityAmbiguityErrorID                = 51
	entityMissingKeyErrorID               = 52
	unknownPolicyUpdateOperationErrorID   = 53
	invalidPolicyUpdatePathElementErrorID = 54
)

type externalError struct {
	errorLink
	err error
}

func newExternalError(err error) *externalError {
	return &externalError{
		errorLink: errorLink{id: externalErrorID},
		err:       err}
}

func (e *externalError) Error() string {
	return e.errorf("%s", e.err)
}

type rootKeysError struct {
	errorLink
	keys map[interface{}]interface{}
}

func newRootKeysError(keys map[interface{}]interface{}) *rootKeysError {
	return &rootKeysError{
		errorLink: errorLink{id: rootKeysErrorID},
		keys:      keys}
}

func (e *rootKeysError) Error() string {
	keys := make([]string, len(e.keys))
	i := 0
	for key := range e.keys {
		if s, ok := key.(string); ok {
			keys[i] = fmt.Sprintf("%q", s)
		} else {
			keys[i] = fmt.Sprintf("%v", key)
		}
		i++
	}
	s := strings.Join(keys, ", ")

	return e.errorf("Expected attribute definitions and policies but got: %s", s)
}

type stringError struct {
	errorLink
	v    interface{}
	desc string
}

func newStringError(v interface{}, desc string) *stringError {
	return &stringError{
		errorLink: errorLink{id: stringErrorID},
		v:         v,
		desc:      desc}
}

func (e *stringError) Error() string {
	return e.errorf("Expected %s but got %T", e.desc, e.v)
}

type missingStringError struct {
	errorLink
	desc string
}

func newMissingStringError(desc string) *missingStringError {
	return &missingStringError{
		errorLink: errorLink{id: missingStringErrorID},
		desc:      desc}
}

func (e *missingStringError) Error() string {
	return e.errorf("Missing %s", e.desc)
}

type mapError struct {
	errorLink
	v    interface{}
	desc string
}

func newMapError(v interface{}, desc string) *mapError {
	return &mapError{
		errorLink: errorLink{id: mapErrorID},
		v:         v,
		desc:      desc}
}

func (e *mapError) Error() string {
	return e.errorf("Expected %s but got %T", e.desc, e.v)
}

type missingMapError struct {
	errorLink
	desc string
}

func newMissingMapError(desc string) *missingMapError {
	return &missingMapError{
		errorLink: errorLink{id: missingMapErrorID},
		desc:      desc}
}

func (e *missingMapError) Error() string {
	return e.errorf("Missing %s", e.desc)
}

type listError struct {
	errorLink
	v    interface{}
	desc string
}

func newListError(v interface{}, desc string) *listError {
	return &listError{
		errorLink: errorLink{id: listErrorID},
		v:         v,
		desc:      desc}
}

func (e *listError) Error() string {
	return e.errorf("Expected %s but got %T", e.desc, e.v)
}

type missingListError struct {
	errorLink
	desc string
}

func newMissingListError(desc string) *missingListError {
	return &missingListError{
		errorLink: errorLink{id: missingListErrorID},
		desc:      desc}
}

func (e *missingListError) Error() string {
	return e.errorf("Missing %s", e.desc)
}

type attributeTypeError struct {
	errorLink
	t string
}

func newAttributeTypeError(t string) *attributeTypeError {
	return &attributeTypeError{
		errorLink: errorLink{id: attributeTypeErrorID},
		t:         t}
}

func (e *attributeTypeError) Error() string {
	return e.errorf("Expected attribute data type but got \"%s\"", e.t)
}

type policyAmbiguityError struct {
	errorLink
}

func newPolicyAmbiguityError() *policyAmbiguityError {
	return &policyAmbiguityError{
		errorLink: errorLink{id: policyAmbiguityErrorID}}
}

func (e *policyAmbiguityError) Error() string {
	return e.errorf("Expected rules (for policy) or policies (for policy set) but got both")
}

type policyMissingKeyError struct {
	errorLink
}

func newPolicyMissingKeyError() *policyMissingKeyError {
	return &policyMissingKeyError{
		errorLink: errorLink{id: policyMissingKeyErrorID}}
}

func (e *policyMissingKeyError) Error() string {
	return e.errorf("Expected rules (for policy) or policies (for policy set) but got nothing")
}

type unknownRCAError struct {
	errorLink
	alg string
}

func newUnknownRCAError(alg string) *unknownRCAError {
	return &unknownRCAError{
		errorLink: errorLink{id: unknownRCAErrorID},
		alg:       alg}
}

func (e *unknownRCAError) Error() string {
	return e.errorf("Unknown rule combinig algorithm \"%s\"", e.alg)
}

type missingRCAError struct {
	errorLink
}

func newMissingRCAError() *missingRCAError {
	return &missingRCAError{
		errorLink: errorLink{id: missingRCAErrorID}}
}

func (e *missingRCAError) Error() string {
	return e.errorf("Missing policy combinig algorithm")
}

type invalidRCAError struct {
	errorLink
	v interface{}
}

func newInvalidRCAError(v interface{}) *invalidRCAError {
	return &invalidRCAError{
		errorLink: errorLink{id: invalidRCAErrorID},
		v:         v}
}

func (e *invalidRCAError) Error() string {
	return e.errorf("Expected string or map as policy combinig algorithm but got %T", e.v)
}

type missingMapRCAParamError struct {
	errorLink
}

func newMissingMapRCAParamError() *missingMapRCAParamError {
	return &missingMapRCAParamError{
		errorLink: errorLink{id: missingMapRCAParamErrorID}}
}

func (e *missingMapRCAParamError) Error() string {
	return e.errorf("Missing map parameter")
}

type missingDefaultRuleRCAError struct {
	errorLink
	ID string
}

func newMissingDefaultRuleRCAError(ID string) *missingDefaultRuleRCAError {
	return &missingDefaultRuleRCAError{
		errorLink: errorLink{id: missingDefaultRuleRCAErrorID},
		ID:        ID}
}

func (e *missingDefaultRuleRCAError) Error() string {
	return e.errorf("No rule with ID %q to use as default rule", e.ID)
}

type missingErrorRuleRCAError struct {
	errorLink
	ID string
}

func newMissingErrorRuleRCAError(ID string) *missingErrorRuleRCAError {
	return &missingErrorRuleRCAError{
		errorLink: errorLink{id: missingErrorRuleRCAErrorID},
		ID:        ID}
}

func (e *missingErrorRuleRCAError) Error() string {
	return e.errorf("No rule with ID %q to use as on error rule", e.ID)
}

type unknownMapperPCAOrder struct {
	errorLink
	ord string
}

func newUnknownMapperPCAOrder(ord string) *unknownMapperPCAOrder {
	return &unknownMapperPCAOrder{
		errorLink: errorLink{id: unknownMapperPCAOrderID},
		ord:       ord}
}

func (e *unknownMapperPCAOrder) Error() string {
	return e.errorf("Unknown policy ordering for mapper \"%s\"", e.ord)
}

type unknownMapperRCAOrder struct {
	errorLink
	ord string
}

func newUnknownMapperRCAOrder(ord string) *unknownMapperRCAOrder {
	return &unknownMapperRCAOrder{
		errorLink: errorLink{id: unknownMapperRCAOrderID},
		ord:       ord}
}

func (e *unknownMapperRCAOrder) Error() string {
	return e.errorf("Unknown rule ordering for mapper \"%s\"", e.ord)
}

type notImplementedRCAError struct {
	errorLink
	ID string
}

func newNotImplementedRCAError(ID string) *notImplementedRCAError {
	return &notImplementedRCAError{
		errorLink: errorLink{id: notImplementedRCAErrorID},
		ID:        ID}
}

func (e *notImplementedRCAError) Error() string {
	return e.errorf("Parsing for %q rule combinig algorithm hasn't been implemented yet", e.ID)
}

type unknownPCAError struct {
	errorLink
	alg string
}

func newUnknownPCAError(alg string) *unknownPCAError {
	return &unknownPCAError{
		errorLink: errorLink{id: unknownPCAErrorID},
		alg:       alg}
}

func (e *unknownPCAError) Error() string {
	return e.errorf("Unknown policy combinig algorithm \"%s\"", e.alg)
}

type missingPCAError struct {
	errorLink
}

func newMissingPCAError() *missingPCAError {
	return &missingPCAError{
		errorLink: errorLink{id: missingPCAErrorID}}
}

func (e *missingPCAError) Error() string {
	return e.errorf("Missing policy combinig algorithm")
}

type invalidPCAError struct {
	errorLink
	v interface{}
}

func newInvalidPCAError(v interface{}) *invalidPCAError {
	return &invalidPCAError{
		errorLink: errorLink{id: invalidPCAErrorID},
		v:         v}
}

func (e *invalidPCAError) Error() string {
	return e.errorf("Expected string or map as policy combinig algorithm but got %T", e.v)
}

type missingMapPCAParamError struct {
	errorLink
}

func newMissingMapPCAParamError() *missingMapPCAParamError {
	return &missingMapPCAParamError{
		errorLink: errorLink{id: missingMapPCAParamErrorID}}
}

func (e *missingMapPCAParamError) Error() string {
	return e.errorf("Missing map parameter")
}

type missingDefaultPolicyPCAError struct {
	errorLink
	ID string
}

func newMissingDefaultPolicyPCAError(ID string) *missingDefaultPolicyPCAError {
	return &missingDefaultPolicyPCAError{
		errorLink: errorLink{id: missingDefaultPolicyPCAErrorID},
		ID:        ID}
}

func (e *missingDefaultPolicyPCAError) Error() string {
	return e.errorf("No policy with ID %q to use as default policy", e.ID)
}

type missingErrorPolicyPCAError struct {
	errorLink
	ID string
}

func newMissingErrorPolicyPCAError(ID string) *missingErrorPolicyPCAError {
	return &missingErrorPolicyPCAError{
		errorLink: errorLink{id: missingErrorPolicyPCAErrorID},
		ID:        ID}
}

func (e *missingErrorPolicyPCAError) Error() string {
	return e.errorf("No policy with ID %q to use as on error policy", e.ID)
}

type notImplementedPCAError struct {
	errorLink
	ID string
}

func newNotImplementedPCAError(ID string) *notImplementedPCAError {
	return &notImplementedPCAError{
		errorLink: errorLink{id: notImplementedPCAErrorID},
		ID:        ID}
}

func (e *notImplementedPCAError) Error() string {
	return e.errorf("Parsing for %q policy combinig algorithm hasn't been implemented yet", e.ID)
}

type mapperArgumentTypeError struct {
	errorLink
	actual int
}

func newMapperArgumentTypeError(actual int) *mapperArgumentTypeError {
	return &mapperArgumentTypeError{
		errorLink: errorLink{id: mapperArgumentTypeErrorID},
		actual:    actual}
}

func (e *mapperArgumentTypeError) Error() string {
	return e.errorf("Expected %s, %s or %s as argument but got %s", pdp.TypeNames[pdp.TypeString], pdp.TypeNames[pdp.TypeSetOfStrings], pdp.TypeNames[pdp.TypeListOfStrings], pdp.TypeNames[e.actual])
}

type conditionTypeError struct {
	errorLink
	t int
}

func newConditionTypeError(t int) *conditionTypeError {
	return &conditionTypeError{
		errorLink: errorLink{id: conditionTypeErrorID},
		t:         t}
}

func (e *conditionTypeError) Error() string {
	return e.errorf("Expected %q as condition expression result but got %q", pdp.TypeNames[pdp.TypeBoolean], pdp.TypeNames[e.t])
}

type unknownEffectError struct {
	errorLink
	e string
}

func newUnknownEffectError(e string) *unknownEffectError {
	return &unknownEffectError{
		errorLink: errorLink{id: unknownEffectErrorID},
		e:         e}
}

func (e *unknownEffectError) Error() string {
	return e.errorf("Unknown rule effect %q", e.e)
}

type noSMPItemsError struct {
	errorLink
	desc string
	n    int
}

func newNoSMPItemsError(desc string, n int) *noSMPItemsError {
	return &noSMPItemsError{
		errorLink: errorLink{id: noSMPItemsErrorID},
		desc:      desc,
		n:         n}
}

func (e *noSMPItemsError) Error() string {
	return e.errorf("Expected at least one entry in %s got %d", e.desc, e.n)
}

type tooManySMPItemsError struct {
	errorLink
	desc string
	n    int
}

func newTooManySMPItemsError(desc string, n int) *tooManySMPItemsError {
	return &tooManySMPItemsError{
		errorLink: errorLink{id: tooManySMPItemsErrorID},
		desc:      desc,
		n:         n}
}

func (e *tooManySMPItemsError) Error() string {
	return e.errorf("Expected only one entry in %s got %d", e.desc, e.n)
}

type unknownMatchFunctionError struct {
	errorLink
	ID string
}

func newUnknownMatchFunctionError(ID string) *unknownMatchFunctionError {
	return &unknownMatchFunctionError{
		errorLink: errorLink{id: unknownMatchFunctionErrorID},
		ID:        ID}
}

func (e *unknownMatchFunctionError) Error() string {
	return e.errorf("Unknown match function %q", e.ID)
}

type matchFunctionCastError struct {
	errorLink
	ID     string
	first  int
	second int
}

func newMatchFunctionCastError(ID string, first, second int) *matchFunctionCastError {
	return &matchFunctionCastError{
		errorLink: errorLink{id: matchFunctionCastErrorID},
		ID:        ID,
		first:     first,
		second:    second}
}

func (e *matchFunctionCastError) Error() string {
	return e.errorf("No function %s for arguments %s and %s", e.ID, pdp.TypeNames[e.first], pdp.TypeNames[e.second])
}

type matchFunctionArgsNumberError struct {
	errorLink
	n int
}

func newMatchFunctionArgsNumberError(n int) *matchFunctionArgsNumberError {
	return &matchFunctionArgsNumberError{
		errorLink: errorLink{id: matchFunctionArgsNumberErrorID},
		n:         n}
}

func (e *matchFunctionArgsNumberError) Error() string {
	return e.errorf("Expected two arguments got %d", e.n)
}

type invalidMatchFunctionArgError struct {
	errorLink
	expr pdp.Expression
}

func newInvalidMatchFunctionArgError(expr pdp.Expression) *invalidMatchFunctionArgError {
	return &invalidMatchFunctionArgError{
		errorLink: errorLink{id: invalidMatchFunctionArgErrorID},
		expr:      expr}
}

func (e *invalidMatchFunctionArgError) Error() string {
	return e.errorf("Expected one immediate value and one attribute got %T", e.expr)
}

type matchFunctionBothValuesError struct {
	errorLink
}

func newMatchFunctionBothValuesError() *matchFunctionBothValuesError {
	return &matchFunctionBothValuesError{
		errorLink: errorLink{id: matchFunctionBothValuesErrorID}}
}

func (e *matchFunctionBothValuesError) Error() string {
	return e.errorf("Expected one immediate value and one attribute got both immediate values")
}

type matchFunctionBothAttrsError struct {
	errorLink
}

func newMatchFunctionBothAttrsError() *matchFunctionBothAttrsError {
	return &matchFunctionBothAttrsError{
		errorLink: errorLink{id: matchFunctionBothAttrsErrorID}}
}

func (e *matchFunctionBothAttrsError) Error() string {
	return e.errorf("Expected one immediate value and one attribute got both immediate values")
}

type unknownFunctionError struct {
	errorLink
	ID string
}

func newUnknownFunctionError(ID string) *unknownFunctionError {
	return &unknownFunctionError{
		errorLink: errorLink{id: unknownFunctionErrorID},
		ID:        ID}
}

func (e *unknownFunctionError) Error() string {
	return e.errorf("Unknown function %q", e.ID)
}

type functionCastError struct {
	errorLink
	ID    string
	exprs []pdp.Expression
}

func newFunctionCastError(ID string, exprs []pdp.Expression) *functionCastError {
	return &functionCastError{
		errorLink: errorLink{id: functionCastErrorID},
		ID:        ID,
		exprs:     exprs}
}

func (e *functionCastError) Error() string {
	args := ""
	if len(e.exprs) > 1 {
		t := make([]string, len(e.exprs))
		for i, e := range e.exprs {
			t[i] = pdp.TypeNames[e.GetResultType()]
		}
		args = fmt.Sprintf("%d arguments of following types \"%s\"", len(e.exprs), strings.Join(t, "\", \""))
	} else if len(e.exprs) > 0 {
		args = fmt.Sprintf("argument of type \"%s\"", pdp.TypeNames[e.exprs[0].GetResultType()])
	} else {
		args = "no arguments"
	}

	return e.errorf("Can't find function %s which takes %s", e.ID, args)
}

type unknownAttributeError struct {
	errorLink
	ID string
}

func newUnknownAttributeError(ID string) *unknownAttributeError {
	return &unknownAttributeError{
		errorLink: errorLink{id: unknownAttributeErrorID},
		ID:        ID}
}

func (e *unknownAttributeError) Error() string {
	return e.errorf("Unknown attribute %q", e.ID)
}

type unknownTypeError struct {
	errorLink
	t string
}

func newUnknownTypeError(t string) *unknownTypeError {
	return &unknownTypeError{
		errorLink: errorLink{id: unknownTypeErrorID},
		t:         t}
}

func (e *unknownTypeError) Error() string {
	return e.errorf("Unknown value type %q", e.t)
}

type invalidTypeError struct {
	errorLink
	t int
}

func newInvalidTypeError(t int) *invalidTypeError {
	return &invalidTypeError{
		errorLink: errorLink{id: invalidTypeErrorID},
		t:         t}
}

func (e *invalidTypeError) Error() string {
	return e.errorf("Can't make value of %q type", pdp.TypeNames[e.t])
}

type missingContentError struct {
	errorLink
}

func newMissingContentError() *missingContentError {
	return &missingContentError{
		errorLink: errorLink{id: missingContentErrorID}}
}

func (e *missingContentError) Error() string {
	return e.errorf("Missing value content")
}

type notImplementedValueTypeError struct {
	errorLink
	t int
}

func newNotImplementedValueTypeError(t int) *notImplementedValueTypeError {
	return &notImplementedValueTypeError{
		errorLink: errorLink{id: notImplementedValueTypeErrorID},
		t:         t}
}

func (e *notImplementedValueTypeError) Error() string {
	return e.errorf("Parsing for type %s hasn't been implemented yet", pdp.TypeNames[e.t])
}

type invalidAddressError struct {
	errorLink
	s string
}

func newInvalidAddressError(s string) *invalidAddressError {
	return &invalidAddressError{
		errorLink: errorLink{id: invalidAddressErrorID},
		s:         s}
}

func (e *invalidAddressError) Error() string {
	return e.errorf("Expected value of address type but got %q", e.s)
}

type invalidNetworkError struct {
	errorLink
	s   string
	err error
}

func newInvalidNetworkError(s string, err error) *invalidNetworkError {
	return &invalidNetworkError{
		errorLink: errorLink{id: invalidNetworkErrorID},
		s:         s,
		err:       err}
}

func (e *invalidNetworkError) Error() string {
	return e.errorf("Expected value of network type but got %q (%v)", e.s, e.err)
}

type invalidDomainError struct {
	errorLink
	s   string
	err error
}

func newInvalidDomainError(s string, err error) *invalidDomainError {
	return &invalidDomainError{
		errorLink: errorLink{id: invalidDomainErrorID},
		s:         s,
		err:       err}
}

func (e *invalidDomainError) Error() string {
	return e.errorf("Expected value of domain type but got %q (%v)", e.s, e.err)
}

type selectorURIError struct {
	errorLink
	uri string
	err error
}

func newSelectorURIError(uri string, err error) *selectorURIError {
	return &selectorURIError{
		errorLink: errorLink{id: selectorURIErrorID},
		uri:       uri,
		err:       err}
}

func (e *selectorURIError) Error() string {
	return e.errorf("Expected seletor URI but got %q (%s)", e.uri, e.err)
}

type selectorLocationError struct {
	errorLink
	loc string
	uri string
}

func newSelectorLocationError(loc, uri string) *selectorLocationError {
	return &selectorLocationError{
		errorLink: errorLink{id: selectorLocationErrorID},
		loc:       loc,
		uri:       uri}
}

func (e *selectorLocationError) Error() string {
	return e.errorf("Expected selector location in form of <Content-ID>/<Item-ID> got %q (%s)", e.loc, e.uri)
}

type unsupportedSelectorSchemeError struct {
	errorLink
	scheme string
	uri    string
}

func newUnsupportedSelectorSchemeError(scheme, uri string) *unsupportedSelectorSchemeError {
	return &unsupportedSelectorSchemeError{
		errorLink: errorLink{id: unsupportedSelectorSchemeErrorID},
		scheme:    scheme,
		uri:       uri}
}

func (e *unsupportedSelectorSchemeError) Error() string {
	return e.errorf("Unsupported selector scheme %q (%s)", e.scheme, e.uri)
}

type entityAmbiguityError struct {
	errorLink
	fields []string
}

func newEntityAmbiguityError(fields []string) *entityAmbiguityError {
	return &entityAmbiguityError{
		errorLink: errorLink{id: entityAmbiguityErrorID},
		fields:    fields}
}

func (e *entityAmbiguityError) Error() string {
	return e.errorf("Expected rules (for policy), policies (for policy set) or effect (for rule) but got %s", strings.Join(e.fields, ", "))
}

type entityMissingKeyError struct {
	errorLink
}

func newEntityMissingKeyError() *entityMissingKeyError {
	return &entityMissingKeyError{
		errorLink: errorLink{id: entityMissingKeyErrorID}}
}

func (e *entityMissingKeyError) Error() string {
	return e.errorf("Expected rules (for policy), policies (for policy set) or effect (for rule) but got nothing")
}

type unknownPolicyUpdateOperationError struct {
	errorLink
	op string
}

func newUnknownPolicyUpdateOperationError(op string) *unknownPolicyUpdateOperationError {
	return &unknownPolicyUpdateOperationError{
		errorLink: errorLink{id: unknownPolicyUpdateOperationErrorID},
		op:        op}
}

func (e *unknownPolicyUpdateOperationError) Error() string {
	return e.errorf("Unknown policy update operation %q", e.op)
}

type invalidPolicyUpdatePathElementError struct {
	errorLink
	v   interface{}
	idx int
}

func newInvalidPolicyUpdatePathElementError(v interface{}, idx int) *invalidPolicyUpdatePathElementError {
	return &invalidPolicyUpdatePathElementError{
		errorLink: errorLink{id: invalidPolicyUpdatePathElementErrorID},
		v:         v,
		idx:       idx}
}

func (e *invalidPolicyUpdatePathElementError) Error() string {
	return e.errorf("Expected string as %d path element but got %T", e.idx, e.v)
}
