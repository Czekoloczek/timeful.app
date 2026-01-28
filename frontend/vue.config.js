const { defineConfig } = require("@vue/cli-service")
const { DefinePlugin } = require("webpack")
const { execSync } = require("child_process")
const pkg = require("./package.json")

const getGitSha = () => {
  try {
    return execSync("git rev-parse --short HEAD").toString().trim()
  } catch (error) {
    return "unknown"
  }
}
module.exports = defineConfig({
  configureWebpack: {
    plugins: [
      new DefinePlugin({
        "process.env.VUE_APP_VERSION": JSON.stringify(pkg.version),
        "process.env.VUE_APP_COMMIT": JSON.stringify(getGitSha()),
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
