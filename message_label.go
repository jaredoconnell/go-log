package log

import "strings"

// Labels are applied to a message to indicate where they are coming from or other relevant data.
type Labels map[string]string

func (l Labels) String() string {
    labelList := make([]string, len(l))
    i := 0
    for k, v := range l {
        labelList[i] = k + "=" + v
        i++
    }
    return strings.Join(labelList, ";")
}
