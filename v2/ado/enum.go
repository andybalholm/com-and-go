package ado

const (
	OpenUnspecified = -1
	OpenForwardOnly = 0
	OpenKeyset      = 1
	OpenDynamic     = 2
	OpenStatic      = 3
)

const (
	HoldRecords    = 0x100
	MovePrevious   = 0x200
	AddNew         = 0x1000400
	Delete         = 0x1000800
	Update         = 0x1008000
	Bookmark       = 0x2000
	ApproxPosition = 0x4000
	UpdateBatch    = 0x10000
	Resync         = 0x20000
	Notify         = 0x40000
	Find           = 0x80000
	Seek           = 0x400000
	Index          = 0x800000
)

const (
	LockUnspecified     = -1
	LockReadOnly        = 1
	LockPessimistic     = 2
	LockOptimistic      = 3
	LockBatchOptimistic = 4
)

const (
	OptionUnspecified     = -1
	AsyncExecute          = 0x10
	AsyncFetch            = 0x20
	AsyncFetchNonBlocking = 0x40
	ExecuteNoRecords      = 0x80
	ExecuteStream         = 0x400
	ExecuteRecord         = 0x800
)

const (
	ConnectUnspecified = -1
	AsyncConnect       = 0x10
)

const (
	StateClosed     = 0
	StateOpen       = 0x1
	StateConnecting = 0x2
	StateExecuting  = 0x4
	StateFetching   = 0x8
)

const (
	UseNone        = 1
	UseServer      = 2
	UseClient      = 3
	UseClientBatch = 3
)

const (
	Empty            = 0
	TinyInt          = 16
	SmallInt         = 2
	Integer          = 3
	BigInt           = 20
	UnsignedTinyInt  = 17
	UnsignedSmallInt = 18
	UnsignedInt      = 19
	UnsignedBigInt   = 21
	Single           = 4
	Double           = 5
	Currency         = 6
	Decimal          = 14
	Numeric          = 131
	Boolean          = 11
	Error            = 10
	UserDefined      = 132
	Variant          = 12
	IDispatch        = 9
	IUnknown         = 13
	GUID             = 72
	Date             = 7
	DBDate           = 133
	DBTime           = 134
	DBTimeStamp      = 135
	BSTR             = 8
	Char             = 129
	VarChar          = 200
	LongVarChar      = 201
	WChar            = 130
	VarWChar         = 202
	LongVarWChar     = 203
	Binary           = 128
	VarBinary        = 204
	LongVarBinary    = 205
	Chapter          = 136
	FileTime         = 64
	PropVariant      = 138
	VarNumeric       = 139
	Array            = 0x2000
)

const (
	FldUnspecified      = -1
	FldMayDefer         = 0x2
	FldUpdatable        = 0x4
	FldUnknownUpdatable = 0x8
	FldFixed            = 0x10
	FldIsNullable       = 0x20
	FldMayBeNull        = 0x40
	FldLong             = 0x80
	FldRowID            = 0x100
	FldRowVersion       = 0x200
	FldCacheDeferred    = 0x1000
	FldIsChapter        = 0x2000
	FldNegativeScale    = 0x4000
	FldKeyColumn        = 0x8000
	FldIsRowURL         = 0x10000
	FldIsDefaultStream  = 0x20000
	FldIsCollection     = 0x40000
)

const (
	EditNone       = 0
	EditInProgress = 0x1
	EditAdd        = 0x2
	EditDelete     = 0x4
)

const (
	RecOK                   = 0
	RecNew                  = 0x1
	RecModified             = 0x2
	RecDeleted              = 0x4
	RecUnmodified           = 0x8
	RecInvalid              = 0x10
	RecMultipleChanges      = 0x40
	RecPendingChanges       = 0x80
	RecCanceled             = 0x100
	RecCantRelease          = 0x400
	RecConcurrencyViolation = 0x800
	RecIntegrityViolation   = 0x1000
	RecMaxChangesExceeded   = 0x2000
	RecObjectOpen           = 0x4000
	RecOutOfMemory          = 0x8000
	RecPermissionDenied     = 0x10000
	RecSchemaViolation      = 0x20000
	RecDBDeleted            = 0x40000
)

const (
	GetRowsRest = -1
)

const (
	PosUnknown = -1
	PosBOF     = -2
	PosEOF     = -3
)

const (
	BookmarkCurrent = 0
	BookmarkFirst   = 1
	BookmarkLast    = 2
)

const (
	MarshalAll          = 0
	MarshalModifiedOnly = 1
)

const (
	AffectCurrent     = 1
	AffectGroup       = 2
	AffectAll         = 3
	AffectAllChapters = 4
)

const (
	ResyncUnderlyingValues = 1
	ResyncAllValues        = 2
)

const (
	CompareLessThan      = 0
	CompareEqual         = 1
	CompareGreaterThan   = 2
	CompareNotEqual      = 3
	CompareNotComparable = 4
)

const (
	FilterNone               = 0
	FilterPendingRecords     = 1
	FilterAffectedRecords    = 2
	FilterFetchedRecords     = 3
	FilterPredicate          = 4
	FilterConflictingRecords = 5
)

const (
	SearchForward  = 1
	SearchBackward = -1
)

const (
	PersistADTG = 0
	PersistXML  = 1
)

const (
	ClipString = 2
)

const (
	PromptAlways           = 1
	PromptComplete         = 2
	PromptCompleteRequired = 3
	PromptNever            = 4
)

const (
	ModeUnknown        = 0
	ModeRead           = 1
	ModeWrite          = 2
	ModeReadWrite      = 3
	ModeShareDenyRead  = 4
	ModeShareDenyWrite = 8
	ModeShareExclusive = 0xc
	ModeShareDenyNone  = 0x10
	ModeRecursive      = 0x400000
)

const (
	CreateCollection    = 0x2000
	CreateStructDoc     = -0x80000000
	CreateNonCollection = 0
	OpenIfExists        = 0x2000000
	CreateOverwrite     = 0x4000000
	FailIfNotExists     = -1
)

const (
	OpenRecordUnspecified = -1
	OpenSource            = 0x800000
	OpenOutput            = 0x800000
	OpenAsync             = 0x1000
	DelayFetchStream      = 0x4000
	DelayFetchFields      = 0x8000
	OpenExecuteCommand    = 0x10000
)

const (
	XactUnspecified     = -1
	XactChaos           = 0x10
	XactReadUncommitted = 0x100
	XactBrowse          = 0x100
	XactCursorStability = 0x1000
	XactReadCommitted   = 0x1000
	XactRepeatableRead  = 0x10000
	XactSerializable    = 0x100000
	XactIsolated        = 0x100000
)

const (
	XactCommitRetaining = 0x20000
	XactAbortRetaining  = 0x40000
	XactAsyncPhaseOne   = 0x80000
	XactSyncPhaseOne    = 0x100000
)

const (
	PropNotSupported = 0
	PropRequired     = 0x1
	PropOptional     = 0x2
	PropRead         = 0x200
	PropWrite        = 0x400
)

const (
	ParamSigned   = 0x10
	ParamNullable = 0x40
	ParamLong     = 0x80
)

const (
	ParamUnknown     = 0
	ParamInput       = 0x1
	ParamOutput      = 0x2
	ParamInputOutput = 0x3
	ParamReturnValue = 0x4
)

const (
	CmdUnspecified = -1
	CmdUnknown     = 0x8
	CmdText        = 0x1
	CmdTable       = 0x2
	CmdStoredProc  = 0x4
	CmdFile        = 0x100
	CmdTableDirect = 0x200
)

const (
	StatusOK             = 0x1
	StatusErrorsOccurred = 0x2
	StatusCantDeny       = 0x3
	StatusCancel         = 0x4
	StatusUnwantedEvent  = 0x5
)

const (
	RsnAddNew       = 1
	RsnDelete       = 2
	RsnUpdate       = 3
	RsnUndoUpdate   = 4
	RsnUndoAddNew   = 5
	RsnUndoDelete   = 6
	RsnRequery      = 7
	RsnResynch      = 8
	RsnClose        = 9
	RsnMove         = 10
	RsnFirstChange  = 11
	RsnMoveFirst    = 12
	RsnMoveNext     = 13
	RsnMovePrevious = 14
	RsnMoveLast     = 15
)

const (
	SchemaProviderSpecific       = -1
	SchemaAsserts                = 0
	SchemaCatalogs               = 1
	SchemaCharacterSets          = 2
	SchemaCollations             = 3
	SchemaColumns                = 4
	SchemaCheckConstraints       = 5
	SchemaConstraintColumnUsage  = 6
	SchemaConstraintTableUsage   = 7
	SchemaKeyColumnUsage         = 8
	SchemaReferentialContraints  = 9
	SchemaReferentialConstraints = 9
	SchemaTableConstraints       = 10
	SchemaColumnsDomainUsage     = 11
	SchemaIndexes                = 12
	SchemaColumnPrivileges       = 13
	SchemaTablePrivileges        = 14
	SchemaUsagePrivileges        = 15
	SchemaProcedures             = 16
	SchemaSchemata               = 17
	SchemaSQLLanguages           = 18
	SchemaStatistics             = 19
	SchemaTables                 = 20
	SchemaTranslations           = 21
	SchemaProviderTypes          = 22
	SchemaViews                  = 23
	SchemaViewColumnUsage        = 24
	SchemaViewTableUsage         = 25
	SchemaProcedureParameters    = 26
	SchemaForeignKeys            = 27
	SchemaPrimaryKeys            = 28
	SchemaProcedureColumns       = 29
	SchemaDBInfoKeywords         = 30
	SchemaDBInfoLiterals         = 31
	SchemaCubes                  = 32
	SchemaDimensions             = 33
	SchemaHierarchies            = 34
	SchemaLevels                 = 35
	SchemaMeasures               = 36
	SchemaProperties             = 37
	SchemaMembers                = 38
	SchemaTrustees               = 39
	SchemaFunctions              = 40
	SchemaActions                = 41
	SchemaCommands               = 42
	SchemaSets                   = 43
)

const (
	FieldOK                   = 0
	FieldCantConvertValue     = 2
	FieldIsNull               = 3
	FieldTruncated            = 4
	FieldSignMismatch         = 5
	FieldDataOverflow         = 6
	FieldCantCreate           = 7
	FieldUnavailable          = 8
	FieldPermissionDenied     = 9
	FieldIntegrityViolation   = 10
	FieldSchemaViolation      = 11
	FieldBadStatus            = 12
	FieldDefault              = 13
	FieldIgnore               = 15
	FieldDoesNotExist         = 16
	FieldInvalidURL           = 17
	FieldResourceLocked       = 18
	FieldResourceExists       = 19
	FieldCannotComplete       = 20
	FieldVolumeNotFound       = 21
	FieldOutOfSpace           = 22
	FieldCannotDeleteSource   = 23
	FieldReadOnly             = 24
	FieldResourceOutOfScope   = 25
	FieldAlreadyExists        = 26
	FieldPendingInsert        = 0x10000
	FieldPendingDelete        = 0x20000
	FieldPendingChange        = 0x40000
	FieldPendingUnknown       = 0x80000
	FieldPendingUnknownDelete = 0x100000
)

const (
	SeekFirstEQ  = 0x1
	SeekLastEQ   = 0x2
	SeekAfterEQ  = 0x4
	SeekAfter    = 0x8
	SeekBeforeEQ = 0x10
	SeekBefore   = 0x20
)

const (
	MoveUnspecified     = -1
	MoveOverWrite       = 1
	MoveDontUpdateLinks = 2
	MoveAllowEmulation  = 4
)

const (
	CopyUnspecified    = -1
	CopyOverWrite      = 1
	CopyAllowEmulation = 4
	CopyNonRecursive   = 2
)

const (
	TypeBinary = 1
	TypeText   = 2
)

const (
	LF   = 10
	CR   = 13
	CRLF = -1
)

const (
	OpenStreamUnspecified = -1
	OpenStreamAsync       = 1
	OpenStreamFromRecord  = 4
)

const (
	WriteChar = 0
	WriteLine = 1
)

const (
	SaveCreateNotExist  = 1
	SaveCreateOverWrite = 2
)

const (
	DefaultStream = -1
	RecordURL     = -2
)

const (
	ReadAll  = -1
	ReadLine = -2
)

const (
	SimpleRecord     = 0
	CollectionRecord = 1
	StructDoc        = 2
)
