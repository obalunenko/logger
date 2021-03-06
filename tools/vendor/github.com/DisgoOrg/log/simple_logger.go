package log

import (
	"fmt"
	"log"
	"os"
)

var _ Logger = (*SimpleLogger)(nil)

var std *SimpleLogger = nil

// These flags define which text to prefix to each log entry generated by the Logger.
// Bits are or'ed together to control what's printed.
// Except the Lmsgprefix flag, there is no
// control over the order they appear (the order listed here)
// or the format they present (as described in the comments).
// The prefix is followed by a colon only when Llongfile or Lshortfile
// is specified.
// For example, flags Ldate | Ltime (or LstdFlags) produce,
//	2009/01/23 01:23:23 message
// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
//goland:noinspection GoUnusedConst
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

// Level are different levels at which the SimpleLogger can log
type Level int

// All Level(s) which SimpleLogger supports
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// String returns the name of the Level
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO "
	case LevelWarn:
		return "WARN "
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelPanic:
		return "PANIC"
	default:
		return ""
	}
}

var (
	enableColors = true
	prefixStyle  = ForegroundColorBrightBlack
	levelStyle   = StyleBold
	textStyle    = ForegroundColorWhite
)

var styles = map[Level]Style{
	LevelDebug: ForegroundColorWhite,
	LevelInfo:  ForegroundColorCyan,
	LevelWarn:  ForegroundColorYellow,
	LevelError: ForegroundColorBrightRed,
	LevelFatal: ForegroundColorRed,
	LevelPanic: ForegroundColorMagenta,
}

//Default returns the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Default() *SimpleLogger {
	if std == nil {
		std = New(log.LstdFlags)
	}
	return std
}

// New returns a newInt SimpleLogger implementation
//goland:noinspection GoUnusedExportedFunction
func New(flags int) *SimpleLogger {
	return &SimpleLogger{
		logger: log.New(os.Stderr, "", flags),
		level:  LevelInfo,
	}
}

// SimpleLogger is a wrapper for the std Logger
type SimpleLogger struct {
	logger *log.Logger
	level  Level
	prefix Style
}

// SetLevel sets the lowest Level to log for
func (l *SimpleLogger) SetLevel(level Level) {
	l.level = level
}

// SetFlags sets the log flags like: Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC, Lmsgprefix,LstdFlags
func (l *SimpleLogger) SetFlags(flags int) {
	l.logger.SetFlags(flags)
}

func (l *SimpleLogger) log(level Level, args ...interface{}) {
	if level < l.level {
		return
	}

	if l.prefix != prefixStyle {
		l.prefix = prefixStyle
		l.logger.SetPrefix(prefixStyle.String())
	}

	args = append(args, "", StyleReset)
	copy(args[2:], args)

	levelStr := level.String() + " "
	textStyleStr := ""
	if enableColors {
		levelStr = levelStyle.And(styles[level]).ApplyClear(levelStr)
		textStyleStr = textStyle.String()
	}
	args[0] = levelStr
	args[1] = textStyleStr

	switch level {
	case LevelFatal:
		l.logger.Fatal(args...)
	case LevelPanic:
		l.logger.Panic(args...)
	default:
		l.logger.Print(args...)
	}
}

func (l *SimpleLogger) logf(level Level, format string, args ...interface{}) {
	l.log(level, fmt.Sprintf(format, args...))
}

// Debug logs on the LevelDebug
func (l *SimpleLogger) Debug(args ...interface{}) {
	l.log(LevelDebug, args...)
}

// Debugf logs on the LevelDebug
func (l *SimpleLogger) Debugf(format string, args ...interface{}) {
	l.logf(LevelDebug, format, args...)
}

// Info logs on the LevelInfo
func (l *SimpleLogger) Info(args ...interface{}) {
	l.log(LevelInfo, args...)
}

// Infof logs on the LevelInfo
func (l *SimpleLogger) Infof(format string, args ...interface{}) {
	l.logf(LevelInfo, format, args...)
}

// Warn logs on the LevelWarn
func (l *SimpleLogger) Warn(args ...interface{}) {
	l.log(LevelWarn, args...)
}

// Warnf logs on the LevelWarn
func (l *SimpleLogger) Warnf(format string, args ...interface{}) {
	l.logf(LevelWarn, format, args...)
}

// Error logs on the LevelError
func (l *SimpleLogger) Error(args ...interface{}) {
	l.log(LevelError, args...)
}

// Errorf logs on the LevelError
func (l *SimpleLogger) Errorf(format string, args ...interface{}) {
	l.logf(LevelError, format, args...)
}

// Fatal logs on the LevelFatal
func (l *SimpleLogger) Fatal(args ...interface{}) {
	l.log(LevelFatal, args...)
}

// Fatalf logs on the LevelFatal
func (l *SimpleLogger) Fatalf(format string, args ...interface{}) {
	l.logf(LevelFatal, format, args...)
}

// Panic logs on the LevelPanic
func (l *SimpleLogger) Panic(args ...interface{}) {
	l.log(LevelPanic, args...)
}

// Panicf logs on the LevelPanic
func (l *SimpleLogger) Panicf(format string, args ...interface{}) {
	l.logf(LevelPanic, format, args...)
}

// SetLevel sets the Level of the default Logger
//goland:noinspection GoUnusedExportedFunction
func SetLevel(level Level) {
	Default().SetLevel(level)
}

// SetLevelColor sets the Style of the given Level
//goland:noinspection GoUnusedExportedFunction
func SetLevelColor(level Level, color Style) {
	styles[level] = color
}

// SetLevelStyle sets the default Style of all Level(s)
//goland:noinspection GoUnusedExportedFunction
func SetLevelStyle(style Style) {
	levelStyle = style
}

// EnableColors enables/disables usage of Style(s) in all Logger(s)
//goland:noinspection GoUnusedExportedFunction
func EnableColors(enable bool) {
	enableColors = enable
}

// SetTextColor sets the Style which is used for text of a log message
//goland:noinspection GoUnusedExportedFunction
func SetTextColor(color Style) {
	textStyle = color
}

// SetFlags sets the log flags like: Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC, Lmsgprefix,LstdFlags of the default Logger
//goland:noinspection GoUnusedExportedFunction
func SetFlags(flags int) {
	Default().SetFlags(flags)
}

// Debug logs on the LevelDebug with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Debug(args ...interface{}) {
	Default().Debug(args...)
}

// Debugf logs on the LevelDebug with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Debugf(format string, args ...interface{}) {
	Default().Debugf(format, args...)
}

// Info logs on the LevelInfo with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Info(args ...interface{}) {
	Default().Info(args...)
}

// Infof logs on the LevelInfo with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Infof(format string, args ...interface{}) {
	Default().Infof(format, args...)
}

// Warn logs on the LevelWarn with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Warn(args ...interface{}) {
	Default().Warn(args...)
}

// Warnf logs on the Level with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Warnf(format string, args ...interface{}) {
	Default().Warnf(format, args...)
}

// Error logs on the LevelError with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Error(args ...interface{}) {
	Default().Error(args...)
}

// Errorf logs on the LevelError with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Errorf(format string, args ...interface{}) {
	Default().Errorf(format, args...)
}

// Fatal logs on the LevelFatal with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Fatal(args ...interface{}) {
	Default().Fatal(args...)
}

// Fatalf logs on the LevelFatal with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Fatalf(format string, args ...interface{}) {
	Default().Fatalf(format, args...)
}

// Panic logs on the LevelPanic with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Panic(args ...interface{}) {
	Default().Panic(args...)
}

// Panicf logs on the LevelPanic with the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Panicf(format string, args ...interface{}) {
	Default().Panicf(format, args...)
}
