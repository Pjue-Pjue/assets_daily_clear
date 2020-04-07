module.exports = {
  lintOnSave: false,
  publicPath: '/',
  outputDir: 'dist', // 默认
  pwa: {
    iconPaths: {
      favicon32: 'favicon.png',
      favicon16: 'favicon.png',
      appleTouchIcon: 'favicon.png',
      maskIcon: 'favicon.png',
      msTileImage: 'favicon.png'
    }
  },
  productionSourceMap: false, // 加速生产环境构建
  devServer: {
    overlay: {
      warnings: true,
      errors: true
    },
    open: true,
    host: '127.0.0.1',
    port: 8868,
    https: false,
    // proxy: null
    proxy: { // 设置开发环境API服务器请求代理
      '/api/*': {
        target: 'http://127.0.0.1:3001/api',
        pathRewrite: { '^/api': '' },
        ws: true,
        changeOrigin: true
      },
    }
    // hotOnly: false,
    // before: app => {}
  },
  // 传递第三方插件选项
  pluginOptions: {},
  // 全局注入通用样式
  css: {}
};
