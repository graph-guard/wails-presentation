import './style.css'
import App from './App.svelte'
import {Environment} from '../wailsjs/runtime/runtime'

Environment().then((env)=> {
	const doc = document.documentElement
	doc.setAttribute('env-arch', env.arch)
	doc.setAttribute('env-build-type', env.buildType)
	doc.setAttribute('env-platform', env.platform)
})

const app = new App({target: document.body})
export default app
