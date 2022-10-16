<script lang='ts'>
import {Fibonacci as go_Fibonacci} from '../wailsjs/go/main/App'

let fibonacciIterations = 50_000_000

function js_Fibonacci(n) {
	var a = 1
	var b = 0
	var temp = 0
	while (n >= 0) {
		temp = a
		a = a + b
		b = temp
		n--
	}
	return b
}
let js_Fibonacci_duration = {start: 0, end: 0}
let go_Fibonacci_duration = {start: 0, end: 0}

function run_js_Fibonacci() {
	js_Fibonacci_duration.start = performance.now()
	js_Fibonacci(fibonacciIterations)
	js_Fibonacci_duration.end = performance.now()
}

async function run_go_Fibonacci() {
	go_Fibonacci_duration.start = performance.now()
	await go_Fibonacci(fibonacciIterations)
	go_Fibonacci_duration.end = performance.now()
}

function readableDuration({start, end}) {
	return formatScientificDuration(end - start)
}

export function fixedNum(val: number, decimals: number = 2) {
	return Number(val.toFixed(decimals))
}

/**
 * formatScientificDuration expects nanoseconds as argument
 */
export function formatScientificDuration(ms: number, decimals = 3) {
	if (ms < 1000) { // if less a microsecond, it's nanoseconds
		return fixedNum(ms, decimals) + ' ms'
	}
	return fixedNum(ms / 1000, decimals).toLocaleString() + ' s'
}
</script>



<section>
	<h1>Fibonacci</h1>
	<input type='number' min=0 bind:value={fibonacciIterations}/>
	<div class='split-content'>
		<div>
			<span class='name'>JavaScript</span>
			<span class='duration'>{readableDuration(js_Fibonacci_duration)}</span>
			<button on:click={run_js_Fibonacci}>Run</button>
		</div>
		<div>
			<span class='name'>Go</span>
			<span class='duration'>{readableDuration(go_Fibonacci_duration)}</span>
			<button on:click={run_go_Fibonacci}>Run</button>
		</div>
	</div>
</section>



<style>
.split-content {
	display: flex;
	justify-content: center;
	flex-flow: row nowrap;
}
.split-content > div {
	padding: 2rem;
}
.split-content > div:not(:last-child) {
	border-right: solid 1px rgba(255, 255, 255, 0.15);
}
.split-content > div:not(:first-child) {
	border-left: solid 1px rgba(0,0,0, 0.25);
}
.split-content > div .name,
.split-content > div .duration {
	display: block;
}
.split-content > div .name {
	font-size: 1.15rem;
}
.split-content > div .duration {
	margin: 0.5rem 0;
	opacity: 0.75;
}
</style>
