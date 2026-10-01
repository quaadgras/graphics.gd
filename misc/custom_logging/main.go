package main

import (
	"errors"
	"fmt"
	"path"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Logger"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/RichTextLabel"
	"graphics.gd/classdb/ScriptBacktrace"
	"graphics.gd/startup"
	"graphics.gd/variant/Callable"
)

type Main struct {
	Control.Extension[Main]

	FlushStdoutOnPrint Label.Instance

	messageCounter       int
	messageRawCounter    int
	messageStderrCounter int
	warningCounter       int
	errorCounter         int
}

func (m *Main) Ready() {
	Engine.Println("Normal message 1.")
	Engine.Raise(errors.New("Error 1."))
	Engine.RaiseWarning("Warning 1.")
	Engine.Raise(errors.New("Error 2."))
	Engine.RaiseWarning("Warning 2.")
	Engine.Println("Normal message 2.")
	Engine.Log("Normal message 1 (stderr).")
	Engine.Log("Normal message 2 (stderr).")
	Engine.Print("Normal message 1 (raw). ")
	Engine.Print("Normal message 2 (raw).\n--------\n")

	flush := false
	switch setting := ProjectSettings.GetSettingWithOverride("application/run/flush_stdout_on_print").(type) {
	case bool:
		flush = setting
	case int:
		flush = setting != 0
	case int64:
		flush = setting != 0
	}
	if flush {
		m.FlushStdoutOnPrint.SetText("Flush stdout on print: Yes (?)")
	} else {
		m.FlushStdoutOnPrint.SetText("Flush stdout on print: No (?)")
	}
}

func (m *Main) OnPrintMessagePressed() {
	m.messageCounter++
	Engine.Println(fmt.Sprintf("Printing message #%d.", m.messageCounter))
}

func (m *Main) OnPrintMessageRawPressed() {
	m.messageRawCounter++
	Engine.Print(fmt.Sprintf("Printing message #%d (raw). ", m.messageRawCounter))
}

func (m *Main) OnPrintMessageStderrPressed() {
	m.messageStderrCounter++
	Engine.Log(fmt.Sprintf("Printing message #%d (stderr).", m.messageStderrCounter))
}

func (m *Main) OnPrintWarningPressed() {
	m.warningCounter++
	Engine.RaiseWarning(fmt.Sprintf("Printing warning #%d.", m.warningCounter))
}

func (m *Main) OnPrintErrorPressed() {
	m.errorCounter++
	Engine.Raise(fmt.Errorf("Printing error #%d.", m.errorCounter))
}

func (m *Main) OnOpenLogsFolderPressed() {
	logPath, _ := ProjectSettings.GetSettingWithOverride("debug/file_logging/log_path").(string)
	OS.ShellOpen(ProjectSettings.GlobalizePath(path.Dir(logPath)))
}

func (m *Main) OnCrashEnginePressed() {
	OS.Crash("Crashing the engine on user request (the Crash Engine button was pressed). Do not report this as a bug.")
}

// CustomLoggerUI is autoloaded (see project.godot) and installs a CustomLogger
// that mirrors all engine log output into its RichTextLabel.
//
// The GDScript original registers the logger in _init to catch messages as
// early as possible; the earliest hook available to a Go scene class is when
// it enters the tree, so a few startup messages may be missed here.
type CustomLoggerUI struct {
	RichTextLabel.Extension[CustomLoggerUI] `gd:"CustomLoggerUI"`

	logger *CustomLogger
}

func (ui *CustomLoggerUI) EnterTree() {
	ui.logger = new(CustomLogger)
	ui.logger.label = ui.AsRichTextLabel()
	OS.AddLogger(ui.logger.AsLogger())
}

// Removing the logger happens automatically when the project exits by default.
// In case you need to remove a custom logger earlier, you can use OS.RemoveLogger.
// Doing so can also avoid object leak warnings that may be printed on exit.
func (ui *CustomLoggerUI) ExitTree() {
	OS.RemoveLogger(ui.logger.AsLogger())
}

// CustomLogger must be thread-safe, as it may be called from non-main threads.
// We use Callable.Defer to modify the label from the main thread, as thread
// guards could prevent the methods from being called successfully otherwise.
type CustomLogger struct {
	Logger.Extension[CustomLogger] `gd:"CustomLogger"`

	label RichTextLabel.Instance
}

func (l *CustomLogger) LogMessage(message string, _ bool) {
	label := l.label
	Callable.Defer(Callable.New(func() {
		label.AppendText(message)
	}))
}

func (l *CustomLogger) LogError(function, file string, line int, code, rationale string, _ bool, errorType int, backtraces []ScriptBacktrace.Instance) {
	var prefix string
	// The column at which to print the trace. Should match the length of the
	// unformatted text above it.
	var traceIndent int

	switch Logger.ErrorType(errorType) {
	case Logger.ErrorTypeError:
		prefix, traceIndent = "[color=#f54][b]ERROR:[/b]", 6
	case Logger.ErrorTypeWarning:
		prefix, traceIndent = "[color=#fd4][b]WARNING:[/b]", 8
	case Logger.ErrorTypeScript:
		prefix, traceIndent = "[color=#f4f][b]SCRIPT ERROR:[/b]", 13
	case Logger.ErrorTypeShader:
		prefix, traceIndent = "[color=#4bf][b]SHADER ERROR:[/b]", 13
	}

	trace := fmt.Sprintf("%*s %s (%s:%d)", traceIndent, "at:", function, file, line)
	var backtracesText string
	for _, backtrace := range backtraces {
		backtracesText += backtrace.MoreArgs().Format(traceIndent-3, 4) + "\n"
	}

	text := fmt.Sprintf("%s %s %s[/color]\n[color=#999]%s[/color]\n[color=#999]%s[/color]",
		prefix, code, rationale, trace, backtracesText)
	label := l.label
	Callable.Defer(Callable.New(func() {
		label.AppendText(text)
	}))
}

func main() {
	classdb.Register[Main]()
	classdb.Register[CustomLoggerUI]()
	classdb.Register[CustomLogger]()
	startup.Scene()
}
