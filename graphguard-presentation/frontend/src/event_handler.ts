import {v4 as UUID} from 'uuid'
import {EventsOff, EventsOn} from '../wailsjs/runtime/runtime'

type Callback = ()=> void

class EventHandler {
	private _events: {[eventName: string]: {[id: string]: Callback}} = {}

	public on(eventName: string, fn: Callback): {
		unsub: Callback,
		id: string,
	} {
		// initiate the actual listener
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
		return {
			unsub: ()=> {this.remove(eventName, id)},
			id,
		}
	}

	public remove(eventName: string, id: string) {
		delete this._events[eventName][id]
		if (Object.keys(this._events[eventName]).length < 1) {
			EventsOff(eventName)
		}
	}
}

export default new EventHandler()
