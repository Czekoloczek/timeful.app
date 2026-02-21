const { defineConfig } = require("@vue/cli-service")
const { DefinePlugin } = require("webpack")
const { execSync } = require("child_process")
const pkg = require("./package.json")

const getGitSha = () => {
  // Use CI-provided commit SHA if available (e.g. GitHub Actions sets GITHUB_SHA)
  if (process.env.COMMIT_SHA) return process.env.COMMIT_SHA.slice(0, 7)
  if (process.env.GITHUB_SHA) return process.env.GITHUB_SHA.slice(0, 7)
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
