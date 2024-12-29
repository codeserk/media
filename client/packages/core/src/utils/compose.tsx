import React, { FC, ProviderExoticComponent, ReactNode } from 'react'

type Components = [ProviderExoticComponent<any>, Record<string, any>]

interface Props {
  readonly components: Components[]
  readonly children: ReactNode
}

export const Compose: FC<Props> = ({ components, children }) => (
  <>
    {components.reverse().reduce((acc, curr) => {
      const [Provider, props] = Array.isArray(curr) ? [curr[0], curr[1]] : [curr, {}]
      return <Provider {...props}>{acc}</Provider>
    }, children)}
  </>
)
