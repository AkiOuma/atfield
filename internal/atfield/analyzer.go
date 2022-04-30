package atfield

import (
	"go/ast"
	"go/token"
)

const (
	selfPackageName  = "atfield"
	linkStructMethod = "LinkStruct"
	linkFieldMethod  = "LinkField"
)

var excludePackage = map[string]struct{}{
	"github.com/AkiOuma/atfield": {},
}

// analyse packages.Imports to extract information of import packages
func (e *engine) analyseImports() {
	for path, v := range e.pkg.Imports {
		if _, ok := excludePackage[path]; ok {
			continue
		} else {
			if e.linkPackages == nil {
				e.linkPackages = make([]*pkg, 0)
			}
			e.linkPackages = append(e.linkPackages, &pkg{
				Alias: v.Name,
				Dir:   v.PkgPath,
			})
		}
	}
}

func (e *engine) analyseSyntaxes() {
	for _, decl := range e.pkg.Syntax[0].Decls {
		e.analyseDecl(decl)
	}
}

func (e *engine) analyseDecl(d ast.Decl) {
	decl, ok := d.(*ast.GenDecl)
	if !ok {
		return
	}
	if decl.Tok != token.VAR {
		return
	}
	for _, s := range decl.Specs {
		e.analyseSpec(s)
	}
}

func (e *engine) analyseSpec(s ast.Spec) {
	spec, ok := s.(*ast.ValueSpec)
	if !ok && len(spec.Values) > 0 {
		return
	}
	linkStructsCache := make([][2]structId, 0)
	linkFieldSetCache := make(linkFieldSet)
	e.analyseValue(spec.Values[0], &linkStructsCache, linkFieldSetCache)
	// set linkStruct to atfield object
	if linkStructsCache != nil {
		e.linkStructs = append(e.linkStructs, linkStructsCache...)
	}
	// set linkFieldSet to atfield object
	for field, set := range linkFieldSetCache {
		if e.linkFieldSet[field] == nil {
			e.linkFieldSet[field] = make(beLinkedFields)
		}
		e.linkFieldSet[field] = set
	}
}

func (e *engine) analyseValue(
	v ast.Expr,
	linkStructsCache *[][2]structId,
	linkFieldSetCache linkFieldSet,
) {
	switch value := v.(type) {
	case *ast.CallExpr:
		e.analyseFun(value.Fun, value.Args, linkStructsCache, linkFieldSetCache)
	case *ast.Ident:
		// clean cache if calls did not from package atfield
		if value.Name != selfPackageName {
			*linkStructsCache = nil
			linkFieldSetCache = nil
		}
	default:
		return
	}
}

func (e *engine) analyseFun(
	f ast.Expr,
	args []ast.Expr,
	linkStructsCache *[][2]structId,
	linkFieldSetCache linkFieldSet,
) {
	switch fun := f.(type) {
	case *ast.SelectorExpr:
		switch fun.Sel.Name {
		case linkStructMethod:
			buildLinkStructs(args, linkStructsCache)
		case linkFieldMethod:
			buildLinkFields(args, linkFieldSetCache)
		default:
		}
		e.analyseValue(fun.X, linkStructsCache, linkFieldSetCache)
	default:
	}
}

func analyseArgument(a ast.Expr) (res []string) {
	var analyse func(a ast.Expr)
	analyse = func(a ast.Expr) {
		switch arg := a.(type) {
		case *ast.SelectorExpr:
			analyse(arg.X)
			res = append(res, arg.Sel.Name)
		case *ast.CompositeLit:
			analyse(arg.Type)
		case *ast.Ident:
			res = append(res, arg.Name)
		}
	}
	analyse(a)
	return
}

func buildLinkStructs(args []ast.Expr, linkStructsCache *[][2]structId) {
	size := len(args)
	if size < 2 {
		return
	}
	for head := 0; head < size-1; head++ {
		for tail := head + 1; tail < size; tail++ {
			argX, argY := analyseArgument(args[head]), analyseArgument(args[tail])
			*linkStructsCache = append(*linkStructsCache, [2]structId{
				{PkgName: argX[0], StructName: argX[1]},
				{PkgName: argY[0], StructName: argY[1]},
			})
		}
	}
}

func buildLinkFields(args []ast.Expr, linkFieldSetCache linkFieldSet) {
	argSlice := make([][]string, 0, len(args))
	for _, v := range args {
		argSlice = append(argSlice, analyseArgument(v))
	}
	for _, a1 := range argSlice {
		fieldDir1 := fieldDir{
			StructId:  structId{PkgName: a1[0], StructName: a1[1]},
			FieldName: a1[2],
		}
		for _, a2 := range argSlice {
			fieldDir2 := fieldDir{
				StructId:  structId{PkgName: a2[0], StructName: a2[1]},
				FieldName: a2[2],
			}
			if fieldDir1 == fieldDir2 {
				continue
			}
			if linkFieldSetCache[fieldDir1] == nil {
				linkFieldSetCache[fieldDir1] = make(beLinkedFields)
			}
			linkFieldSetCache[fieldDir1][fieldDir2.StructId] = fieldDir2.FieldName
		}
	}
}
