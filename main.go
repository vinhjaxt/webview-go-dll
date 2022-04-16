package main

// cd D:\home\go-packages\pkg\mod\github.com\webview\webview@v0.0.0-20220415202704-5f6562f358d9
// Open:script\build.bat
// Replace: WebView2Loader.dll.lib => WebView2LoaderStatic.lib
// Run: script\build.bat
// Copy: D:\home\go-packages\pkg\mod\github.com\webview\webview@v0.0.0-20220415202704-5f6562f358d9\script\Microsoft.Web.WebView2.1.0.1150.38\build\native\x64\WebView2LoaderStatic.lib
// Copy: D:\home\go-packages\pkg\mod\github.com\webview\webview@v0.0.0-20220415202704-5f6562f358d9\script\Microsoft.Web.WebView2.1.0.1150.38\build\native\x86\WebView2LoaderStatic.lib
// copy "%src_dir%\build\webview.obj" "%src_dir%\build\webview_x64.obj"
// copy "%src_dir%\build\webview.obj" "%src_dir%\build\webview_x86.obj"

// set CGO_ENABLED=1&&go build -ldflags="-H windowsgui -linkmode external -extldflags -static"
// set CGO_ENABLED=1&&go build -ldflags="-H windowsgui"
import (
	"log"

	"github.com/webview/webview"
)

// CheckNetIsolation.exe LoopbackExempt -a -n="Microsoft.Win32WebViewHost_cw5n1h2txyewy"
func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	// w.SetSize(800, 600, webview.HintNone)
	w.Bind("noop", func() string {
		log.Println("hello")
		return "hello"
	})
	w.Bind("add", func(a, b int) int {
		return a + b
	})
	w.Bind("quit", func() {
		w.Terminate()
	})
	w.Navigate(`data:text/html,
		<!doctype html>
		<html>
			<body>hello</body>
			<script>
				window.onload = function() {
					document.body.innerText = "hello, "+navigator.userAgent;
					noop().then(function(res) {
						console.log('noop res', res);
						add(1, 2).then(function(res) {
							console.log('add res', res);
							// quit();
						});
					});
				};
			</script>
		</html>
	)`)
	w.Run()
}
