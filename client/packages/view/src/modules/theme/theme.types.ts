import { SignalOrValue } from 'core/modules/signals/signals.types'

import { Color } from '../color/color.types'
import { ThemeFont } from '../text/font.types'

export type ThemeDimensionType = Color | string | number | ThemeFont
export type ThemeDimension<TValue extends ThemeDimensionType = ThemeDimensionType> = Record<
  string,
  TValue
>
export type ThemeDimensionKey<TThemeDimension extends ThemeDimension> = keyof TThemeDimension
export type ThemeDimensionValue<TThemeDimension extends ThemeDimension> =
  TThemeDimension[ThemeDimensionKey<TThemeDimension>]

export interface Theme {
  readonly color?: ThemeDimension<string>
  readonly aspectRatio?: ThemeDimension<string | number>
  readonly space?: ThemeDimension<number>
  readonly size?: ThemeDimension<number>
  readonly borderRadius?: ThemeDimension<number>
  // Text
  readonly fontFamily?: ThemeDimension<ThemeFont>
  readonly fontSize?: ThemeDimension<number>
}
export type ThemeScaleKey =
  | 'color'
  | 'aspectRatio'
  | 'space'
  | 'size'
  | 'borderRadius'
  | 'fontFamily'
  | 'fontSize'

export type ThemeScale<
  TTheme extends Theme,
  TKey extends ThemeScaleKey,
> = TTheme[TKey] extends ThemeDimension ? ThemeDimensionKey<TTheme[TKey]> : never

export type ThemeValue<
  TTheme extends Theme,
  TKey extends ThemeScaleKey,
> = TTheme[TKey] extends ThemeDimension ? ThemeDimensionValue<TTheme[TKey]> : never

export type StyleDimensions<TStyle extends object> = Partial<Record<keyof TStyle, ThemeScaleKey>>
export type StyleDimensionsValue<
  TTheme extends Theme,
  TStyle extends object,
  TStyleDimensions extends StyleDimensions<TStyle>,
  TKey extends keyof TStyle,
> = TStyleDimensions[TKey] extends ThemeScaleKey
  ? ThemeScale<TTheme, TStyleDimensions[TKey]> extends undefined
    ? Exclude<TStyle[TKey], 'undefined'>
    : ThemeScale<TTheme, TStyleDimensions[TKey]>
  : Exclude<TStyle[TKey], 'undefined'>

export type ThemeStyle<
  TTheme extends Theme,
  TStyle extends object,
  TDimensions extends StyleDimensions<TStyle>,
> = {
  [K in keyof TStyle]?: SignalOrValue<StyleDimensionsValue<TTheme, TStyle, TDimensions, K>>
}
