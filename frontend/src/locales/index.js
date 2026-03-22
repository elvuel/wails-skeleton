import enUS from './en-US'
import zhCN from './zh-CN'

export const DEFAULT_LOCALE = 'zh-CN'
export const FALLBACK_LOCALE = 'en-US'

export const locales = {
  'zh-CN': zhCN,
  'en-US': enUS
}

export const SUPPORTED_LOCALES = Object.freeze(Object.keys(locales))

export function resolveLocale(locale) {
  return locales[locale] ?? locales[FALLBACK_LOCALE]
}
