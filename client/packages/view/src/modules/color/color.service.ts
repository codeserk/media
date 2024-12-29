import { argbFromHex, Hct, hexFromArgb } from '@material/material-color-utilities'

import { Color } from './color.types'

/**
 * Darkens a hex color
 * @param color hex color (#123123)
 * @param percent hex percent (10% -> 10)
 * @returns darkened hex color
 */
export function darkenColor(color: string, percent: number): Color {
  let R = parseInt(color.substring(1, 3), 16)
  let G = parseInt(color.substring(3, 5), 16)
  let B = parseInt(color.substring(5, 7), 16)

  R = Math.floor((R * (100 - percent)) / 100)
  G = Math.floor((G * (100 - percent)) / 100)
  B = Math.floor((B * (100 - percent)) / 100)

  R = R < 255 ? R : 255
  G = G < 255 ? G : 255
  B = B < 255 ? B : 255

  const RR = R.toString(16).length == 1 ? '0' + R.toString(16) : R.toString(16)
  const GG = G.toString(16).length == 1 ? '0' + G.toString(16) : G.toString(16)
  const BB = B.toString(16).length == 1 ? '0' + B.toString(16) : B.toString(16)

  return `#${RR}${GG}${BB}`
}

/**
 * Lightens a hex color
 * @param color hex color (#123123)
 * @param percent hex percent (10% -> 10)
 * @returns lightened hex color
 */
export function lightenColor(color: string, percent: number) {
  return darkenColor(color, -percent)
}

export enum ColorShade {
  Shade100 = 100,
  Shade200 = 200,
  Shade300 = 300,
  Shade400 = 400,
  Shade500 = 500,
  Shade600 = 600,
  Shade700 = 700,
  Shade800 = 800,
  Shade900 = 900,
}

/**
 * Creates shades for a given color
 * @param name
 * @param color
 * @param position Tone position of the given color
 * @param steps Tone variations for each shade
 * @returns shades
 */
export function createColorShades<TName extends string>(
  name: TName,
  color: string,
  position = 5,
  steps = 3,
): Record<TName | `${TName}-${ColorShade}`, Color> {
  const hct = Hct.fromInt(argbFromHex(color.toLocaleLowerCase()))
  const result: Record<`${TName}-${ColorShade}`, string> = Array.from(Array(9).keys()).reduce(
    (result, i) => {
      const key = `${name}-${i + 1}00` as `${TName}-${ColorShade}`

      const diff = position - i

      const newTone = Math.min(100, Math.max(0, hct.tone + diff * steps))
      const newHct = Hct.fromInt(hct.toInt())
      newHct.tone = newTone
      result[key] = hexFromArgb(newHct.toInt()) as Color

      return result
    },
    {} as Record<`${TName}-${ColorShade}`, Color>,
  )

  return { [name]: color, ...result } as Record<TName | `${TName}-${ColorShade}`, Color>
}
