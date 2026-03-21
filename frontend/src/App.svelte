<script>
  import { onMount } from 'svelte'
  import { bootstrap, hideWindow, quitApp, setLocale, showWindow } from './lib/backend'

  const copy = {
    'zh-CN': {
      badge: 'Wails + Svelte Skeleton',
      title: '桌面骨架已就绪',
      subtitle: '内置多语言和窗口控制入口，适合直接往业务层扩展。',
      status: '运行状态',
      features: '能力面板',
      actions: '窗口动作',
      tips: '说明',
      appStatus: '应用状态',
      locale: '当前语言',
      supportedLocales: '支持语言',
      localeHint: '切换语言会立即更新界面文案',
      showWindow: '显示窗口',
      hideWindow: '隐藏窗口',
      quit: '退出应用',
      startupError: '启动错误',
      noError: '无',
      notes: [
        '当前示例聚焦 Wails 与 Svelte 的基础通信。',
        '语言切换调用后端方法并实时刷新界面内容。',
        '这个首页仅提供骨架状态，你可以在此基础上继续拆业务模块。'
      ]
    },
    'en-US': {
      badge: 'Wails + Svelte Skeleton',
      title: 'Desktop skeleton is ready',
      subtitle: 'Built-in i18n and window controls so you can move straight into app features.',
      status: 'Runtime Status',
      features: 'Feature Panel',
      actions: 'Window Actions',
      tips: 'Notes',
      appStatus: 'App Status',
      locale: 'Current Locale',
      supportedLocales: 'Supported Locales',
      localeHint: 'Switching locale updates UI copy immediately',
      showWindow: 'Show Window',
      hideWindow: 'Hide Window',
      quit: 'Quit App',
      startupError: 'Startup Error',
      noError: 'None',
      notes: [
        'This example focuses on the baseline Wails and Svelte bridge.',
        'Locale changes call backend methods and refresh the UI in place.',
        'This landing page is intentionally minimal and meant to be replaced by your domain modules.'
      ]
    }
  }

  const fallbackState = {
    appName: 'wails-skeleton',
    locale: 'zh-CN',
    supportedLocales: ['zh-CN', 'en-US'],
    message: '',
    startupError: ''
  }

  let state = { ...fallbackState }
  let locale = fallbackState.locale
  let busy = false

  $: text = copy[locale] ?? copy['en-US']

  onMount(async () => {
    await loadState()
  })

  async function loadState() {
    busy = true
    try {
      const next = await bootstrap()
      state = { ...fallbackState, ...next }
      locale = state.locale || fallbackState.locale
    } catch (error) {
      state = {
        ...state,
        startupError: error?.message ?? String(error),
        message: error?.message ?? String(error)
      }
    } finally {
      busy = false
    }
  }

  async function handleLocaleChange(nextLocale) {
    if (nextLocale === locale) {
      return
    }

    busy = true
    try {
      const next = await setLocale(nextLocale)
      state = { ...state, ...next }
      locale = state.locale
    } catch (error) {
      state = {
        ...state,
        startupError: error?.message ?? String(error)
      }
    } finally {
      busy = false
    }
  }
</script>

<svelte:head>
  <title>{state.appName}</title>
</svelte:head>

<main class="shell">
  <section class="hero">
    <div class="hero-copy">
      <p class="badge">{text.badge}</p>
      <h1>{text.title}</h1>
      <p class="subtitle">{text.subtitle}</p>
    </div>

    <div class="locale-switcher">
      {#each state.supportedLocales as entry}
        <button
          class:active={entry === locale}
          disabled={busy}
          on:click={() => handleLocaleChange(entry)}
          type="button"
        >
          {entry}
        </button>
      {/each}
    </div>
  </section>

  <section class="grid">
    <article class="panel panel-accent">
      <div class="panel-header">
        <span>{text.status}</span>
        <span class="muted">{state.message || '...'}</span>
      </div>

      <div class="metric-grid">
        <div class="metric">
          <span>{text.appStatus}</span>
          <strong>{state.message || 'ready'}</strong>
        </div>

        <div class="metric">
          <span>{text.locale}</span>
          <strong>{locale}</strong>
        </div>

        <div class="metric">
          <span>{text.supportedLocales}</span>
          <strong>{(state.supportedLocales || []).join(', ') || '-'}</strong>
        </div>
      </div>
    </article>

    <article class="panel">
      <div class="panel-header">
        <span>{text.features}</span>
        <span class="muted">{state.appName}</span>
      </div>

      <div class="detail-row">
        <span>{text.status}</span>
        <code>{state.message || 'ready'}</code>
      </div>

      <div class="detail-row">
        <span>{text.startupError}</span>
        <code>{state.startupError || text.noError}</code>
      </div>

      <div class="detail-row">
        <span>{text.locale}</span>
        <code>{text.localeHint}</code>
      </div>
    </article>

    <article class="panel">
      <div class="panel-header">
        <span>{text.actions}</span>
        <span class="muted">{busy ? 'busy' : 'ready'}</span>
      </div>

      <div class="actions">
        <button disabled={busy} on:click={showWindow} type="button">{text.showWindow}</button>
        <button class="ghost" disabled={busy} on:click={hideWindow} type="button">{text.hideWindow}</button>
        <button class="danger" disabled={busy} on:click={quitApp} type="button">{text.quit}</button>
      </div>
    </article>

    <article class="panel">
      <div class="panel-header">
        <span>{text.tips}</span>
        <span class="muted">{state.appName}</span>
      </div>

      <ul class="notes">
        {#each text.notes as note}
          <li>{note}</li>
        {/each}
      </ul>
    </article>
  </section>
</main>
