package net

type (
	TableIndex struct {
		index uint32
		table int
	}
	AssemblyRefIndex  TableIndex
	BlobIndex         TableIndex
	EventIndex        TableIndex
	FieldIndex        TableIndex
	GenericParamIndex TableIndex
	GuidIndex         TableIndex
	MethodDefIndex    TableIndex
	ModuleRefIndex    TableIndex
	ParamIndex        TableIndex
	PropertyIndex     TableIndex
	StringIndex       TableIndex
	TypeDefIndex      TableIndex
)

// II.24.2.6
type (
	EncodedIndex                    TableIndex
	CustomAttributeTypeEncodedIndex EncodedIndex
	HasConstantEncodedIndex         EncodedIndex
	HasCustomAttributeEncodedIndex  EncodedIndex
	HasDeclSecurityEncodedIndex     EncodedIndex
	HasFieldMarshalEncodedIndex     EncodedIndex
	HasSemanticsEncodedIndex        EncodedIndex
	ImplementationEncodedIndex      EncodedIndex
	MemberForwardedEncodedIndex     EncodedIndex
	MemberRefParentEncodedIndex     EncodedIndex
	MethodDefOrRefEncodedIndex      EncodedIndex
	ResolutionScopeEncodedIndex     EncodedIndex
	TypeDefOrRefEncodedIndex        EncodedIndex
	TypeOrMethodDefEncodedIndex     EncodedIndex
)

const (
	idx_TypeDefOrRef = iota
	idx_HasConstant
	idx_HasCustomAttribute
	idx_HasFieldMarshal
	idx_HasDeclSecurity
	idx_MemberRefParent
	idx_HasSemantics
	idx_MethodDefOrRef
	idx_MemberForwarded
	idx_Implementation
	idx_CustomAttributeType
	idx_ResolutionScope
	idx_TypeOrMethodDef
)

var idx_name_lut = map[string]int{
	"AssemblyRefIndex":                id_AssemblyRef,
	"EventIndex":                      id_Event,
	"FieldIndex":                      id_Field,
	"GenericParamIndex":               id_GenericParam,
	"MethodDefIndex":                  id_MethodDef,
	"ModuleRefIndex":                  id_ModuleRef,
	"ParamIndex":                      id_Param,
	"PropertyIndex":                   id_Property,
	"TypeDefIndex":                    id_TypeDef,
	"TypeDefOrRefEncodedIndex":        idx_TypeDefOrRef,
	"HasConstantEncodedIndex":         idx_HasConstant,
	"HasCustomAttributeEncodedIndex":  idx_HasCustomAttribute,
	"HasFieldMarshalEncodedIndex":     idx_HasFieldMarshal,
	"HasDeclSecurityEncodedIndex":     idx_HasDeclSecurity,
	"MemberRefParentEncodedIndex":     idx_MemberRefParent,
	"HasSemanticsEncodedIndex":        idx_HasSemantics,
	"MethodDefOrRefEncodedIndex":      idx_MethodDefOrRef,
	"MemberForwardedEncodedIndex":     idx_MemberForwarded,
	"ImplementationEncodedIndex":      idx_Implementation,
	"CustomAttributeTypeEncodedIndex": idx_CustomAttributeType,
	"ResolutionScopeEncodedIndex":     idx_ResolutionScope,
	"TypeOrMethodDefEncodedIndex":     idx_TypeOrMethodDef,
}

var enc_lut = [][]int{
	// TypeDefOrRef
	[]int{
		id_TypeDef,
		id_TypeRef,
		id_TypeSpec},
	// HasConstant
	[]int{
		id_Field,
		id_Param,
		id_Property},
	// HasCustomAttribute
	[]int{
		id_MethodDef,
		id_Field,
		id_TypeRef,
		id_TypeDef,
		id_Param,
		id_InterfaceImpl,
		id_MemberRef,
		id_Module,
		id_DeclSecurity,
		id_Property,
		id_Event,
		id_StandAloneSig,
		id_ModuleRef,
		id_TypeSpec,
		id_Assembly,
		id_AssemblyRef,
		id_File,
		id_ExportedType,
		id_ManifestResource,
		id_GenericParam,
		id_GenericParamConstraint,
		id_MethodSpec},
	// HasFieldMarshal
	[]int{
		id_Field,
		id_Param},
	// HasDeclSecurity
	[]int{
		id_TypeDef,
		id_MethodDef,
		id_Assembly},
	// MemberRefParent
	[]int{
		id_TypeDef,
		id_TypeRef,
		id_ModuleRef,
		id_MethodDef,
		id_TypeSpec},
	// HasSemantics
	[]int{
		id_Event,
		id_Property},
	// MethodDefOrRef
	[]int{
		id_MethodDef,
		id_MemberRef},
	// MemberForwarded
	[]int{
		id_Field,
		id_MethodDef},
	// Implementation
	[]int{
		id_File,
		id_AssemblyRef,
		id_ExportedType},
	// CustomAttributeType
	[]int{
		id_nullTable,
		id_nullTable,
		id_MethodDef,
		id_MemberRef,
		id_nullTable},
	// ResolutionScope
	[]int{
		id_Module,
		id_ModuleRef,
		id_AssemblyRef,
		id_TypeRef},
	// TypeOrMethodDef
	[]int{
		id_TypeDef,
		id_MethodDef,
	},
}
