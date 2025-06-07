package shi

type Call struct {
	returnPC PC
	sloc Sloc
	target *ScriptMethod
}
