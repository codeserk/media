module.exports = function (api) {
  api.cache(true)
  return {
    presets: ['babel-preset-expo'],
    plugins: ['module:@preact-signals/safe-react/babel'],
  }
}
