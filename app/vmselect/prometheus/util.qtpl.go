// Code generated by qtc from "util.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/prometheus/util.qtpl:1
package prometheus

//line app/vmselect/prometheus/util.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/querytracer"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/storage"
)

//line app/vmselect/prometheus/util.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/util.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/prometheus/util.qtpl:8
func streammetricNameObject(qw422016 *qt422016.Writer, mn *storage.MetricName) {
//line app/vmselect/prometheus/util.qtpl:8
	qw422016.N().S(`{`)
//line app/vmselect/prometheus/util.qtpl:10
	if len(mn.MetricGroup) > 0 {
//line app/vmselect/prometheus/util.qtpl:10
		qw422016.N().S(`"__name__":`)
//line app/vmselect/prometheus/util.qtpl:11
		qw422016.N().QZ(mn.MetricGroup)
//line app/vmselect/prometheus/util.qtpl:11
		if len(mn.Tags) > 0 {
//line app/vmselect/prometheus/util.qtpl:11
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/util.qtpl:11
		}
//line app/vmselect/prometheus/util.qtpl:12
	}
//line app/vmselect/prometheus/util.qtpl:13
	for j := range mn.Tags {
//line app/vmselect/prometheus/util.qtpl:14
		tag := &mn.Tags[j]

//line app/vmselect/prometheus/util.qtpl:15
		qw422016.N().QZ(tag.Key)
//line app/vmselect/prometheus/util.qtpl:15
		qw422016.N().S(`:`)
//line app/vmselect/prometheus/util.qtpl:15
		qw422016.N().QZ(tag.Value)
//line app/vmselect/prometheus/util.qtpl:15
		if j+1 < len(mn.Tags) {
//line app/vmselect/prometheus/util.qtpl:15
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/util.qtpl:15
		}
//line app/vmselect/prometheus/util.qtpl:16
	}
//line app/vmselect/prometheus/util.qtpl:16
	qw422016.N().S(`}`)
//line app/vmselect/prometheus/util.qtpl:18
}

//line app/vmselect/prometheus/util.qtpl:18
func writemetricNameObject(qq422016 qtio422016.Writer, mn *storage.MetricName) {
//line app/vmselect/prometheus/util.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/util.qtpl:18
	streammetricNameObject(qw422016, mn)
//line app/vmselect/prometheus/util.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/util.qtpl:18
}

//line app/vmselect/prometheus/util.qtpl:18
func metricNameObject(mn *storage.MetricName) string {
//line app/vmselect/prometheus/util.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/util.qtpl:18
	writemetricNameObject(qb422016, mn)
//line app/vmselect/prometheus/util.qtpl:18
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/util.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/util.qtpl:18
	return qs422016
//line app/vmselect/prometheus/util.qtpl:18
}

//line app/vmselect/prometheus/util.qtpl:20
func streammetricRow(qw422016 *qt422016.Writer, timestamp int64, value float64) {
//line app/vmselect/prometheus/util.qtpl:20
	qw422016.N().S(`[`)
//line app/vmselect/prometheus/util.qtpl:21
	qw422016.N().F(float64(timestamp) / 1e3)
//line app/vmselect/prometheus/util.qtpl:21
	qw422016.N().S(`,"`)
//line app/vmselect/prometheus/util.qtpl:21
	qw422016.N().F(value)
//line app/vmselect/prometheus/util.qtpl:21
	qw422016.N().S(`"]`)
//line app/vmselect/prometheus/util.qtpl:22
}

//line app/vmselect/prometheus/util.qtpl:22
func writemetricRow(qq422016 qtio422016.Writer, timestamp int64, value float64) {
//line app/vmselect/prometheus/util.qtpl:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/util.qtpl:22
	streammetricRow(qw422016, timestamp, value)
//line app/vmselect/prometheus/util.qtpl:22
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/util.qtpl:22
}

//line app/vmselect/prometheus/util.qtpl:22
func metricRow(timestamp int64, value float64) string {
//line app/vmselect/prometheus/util.qtpl:22
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/util.qtpl:22
	writemetricRow(qb422016, timestamp, value)
//line app/vmselect/prometheus/util.qtpl:22
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/util.qtpl:22
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/util.qtpl:22
	return qs422016
//line app/vmselect/prometheus/util.qtpl:22
}

//line app/vmselect/prometheus/util.qtpl:24
func streamvaluesWithTimestamps(qw422016 *qt422016.Writer, values []float64, timestamps []int64) {
//line app/vmselect/prometheus/util.qtpl:25
	if len(values) == 0 {
//line app/vmselect/prometheus/util.qtpl:25
		qw422016.N().S(`[]`)
//line app/vmselect/prometheus/util.qtpl:27
		return
//line app/vmselect/prometheus/util.qtpl:28
	}
//line app/vmselect/prometheus/util.qtpl:28
	qw422016.N().S(`[`)
//line app/vmselect/prometheus/util.qtpl:30
	/* inline metricRow call here for the sake of performance optimization */

//line app/vmselect/prometheus/util.qtpl:30
	qw422016.N().S(`[`)
//line app/vmselect/prometheus/util.qtpl:31
	qw422016.N().F(float64(timestamps[0]) / 1e3)
//line app/vmselect/prometheus/util.qtpl:31
	qw422016.N().S(`,"`)
//line app/vmselect/prometheus/util.qtpl:31
	qw422016.N().F(values[0])
//line app/vmselect/prometheus/util.qtpl:31
	qw422016.N().S(`"]`)
//line app/vmselect/prometheus/util.qtpl:33
	timestamps = timestamps[1:]
	values = values[1:]

//line app/vmselect/prometheus/util.qtpl:36
	if len(values) > 0 {
//line app/vmselect/prometheus/util.qtpl:38
		// Remove bounds check inside the loop below
		_ = timestamps[len(values)-1]

//line app/vmselect/prometheus/util.qtpl:41
		for i, v := range values {
//line app/vmselect/prometheus/util.qtpl:42
			/* inline metricRow call here for the sake of performance optimization */

//line app/vmselect/prometheus/util.qtpl:42
			qw422016.N().S(`,[`)
//line app/vmselect/prometheus/util.qtpl:43
			qw422016.N().F(float64(timestamps[i]) / 1e3)
//line app/vmselect/prometheus/util.qtpl:43
			qw422016.N().S(`,"`)
//line app/vmselect/prometheus/util.qtpl:43
			qw422016.N().F(v)
//line app/vmselect/prometheus/util.qtpl:43
			qw422016.N().S(`"]`)
//line app/vmselect/prometheus/util.qtpl:44
		}
//line app/vmselect/prometheus/util.qtpl:45
	}
//line app/vmselect/prometheus/util.qtpl:45
	qw422016.N().S(`]`)
//line app/vmselect/prometheus/util.qtpl:47
}

//line app/vmselect/prometheus/util.qtpl:47
func writevaluesWithTimestamps(qq422016 qtio422016.Writer, values []float64, timestamps []int64) {
//line app/vmselect/prometheus/util.qtpl:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/util.qtpl:47
	streamvaluesWithTimestamps(qw422016, values, timestamps)
//line app/vmselect/prometheus/util.qtpl:47
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/util.qtpl:47
}

//line app/vmselect/prometheus/util.qtpl:47
func valuesWithTimestamps(values []float64, timestamps []int64) string {
//line app/vmselect/prometheus/util.qtpl:47
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/util.qtpl:47
	writevaluesWithTimestamps(qb422016, values, timestamps)
//line app/vmselect/prometheus/util.qtpl:47
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/util.qtpl:47
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/util.qtpl:47
	return qs422016
//line app/vmselect/prometheus/util.qtpl:47
}

//line app/vmselect/prometheus/util.qtpl:49
func streamdumpQueryTrace(qw422016 *qt422016.Writer, qt *querytracer.Tracer) {
//line app/vmselect/prometheus/util.qtpl:50
	traceJSON := qt.ToJSON()

//line app/vmselect/prometheus/util.qtpl:51
	if traceJSON != "" {
//line app/vmselect/prometheus/util.qtpl:51
		qw422016.N().S(`,"trace":`)
//line app/vmselect/prometheus/util.qtpl:51
		qw422016.N().S(traceJSON)
//line app/vmselect/prometheus/util.qtpl:51
	}
//line app/vmselect/prometheus/util.qtpl:52
}

//line app/vmselect/prometheus/util.qtpl:52
func writedumpQueryTrace(qq422016 qtio422016.Writer, qt *querytracer.Tracer) {
//line app/vmselect/prometheus/util.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/util.qtpl:52
	streamdumpQueryTrace(qw422016, qt)
//line app/vmselect/prometheus/util.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/util.qtpl:52
}

//line app/vmselect/prometheus/util.qtpl:52
func dumpQueryTrace(qt *querytracer.Tracer) string {
//line app/vmselect/prometheus/util.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/util.qtpl:52
	writedumpQueryTrace(qb422016, qt)
//line app/vmselect/prometheus/util.qtpl:52
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/util.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/util.qtpl:52
	return qs422016
//line app/vmselect/prometheus/util.qtpl:52
}
