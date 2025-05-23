// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.865
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func IndexPage() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Geolocation Track Visualization</title><link rel=\"stylesheet\" href=\"https://unpkg.com/leaflet@1.9.4/dist/leaflet.css\"><link rel=\"stylesheet\" href=\"/static/style.css\"><script src=\"https://unpkg.com/htmx.org@1.9.10\"></script><script src=\"https://unpkg.com/leaflet@1.9.4/dist/leaflet.js\"></script></head><body><div id=\"map\"></div><div class=\"controls\"><div class=\"speed-control\"><label for=\"speedSlider\">Animation Speed:</label> <input type=\"range\" id=\"speedSlider\" min=\"-25\" max=\"25\" step=\"0.5\" value=\"1\" hx-post=\"/api/control\" hx-trigger=\"change\" hx-vals=\"{&#34;action&#34;: &#34;speed&#34;, &#34;value&#34;: &#34;this.value&#34;}\"> <span id=\"speedValue\">0x</span></div><div class=\"button-controls\"><button id=\"pauseButton\" hx-post=\"/api/control\" hx-trigger=\"click\" hx-vals=\"{&#34;action&#34;: &#34;pause&#34;}\">Pause</button> <button id=\"resetButton\" hx-post=\"/api/control\" hx-trigger=\"click\" hx-vals=\"{&#34;action&#34;: &#34;reset&#34;}\">Reset</button></div><div class=\"track-select\"><select id=\"trackSelect\" hx-get=\"/api/tracks\" hx-trigger=\"change\"><option value=\"1\">1</option> <option value=\"2\">2</option> <option value=\"3\">3</option></select></div></div><script src=\"/static/script.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
