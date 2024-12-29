import { signalValue } from 'core/modules/signals/signals'

import { StyleDimensions, Theme, ThemeScaleKey, ThemeStyle } from './theme.types'

export function compileThemeStyle<TTheme extends Theme, TStyle extends object>(
  theme: TTheme,
  dimensions: StyleDimensions<TStyle>,
  style: ThemeStyle<TTheme, TStyle, StyleDimensions<TStyle>>,
): TStyle {
  const result: TStyle = {} as TStyle
  for (const k in style) {
    const key = k as keyof typeof style
    const value = signalValue(style[key])
    result[key] = compileThemeStyleKey<TTheme, TStyle>(theme, dimensions, key, value)
  }

  return result
}

function compileThemeStyleKey<
  TTheme extends Theme,
  TStyle extends object,
  TKey extends keyof ThemeStyle<TTheme, TStyle, StyleDimensions<TStyle>> = keyof ThemeStyle<
    TTheme,
    TStyle,
    StyleDimensions<TStyle>
  >,
>(theme: TTheme, dimensions: StyleDimensions<TStyle>, key: TKey, value: any): any {
  if (!(key in dimensions)) {
    return value
  }

  const themeScale = dimensions[key as keyof typeof dimensions] as ThemeScaleKey
  const themeDimensions = theme[themeScale]
  if (!themeDimensions) {
    return value
  }

  if (!(value in themeDimensions)) {
    return value
  }

  return themeDimensions[value]
}
