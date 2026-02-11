package logs

import ("strings"
        "unicode/utf8"
)


// Application identifies the application emitting the given log.
func Application(log string) string {
	if log == "" {
        return ""
    }
    
    applications := map[rune]string {
        '❗' : "recommendation",
        '🔍' : "search",
        '☀' : "weather",
    }
    
	var app string = ""
    for _, letter := range log {
     	app, _ = applications[letter]
    	if app != "" {
        	break
    	}   
    }

    if app == "" {
        return "default"
    }
    
    return app
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	if log == "" {
        return ""
    }
    
    return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if log == "" || limit == 0{
        return false
    }
    
    return utf8.RuneCountInString(log) <= limit
}
