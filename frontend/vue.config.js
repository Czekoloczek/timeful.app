const { defineConfig } = require("@vue/cli-service")
const { DefinePlugin } = require("webpack")
const pkg = require("./package.json")
module.exports = defineConfig({
  configureWebpack: {
    plugins: [
      new DefinePlugin({
        "process.env.VUE_APP_VERSION": JSON.stringify(pkg.version),
      }),
    ],
  },
  transpileDependencies: ["vuetify"],
  // publicPath: "/dist",
  // pluginOptions: {
  //   webpackBundleAnalyzer: {
  //     openAnalyzer: false,
  //   },
  // },
})
