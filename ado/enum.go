package ado

type CursorTypeEnum int32

const (
	OpenUnspecified CursorTypeEnum = -1
	OpenForwardOnly CursorTypeEnum = 0
	OpenKeyset      CursorTypeEnum = 1
	OpenDynamic     CursorTypeEnum = 2
	OpenStatic      CursorTypeEnum = 3
)

type CursorOptionEnum int32

const (
	HoldRecords    CursorOptionEnum = 0x100
	MovePrevious   CursorOptionEnum = 0x200
	AddNew         CursorOptionEnum = 0x1000400
	Delete         CursorOptionEnum = 0x1000800
	Update         CursorOptionEnum = 0x1008000
	Bookmark       CursorOptionEnum = 0x2000
	ApproxPosition CursorOptionEnum = 0x4000
	UpdateBatch    CursorOptionEnum = 0x10000
	Resync         CursorOptionEnum = 0x20000
	Notify         CursorOptionEnum = 0x40000
	Find           CursorOptionEnum = 0x80000
	Seek           CursorOptionEnum = 0x400000
	Index          CursorOptionEnum = 0x800000
)

type LockTypeEnum int32

const (
	LockUnspecified     LockTypeEnum = -1
	LockReadOnly        LockTypeEnum = 1
	LockPessimistic     LockTypeEnum = 2
	LockOptimistic      LockTypeEnum = 3
	LockBatchOptimistic LockTypeEnum = 4
)

type ExecuteOptionEnum int32

const (
	OptionUnspecified     ExecuteOptionEnum = -1
	AsyncExecute          ExecuteOptionEnum = 0x10
	AsyncFetch            ExecuteOptionEnum = 0x20
	AsyncFetchNonBlocking ExecuteOptionEnum = 0x40
	ExecuteNoRecords      ExecuteOptionEnum = 0x80
	ExecuteStream         ExecuteOptionEnum = 0x400
	ExecuteRecord         ExecuteOptionEnum = 0x800
)

type ConnectOptionEnum int32

const (
	ConnectUnspecified ConnectOptionEnum = -1
	AsyncConnect       ConnectOptionEnum = 0x10
)

type ObjectStateEnum int32

const (
	StateClosed     ObjectStateEnum = 0
	StateOpen       ObjectStateEnum = 0x1
	StateConnecting ObjectStateEnum = 0x2
	StateExecuting  ObjectStateEnum = 0x4
	StateFetching   ObjectStateEnum = 0x8
)

type CursorLocationEnum int32

const (
	UseNone        CursorLocationEnum = 1
	UseServer      CursorLocationEnum = 2
	UseClient      CursorLocationEnum = 3
	UseClientBatch CursorLocationEnum = 3
)

type DataTypeEnum int32

const (
	Empty            DataTypeEnum = 0
	TinyInt          DataTypeEnum = 16
	SmallInt         DataTypeEnum = 2
	Integer          DataTypeEnum = 3
	BigInt           DataTypeEnum = 20
	UnsignedTinyInt  DataTypeEnum = 17
	UnsignedSmallInt DataTypeEnum = 18
	UnsignedInt      DataTypeEnum = 19
	UnsignedBigInt   DataTypeEnum = 21
	Single           DataTypeEnum = 4
	Double           DataTypeEnum = 5
	Currency         DataTypeEnum = 6
	Decimal          DataTypeEnum = 14
	Numeric          DataTypeEnum = 131
	Boolean          DataTypeEnum = 11
	Error            DataTypeEnum = 10
	UserDefined      DataTypeEnum = 132
	Variant          DataTypeEnum = 12
	IDispatch        DataTypeEnum = 9
	IUnknown         DataTypeEnum = 13
	GUID             DataTypeEnum = 72
	Date             DataTypeEnum = 7
	DBDate           DataTypeEnum = 133
	DBTime           DataTypeEnum = 134
	DBTimeStamp      DataTypeEnum = 135
	BSTR             DataTypeEnum = 8
	Char             DataTypeEnum = 129
	VarChar          DataTypeEnum = 200
	LongVarChar      DataTypeEnum = 201
	WChar            DataTypeEnum = 130
	VarWChar         DataTypeEnum = 202
	LongVarWChar     DataTypeEnum = 203
	Binary           DataTypeEnum = 128
	VarBinary        DataTypeEnum = 204
	LongVarBinary    DataTypeEnum = 205
	Chapter          DataTypeEnum = 136
	FileTime         DataTypeEnum = 64
	PropVariant      DataTypeEnum = 138
	VarNumeric       DataTypeEnum = 139
	Array            DataTypeEnum = 0x2000
)

type FieldAttributeEnum int32

const (
	FldUnspecified      FieldAttributeEnum = -1
	FldMayDefer         FieldAttributeEnum = 0x2
	FldUpdatable        FieldAttributeEnum = 0x4
	FldUnknownUpdatable FieldAttributeEnum = 0x8
	FldFixed            FieldAttributeEnum = 0x10
	FldIsNullable       FieldAttributeEnum = 0x20
	FldMayBeNull        FieldAttributeEnum = 0x40
	FldLong             FieldAttributeEnum = 0x80
	FldRowID            FieldAttributeEnum = 0x100
	FldRowVersion       FieldAttributeEnum = 0x200
	FldCacheDeferred    FieldAttributeEnum = 0x1000
	FldIsChapter        FieldAttributeEnum = 0x2000
	FldNegativeScale    FieldAttributeEnum = 0x4000
	FldKeyColumn        FieldAttributeEnum = 0x8000
	FldIsRowURL         FieldAttributeEnum = 0x10000
	FldIsDefaultStream  FieldAttributeEnum = 0x20000
	FldIsCollection     FieldAttributeEnum = 0x40000
)

type EditModeEnum int32

const (
	EditNone       EditModeEnum = 0
	EditInProgress EditModeEnum = 0x1
	EditAdd        EditModeEnum = 0x2
	EditDelete     EditModeEnum = 0x4
)

type RecordStatusEnum int32

const (
	RecOK                   RecordStatusEnum = 0
	RecNew                  RecordStatusEnum = 0x1
	RecModified             RecordStatusEnum = 0x2
	RecDeleted              RecordStatusEnum = 0x4
	RecUnmodified           RecordStatusEnum = 0x8
	RecInvalid              RecordStatusEnum = 0x10
	RecMultipleChanges      RecordStatusEnum = 0x40
	RecPendingChanges       RecordStatusEnum = 0x80
	RecCanceled             RecordStatusEnum = 0x100
	RecCantRelease          RecordStatusEnum = 0x400
	RecConcurrencyViolation RecordStatusEnum = 0x800
	RecIntegrityViolation   RecordStatusEnum = 0x1000
	RecMaxChangesExceeded   RecordStatusEnum = 0x2000
	RecObjectOpen           RecordStatusEnum = 0x4000
	RecOutOfMemory          RecordStatusEnum = 0x8000
	RecPermissionDenied     RecordStatusEnum = 0x10000
	RecSchemaViolation      RecordStatusEnum = 0x20000
	RecDBDeleted            RecordStatusEnum = 0x40000
)

type GetRowsOptionEnum int32

const (
	GetRowsRest GetRowsOptionEnum = -1
)

type PositionEnum int

const (
	PosUnknown PositionEnum = -1
	PosBOF     PositionEnum = -2
	PosEOF     PositionEnum = -3
)

type BookmarkEnum int32

const (
	BookmarkCurrent BookmarkEnum = 0
	BookmarkFirst   BookmarkEnum = 1
	BookmarkLast    BookmarkEnum = 2
)

type MarshalOptionsEnum int32

const (
	MarshalAll          MarshalOptionsEnum = 0
	MarshalModifiedOnly MarshalOptionsEnum = 1
)

type AffectEnum int32

const (
	AffectCurrent     AffectEnum = 1
	AffectGroup       AffectEnum = 2
	AffectAll         AffectEnum = 3
	AffectAllChapters AffectEnum = 4
)

type ResyncEnum int32

const (
	ResyncUnderlyingValues ResyncEnum = 1
	ResyncAllValues        ResyncEnum = 2
)

type CompareEnum int32

const (
	CompareLessThan      CompareEnum = 0
	CompareEqual         CompareEnum = 1
	CompareGreaterThan   CompareEnum = 2
	CompareNotEqual      CompareEnum = 3
	CompareNotComparable CompareEnum = 4
)

type FilterGroupEnum int32

const (
	FilterNone               FilterGroupEnum = 0
	FilterPendingRecords     FilterGroupEnum = 1
	FilterAffectedRecords    FilterGroupEnum = 2
	FilterFetchedRecords     FilterGroupEnum = 3
	FilterPredicate          FilterGroupEnum = 4
	FilterConflictingRecords FilterGroupEnum = 5
)

type SearchDirectionEnum int32

const (
	SearchForward  SearchDirectionEnum = 1
	SearchBackward SearchDirectionEnum = -1
)

type PersistFormatEnum int32

const (
	PersistADTG PersistFormatEnum = 0
	PersistXML  PersistFormatEnum = 1
)

type StringFormatEnum int32

const (
	ClipString StringFormatEnum = 2
)

type ConnectPromptEnum int32

const (
	PromptAlways           ConnectPromptEnum = 1
	PromptComplete         ConnectPromptEnum = 2
	PromptCompleteRequired ConnectPromptEnum = 3
	PromptNever            ConnectPromptEnum = 4
)

type ConnectModeEnum int32

const (
	ModeUnknown        ConnectModeEnum = 0
	ModeRead           ConnectModeEnum = 1
	ModeWrite          ConnectModeEnum = 2
	ModeReadWrite      ConnectModeEnum = 3
	ModeShareDenyRead  ConnectModeEnum = 4
	ModeShareDenyWrite ConnectModeEnum = 8
	ModeShareExclusive ConnectModeEnum = 0xc
	ModeShareDenyNone  ConnectModeEnum = 0x10
	ModeRecursive      ConnectModeEnum = 0x400000
)

type RecordCreateOptionsEnum int32

const (
	CreateCollection    RecordCreateOptionsEnum = 0x2000
	CreateStructDoc     RecordCreateOptionsEnum = -0x80000000
	CreateNonCollection RecordCreateOptionsEnum = 0
	OpenIfExists        RecordCreateOptionsEnum = 0x2000000
	CreateOverwrite     RecordCreateOptionsEnum = 0x4000000
	FailIfNotExists     RecordCreateOptionsEnum = -1
)

type RecordOpenOptionsEnum int32

const (
	OpenRecordUnspecified RecordOpenOptionsEnum = -1
	OpenSource            RecordOpenOptionsEnum = 0x800000
	OpenOutput            RecordOpenOptionsEnum = 0x800000
	OpenAsync             RecordOpenOptionsEnum = 0x1000
	DelayFetchStream      RecordOpenOptionsEnum = 0x4000
	DelayFetchFields      RecordOpenOptionsEnum = 0x8000
	OpenExecuteCommand    RecordOpenOptionsEnum = 0x10000
)

type IsolationLevelEnum int32

const (
	XactUnspecified     IsolationLevelEnum = -1
	XactChaos           IsolationLevelEnum = 0x10
	XactReadUncommitted IsolationLevelEnum = 0x100
	XactBrowse          IsolationLevelEnum = 0x100
	XactCursorStability IsolationLevelEnum = 0x1000
	XactReadCommitted   IsolationLevelEnum = 0x1000
	XactRepeatableRead  IsolationLevelEnum = 0x10000
	XactSerializable    IsolationLevelEnum = 0x100000
	XactIsolated        IsolationLevelEnum = 0x100000
)

type XactAttributeEnum int32

const (
	XactCommitRetaining XactAttributeEnum = 0x20000
	XactAbortRetaining  XactAttributeEnum = 0x40000
	XactAsyncPhaseOne   XactAttributeEnum = 0x80000
	XactSyncPhaseOne    XactAttributeEnum = 0x100000
)

type PropertyAttributesEnum int32

const (
	PropNotSupported PropertyAttributesEnum = 0
	PropRequired     PropertyAttributesEnum = 0x1
	PropOptional     PropertyAttributesEnum = 0x2
	PropRead         PropertyAttributesEnum = 0x200
	PropWrite        PropertyAttributesEnum = 0x400
)

type ParameterAttributesEnum int32

const (
	ParamSigned   ParameterAttributesEnum = 0x10
	ParamNullable ParameterAttributesEnum = 0x40
	ParamLong     ParameterAttributesEnum = 0x80
)

type ParameterDirectionEnum int32

const (
	ParamUnknown     ParameterDirectionEnum = 0
	ParamInput       ParameterDirectionEnum = 0x1
	ParamOutput      ParameterDirectionEnum = 0x2
	ParamInputOutput ParameterDirectionEnum = 0x3
	ParamReturnValue ParameterDirectionEnum = 0x4
)

type CommandTypeEnum int32

const (
	CmdUnspecified CommandTypeEnum = -1
	CmdUnknown     CommandTypeEnum = 0x8
	CmdText        CommandTypeEnum = 0x1
	CmdTable       CommandTypeEnum = 0x2
	CmdStoredProc  CommandTypeEnum = 0x4
	CmdFile        CommandTypeEnum = 0x100
	CmdTableDirect CommandTypeEnum = 0x200
)

type EventStatusEnum int32

const (
	StatusOK             EventStatusEnum = 0x1
	StatusErrorsOccurred EventStatusEnum = 0x2
	StatusCantDeny       EventStatusEnum = 0x3
	StatusCancel         EventStatusEnum = 0x4
	StatusUnwantedEvent  EventStatusEnum = 0x5
)

type EventReasonEnum int32

const (
	RsnAddNew       EventReasonEnum = 1
	RsnDelete       EventReasonEnum = 2
	RsnUpdate       EventReasonEnum = 3
	RsnUndoUpdate   EventReasonEnum = 4
	RsnUndoAddNew   EventReasonEnum = 5
	RsnUndoDelete   EventReasonEnum = 6
	RsnRequery      EventReasonEnum = 7
	RsnResynch      EventReasonEnum = 8
	RsnClose        EventReasonEnum = 9
	RsnMove         EventReasonEnum = 10
	RsnFirstChange  EventReasonEnum = 11
	RsnMoveFirst    EventReasonEnum = 12
	RsnMoveNext     EventReasonEnum = 13
	RsnMovePrevious EventReasonEnum = 14
	RsnMoveLast     EventReasonEnum = 15
)

type SchemaEnum int32

const (
	SchemaProviderSpecific       SchemaEnum = -1
	SchemaAsserts                SchemaEnum = 0
	SchemaCatalogs               SchemaEnum = 1
	SchemaCharacterSets          SchemaEnum = 2
	SchemaCollations             SchemaEnum = 3
	SchemaColumns                SchemaEnum = 4
	SchemaCheckConstraints       SchemaEnum = 5
	SchemaConstraintColumnUsage  SchemaEnum = 6
	SchemaConstraintTableUsage   SchemaEnum = 7
	SchemaKeyColumnUsage         SchemaEnum = 8
	SchemaReferentialContraints  SchemaEnum = 9
	SchemaReferentialConstraints SchemaEnum = 9
	SchemaTableConstraints       SchemaEnum = 10
	SchemaColumnsDomainUsage     SchemaEnum = 11
	SchemaIndexes                SchemaEnum = 12
	SchemaColumnPrivileges       SchemaEnum = 13
	SchemaTablePrivileges        SchemaEnum = 14
	SchemaUsagePrivileges        SchemaEnum = 15
	SchemaProcedures             SchemaEnum = 16
	SchemaSchemata               SchemaEnum = 17
	SchemaSQLLanguages           SchemaEnum = 18
	SchemaStatistics             SchemaEnum = 19
	SchemaTables                 SchemaEnum = 20
	SchemaTranslations           SchemaEnum = 21
	SchemaProviderTypes          SchemaEnum = 22
	SchemaViews                  SchemaEnum = 23
	SchemaViewColumnUsage        SchemaEnum = 24
	SchemaViewTableUsage         SchemaEnum = 25
	SchemaProcedureParameters    SchemaEnum = 26
	SchemaForeignKeys            SchemaEnum = 27
	SchemaPrimaryKeys            SchemaEnum = 28
	SchemaProcedureColumns       SchemaEnum = 29
	SchemaDBInfoKeywords         SchemaEnum = 30
	SchemaDBInfoLiterals         SchemaEnum = 31
	SchemaCubes                  SchemaEnum = 32
	SchemaDimensions             SchemaEnum = 33
	SchemaHierarchies            SchemaEnum = 34
	SchemaLevels                 SchemaEnum = 35
	SchemaMeasures               SchemaEnum = 36
	SchemaProperties             SchemaEnum = 37
	SchemaMembers                SchemaEnum = 38
	SchemaTrustees               SchemaEnum = 39
	SchemaFunctions              SchemaEnum = 40
	SchemaActions                SchemaEnum = 41
	SchemaCommands               SchemaEnum = 42
	SchemaSets                   SchemaEnum = 43
)

type FieldStatusEnum int32

const (
	FieldOK                   FieldStatusEnum = 0
	FieldCantConvertValue     FieldStatusEnum = 2
	FieldIsNull               FieldStatusEnum = 3
	FieldTruncated            FieldStatusEnum = 4
	FieldSignMismatch         FieldStatusEnum = 5
	FieldDataOverflow         FieldStatusEnum = 6
	FieldCantCreate           FieldStatusEnum = 7
	FieldUnavailable          FieldStatusEnum = 8
	FieldPermissionDenied     FieldStatusEnum = 9
	FieldIntegrityViolation   FieldStatusEnum = 10
	FieldSchemaViolation      FieldStatusEnum = 11
	FieldBadStatus            FieldStatusEnum = 12
	FieldDefault              FieldStatusEnum = 13
	FieldIgnore               FieldStatusEnum = 15
	FieldDoesNotExist         FieldStatusEnum = 16
	FieldInvalidURL           FieldStatusEnum = 17
	FieldResourceLocked       FieldStatusEnum = 18
	FieldResourceExists       FieldStatusEnum = 19
	FieldCannotComplete       FieldStatusEnum = 20
	FieldVolumeNotFound       FieldStatusEnum = 21
	FieldOutOfSpace           FieldStatusEnum = 22
	FieldCannotDeleteSource   FieldStatusEnum = 23
	FieldReadOnly             FieldStatusEnum = 24
	FieldResourceOutOfScope   FieldStatusEnum = 25
	FieldAlreadyExists        FieldStatusEnum = 26
	FieldPendingInsert        FieldStatusEnum = 0x10000
	FieldPendingDelete        FieldStatusEnum = 0x20000
	FieldPendingChange        FieldStatusEnum = 0x40000
	FieldPendingUnknown       FieldStatusEnum = 0x80000
	FieldPendingUnknownDelete FieldStatusEnum = 0x100000
)

type SeekEnum int32

const (
	SeekFirstEQ  SeekEnum = 0x1
	SeekLastEQ   SeekEnum = 0x2
	SeekAfterEQ  SeekEnum = 0x4
	SeekAfter    SeekEnum = 0x8
	SeekBeforeEQ SeekEnum = 0x10
	SeekBefore   SeekEnum = 0x20
)

type MoveRecordOptionsEnum int32

const (
	MoveUnspecified     MoveRecordOptionsEnum = -1
	MoveOverWrite       MoveRecordOptionsEnum = 1
	MoveDontUpdateLinks MoveRecordOptionsEnum = 2
	MoveAllowEmulation  MoveRecordOptionsEnum = 4
)

type CopyRecordOptionsEnum int32

const (
	CopyUnspecified    CopyRecordOptionsEnum = -1
	CopyOverWrite      CopyRecordOptionsEnum = 1
	CopyAllowEmulation CopyRecordOptionsEnum = 4
	CopyNonRecursive   CopyRecordOptionsEnum = 2
)

type StreamTypeEnum int32

const (
	TypeBinary StreamTypeEnum = 1
	TypeText   StreamTypeEnum = 2
)

type LineSeparatorEnum int32

const (
	LF   LineSeparatorEnum = 10
	CR   LineSeparatorEnum = 13
	CRLF LineSeparatorEnum = -1
)

type StreamOpenOptionsEnum int32

const (
	OpenStreamUnspecified StreamOpenOptionsEnum = -1
	OpenStreamAsync       StreamOpenOptionsEnum = 1
	OpenStreamFromRecord  StreamOpenOptionsEnum = 4
)

type StreamWriteEnum int32

const (
	WriteChar StreamWriteEnum = 0
	WriteLine StreamWriteEnum = 1
)

type SaveOptionsEnum int32

const (
	SaveCreateNotExist  SaveOptionsEnum = 1
	SaveCreateOverWrite SaveOptionsEnum = 2
)

type FieldEnum int32

const (
	DefaultStream FieldEnum = -1
	RecordURL     FieldEnum = -2
)

type StreamReadEnum int32

const (
	ReadAll  StreamReadEnum = -1
	ReadLine StreamReadEnum = -2
)

type RecordTypeEnum int32

const (
	SimpleRecord     RecordTypeEnum = 0
	CollectionRecord RecordTypeEnum = 1
	StructDoc        RecordTypeEnum = 2
)
