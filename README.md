# Wails presentation

# Setup

It's increadably easy! Just installing all the dependencies and Wails itself,
which is described in their docs. After that you just inititialize a new project by
`wails init` with the desired frontend framework (Svelte, Vue, React, Preact, Lit, or even just Vanilla).
Then it's possible to just write some Go and frontend code or
directly run `wails dev` for a dev environment, which opens a native window containing the frontend,
or you could do `wails build` to get the final binary executable.

# The Go base
```go
err := wails.Run(&options.App{
	Title:            "Window title",
	Width:            1024,
	Height:           768,
	MinWidth:         400,
	MinHeight:        400,
	Assets:           assets,
	BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
	OnStartup:        app.startup,
	Bind: []interface{}{
		app,
	},
	Mac: &mac.Options{
		Appearance:           mac.NSAppearanceNameDarkAqua,
		WebviewIsTransparent: true,
		WindowIsTranslucent:  true,
		About: &mac.AboutInfo{
			Title:   "GraphGuard Presentation",
			Message: "© 2022 GraphGuard",
			Icon:    icon,
		},
	},
	Windows: &windows.Options{
		WebviewIsTransparent: true,
		WindowIsTranslucent: true,
		BackdropType: windows.Auto,
		Theme: windows.SystemDefault,
	},
	/*Linux: &linux.Options{
		Icon: [...], // the only availabe option for linux
	},*/
})
```

Only certain options can be dynamically updated later, like height, width...

# The frontend base

```svelte
<script lang='ts'>
import SectionFibonacci from './SectionFibonacci.svelte'
import SectionEvents from './SectionEvents.svelte'
</script>

<main>
	<svg id='Logo'>[...]</svg>
	<SectionFibonacci/>
	<SectionEvents/>
</main>
```


# How we implemented events
In the frontend events are handled kinda strage. There is an runtime endpoint called "EventsOn",
by which you can listen on a certain event. Another endpoint called "EventsOff" is there to
unsubscribe from the event by the given name. It doesn't only remove one listener, but of all of them. To keep the experience like we know it in JavaScript, we created our own event handler, which is only registering and removing certain listeners.

## Events handler code snippet

```typescript
import {v4 as UUID} from 'uuid'
import {EventsOff, EventsOn} from '../wailsjs/runtime/runtime'

type Callback = ()=> void

class EventHandler {
	private _events: {[eventName: string]: {[id: string]: Callback}} = {}

	public on(eventName: string, fn: Callback): Callback {
		if (
			!this._events[eventName] ||
			Object.keys(this._events[eventName]).length < 1
		) {
			this._events[eventName] = {}
			EventsOn(eventName, (...args)=> {
				const handlers = Object.values(this._events[eventName])
				for (fn of handlers) {
					fn.apply(this, args)
				}
			})
		}

		const id = UUID()
		this._events[eventName][id] = fn
		return ()=> {this.remove(eventName, id)}
	}

	public remove(eventName: string, id: string) {
		delete this._events[eventName][id]
		if (Object.keys(this._events[eventName]).length < 1) {
			EventsOff(eventName)
		}
	}
}

export default new EventHandler()
```


# Gotchas

Wails is still in a early stage, as it is missing quite a lot essential features,
e.g. native context menus, zoom, notifications, and much more.

## Currently cross-compilation is unsupported
It's possible through build pipes in GitHub Actions, or even by Docker.

## Safari is the WebView on macOS
Every frontend developer knows what that means... we all love Safari.
Honsetly it would be great to strictly set it to like Firefox or Chrome.
As we tested on Linux (Ubuntu), we encountered that it's powered on Safari too.
On Windows it's as expected Edge ("Chrome").

## WebAssembly
Wails could potentionally compile to WebAssembly,
just without the OS functionalities.
Like what if you offer a standalone version,
but also use this app for your cloud service in the browser.

## No TypeScript support for models
The conversion between Go structs and JSON. It generates JS Classes,
where the constructor takes an untyped Object and just maps the values by the keys.
First off all: Why JS Classes? What's their purpose?
Secondly: Why to do this at all, when the constructor just takes an untyped object?

## Native dialogs (modals)
Native dialogs are too OS specific and therefore inconsistent.
When defining 4 buttons (the total amount for macOS) then only macOS will show them,
Linux will have just a default "Ok" button.

## Partly implemented toolbar menu
For now there are only two predefined menus: "About" and "Edit". Problem is,
that both only work on macOS. The implementation is all up on you.

## Performance overhead calling Go
Calls to Go code go through websockets and incurring a cost.
Calling a synchronous JavaScript function calculating the fibonacci number at 1.000
iterations will be faster than calling an analogous Go function due to the overhead.
For more complex computations Go can be more efficient.

## No access to the DOM from Go
There is no access point to the DOM from Go. It's not possible to do something with the frontend, like with WebAssambly. Such desires would need to be handled by events.

## Transparent window style only for Win11 and macOS
As all features are OS dependent, it's hard to keep an universal look. The Transparent window feature is unfortunately unavailable for linux. On Windows 11 it's even possible to use different bluring techniques: Acrylic, the frosty looking like used in Windows 10, and Mica, which is used in Windows 11.

## Invisible toolbar on transparent window in Windows 10
As seen on the screenshot, there is no text, it's invisible, but if you click it, it opens the corresponding menu.
