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

<main class="p-8 max-[900px]:p-[18px]">
  <section
    class="grid grid-cols-[minmax(0,1fr)_auto] items-start gap-6 rounded-[28px] border border-[rgba(66,40,14,0.12)] bg-[rgba(255,250,240,0.82)] p-8 shadow-[0_28px_80px_rgba(87,46,18,0.14)] backdrop-blur-[16px] max-[900px]:grid-cols-1 max-[900px]:p-6"
  >
    <div>
      <p class="m-0 text-[0.82rem] font-extrabold uppercase tracking-[0.18em] text-[#8f3811]">{text.badge}</p>
      <h1 class="mb-3 mt-2 text-[clamp(2.4rem,4vw,4.4rem)] leading-[0.94] tracking-[-0.05em]">{text.title}</h1>
      <p class="m-0 max-w-[54rem] text-[1.02rem] leading-[1.7] text-[#6b5a4b]">{text.subtitle}</p>
    </div>

    <div class="flex flex-wrap gap-[10px]">
      {#each state.supportedLocales as entry}
        <button
          class={`cursor-pointer rounded-full border-0 px-[18px] py-3 transition-[transform,background-color,color,opacity] duration-150 ease-out hover:-translate-y-px disabled:cursor-not-allowed disabled:opacity-[0.55] ${
            entry === locale ? 'bg-[#bf5a2b] text-[#fffaf2]' : 'bg-[rgba(35,23,15,0.06)] text-[#23170f]'
          }`}
          disabled={busy}
          onclick={() => handleLocaleChange(entry)}
          type="button"
        >
          {entry}
        </button>
      {/each}
    </div>
  </section>

  <section class="mt-5 grid grid-cols-2 gap-5 max-[900px]:grid-cols-1">
    <article
      class="rounded-3xl border border-[rgba(66,40,14,0.12)] bg-[linear-gradient(160deg,rgba(191,90,43,0.11),rgba(255,255,255,0.74)),rgba(255,255,255,0.72)] p-6 shadow-[0_18px_36px_rgba(96,53,23,0.08)]"
    >
      <div class="mb-[18px] flex items-center justify-between gap-3 text-[0.92rem] font-extrabold uppercase tracking-[0.08em]">
        <span>{text.status}</span>
        <span class="text-[0.78rem] font-bold text-[#6b5a4b]">{state.message || '...'}</span>
      </div>

      <div class="grid grid-cols-2 gap-[14px] max-[900px]:grid-cols-1">
        <div class="flex flex-col gap-2 rounded-[18px] bg-[rgba(255,255,255,0.64)] p-4">
          <span class="text-[0.82rem] uppercase tracking-[0.08em] text-[#6b5a4b]">{text.appStatus}</span>
          <strong class="text-[1.02rem] leading-[1.5]">{state.message || 'ready'}</strong>
        </div>

        <div class="flex flex-col gap-2 rounded-[18px] bg-[rgba(255,255,255,0.64)] p-4">
          <span class="text-[0.82rem] uppercase tracking-[0.08em] text-[#6b5a4b]">{text.locale}</span>
          <strong class="text-[1.02rem] leading-[1.5]">{locale}</strong>
        </div>

        <div class="flex flex-col gap-2 rounded-[18px] bg-[rgba(255,255,255,0.64)] p-4">
          <span class="text-[0.82rem] uppercase tracking-[0.08em] text-[#6b5a4b]">{text.supportedLocales}</span>
          <strong class="text-[1.02rem] leading-[1.5]">{(state.supportedLocales || []).join(', ') || '-'}</strong>
        </div>
      </div>
    </article>

    <article class="rounded-3xl border border-[rgba(66,40,14,0.12)] bg-[rgba(255,255,255,0.72)] p-6 shadow-[0_18px_36px_rgba(96,53,23,0.08)]">
      <div class="mb-[18px] flex items-center justify-between gap-3 text-[0.92rem] font-extrabold uppercase tracking-[0.08em]">
        <span>{text.features}</span>
        <span class="text-[0.78rem] font-bold text-[#6b5a4b]">{state.appName}</span>
      </div>

      <div class="space-y-3">
        <div class="flex flex-col gap-2 rounded-[18px] bg-[rgba(255,255,255,0.64)] p-4">
          <span class="text-[0.82rem] uppercase tracking-[0.08em] text-[#6b5a4b]">{text.status}</span>
          <code class="break-words text-[#23170f]">{state.message || 'ready'}</code>
        </div>

        <div class="flex flex-col gap-2 rounded-[18px] bg-[rgba(255,255,255,0.64)] p-4">
          <span class="text-[0.82rem] uppercase tracking-[0.08em] text-[#6b5a4b]">{text.startupError}</span>
          <code class="break-words text-[#23170f]">{state.startupError || text.noError}</code>
        </div>

        <div class="flex flex-col gap-2 rounded-[18px] bg-[rgba(255,255,255,0.64)] p-4">
          <span class="text-[0.82rem] uppercase tracking-[0.08em] text-[#6b5a4b]">{text.locale}</span>
          <code class="break-words text-[#23170f]">{text.localeHint}</code>
        </div>
      </div>
    </article>

    <article class="rounded-3xl border border-[rgba(66,40,14,0.12)] bg-[rgba(255,255,255,0.72)] p-6 shadow-[0_18px_36px_rgba(96,53,23,0.08)]">
      <div class="mb-[18px] flex items-center justify-between gap-3 text-[0.92rem] font-extrabold uppercase tracking-[0.08em]">
        <span>{text.actions}</span>
        <span class="text-[0.78rem] font-bold text-[#6b5a4b]">{busy ? 'busy' : 'ready'}</span>
      </div>

      <div class="flex flex-wrap gap-3">
        <button
          class="cursor-pointer rounded-full border-0 bg-[#bf5a2b] px-[18px] py-3 text-[#fffaf2] transition-[transform,background-color,color,opacity] duration-150 ease-out hover:-translate-y-px disabled:cursor-not-allowed disabled:opacity-[0.55]"
          disabled={busy}
          onclick={showWindow}
          type="button"
        >
          {text.showWindow}
        </button>
        <button
          class="cursor-pointer rounded-full border-0 bg-[rgba(35,23,15,0.08)] px-[18px] py-3 text-[#23170f] transition-[transform,background-color,color,opacity] duration-150 ease-out hover:-translate-y-px disabled:cursor-not-allowed disabled:opacity-[0.55]"
          disabled={busy}
          onclick={hideWindow}
          type="button"
        >
          {text.hideWindow}
        </button>
        <button
          class="cursor-pointer rounded-full border-0 bg-[#8a1c1c] px-[18px] py-3 text-[#fffaf2] transition-[transform,background-color,color,opacity] duration-150 ease-out hover:-translate-y-px disabled:cursor-not-allowed disabled:opacity-[0.55]"
          disabled={busy}
          onclick={quitApp}
          type="button"
        >
          {text.quit}
        </button>
      </div>
    </article>

    <article class="rounded-3xl border border-[rgba(66,40,14,0.12)] bg-[rgba(255,255,255,0.72)] p-6 shadow-[0_18px_36px_rgba(96,53,23,0.08)]">
      <div class="mb-[18px] flex items-center justify-between gap-3 text-[0.92rem] font-extrabold uppercase tracking-[0.08em]">
        <span>{text.tips}</span>
        <span class="text-[0.78rem] font-bold text-[#6b5a4b]">{state.appName}</span>
      </div>

      <ul class="m-0 list-disc space-y-[10px] pl-[18px] leading-[1.7] text-[#6b5a4b]">
        {#each text.notes as note}
          <li>{note}</li>
        {/each}
      </ul>
    </article>
  </section>
</main>
