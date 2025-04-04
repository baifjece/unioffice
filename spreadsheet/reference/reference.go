//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package reference ;import (_df "errors";_a "fmt";_f "github.com/unidoc/unioffice/v2/spreadsheet/update";_g "regexp";_dg "strconv";_de "strings";);

// ParseCellReference parses a cell reference of the form 'A10' and splits it
// into column/row segments.
func ParseCellReference (s string )(CellReference ,error ){s =_de .TrimSpace (s );if len (s )< 2{return CellReference {},_df .New ("\u0063\u0065\u006c\u006c\u0020\u0072\u0065\u0066e\u0072\u0065\u006ece\u0020\u006d\u0075\u0073\u0074\u0020h\u0061\u0076\u0065\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0077o\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0073");
};_cb :=CellReference {};_e ,_ea ,_dfg :=_dff (s );if _dfg !=nil {return CellReference {},_dfg ;};if _e !=""{_cb .SheetName =_e ;};if s [0]=='$'{_cb .AbsoluteColumn =true ;_ea =_ea [1:];};_db :=-1;_ca :for _b :=0;_b < len (_ea );_b ++{switch {case _ea [_b ]>='0'&&_ea [_b ]<='9'||_ea [_b ]=='$':_db =_b ;
break _ca ;};};switch _db {case 0:return CellReference {},_a .Errorf ("\u006e\u006f\u0020\u006cet\u0074\u0065\u0072\u0020\u0070\u0072\u0065\u0066\u0069\u0078\u0020\u0069\u006e\u0020%\u0073",_ea );case -1:return CellReference {},_a .Errorf ("\u006eo\u0020d\u0069\u0067\u0069\u0074\u0073\u0020\u0069\u006e\u0020\u0025\u0073",_ea );
};_cb .Column =_ea [0:_db ];if _ea [_db ]=='$'{_cb .AbsoluteRow =true ;_db ++;};_cb .ColumnIdx =ColumnToIndex (_cb .Column );_ad ,_dfg :=_dg .ParseUint (_ea [_db :],10,32);if _dfg !=nil {return CellReference {},_a .Errorf ("e\u0072\u0072\u006f\u0072 p\u0061r\u0073\u0069\u006e\u0067\u0020r\u006f\u0077\u003a\u0020\u0025\u0073",_dfg );
};if _ad ==0{return CellReference {},_a .Errorf ("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0072\u006f\u0077\u003a \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0030");};_cb .RowIdx =uint32 (_ad );
return _cb ,nil ;};

// ParseRangeReference splits a range reference of the form "A1:B5" into its
// components.
func ParseRangeReference (s string )(_ff ,_eec CellReference ,_fge error ){_fb ,_ef ,_fge :=_dff (s );if _fge !=nil {return CellReference {},CellReference {},_fge ;};_gd :=_de .Split (_ef ,"\u003a");if len (_gd )!=2{return CellReference {},CellReference {},_df .New ("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074");
};if _fb !=""{_gd [0]=_fb +"\u0021"+_gd [0];_gd [1]=_fb +"\u0021"+_gd [1];};_eeg ,_fge :=ParseCellReference (_gd [0]);if _fge !=nil {return CellReference {},CellReference {},_fge ;};_bge ,_fge :=ParseCellReference (_gd [1]);if _fge !=nil {return CellReference {},CellReference {},_fge ;
};return _eeg ,_bge ,nil ;};

// ColumnToIndex maps a column to a zero based index (e.g. A = 0, B = 1, AA = 26)
func ColumnToIndex (col string )uint32 {col =_de .ToUpper (col );_cbf :=uint32 (0);for _ ,_ce :=range col {_cbf *=26;_cbf +=uint32 (_ce -'A'+1);};return _cbf -1;};

// String returns a string representation of ColumnReference.
func (_eaa ColumnReference )String ()string {_eg :=make ([]byte ,0,4);if _eaa .AbsoluteColumn {_eg =append (_eg ,'$');};_eg =append (_eg ,_eaa .Column ...);return string (_eg );};

// Update updates reference to point one of the neighboring columns with respect to the update type after removing a row/column.
func (_eab *ColumnReference )Update (updateType _f .UpdateAction )*ColumnReference {switch updateType {case _f .UpdateActionRemoveColumn :_dd :=_eab ;_dd .ColumnIdx =_eab .ColumnIdx -1;_dd .Column =IndexToColumn (_dd .ColumnIdx );return _dd ;default:return _eab ;
};};

// ColumnReference is a parsed reference to a column.  Input is of the form 'A',
// '$C', etc.
type ColumnReference struct{ColumnIdx uint32 ;Column string ;AbsoluteColumn bool ;SheetName string ;};

// ParseColumnReference parses a column reference of the form 'Sheet1!A' and splits it
// into sheet name and column segments.
func ParseColumnReference (s string )(ColumnReference ,error ){s =_de .TrimSpace (s );if len (s )< 1{return ColumnReference {},_df .New ("\u0063\u006f\u006c\u0075\u006d\u006e \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065\u0020a\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0063\u0068a\u0072a\u0063\u0074\u0065\u0072");
};_ba :=ColumnReference {};_dfd ,_ee ,_fg :=_dff (s );if _fg !=nil {return ColumnReference {},_fg ;};if _dfd !=""{_ba .SheetName =_dfd ;};if _ee [0]=='$'{_ba .AbsoluteColumn =true ;_ee =_ee [1:];};if !_cd .MatchString (_ee ){return ColumnReference {},_df .New ("\u0063\u006f\u006c\u0075\u006dn\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0041\u0020\u0061\u006e\u0064\u0020\u005a\u005a");
};_ba .Column =_ee ;_ba .ColumnIdx =ColumnToIndex (_ba .Column );return _ba ,nil ;};

// IndexToColumn maps a column number to a column name (e.g. 0 = A, 1 = B, 26 = AA)
func IndexToColumn (col uint32 )string {var _cf [64+1]byte ;_ge :=len (_cf );_fc :=col ;const _fa =26;for _fc >=_fa {_ge --;_gfb :=_fc /_fa ;_cf [_ge ]=byte ('A'+uint (_fc -_gfb *_fa ));_fc =_gfb -1;};_ge --;_cf [_ge ]=byte ('A'+uint (_fc ));return string (_cf [_ge :]);
};

// String returns a string representation of CellReference.
func (_c CellReference )String ()string {_ag :=make ([]byte ,0,4);if _c .AbsoluteColumn {_ag =append (_ag ,'$');};_ag =append (_ag ,_c .Column ...);if _c .AbsoluteRow {_ag =append (_ag ,'$');};_ag =_dg .AppendInt (_ag ,int64 (_c .RowIdx ),10);return string (_ag );
};var _cd =_g .MustCompile ("^\u005b\u0061\u002d\u007aA-\u005a]\u0028\u005b\u0061\u002d\u007aA\u002d\u005a\u005d\u003f\u0029\u0024");

// Update updates reference to point one of the neighboring cells with respect to the update type after removing a row/column.
func (_gf *CellReference )Update (updateType _f .UpdateAction )*CellReference {switch updateType {case _f .UpdateActionRemoveColumn :_bf :=_gf ;_bf .ColumnIdx =_gf .ColumnIdx -1;_bf .Column =IndexToColumn (_bf .ColumnIdx );return _bf ;default:return _gf ;
};};func _dff (_agb string )(string ,string ,error ){_fd :="";_bg :=_de .LastIndex (_agb ,"\u0021");if _bg > -1{_fd =_agb [:_bg ];_agb =_agb [_bg +1:];if _fd ==""{return "","",_df .New ("\u0049n\u0076a\u006c\u0069\u0064\u0020\u0073h\u0065\u0065t\u0020\u006e\u0061\u006d\u0065");
};};return _fd ,_agb ,nil ;};

// ParseColumnRangeReference splits a range reference of the form "A:B" into its
// components.
func ParseColumnRangeReference (s string )(_dbf ,_gde ColumnReference ,_efa error ){_ged :="";_geda :=_de .Split (s ,"\u0021");if len (_geda )==2{_ged =_geda [0];s =_geda [1];};_ac :=_de .Split (s ,"\u003a");if len (_ac )!=2{return ColumnReference {},ColumnReference {},_df .New ("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074");
};if _ged !=""{_ac [0]=_ged +"\u0021"+_ac [0];_ac [1]=_ged +"\u0021"+_ac [1];};_fce ,_efa :=ParseColumnReference (_ac [0]);if _efa !=nil {return ColumnReference {},ColumnReference {},_efa ;};_eabg ,_efa :=ParseColumnReference (_ac [1]);if _efa !=nil {return ColumnReference {},ColumnReference {},_efa ;
};return _fce ,_eabg ,nil ;};

// CellReference is a parsed reference to a cell.  Input is of the form 'A1',
// '$C$2', etc.
type CellReference struct{RowIdx uint32 ;ColumnIdx uint32 ;Column string ;AbsoluteColumn bool ;AbsoluteRow bool ;SheetName string ;};