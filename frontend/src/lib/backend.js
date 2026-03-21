function getAppBridge() {
	const app = window?.go?.main?.App
	if (!app) {
		throw new Error('Wails backend bridge is unavailable. Start the app with `wails dev` or `wails build`.')
	}

	return app
}

function invoke(method, ...args) {
	const bridge = getAppBridge()
	const target = bridge?.[method]

	if (typeof target !== 'function') {
		throw new Error(`Backend method is unavailable: ${method}`)
	}

	return target(...args)
}

export function bootstrap() {
	return invoke('Bootstrap')
}

export function setLocale(locale) {
	return invoke('SetLocale', locale)
}

export function showWindow() {
	return invoke('ShowWindow')
}

export function hideWindow() {
	return invoke('HideWindow')
}

export function quitApp() {
	return invoke('QuitApp')
}
