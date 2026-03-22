import enUS from './en-US'
import zhCN from './zh-CN'

export const DEFAULT_LOCALE = 'zh-CN'
export const FALLBACK_LOCALE = 'en-US'

export const locales = {
  'zh-CN': zhCN,
  'en-US': enUS
}

export const SUPPORTED_LOCALES = Object.freeze(Object.keys(locales))

const LOCALE_RESOLVER_KEY = Symbol.for('wails-skeleton.locale-resolver')
const globalScope = globalThis

function createLocaleResolver() {
  return {
    resolve(locale) {
      return locales[locale] ?? locales[FALLBACK_LOCALE]
    }
  }
}

const localeResolver =
  globalScope[LOCALE_RESOLVER_KEY] ?? (globalScope[LOCALE_RESOLVER_KEY] = createLocaleResolver())

export function resolveLocale(locale) {
  return localeResolver.resolve(locale)
}
