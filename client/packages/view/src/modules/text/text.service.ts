import { TextStyle } from 'react-native'

import { compileThemeStyle } from '../theme/theme.service'
import { StyleDimensions, Theme, ThemeStyle } from '../theme/theme.types'
import { textStyleDimensions } from './text.types'

export function compileThemeTextStyle<TTheme extends Theme>(
  theme: TTheme,
  style: ThemeStyle<TTheme, TextStyle, StyleDimensions<TextStyle>>,
): TextStyle {
  return compileThemeStyle(theme, textStyleDimensions, style)
}
