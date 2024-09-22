// Code generated by _tools/gen_builtin.go; DO NOT EDIT.

package gojq

func init() {
	builtinFuncDefs = map[string][]*FuncDef{
		"IN": {{Name: "IN", Args: []string{"s"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "any", Args: []*Query{{Left: &Query{Func: "s"}, Op: OpEq, Right: &Query{Func: "."}}, {Func: "."}}}}}}, {Name: "IN", Args: []string{"src", "s"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "any", Args: []*Query{{Left: &Query{Func: "src"}, Op: OpEq, Right: &Query{Func: "s"}}, {Func: "."}}}}}}},
		"INDEX": {{Name: "INDEX", Args: []string{"stream", "idx_expr"}, Body: &Query{Term: &Term{Type: TermTypeReduce, Reduce: &Reduce{Query: &Query{Func: "stream"}, Pattern: &Pattern{Name: "$row"}, Start: &Query{Term: &Term{Type: TermTypeObject, Object: &Object{}}}, Update: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Left: &Query{Func: "$row"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "idx_expr"}, Op: OpPipe, Right: &Query{Func: "tostring"}}}}}}, Op: OpAssign, Right: &Query{Func: "$row"}}}}}}, {Name: "INDEX", Args: []string{"idx_expr"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "INDEX", Args: []*Query{{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Iter: true}}}}, {Func: "idx_expr"}}}}}}},
		"JOIN": {{Name: "JOIN", Args: []string{"$idx", "idx_expr"}, Body: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Iter: true}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$idx"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Func: "idx_expr"}}}}}}}}}}}}}}}, {Name: "JOIN", Args: []string{"$idx", "stream", "idx_expr"}, Body: &Query{Left: &Query{Func: "stream"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$idx"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Func: "idx_expr"}}}}}}}}}}}}, {Name: "JOIN", Args: []string{"$idx", "stream", "idx_expr", "join_expr"}, Body: &Query{Left: &Query{Func: "stream"}, Op: OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$idx"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Func: "idx_expr"}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "join_expr"}}}}},
		"_assign": {},
		"_modify": {},
		"add": {{Name: "add", Args: []string{"f"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}, Op: OpPipe, Right: &Query{Func: "add"}}}},
		"all": {{Name: "all", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "all", Args: []*Query{{Func: "."}}}}}}, {Name: "all", Args: []string{"y"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "all", Args: []*Query{{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Iter: true}}}}, {Func: "y"}}}}}}, {Name: "all", Args: []string{"g", "y"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "isempty", Args: []*Query{{Left: &Query{Func: "g"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "y"}, Op: OpPipe, Right: &Query{Func: "not"}}}}}}}}}}}}},
		"any": {{Name: "any", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "any", Args: []*Query{{Func: "."}}}}}}, {Name: "any", Args: []string{"y"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "any", Args: []*Query{{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Iter: true}}}}, {Func: "y"}}}}}}, {Name: "any", Args: []string{"g", "y"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "isempty", Args: []*Query{{Left: &Query{Func: "g"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Func: "y"}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "not"}}}},
		"arrays": {{Name: "arrays", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "array"}}}}}}}}}},
		"booleans": {{Name: "booleans", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "boolean"}}}}}}}}}},
		"capture": {{Name: "capture", Args: []string{"$re"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "capture", Args: []*Query{{Func: "$re"}, {Func: "null"}}}}}}, {Name: "capture", Args: []string{"$re", "$flags"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "match", Args: []*Query{{Func: "$re"}, {Func: "$flags"}}}}}, Op: OpPipe, Right: &Query{Func: "_capture"}}}},
		"combinations": {{Name: "combinations", Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "length"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{}}}, Else: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, SuffixList: []*Suffix{{Iter: true}, {Bind: &Bind{Patterns: []*Pattern{{Name: "$x"}}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "$x"}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}, IsSlice: true}}}, Op: OpPipe, Right: &Query{Func: "combinations"}}}}}}}}}}}}}}, {Name: "combinations", Args: []string{"n"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "limit", Args: []*Query{{Func: "n"}, {Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "repeat", Args: []*Query{{Func: "."}}}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "combinations"}}}},
		"del": {{Name: "del", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "delpaths", Args: []*Query{{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{{Func: "f"}}}}}}}}}}}}}},
		"finites": {{Name: "finites", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Func: "isfinite"}}}}}}},
		"first": {{Name: "first", Body: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}}, {Name: "first", Args: []string{"g"}, Body: &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{Ident: "$out", Body: &Query{Left: &Query{Func: "g"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeBreak, Break: "$out"}}}}}}}}},
		"fromdate": {{Name: "fromdate", Body: &Query{Func: "fromdateiso8601"}}},
		"fromdateiso8601": {{Name: "fromdateiso8601", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "strptime", Args: []*Query{{Term: &Term{Type: TermTypeString, Str: &String{Str: "%Y-%m-%dT%H:%M:%S%z"}}}}}}}, Op: OpPipe, Right: &Query{Func: "mktime"}}}},
		"fromstream": {{Name: "fromstream", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeForeach, Foreach: &Foreach{Query: &Query{Func: "f"}, Pattern: &Pattern{Name: "$pv"}, Start: &Query{Func: "null"}, Update: &Query{Left: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "e"}}}, Then: &Query{Func: "null"}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$pv"}, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Array: []*Pattern{{Name: "$p"}, {Name: "$v"}}}}, Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "$pv"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "length"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "2"}}}}, Then: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "setpath", Args: []*Query{{Left: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "v"}}}}}}, Op: OpAdd, Right: &Query{Func: "$p"}}, {Func: "$v"}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "setpath", Args: []*Query{{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "e"}}}}}}, {Left: &Query{Func: "$p"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "length"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}}}}}, Else: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "setpath", Args: []*Query{{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "e"}}}}}}, {Left: &Query{Func: "$p"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "length"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}}}}}}}}}}}}, Extract: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "e"}}}, Then: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "v"}}}, Else: &Query{Func: "empty"}}}}}}}}},
		"group_by": {{Name: "group_by", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_group_by", Args: []*Query{{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}}}}}}}}}}},
		"gsub": {{Name: "gsub", Args: []string{"$re", "str"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "sub", Args: []*Query{{Func: "$re"}, {Func: "str"}, {Term: &Term{Type: TermTypeString, Str: &String{Str: "g"}}}}}}}}, {Name: "gsub", Args: []string{"$re", "str", "$flags"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "sub", Args: []*Query{{Func: "$re"}, {Func: "str"}, {Left: &Query{Func: "$flags"}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "g"}}}}}}}}}},
		"in": {{Name: "in", Args: []string{"xs"}, Body: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$x"}}, Body: &Query{Left: &Query{Func: "xs"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "has", Args: []*Query{{Func: "$x"}}}}}}}}}}}}},
		"inputs": {{Name: "inputs", Body: &Query{Term: &Term{Type: TermTypeTry, Try: &Try{Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "repeat", Args: []*Query{{Func: "input"}}}}}, Catch: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "."}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "break"}}}}, Then: &Query{Func: "empty"}, Else: &Query{Func: "error"}}}}}}}}},
		"inside": {{Name: "inside", Args: []string{"xs"}, Body: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$x"}}, Body: &Query{Left: &Query{Func: "xs"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "contains", Args: []*Query{{Func: "$x"}}}}}}}}}}}}},
		"isempty": {{Name: "isempty", Args: []string{"g"}, Body: &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{Ident: "$out", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Func: "g"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "false"}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeBreak, Break: "$out"}}}}}}, Op: OpComma, Right: &Query{Func: "true"}}}}}}},
		"iterables": {{Name: "iterables", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "type"}, Op: OpPipe, Right: &Query{Left: &Query{Left: &Query{Func: "."}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "array"}}}}, Op: OpOr, Right: &Query{Left: &Query{Func: "."}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "object"}}}}}}}}}}}},
		"last": {{Name: "last", Body: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeUnary, Unary: &Unary{Op: OpSub, Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}}}},
		"limit": {{Name: "limit", Args: []string{"$n", "g"}, Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "$n"}, Op: OpGt, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{Ident: "$out", Body: &Query{Term: &Term{Type: TermTypeForeach, Foreach: &Foreach{Query: &Query{Func: "g"}, Pattern: &Pattern{Name: "$item"}, Start: &Query{Func: "$n"}, Update: &Query{Left: &Query{Func: "."}, Op: OpSub, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}, Extract: &Query{Left: &Query{Func: "$item"}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "."}, Op: OpLe, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Term: &Term{Type: TermTypeBreak, Break: "$out"}}, Else: &Query{Func: "empty"}}}}}}}}}}}, Elif: []*IfElif{{Cond: &Query{Left: &Query{Func: "$n"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Func: "empty"}}}, Else: &Query{Func: "g"}}}}}},
		"map": {{Name: "map", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Iter: true}}}}, Op: OpPipe, Right: &Query{Func: "f"}}}}}}},
		"map_values": {{Name: "map_values", Args: []string{"f"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Iter: true}}}}, Op: OpModify, Right: &Query{Func: "f"}}}},
		"match": {{Name: "match", Args: []string{"$re"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "match", Args: []*Query{{Func: "$re"}, {Func: "null"}}}}}}, {Name: "match", Args: []string{"$re", "$flags"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_match", Args: []*Query{{Func: "$re"}, {Func: "$flags"}, {Func: "false"}}}, SuffixList: []*Suffix{{Iter: true}}}}}},
		"max_by": {{Name: "max_by", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_max_by", Args: []*Query{{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}}}}}}}}}}},
		"min_by": {{Name: "min_by", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_min_by", Args: []*Query{{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}}}}}}}}}}},
		"normals": {{Name: "normals", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Func: "isnormal"}}}}}}},
		"not": {{Name: "not", Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Func: "."}, Then: &Query{Func: "false"}, Else: &Query{Func: "true"}}}}}},
		"nth": {{Name: "nth", Args: []string{"$n"}, Body: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Func: "$n"}}}}}, {Name: "nth", Args: []string{"$n", "g"}, Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "$n"}, Op: OpLt, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "error", Args: []*Query{{Term: &Term{Type: TermTypeString, Str: &String{Str: "nth doesn't support negative indices"}}}}}}}, Else: &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{Ident: "$out", Body: &Query{Term: &Term{Type: TermTypeForeach, Foreach: &Foreach{Query: &Query{Func: "g"}, Pattern: &Pattern{Name: "$item"}, Start: &Query{Left: &Query{Func: "$n"}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}, Update: &Query{Left: &Query{Func: "."}, Op: OpSub, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "1"}}}, Extract: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "."}, Op: OpLe, Right: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}, Then: &Query{Left: &Query{Func: "$item"}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeBreak, Break: "$out"}}}, Else: &Query{Func: "empty"}}}}}}}}}}}}}}},
		"nulls": {{Name: "nulls", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "."}, Op: OpEq, Right: &Query{Func: "null"}}}}}}}},
		"numbers": {{Name: "numbers", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "number"}}}}}}}}}},
		"objects": {{Name: "objects", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "object"}}}}}}}}}},
		"paths": {{Name: "paths", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{{Func: ".."}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "."}, Op: OpNe, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{}}}}}}}}}}, {Name: "paths", Args: []string{"f"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{{Left: &Query{Func: ".."}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Func: "f"}}}}}}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "."}, Op: OpNe, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{}}}}}}}}}}},
		"pick": {{Name: "pick", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$v"}}, Body: &Query{Term: &Term{Type: TermTypeReduce, Reduce: &Reduce{Query: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{{Func: "f"}}}}}, Pattern: &Pattern{Name: "$p"}, Start: &Query{Func: "null"}, Update: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "setpath", Args: []*Query{{Func: "$p"}, {Left: &Query{Func: "$v"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "getpath", Args: []*Query{{Func: "$p"}}}}}}}}}}}}}}}}}}}},
		"range": {{Name: "range", Args: []string{"$end"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_range", Args: []*Query{{Term: &Term{Type: TermTypeNumber, Number: "0"}}, {Func: "$end"}, {Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}}, {Name: "range", Args: []string{"$start", "$end"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_range", Args: []*Query{{Func: "$start"}, {Func: "$end"}, {Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}}, {Name: "range", Args: []string{"$start", "$end", "$step"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_range", Args: []*Query{{Func: "$start"}, {Func: "$end"}, {Func: "$step"}}}}}}},
		"recurse": {{Name: "recurse", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "recurse", Args: []*Query{{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Iter: true}, {Optional: true}}}}}}}}}, {Name: "recurse", Args: []string{"f"}, Body: &Query{FuncDefs: []*FuncDef{{Name: "r", Body: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Func: "f"}, Op: OpPipe, Right: &Query{Func: "r"}}}}}}}, Func: "r"}}, {Name: "recurse", Args: []string{"f", "cond"}, Body: &Query{FuncDefs: []*FuncDef{{Name: "r", Body: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Func: "f"}, Op: OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Func: "cond"}}}}}, Op: OpPipe, Right: &Query{Func: "r"}}}}}}}}, Func: "r"}}},
		"repeat": {{Name: "repeat", Args: []string{"f"}, Body: &Query{FuncDefs: []*FuncDef{{Name: "_repeat", Body: &Query{Left: &Query{Func: "f"}, Op: OpComma, Right: &Query{Func: "_repeat"}}}}, Func: "_repeat"}}},
		"scalars": {{Name: "scalars", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "type"}, Op: OpPipe, Right: &Query{Left: &Query{Left: &Query{Func: "."}, Op: OpNe, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "array"}}}}, Op: OpAnd, Right: &Query{Left: &Query{Func: "."}, Op: OpNe, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "object"}}}}}}}}}}}},
		"scan": {{Name: "scan", Args: []string{"$re"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "scan", Args: []*Query{{Func: "$re"}, {Func: "null"}}}}}}, {Name: "scan", Args: []string{"$re", "$flags"}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "match", Args: []*Query{{Func: "$re"}, {Left: &Query{Func: "$flags"}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "g"}}}}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "captures"}}}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{}}}}, Then: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "string"}}}, Else: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "captures"}, SuffixList: []*Suffix{{Iter: true}, {Index: &Index{Name: "string"}}}}}}}}}}}}}},
		"select": {{Name: "select", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Func: "f"}, Then: &Query{Func: "."}, Else: &Query{Func: "empty"}}}}}},
		"sort_by": {{Name: "sort_by", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_sort_by", Args: []*Query{{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}}}}}}}}}}},
		"splits": {{Name: "splits", Args: []string{"$re"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "splits", Args: []*Query{{Func: "$re"}, {Func: "null"}}}}}}, {Name: "splits", Args: []string{"$re", "$flags"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "split", Args: []*Query{{Func: "$re"}, {Func: "$flags"}}}, SuffixList: []*Suffix{{Iter: true}}}}}},
		"strings": {{Name: "strings", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "string"}}}}}}}}}},
		"sub": {{Name: "sub", Args: []string{"$re", "str"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "sub", Args: []*Query{{Func: "$re"}, {Func: "str"}, {Func: "null"}}}}}}, {Name: "sub", Args: []string{"$re", "str", "$flags"}, Body: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$str"}}, Body: &Query{FuncDefs: []*FuncDef{{Name: "_sub", Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "matches"}}}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{}}}}, Then: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$str"}, SuffixList: []*Suffix{{Index: &Index{End: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "offset"}}}, IsSlice: true}}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "string"}}}}, Else: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "matches"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Term: &Term{Type: TermTypeUnary, Unary: &Unary{Op: OpSub, Term: &Term{Type: TermTypeNumber, Number: "1"}}}}}}, {Bind: &Bind{Patterns: []*Pattern{{Name: "$r"}}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeObject, Object: &Object{KeyVals: []*ObjectKeyVal{{Key: "string", Val: &Query{Left: &Query{Left: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Func: "$r"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "_capture"}, Op: OpPipe, Right: &Query{Func: "str"}}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$str"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$r"}, SuffixList: []*Suffix{{Index: &Index{Name: "offset"}}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$r"}, SuffixList: []*Suffix{{Index: &Index{Name: "length"}}}}}}, End: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "offset"}}}, IsSlice: true}}}}}}, Op: OpAdd, Right: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "string"}}}}}, {Key: "offset", Val: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "$r"}, SuffixList: []*Suffix{{Index: &Index{Name: "offset"}}}}}}, {Key: "matches", Val: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Name: "matches"}, SuffixList: []*Suffix{{Index: &Index{End: &Query{Term: &Term{Type: TermTypeUnary, Unary: &Unary{Op: OpSub, Term: &Term{Type: TermTypeNumber, Number: "1"}}}}, IsSlice: true}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "_sub"}}}}}}}}}}}}, Left: &Query{Term: &Term{Type: TermTypeObject, Object: &Object{KeyVals: []*ObjectKeyVal{{Key: "string", Val: &Query{Term: &Term{Type: TermTypeString, Str: &String{}}}}, {Key: "matches", Val: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "match", Args: []*Query{{Func: "$re"}, {Func: "$flags"}}}}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "_sub"}}}}}}}}},
		"test": {{Name: "test", Args: []string{"$re"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "test", Args: []*Query{{Func: "$re"}, {Func: "null"}}}}}}, {Name: "test", Args: []string{"$re", "$flags"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_match", Args: []*Query{{Func: "$re"}, {Func: "$flags"}, {Func: "true"}}}}}}},
		"todate": {{Name: "todate", Body: &Query{Func: "todateiso8601"}}},
		"todateiso8601": {{Name: "todateiso8601", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "strftime", Args: []*Query{{Term: &Term{Type: TermTypeString, Str: &String{Str: "%Y-%m-%dT%H:%M:%SZ"}}}}}}}}},
		"tostream": {{Name: "tostream", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{{FuncDefs: []*FuncDef{{Name: "r", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Iter: true}, {Optional: true}}}}, Op: OpPipe, Right: &Query{Func: "r"}}}}, Op: OpComma, Right: &Query{Func: "."}}}}, Func: "r"}}}, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$p"}}, Body: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "getpath", Args: []*Query{{Func: "$p"}}}}}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeReduce, Reduce: &Reduce{Query: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "path", Args: []*Query{{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Iter: true}, {Optional: true}}}}}}}}, Pattern: &Pattern{Name: "$q"}, Start: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "$p"}, Op: OpComma, Right: &Query{Func: "."}}}}}, Update: &Query{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Func: "$p"}, Op: OpAdd, Right: &Query{Func: "$q"}}}}}}}}}}}}}}}},
		"truncate_stream": {{Name: "truncate_stream", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$n"}}, Body: &Query{Left: &Query{Func: "null"}, Op: OpPipe, Right: &Query{Left: &Query{Func: "f"}, Op: OpPipe, Right: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}, Op: OpPipe, Right: &Query{Left: &Query{Func: "length"}, Op: OpGt, Right: &Query{Func: "$n"}}}, Then: &Query{Left: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: TermTypeNumber, Number: "0"}}}}}, Op: OpModify, Right: &Query{Term: &Term{Type: TermTypeIndex, Index: &Index{Start: &Query{Func: "$n"}, IsSlice: true}}}}, Else: &Query{Func: "empty"}}}}}}}}}}}}},
		"unique_by": {{Name: "unique_by", Args: []string{"f"}, Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "_unique_by", Args: []*Query{{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{{Term: &Term{Type: TermTypeArray, Array: &Array{Query: &Query{Func: "f"}}}}}}}}}}}}}},
		"until": {{Name: "until", Args: []string{"cond", "next"}, Body: &Query{FuncDefs: []*FuncDef{{Name: "_until", Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Func: "cond"}, Then: &Query{Func: "."}, Else: &Query{Left: &Query{Func: "next"}, Op: OpPipe, Right: &Query{Func: "_until"}}}}}}}, Func: "_until"}}},
		"values": {{Name: "values", Body: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Func: "."}, Op: OpNe, Right: &Query{Func: "null"}}}}}}}},
		"walk": {{Name: "walk", Args: []string{"f"}, Body: &Query{FuncDefs: []*FuncDef{{Name: "_walk", Body: &Query{Left: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "array"}}}}, Then: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{{Func: "_walk"}}}}}, Elif: []*IfElif{{Cond: &Query{Left: &Query{Func: "type"}, Op: OpEq, Right: &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: "object"}}}}, Then: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map_values", Args: []*Query{{Func: "_walk"}}}}}}}}}}, Op: OpPipe, Right: &Query{Func: "f"}}}}, Func: "_walk"}}},
		"while": {{Name: "while", Args: []string{"cond", "update"}, Body: &Query{FuncDefs: []*FuncDef{{Name: "_while", Body: &Query{Term: &Term{Type: TermTypeIf, If: &If{Cond: &Query{Func: "cond"}, Then: &Query{Left: &Query{Func: "."}, Op: OpComma, Right: &Query{Term: &Term{Type: TermTypeQuery, Query: &Query{Left: &Query{Func: "update"}, Op: OpPipe, Right: &Query{Func: "_while"}}}}}, Else: &Query{Func: "empty"}}}}}}, Func: "_while"}}},
		"with_entries": {{Name: "with_entries", Args: []string{"f"}, Body: &Query{Left: &Query{Func: "to_entries"}, Op: OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: TermTypeFunc, Func: &Func{Name: "map", Args: []*Query{{Func: "f"}}}}}, Op: OpPipe, Right: &Query{Func: "from_entries"}}}}},
	}
}
