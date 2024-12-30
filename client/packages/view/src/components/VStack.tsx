import { FC } from 'react'

import { View, ViewProps } from './View'

export const VStack: FC<ViewProps> = (props) => {
  return <View flexDirection="column" {...props} />
}
