package shi

type Call struct {
	returnPc PC
	sloc     Sloc
	target   *ScriptMethod
}
