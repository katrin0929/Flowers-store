export default {
  // Глобальная настройка заголовка страницы
  head: {
    title: 'Мой сайт',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Глобальные CSS файлы
  css: ['./assets/css/index.css'],

  image: {
    provider: 'none',
  },

  // Плагины для загрузки перед инициализацией приложения
  plugins: [],

  // Модули Nuxt.js
  modules: [],

  // Параметры сборки
  build: {},

  compatibilityDate: '2025-04-03'
};