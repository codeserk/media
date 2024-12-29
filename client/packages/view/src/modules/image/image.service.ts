import { ImageStyle } from 'react-native'

import { compileThemeStyle } from '../theme/theme.service'
import { StyleDimensions, Theme, ThemeStyle } from '../theme/theme.types'
import { imageStyleDimensions } from './image.types'

export function compileThemeImageStyle<TTheme extends Theme>(
  theme: TTheme,
  style: ThemeStyle<TTheme, ImageStyle, StyleDimensions<ImageStyle>>,
): ImageStyle {
  return compileThemeStyle(theme, imageStyleDimensions, style)
}
