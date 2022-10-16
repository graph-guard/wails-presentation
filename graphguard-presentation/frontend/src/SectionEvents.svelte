<script lang='ts'>
import EventHandler from './event_handler'
import {Events_HelloWorld_FireOnce, Events_HelloWorld_Start, Events_HelloWorld_Stop} from '../wailsjs/go/main/App'
import {writable} from 'svelte/store'
import {v4 as UUID} from 'uuid'
let helloWorldEventRunning = false

function toggleHelloWorldEvent() {
	if (helloWorldEventRunning) {
		helloWorldEventRunning = false
		Events_HelloWorld_Stop()
		return
	}

	helloWorldEventRunning = true
	Events_HelloWorld_Start()
}

let listeners = writable<{
	[id: string]: {unsub: ()=> void, counter: number}
}>({})

function startListener() {
	if (Object.keys($listeners).length > 4) {return}

	const id = UUID()
	listeners.update(function($) {
		$[id] = ({
			counter: 0,
			unsub: EventHandler.on('HelloWorld', (...args)=> {
				listeners.update(function($) {
					$[id].counter++
					return $
				})
			}).unsub,
		})
		return $
	})
}

function removeListener(id: string) {
	listeners.update(($)=> {
		$[id].unsub()
		delete $[id]
		return $
	})
}
</script>



<section>
	<h1>Events</h1>

	<section>
		<h3>Hello World Event</h3>
		<button on:click={Events_HelloWorld_FireOnce}>
			Fire once
		</button>
		<button on:click={toggleHelloWorldEvent}>
			{!helloWorldEventRunning ? '▶️ Start fire events' : '🟥 Stop'}
		</button>
		<button on:click={startListener} disabled={Object.keys($listeners).length > 5}>
			Add a listener
		</button>
		<div class='listeners'>
			{#each Object.keys($listeners) as lnrID}
				<div>
					<button on:click={()=> removeListener(lnrID)}>X</button>
					<span>{$listeners[lnrID].counter}</span>
				</div>
			{/each}
		</div>
	</section>
</section>



<style>
.listeners {
	display: flex;
	margin: auto;
	gap: 0.5rem;
	padding: 0.5rem;
	justify-content: center;
	align-items: center;
}
.listeners > div {
	background-color: #fff;
	border-radius: 0.25rem;
	box-shadow: 0 1px 3px rgba(0,0,0, 0.5);
}
.listeners > div > button {
	background-color: transparent;
	padding: 0.5rem;
	border-radius: 0.25rem;
	box-shadow: none;
}
.listeners > div > button:hover {
	background-color: rgba(0,0,0, 0.15);
}
.listeners > div > button:active {
	background-color: rgba(0,0,0, 0.25);
}
.listeners > div > span {
	padding: 0.5rem;
	color: #000;
}
</style>
