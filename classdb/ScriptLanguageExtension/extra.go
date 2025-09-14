package ScriptLanguageExtension

import gd "graphics.gd/internal"

type ProfilingInfo = gd.ScriptLanguageExtensionProfilingInfo

type TemplateLocation int

const (
	TemplateBuiltIn TemplateLocation = iota
	TemplateEditor
	TemplateProject
)
