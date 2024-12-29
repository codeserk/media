import { FlexStyle, ShadowStyleIOS, ViewStyle } from 'react-native'

import { compileThemeStyle } from '../theme/theme.service'
import { StyleDimensions, Theme, ThemeStyle } from '../theme/theme.types'
import { flexStyleDimensions, shadowStyleDimensions, viewStyleDimensions } from './view.types'

export function compileThemeFlexStyle<TTheme extends Theme>(
  theme: TTheme,
  style: ThemeStyle<TTheme, FlexStyle, StyleDimensions<FlexStyle>>,
): FlexStyle {
  return compileThemeStyle(theme, flexStyleDimensions, style)
}

export function compileThemeShadowStyle<TTheme extends Theme>(
  theme: TTheme,
  style: ThemeStyle<TTheme, ShadowStyleIOS, StyleDimensions<ShadowStyleIOS>>,
): ShadowStyleIOS {
  return compileThemeStyle(theme, shadowStyleDimensions, style)
}

export function compileThemeViewStyle<TTheme extends Theme>(
  theme: TTheme,
  style: ThemeStyle<TTheme, ViewStyle, StyleDimensions<ViewStyle>>,
): ViewStyle {
  return compileThemeStyle(theme, viewStyleDimensions, style)
}
